<template>
  <section class="flex flex-column flex-align-center p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <form style="width: 100%; min-width: 200px; margin-bottom: 0px;" v-if="userData">
        <div class="form-group">
          <label for="input-name">{{ $t('user.name') }}</label>
          <b-form-input type="text" v-bind:disabled="!isEdit || loading" v-model.trim="userData.name"></b-form-input>
        </div>
        <div class="form-group">
          <label for="input-surname">{{ $t('user.surname') }}</label>
          <b-form-input type="text" v-bind:disabled="!isEdit || loading" v-model.trim="userData.surname"></b-form-input>
        </div>
        <div class="form-group">
          <label for="input-email">{{ $t('user.email') }}</label>
          <b-form-input type="text" v-bind:disabled="!isEdit || loading" v-model.trim="userData.email"></b-form-input>
        </div>
        <div class="form-group">
          <label for="input-phone">{{ $t('user.phone') }}</label>
          <b-form-input type="text" v-bind:disabled="!isEdit || loading" v-model.trim="userData.phone"></b-form-input>
        </div>

        <button type="submit" class="btn btn-outline-primary btn-lg btn-block mt-4" v-if="!isEdit" @click="callEditUser">
          {{ $t('buttons.edit') }}
        </button>

        <button type="submit" class="btn btn-outline-primary btn-lg btn-block mt-4" v-if="isEdit" @click="callUpdateUser">
          {{ $t('buttons.save') }}
        </button>

        <button type="submit" class="btn btn-outline-danger btn-lg btn-block mt-4" @click="callDeleteUser">
          {{ $t('buttons.delete') }}
        </button>
      </form>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { getUser, updateUser, deleteUser } from '../../../services/api';

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
      userData: {},
    }
  },
  computed: {
    ...mapState([]),
    ...mapGetters([]),
  },
  watch: {
    '$route': {
      handler(to, from) {
        if (from.params.uuid !== to.params.uuid) {
          this.callGetUser();
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
  async mounted() {
    await this.callGetUser();
  },
  methods: {
    replaceRoute(editMode) {
      this.$router.replace({ params: { uuid: this.uuid, isEdit: editMode} });
    },
    callEditUser() {
      this.replaceRoute(true);
    },
    async callGetUser() {
      if (!this.uuid) {
        this.notifyError('User Not Found');
        this.$router.push({ name: 'UserList' });
        return;
      }

      this.loading = true;
      try {
        const response = await getUser(this.uuid);
        this.userData = response;

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);
        this.$router.push({ name: 'UserList' });

      } finally {
        this.loading = false;
      }
    },
    async callUpdateUser() {
      if (!this.uuid) {
        this.notifyError('User Not Found');
        this.$router.push({ name: 'UserList' });
        return;
      }

      this.loading = true;
      try {
        await updateUser(this.uuid, {
          name: this.userData.name,
          surname: this.userData.surname,
          email: this.userData.email,
          phone: this.userData.phone,
        });

      } catch (err) {
        console.log(err);
        this.notifyError(err.response.data.message);

      } finally {
        this.loading = false;
        this.replaceRoute(false);
      }
    },
    async callDeleteUser() {
      if (!this.uuid) {
        this.notifyError('User Not Found');
        this.$router.push({ name: 'UserList' });
        return;
      }

      try {
        await deleteUser(this.uuid);
        this.$router.push({ name: 'UserList' });

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
