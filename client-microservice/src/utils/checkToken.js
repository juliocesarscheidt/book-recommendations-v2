import store from '../store/store.js';
import router from '../router.js';
import { getCurrentUserInfo } from '../services/';

const setNullUserToken = () => {
  store.dispatch('setUser', null);
  store.dispatch('setToken', null);
  router.push({ name: 'SignIn' });
}

const checkToken = async () => {
  if (localStorage.getItem('token')) {
    try {
      const token = localStorage.getItem('token');
      store.dispatch('setToken', token);
      const user = await getCurrentUserInfo();
      store.dispatch('setUser', user);
    } catch (err) {
      console.error(err);
      setNullUserToken();
    }
  } else {
    setNullUserToken();
  }
};

export default checkToken;
