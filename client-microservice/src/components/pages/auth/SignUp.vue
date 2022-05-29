<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <form style="width: 100%; min-width: 250px; margin-bottom: 0px;" @keyup.enter="userSignUp">
        <div class="form-group">
          <label for="input-name">{{ $t('user.name') }}</label>
          <input type="text" class="form-control" v-model.trim="name">
        </div>
        <div class="form-group">
          <label for="input-surname">{{ $t('user.surname') }}</label>
          <input type="text" class="form-control" v-model.trim="surname">
        </div>
        <div class="form-group">
          <label for="input-email">{{ $t('user.email') }}</label>
          <input type="email" class="form-control" v-model.trim="email">
        </div>
        <div class="form-group">
          <label for="input-phone">{{ $t('user.phone') }}</label>
          <input type="phone" class="form-control" v-model.trim="phone" v-format-phone="phone">
        </div>
        <div class="form-group">
          <label for="input-password">{{ $t('user.password') }}</label>
          <input type="password" class="form-control" v-model.trim="password">
        </div>
        <button type="button" class="btn btn-outline-primary btn-lg btn-block mt-4" @click="userSignUp">
          {{ $t('buttons.signup') }}
        </button>
      </form>
    </article>

    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 250px;">
        <p>
          {{ $t('auth.or_signin') }}
        </p>
        <button class="btn btn-outline-secondary btn-lg btn-block" @click="goUserSignIn">
          {{ $t('buttons.signin') }}
        </button>
      </div>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { signUp, getCurrentUserInfo } from '../../../services/';

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
      name: '',
      surname: '',
      email: '',
      phone: '',
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
    removeNonNumericDigits(val) {
      return val.replace(/\D/g, '');
    },
    goUserSignIn() {
      this.$router.push({ name: 'SignIn' });
    },
    async userSignUp() {
      if (!this.name || !this.surname || !this.email || !this.phone || !this.password) {
        alert('Invalid Entries');
        return;
      }

      try {
        const token = await signUp(this.name, this.surname, this.email, this.removeNonNumericDigits(this.phone), this.password);
        this.$store.dispatch('setToken', token);

        const user = await getCurrentUserInfo();
        this.$store.dispatch('setUser', user);

        this.$router.push({ name: 'Home' });

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
