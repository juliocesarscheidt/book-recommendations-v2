<template>
  <section class="flex flex-column flex-align-center p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <form style="width: 100%; min-width: 200px; margin-bottom: 0px;">
        <div class="form-group">
          <label for="input-title">{{ $t('book.title') }}</label>
          <input type="text" class="form-control" id="input-title" placeholder="Name" v-model.trim="title">
        </div>
        <div class="form-group">
          <label for="input-author">{{ $t('book.author') }}</label>
          <input type="text" class="form-control" id="input-author" placeholder="Author" v-model.trim="author">
        </div>
        <div class="form-group">
          <label for="input-genre">{{ $t('book.genre') }}</label>
          <input type="text" class="form-control" id="input-genre" placeholder="Genre" v-model.trim="genre">
        </div>
        <div class="form-group">
          <label for="input-image">{{ $t('book.image') }}</label>
          <input type="text" class="form-control" id="input-image" placeholder="Image" v-model.trim="image">
        </div>
        <button type="submit" class="btn btn-outline-primary btn-lg btn-block mt-4" @click="createBookFn">
          {{ $t('buttons.save') }}
        </button>
      </form>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { createBook } from '../../../services/api';

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
      title: 'Clean Architecture',
      author: 'Robert Martin',
      genre: 'Software',
      image: 'https://images-na.ssl-images-amazon.com/images/I/41-sN-mzwKL._SX258_BO1,204,203,200_QL70_ML2_.jpg',
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
    async createBookFn() {
      if (!this.title || !this.author || !this.genre || !this.image) {
        alert('Invalid Entries');
        return;
      }

      try {
        const uuid = await createBook(this.title, this.author, this.genre, this.image);
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
