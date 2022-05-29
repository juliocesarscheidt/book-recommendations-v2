<template>
  <b-navbar toggleable="lg" type="light">
    <b-container>
      <b-navbar-brand v-bind:to="{name: 'Home'}">
        <span style="font-weight: bold;">Book</span>
        <span style="font-weight: lighter;">Recommendations</span>
      </b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav right>
        <b-navbar-nav>
          <b-nav-item v-bind:to="{name: 'Home'}">{{ $t('header.home') }}</b-nav-item>

          <b-nav-item-dropdown v-if="user" v-bind:text="$t('header.users')">
            <b-dropdown-item v-bind:to="{name: 'UserList'}">{{ $t('header.list_users') }}</b-dropdown-item>
            <b-dropdown-item v-bind:to="{name: 'UserCreate'}">{{ $t('header.create_user') }}</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown v-if="user" v-bind:text="$t('header.books')">
            <b-dropdown-item v-bind:to="{name: 'BookList'}">{{ $t('header.list_books') }}</b-dropdown-item>
            <b-dropdown-item v-bind:to="{name: 'BookCreate'}">{{ $t('header.create_book') }}</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item-dropdown v-if="user" v-bind:text="$t('header.recommendations')">
            <b-dropdown-item v-bind:to="{name: 'RecommendationList'}">{{ $t('header.list_recommendations') }}</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-item v-if="user" v-bind:to="{ name: 'UserView', params: { uuid: user.uuid, isEdit: false } }">
            <b>{{ user.name | capitalize }}</b>
          </b-nav-item>

          <b-nav-item-dropdown :text="$t('header.language')" right>
            <b-dropdown-item @click="setLanguage('pt-br')">PR-BR</b-dropdown-item>
            <b-dropdown-item @click="setLanguage('en-us')">EN-US</b-dropdown-item>
          </b-nav-item-dropdown>

          <b-nav-item v-if="user" @click="logout">{{ $t('header.logout') }}</b-nav-item>
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
    ...mapActions(['setLanguage']),
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
</style>