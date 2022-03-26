<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 200px; margin-bottom: 0px;">
        <div class="form-group">
          <label for="input-title">{{ $t('book.title') }}</label>
          <input type="text" class="form-control" v-model.trim="title">
        </div>
        <div class="form-group">
          <label for="input-author">{{ $t('book.author') }}</label>
          <input type="text" class="form-control" v-model.trim="author">
        </div>
        <div class="form-group">
          <label for="input-genre">{{ $t('book.genre') }}</label>
          <input type="text" class="form-control" v-model.trim="genre">
        </div>
        <div class="form-group">
          <label for="input-image">{{ $t('book.image') }}</label>
          <input type="text" class="form-control" v-model.trim="image">
        </div>

        <button type="button" class="btn btn-outline-primary btn-lg btn-block mt-4" @click="callCreateBook">
          {{ $t('buttons.save') }}
        </button>
      </div>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { createBook } from '../../../services/';

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
      title: '',
      author: '',
      genre: '',
      image: '',
    }
  },
  computed: {
    ...mapState([]),
    ...mapGetters([]),
  },
  watch: {
  },
  beforeMount() {
  },
  mounted() {
  },
  methods: {
    async callCreateBook() {
      if (!this.title || !this.author || !this.genre || !this.image) {
        alert('Invalid Entries');
        return;
      }

      try {
        const uuid = await createBook(this.title, this.author, this.genre, this.image);
        this.notifySuccess(this.$t('messages.success.created_with_success'));
        this.$router.push({ name: 'BookView', params: { uuid, isEdit: false } });

      } catch (err) {
        console.error(err);
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
