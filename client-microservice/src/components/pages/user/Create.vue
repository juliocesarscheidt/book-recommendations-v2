<template>
  <section class="flex flex-column flex-align-center pt-5 pb-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <div style="width: 100%; min-width: 250px; margin-bottom: 0px;">
        <div class="form-group">
          <label>{{ $t('user.name') }}</label>
          <input type="text" class="form-control" v-model.trim="name">
        </div>
        <div class="form-group">
          <label>{{ $t('user.surname') }}</label>
          <input type="text" class="form-control" v-model.trim="surname">
        </div>
        <div class="form-group">
          <label>{{ $t('user.email') }}</label>
          <input type="email" class="form-control" v-model.trim="email">
        </div>
        <div class="form-group">
          <label>{{ $t('user.phone') }}</label>
          <input type="phone" class="form-control" v-model.trim="phone" v-format-phone="phone">
        </div>
        <div class="form-group">
          <label>{{ $t('user.password') }}</label>
          <input type="password" class="form-control" v-model.trim="password">
        </div>

        <button type="button" class="btn btn-outline-primary btn-lg btn-block mt-4" @click="createUserFn">
          {{ $t('buttons.save') }}
        </button>

        <button type="button" class="btn btn-outline-secondary btn-lg btn-block mt-4" @click="$router.push({ name: 'UserList' })">
          {{ $t('buttons.return') }}
        </button>
      </div>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { createUser } from '../../../services/';

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
    async createUserFn() {
      if (!this.name || !this.surname || !this.email || !this.phone || !this.password) {
        alert('Invalid Entries');
        return;
      }

      try {
        const uuid = await createUser(this.name, this.surname, this.email, this.removeNonNumericDigits(this.phone), this.password);
        this.notifySuccess(this.$t('messages.success.created_with_success'));
        this.$router.push({ name: 'UserView', params: { uuid, isEdit: false } });

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
