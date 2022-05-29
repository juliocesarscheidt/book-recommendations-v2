<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 200px; margin-bottom: 0px;">
        <div class="form-group">
          <label>{{ $t('book.title') }}</label>
          <input type="text" class="form-control" v-model.trim="title">
        </div>
        <div class="form-group">
          <label>{{ $t('book.author') }}</label>
          <input type="text" class="form-control" v-model.trim="author">
        </div>
        <div class="form-group">
          <label>{{ $t('book.genre') }}</label>
          <input type="text" class="form-control" v-model.trim="genre">
        </div>
        <div class="form-group">
          <label>{{ $t('book.image') }}</label>
          <div class="custom-file mb-2">
            <input type="file" class="custom-file-input" id="inputGroupFile" style="display: none;" value="" accept="image/*" @change="onFileSelected">
            <label class="custom-file-label" for="inputGroupFile">{{ $t('book.choose_file') }}</label>
          </div>
          <div class="flex flex-column flex-justify-center flex-align-center" v-if="imagePreview && imageSelected">
            <span style="text-align: center;">{{ $t('book.image_preview') }}</span>
            <img style="height: auto; max-height: 200px; width: 200px; margin: 0; padding: 0; box-shadow: 0 0 20px 0.25px rgba(0, 0, 0, .25);" v-bind:src="imagePreview" v-bind:alt="$t('book.image_preview')">
            <code style="width: 200px; margin: 0; padding: 0; text-align: center;">{{ imageSelected.name }}</code>
          </div>
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
      imageSelected: '',
      imagePreview: '',
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
      if (!this.title || !this.author || !this.genre || !this.imageSelected) {
        alert('Invalid Entries');
        return;
      }
      try {
        const uuid = await createBook(this.title, this.author, this.genre, this.imageSelected);
        this.notifySuccess(this.$t('messages.success.created_with_success'));
        this.$router.push({ name: 'BookView', params: { uuid, isEdit: false } });

      } catch (err) {
        console.error(err);
        this.notifyError(err.response.data.message);
      }
    },
    onFileSelected (event) {
      this.imageSelected = event.target.files[0];
      const vm = this;
      const readerUrl = new FileReader();
      readerUrl.readAsDataURL(this.imageSelected);
      readerUrl.onload = function(e) {
        console.error(e.target);
        vm.imagePreview = e.target.result;
      }
      readerUrl.onerror = function(e) {
        console.error(e);
        alert('Error uploading file');
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
