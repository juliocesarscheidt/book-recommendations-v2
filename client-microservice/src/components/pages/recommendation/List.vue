<template>
  <section class="flex flex-column flex-align-center p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <b-table
        style="width: 100%; min-width: 200px; margin-bottom: 0px;"
        striped
        hover
        outlined
        fixed
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
            <strong>Loading...</strong>
          </div>
        </template>

        <template #empty="">
          <p style="text-align: center;"><b>Data not Found</b></p>
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
          <a v-bind:href="data.value" target="_blank" title="See Image">{{ data.value | trimLetters(25) }}</a>
        </template>

        <template #head(uuid)="">
          <div class="text-nowrap">Actions</div>
        </template>

        <template #cell(uuid)="data">
          <b-button type="button" variant="primary" class="mr-2" title="View" @click="callGetBook(data.value)">
            <i class="fas fa-eye" aria-hidden="true"></i>
          </b-button>
        </template>
      </b-table>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { listBook, getRecommendations } from '../../../services/api';

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
  async mounted() {
    await this.callRefreshData();
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
        try {
          const response = await getRecommendations(this.user.uuid);
          if (response.recommendations && response.recommendations.length) {
            this.recommendationsData = response.recommendations.map((r) => {
              const book = this.books.find(b => b.uuid === r.book_uuid);
              return {
                uuid: book.uuid,
                title: book.title,
                author: book.author,
                genre: book.genre,
                image: book.image,
              };
            })
          }

        } catch (err) {
          console.log(err);
          this.notifyError(err.response.data.message);
          this.$router.push({ name: 'Home' });
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
      };
    },
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
