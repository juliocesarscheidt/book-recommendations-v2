<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center" style="width: 100%;">
      <button type="button" class="btn btn-outline-secondary btn-lg mb-4 ml-auto" @click="callCreateBook">
        {{ $t('header.create_book') }}
      </button>

      <b-table
        class="lg-2 md-3 sm-4"
        striped
        hover
        outlined
        responsive
        show-empty
        :sortable="false"
        :fields="fields"
        :items="booksData"
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
            <b-button type="button" variant="outline-secondary" class="mr-2" v-bind:title="$t('buttons.edit')" @click="callUpdateBook(data.value)">
              <i class="fas fa-edit" aria-hidden="true"></i>
            </b-button>
            <b-button type="button" variant="outline-danger" class="mr-2" v-bind:title="$t('buttons.delete')" @click="callDeleteBook(data.value)">
              <i class="fas fa-eraser" aria-hidden="true"></i>
            </b-button>
          </div>
        </template>
      </b-table>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import { getBookPresignUrl, listBook, deleteBook, getRate } from '../../../services/';

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
      booksData: [],
      fields: ['title', 'author', 'genre', 'image', 'rate', 'uuid'],
      rates: [],
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
    callCreateBook() {
      this.$router.push({ name: 'BookCreate' });
    },
    callGetBook(uuid) {
      this.$router.push({ name: 'BookView', params: { uuid, isEdit: false } });
    },
    callUpdateBook(uuid) {
      this.$router.push({ name: 'BookView', params: { uuid, isEdit: true } });
    },
    async callDeleteBook(uuid) {
      try {
        await deleteBook(uuid);
        this.notifySuccess(this.$t('messages.success.deleted_with_success'));
        this.callRefreshData();

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
      }
    },
    async callRefreshData() {
      await this.callListBook();
      await this.callGetRate();
    },
    async callListBook() {
      this.loading = true;
      try {
        const response = await listBook();
        this.booksData = response.map((book) => ({ ...book, rate: 0.00 }));

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
        this.$router.push({ name: 'Home' });

      } finally {
        this.loading = false;
      };
    },
    async callGetRate() {
      if (this.user) {
        try {
          const response = await getRate(this.user.uuid);
          this.rates = response.rates;

          if (this.rates && this.rates.length) {
            this.booksData = this.booksData.map((book) => {
              const rate = this.rates.find(b => b.book_uuid === book.uuid);
              return { ...book, rate: rate && parseFloat(rate.rate.toFixed(2)) || 0.00 }
            });
          }

        } catch (err) {
          console.log(err);
          this.notifyError(err.response.data.message);
          this.$router.push({ name: 'Home' });
        }
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
