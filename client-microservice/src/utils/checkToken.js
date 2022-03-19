import store from '../store/store.js';
import { getCurrentUserInfo } from '../services/api';

const checkToken = async () => {
  const token = localStorage.getItem('token');
  store.dispatch('setToken', token);

  if (token) {
    try {
      const user = await getCurrentUserInfo();
      store.dispatch('setUser', user);

    } catch (err) {
      console.error(err);
      store.dispatch('setUser', null);
    }
  } else {
    store.dispatch('setUser', null);
  }
};

export default checkToken;
