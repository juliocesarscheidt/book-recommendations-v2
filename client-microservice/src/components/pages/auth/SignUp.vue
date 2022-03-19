<template>
  <section class="flex flex-column flex-align-center p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <form style="width: 100%; min-width: 200px; margin-bottom: 0px;">
        <div class="form-group">
          <label for="input-name">Name</label>
          <input type="text" class="form-control" id="input-name" placeholder="Name" v-model.trim="name">
        </div>
        <div class="form-group">
          <label for="input-surname">Surname</label>
          <input type="text" class="form-control" id="input-surname" placeholder="Surname" v-model.trim="surname">
        </div>
        <div class="form-group">
          <label for="input-email">Email</label>
          <input type="email" class="form-control" id="input-email" placeholder="Email" v-model.trim="email">
        </div>
        <div class="form-group">
          <label for="input-phone">Phone</label>
          <input type="phone" class="form-control" id="input-phone" placeholder="Phone" v-model.trim="phone">
        </div>
        <div class="form-group">
          <label for="input-password">Password</label>
          <input type="password" class="form-control" id="input-password" placeholder="Password" v-model.trim="password">
        </div>
        <button type="submit" class="btn btn-outline-primary btn-lg btn-block mt-4" @click="userSignUp">
          Sign Up
        </button>
      </form>
    </article>

    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 200px;">
        <p>
          Or sign-in
        </p>
        <button class="btn btn-outline-secondary btn-lg btn-block" @click="goUserSignIn">
          Sign In
        </button>
      </div>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { signUp, getCurrentUserInfo } from '../../../services/api';

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
      name: 'julio',
      surname: 'cesar',
      email: 'julio@mail.com',
      phone: '41995540808',
      password: 'PASSWORD',
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
    goUserSignIn() {
      this.$router.push({ name: 'SignIn' });
    },
    async userSignUp() {
      if (!this.name || !this.surname || !this.email || !this.phone || !this.password) {
        alert('Invalid Entries');
        return;
      }

      try {
        const token = await signUp(this.name, this.surname, this.email, this.phone, this.password);
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
