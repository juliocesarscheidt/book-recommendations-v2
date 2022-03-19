import Vue from 'vue';

import App from './App.vue';
import router from './router';
import store from './store/store';

// plugins
import './plugins/bootstrap';
import './plugins/notification';

// mixins
import './mixins/notification';

// filters
import { formatPhone, capitalize, trimLetters } from './filters/filters';

// utils
import checkToken from './utils/checkToken';

Vue.config.productionTip = process.env.PRODUCTION_TIP || false;
Vue.config.devtools = process.env.DEV_TOOLS || true;

Vue.filter('formatPhone', formatPhone);
Vue.filter('capitalize', capitalize);
Vue.filter('trimLetters', trimLetters);

checkToken();

new Vue({
  router,
  store,
	render: h => h(App),
}).$mount('#app');
