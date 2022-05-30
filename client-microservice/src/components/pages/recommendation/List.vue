<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center" style="width: 100%;">
      <b-table
        class="lg-2 md-3 sm-4"
        striped
        hover
        outlined
        responsive
        show-empty
        :sortable="false"
        :fields="fields"
        :items="recommendationsData"
        :busy="loading"
      >
        <template #table-busy>
          <div class="text-center text-danger my-2">
            <b-spinner class="align-middle"></b-spinner>
            <strong>{{ $t('buttons.loading') }}...</strong>
          </div>
        </template>
        <template #empty="">
          <p class="text-center"><b>{{ $t('messages.info.data_not_found') }}</b></p>
        </template>

        <template #cell(title)="data">
          <i>{{ data.value | capitalize }}</i>
        </template>
        <template #cell(author)="data">
          <i>{{ data.value | capitalize }}</i>
        </template>
        <template #cell(genre)="data">
          <i>{{ data.value | capitalize }}</i>
        </template>

        <template #cell(image)="data">
          <div v-if="data.value" class="text-nowrap text-link" @click="openImageUrl(data.item.uuid)" v-bind:title="$t('book.image_preview')">{{ data.value | trimLetters(25) }}</div>
          <div v-else class="text-nowrap text-link">-</div>
        </template>

        <template #head(title)="">
          <div class="text-nowrap">{{ $t('book.title') }}</div>
        </template>
        <template #head(author)="">
          <div class="text-nowrap">{{ $t('book.author') }}</div>
        </template>
        <template #head(genre)="">
          <div class="text-nowrap">{{ $t('book.genre') }}</div>
        </template>
        <template #head(image)="">
          <div class="text-nowrap">{{ $t('book.image') }}</div>
        </template>
        <template #head(rate)="">
          <div class="text-nowrap">{{ $t('book.rating') }}</div>
        </template>
        <template #head(uuid)="">
          <div class="text-nowrap">{{ $t('buttons.actions') }}</div>
        </template>

        <template #cell(uuid)="data">
          <div style="white-space: nowrap">
            <b-button type="button" variant="outline-primary" class="mr-2" v-bind:title="$t('buttons.view')" @click="callGetBook(data.value)">
              <i class="fas fa-eye" aria-hidden="true"></i>
            </b-button>
          </div>
        </template>
      </b-table>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import { getBookPresignUrl, listBook, getRecommendations } from '../../../services/';

export default {
  components: {
  },
  directives: {
  },
  filters: {
  },
  mixins: [
  ],
  props: {
  },
  data() {
    return {
      loading: false,
      books: [],
      fields: ['title', 'author', 'genre', 'image', 'uuid'],
      recommendationsData: [],
    }
  },
  computed: {
    ...mapState([]),
    ...mapGetters(['user']),
  },
  watch: {
  },
  beforeMount() {
  },
  mounted() {
    this.callRefreshData();
  },
  methods: {
    callGetBook(uuid) {
      this.$router.push({ name: 'BookView', params: { uuid, isEdit: false } });
    },
    async callRefreshData() {
      await this.callListBook();
      await this.callGetRecommendations();
    },
    async callGetRecommendations() {
      if (this.user) {
        this.loading = true;
        try {
          const response = await getRecommendations(this.user.uuid);
          if (response.recommendations && response.recommendations.length) {
            this.recommendationsData = response.recommendations
              .map((r) => {
                const book = this.books.find(b => b.uuid === r.book_uuid);
                if (!book) return null;
                return {
                  uuid: book.uuid,
                  title: book.title,
                  author: book.author,
                  genre: book.genre,
                  image: book.image,
                };
              })
              .filter(b => b !== null);
          }

        } catch (err) {
          console.log(err);
          this.notifyError(err.response.data.message);
          this.$router.push({ name: 'Home' });

        } finally {
          this.loading = false;
        }
      }
    },
    async callListBook() {
      this.loading = true;
      try {
        const response = await listBook();
        this.books = response;

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
        this.$router.push({ name: 'Home' });

      } finally {
        this.loading = false;
      }
    },
    async openImageUrl(uuid) {
      try {
        const url = await getBookPresignUrl(uuid);
        window.open(url, "_blank");
      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
      }
    }
  },
  beforeDestroy() {
  },
  destroyed() {
  },
}
</script>

<style scoped>
section {
  min-height: calc(100vh - 42px);
  width: auto;
}

section > article {
  height: auto;
  padding: 20px;
}

@media only screen and (max-width: 991px) {
}
</style>
