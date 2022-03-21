<template>
  <b-navbar toggleable="lg" type="light">
    <b-container>
      <b-navbar-brand v-bind:to="{name: 'Home'}" class="home-button">
        <span style="font-weight: bold;">
          Book
        </span>
        <span style="font-weight: lighter;">Recommendations</span>
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav right>
        <b-navbar-nav>
          <b-nav-item v-bind:to="{name: 'Home'}" class="mr-4">Home</b-nav-item>

          <b-nav-item-dropdown text="Users">
            <b-dropdown-item v-bind:to="{name: 'UserList'}">List Users</b-dropdown-item>
            <b-dropdown-item v-bind:to="{name: 'UserCreate'}">Create User</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown text="Books">
            <b-dropdown-item v-bind:to="{name: 'BookList'}">List Books</b-dropdown-item>
            <b-dropdown-item v-bind:to="{name: 'BookCreate'}">Create Book</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown text="Recommendations">
            <b-dropdown-item v-bind:to="{name: 'RecommendationList'}">List Recommendations</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto" v-if="user">
          <b-nav-item v-bind:to="{ name: 'UserView', params: { uuid: user.uuid, isEdit: false } }">
            <b>{{ user.name | capitalize }}</b>
          </b-nav-item>

          <b-nav-item @click="logout">Logout</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-container>
  </b-navbar>
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
    ...mapGetters(['user'])
  },
  watch: {
  },
  beforeMount() {
  },
  mounted() {
  },
  methods: {
    logout() {
      this.$store.dispatch('setToken', null);
      this.$store.dispatch('setUser', null);

      this.$router.push({ name: 'SignIn' });
    },
  },
  beforeDestroy() {
  },
  destroyed() {
  },
}
</script>

<style scoped>
.navbar {
  background-color: #efefef;
  /* box-shadow: 0px 0px 8px 2px rgba(0, 0, 0, 0.1) !important; */
}

.home-button-icon {
  transition: ease-in .2s;
}

.home-button:hover > span > .home-button-icon {
  color: #f22;
}
</style>