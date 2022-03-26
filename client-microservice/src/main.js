import Vue from 'vue';
import VueI18n from 'vue-i18n';

import App from './App.vue';
import router from './router';
import store from './store/store';
import locales from './utils/translations/index';

// plugins
import './plugins/bootstrap';
import './plugins/notification';
import './plugins/i18n';

// mixins
import './mixins/notification';

// filters
import { formatPhone, capitalize, trimLetters, currency } from './filters/';

// directives
import { formatPhoneDirective } from './directives/';

// utils
import { getLocalStorageLanguage } from './utils/userLanguage';
import checkToken from './utils/checkToken';

Vue.filter('formatPhone', formatPhone);
Vue.filter('capitalize', capitalize);
Vue.filter('trimLetters', trimLetters);
Vue.filter('currency', currency);

Vue.directive('formatPhone', formatPhoneDirective);

const userlanguage = getLocalStorageLanguage();
const i18n = new VueI18n({ locale: userlanguage, messages: locales });

Vue.config.productionTip = process.env.PRODUCTION_TIP || false;
Vue.config.devtools = process.env.DEV_TOOLS || true;

checkToken();

new Vue({
  router,
  store,
  i18n,
	render: h => h(App),
}).$mount('#app');
