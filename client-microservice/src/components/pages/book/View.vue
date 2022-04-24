<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 200px; margin-bottom: 0px;" v-if="bookData">
        <div class="form-group">
          <label for="input-title">{{ $t('book.title') }}</label>
          <input type="text" class="form-control" v-bind:disabled="!isEdit || loading" v-model.trim="bookData.title"></input>
        </div>
        <div class="form-group">
          <label for="input-author">{{ $t('book.author') }}</label>
          <input type="text" class="form-control" v-bind:disabled="!isEdit || loading" v-model.trim="bookData.author"></input>
        </div>
        <div class="form-group">
          <label for="input-genre">{{ $t('book.genre') }}</label>
          <input type="text" class="form-control" v-bind:disabled="!isEdit || loading" v-model.trim="bookData.genre"></input>
        </div>
        <div class="form-group">
          <label for="input-image">{{ $t('book.image') }}</label>
          <input type="text" class="form-control" v-bind:disabled="!isEdit || loading" v-model.trim="bookData.image"></input>
        </div>
        <div class="form-group">
          <label for="input-rate">{{ $t('book.rating') }} ({{ rate }})</label>
          <star-rating v-model="rate"
            v-bind:show-rating="false"
            v-bind:read-only="loading"
            v-bind:active-on-click="false"
            v-bind:star-size="25"
            v-bind:fixed-points=2
            v-bind:increment=0.01
            v-bind:max-rating=10
            @rating-selected ="callUpsertRate"
          >
          </star-rating>
        </div>

        <button type="button" class="btn btn-outline-primary btn-lg btn-block mt-4" v-if="!isEdit" @click="callEditBook">
          {{ $t('buttons.edit') }}
        </button>
        <button type="button" class="btn btn-outline-primary btn-lg btn-block mt-4" v-if="isEdit" @click="callUpdateBook">
          {{ $t('buttons.save') }}
        </button>
        <button type="button" class="btn btn-outline-danger btn-lg btn-block mt-4" @click="callDeleteBook">
          {{ $t('buttons.delete') }}
        </button>
      </div>
    </article>
  </section>
</template>

<script>
import StarRating from 'vue-star-rating';
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { getBook, updateBook, deleteBook, getRate, upsertRate } from '../../../services/';

export default {
  components: {
    StarRating,
  },
  directives: {
  },
  filters: {
  },
  mixins: [
  ],
  props: {
    uuid: {
      type: [String],
      required: true,
    },
    isEdit: {
      type: [Boolean, String],
      required: true,
      default: false,
    },
  },
  data() {
    return {
      loading: false,
      bookData: {},
      rate: 0,
    }
  },
  computed: {
    ...mapState([]),
    ...mapGetters(['user']),
  },
  watch: {
    '$route': {
      handler(to, from) {
        if (from.params.uuid !== to.params.uuid) {
          this.callRefreshData();
        }
      },
      deep: true,
    },
    'isEdit': {
      handler(to, from) {
        this.isEdit = Boolean(to);
      },
    },
  },
  beforeMount() {
  },
  mounted() {
    this.callRefreshData();
  },
  methods: {
    replaceRoute(editMode) {
      this.$router.replace({ params: { uuid: this.uuid, isEdit: editMode} });
    },
    callEditBook() {
      this.replaceRoute(true);
    },
    async callRefreshData() {
      await this.callGetBook();
      await this.callGetRate();
    },
    async callGetRate() {
      this.loading = true;
      if (this.user) {
        try {
          const response = await getRate(this.user.uuid);
          if (response && response.rates) {
            const rate = response.rates.find(b => b.book_uuid === this.uuid);
            this.rate = rate && parseFloat(rate.rate.toFixed(2)) || 0.00;
          }

        } catch (err) {
          console.log(err);
          this.notifyError(err.response.data.message);

        } finally {
          this.loading = false;
        }
      }
    },
    async callGetBook() {
      if (!this.uuid) {
        this.notifyError('Book Not Found');
        this.$router.push({ name: 'BookList' });
        return;
      }

      this.loading = true;
      try {
        const response = await getBook(this.uuid);
        this.bookData = response;

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
        this.$router.push({ name: 'BookList' });

      } finally {
        this.loading = false;
      }
    },
    async callUpsertRate(rating) {
      console.log('rating', rating);
      console.log('this.rate', this.rate);

      try {
        await upsertRate(this.user.uuid, this.uuid, parseFloat(this.rate.toFixed(2)));
        this.notifySuccess(this.$t('messages.success.updated_with_success'));

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
      }
    },
    async callUpdateBook() {
      if (!this.uuid) {
        this.notifyError('Book Not Found');
        this.$router.push({ name: 'BookList' });
        return;
      }

      this.loading = true;
      try {
        await updateBook(this.uuid, {
          title: this.bookData.title,
          author: this.bookData.author,
          genre: this.bookData.genre,
          image: this.bookData.image,
        });
        this.notifySuccess(this.$t('messages.success.updated_with_success'));

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);

      } finally {
        this.loading = false;
        this.replaceRoute(false);
      }
    },
    async callDeleteBook() {
      if (!this.uuid) {
        this.notifyError('Book Not Found');
        this.$router.push({ name: 'BookList' });
        return;
      }

      try {
        await deleteBook(this.uuid);
        this.notifySuccess(this.$t('messages.success.deleted_with_success'));
        this.$router.push({ name: 'BookList' });

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
      }
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
