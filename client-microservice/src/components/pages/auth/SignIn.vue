<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <form style="width: 100%; min-width: 250px; margin-bottom: 0px;" @keyup.enter="userSignIn">
        <div class="form-group">
          <label for="input-email">{{ $t('user.email') }}</label>
          <input type="email" class="form-control" v-model.trim="email">
        </div>
        <div class="form-group">
          <label for="input-password">{{ $t('user.password') }}</label>
          <input type="password" class="form-control" v-model.trim="password">
        </div>
        <button type="button" class="btn btn-outline-primary btn-lg btn-block mt-4" @click="userSignIn">
          {{ $t('buttons.signin') }}
        </button>
      </form>
    </article>

    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 250px;">
        <p>
          {{ $t('auth.or_signup') }}
        </p>
        <button class="btn btn-outline-secondary btn-lg btn-block" @click="goUserSignUp">
          {{ $t('buttons.signup') }}
        </button>
      </div>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { signIn, getCurrentUserInfo } from '../../../services/';

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
      email: '',
      password: '',
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
    goUserSignUp() {
      this.$router.push({ name: 'SignUp' });
    },
    async userSignIn() {
      if (!this.email || !this.password) {
        alert('Invalid Entries');
        return;
      }

      try {
        const token = await signIn(this.email, this.password);
        this.$store.dispatch('setToken', token);

        const user = await getCurrentUserInfo();
        this.$store.dispatch('setUser', user);

        if (this.$route.query && this.$route.query.redirect) {
          this.$router.push({ path: this.$route.query.redirect });
        } else {
          this.$router.push({ name: 'Home' });
        }

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
