import Vue from 'vue';
import Router from 'vue-router';

import store from './store/store.js';

import Home from './components/pages/Home';

import SignIn from './components/pages/auth/SignIn';
import SignUp from './components/pages/auth/SignUp';

import UserCreate from './components/pages/user/Create';
import UserView from './components/pages/user/View';
import UserList from './components/pages/user/List';

import BookCreate from './components/pages/book/Create';
import BookView from './components/pages/book/View';
import BookList from './components/pages/book/List';

import RecommendationList from './components/pages/recommendation/List';

Vue.use(Router);

const router = new Router({
  mode: 'hash',
  routes: [
    // auth
    {
      path: '/signin',
      name: 'SignIn',
      component: SignIn,
    },
    {
      path: '/signun',
      name: 'SignUp',
      component: SignUp,
    },
    // home
    {
      path: '/home',
      name: 'Home',
      component: Home,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/',
      redirect: '/signin',
    },
    // users
    {
      path: '/user/create',
      name: 'UserCreate',
      component: UserCreate,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/user/:uuid/:isEdit',
      name: 'UserView',
      component: UserView,
      props: true,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/user',
      name: 'UserList',
      component: UserList,
      meta: {
        requiresAuth: true,
      },
    },
    // books
    {
      path: '/book/create',
      name: 'BookCreate',
      component: BookCreate,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/book/:uuid/:isEdit',
      name: 'BookView',
      component: BookView,
      props: true,
      meta: {
        requiresAuth: true,
      },
    },
    {
      path: '/book',
      name: 'BookList',
      component: BookList,
      meta: {
        requiresAuth: true,
      },
    },
    // recommendation
    {
      path: '/recommendation',
      name: 'RecommendationList',
      component: RecommendationList,
      meta: {
        requiresAuth: true,
      },
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.state.token) {
      next({
        name: 'SignIn',
        query: {
          redirect: to.fullPath,
        },
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
