<template>
  <section class="flex flex-column flex-align-center p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <p class="mb-4">
        {{ $t('home.intro_01') }} <b>{{ $t('home.intro_02') }}</b>
      </p>

      <button type="button" class="btn btn-outline-primary btn-lg btn-block mr-2" title="Users" @click="goTo('UserList')">
        {{ $t('header.users') }} <i class="fas fa-users" aria-hidden="true"></i>
      </button>

      <button type="button" class="btn btn-outline-primary btn-lg btn-block mr-2" title="Books" @click="goTo('BookList')">
        {{ $t('header.books') }} <i class="fas fa-book" aria-hidden="true"></i>
      </button>

      <button type="button" class="btn btn-outline-primary btn-lg btn-block mr-2" title="Recommendations" @click="goTo('RecommendationList')">
        {{ $t('header.recommendations') }} <i class="fas fa-star" aria-hidden="true"></i>
      </button>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'

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
    goTo(name) {
      this.$router.push({ name });
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