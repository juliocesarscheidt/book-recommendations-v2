<template>
  <section class="flex flex-column flex-align-center p-5">
    <article class="flex flex-column flex-justify-center flex-align-center">
      <b-table
        style="width: 100%; min-width: 200px; margin-bottom: 0px;"
        striped
        hover
        outlined
        fixed
        responsive
        show-empty
        :sortable="false"
        :fields="fields"
        :items="usersData"
        :busy="loading"
      >
        <template #table-busy>
          <div class="text-center text-danger my-2">
            <b-spinner class="align-middle"></b-spinner>
            <strong>Loading...</strong>
          </div>
        </template>

        <template #empty="">
          <p style="text-align: center;"><b>Data not Found</b></p>
        </template>

        <template #cell(name)="data">
          <i>{{ data.value | capitalize }}</i>
        </template>

        <template #cell(surname)="data">
          <i>{{ data.value | capitalize }}</i>
        </template>

        <template #cell(phone)="data">
          <i>{{ data.value | formatPhone }}</i>
        </template>

        <template #head(uuid)="">
          <div class="text-nowrap">Actions</div>
        </template>

        <template #cell(uuid)="data">
          <b-button type="button" variant="primary" class="mr-2" title="View" @click="callGetUser(data.value)">
            <i class="fas fa-eye" aria-hidden="true"></i>
          </b-button>

          <b-button type="button" variant="secondary" class="mr-2" title="Edit" @click="callUpdateUser(data.value)">
            <i class="fas fa-edit" aria-hidden="true"></i>
          </b-button>

          <b-button type="button" variant="danger" class="mr-2" title="Delete" @click="callDeleteUser(data.value)">
            <i class="fas fa-eraser" aria-hidden="true"></i>
          </b-button>
        </template>
      </b-table>
    </article>
  </section>
</template>

<script>
import { mapState, mapGetters, mapActions, mapMutations } from 'vuex'
import { listUser, deleteUser } from '../../../services/api';

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
      loading: false,
      usersData: [],
      fields: ['name', 'surname', 'email', 'phone', 'uuid'],
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
    this.callListUser();
  },
  methods: {
    callGetUser(uuid) {
      this.$router.push({ name: 'UserView', params: { uuid, isEdit: false } });
    },
    callUpdateUser(uuid) {
      this.$router.push({ name: 'UserView', params: { uuid, isEdit: true } });
    },
    callDeleteUser(uuid) {
      deleteUser(uuid).then((response) => {
        this.callListUser();
      })
      .catch((err) => {
        console.log(err);
        this.notifyError(err.response.data.message);
      });
    },
    callListUser() {
      this.loading = true;
      listUser().then((response) => {
        this.usersData = response;
      })
      .catch((err) => {
        console.log(err);
        this.notifyError(err.response.data.message);
        this.$router.push({ name: 'Home' });
      })
      .finally(() => {
        this.loading = false;
      });
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
