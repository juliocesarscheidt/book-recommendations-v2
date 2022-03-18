<template>
  <section class="flex flex-column flex-space-between p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <form style="width: 100%;">
        <div class="form-group">
          <label for="input-email">Email</label>
          <input type="email" class="form-control" id="input-email" placeholder="Email" v-model.trim="email">
        </div>
        <div class="form-group">
          <label for="input-password">Password</label>
          <input type="password" class="form-control" id="input-password" placeholder="Password" v-model.trim="password">
        </div>
        <button type="submit" class="btn btn-outline-secondary btn-lg btn-block mt-4" @click="userSignIn">Sign In</button>
      </form>

      <button type="button" class="btn btn-outline-secondary btn-lg btn-block mt-4" @click="userSignUp">Sign Up</button>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { signUp, signIn } from '../../services/api';

export default {
  computed: {
    ...mapState([]),
  },
  data() {
    return {
      email: 'julio@mail.com',
      password: 'PASSWORD',
    }
  },
  beforeMount() {
  },
  mounted() {
  },
  methods: {
    async userSignIn() {
      if (!this.email || !this.password) {
        alert('Invalid Entries');
        return;
      }

      const response = await signIn(this.email, this.password);
      console.log('signIn response', response);

      // this.$store.dispatch('setUser', user);
      // this.$router.push({ name: 'Logged' });

      // this.$store.dispatch('setUser', null);
      // this.$router.push({ name: 'Home' });
    },
    async userSignUp() {
      if (!this.email || !this.password) {
        alert('Invalid Entries');
        return;
      }

      const response = await signUp(this.email, this.password);
      console.log('signUp response', response);

      // this.$store.dispatch('setUser', user);
      // this.$router.push({ name: 'Logged' });
    },
  },
  beforeDestroy() {}
}
</script>

<style scoped>
section {
  min-height: calc(100vh - 42px);
  width: auto;
}

section > article {
  min-height: 150px;
  height: auto;
  padding: 20px;
}

.main-article {
  flex-direction: row;
  align-items: center;
}

section > article > div.article-main-text > p {
  font-size: 3rem;
  font-weight: bold;
  max-width: 200px;
}

section > article > div.article-main-image > img {
  max-height: 500px;
  min-height: 400px;
  width: 100%;
}

@media only screen and (max-width: 991px) {
  .main-article {
    flex-direction: column;
    align-items: flex-start;
  }

  section > article > div.article-main-image > img {
    max-height: 300px;
    min-height: 100px;
  }
}
</style>