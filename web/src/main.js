import Vue from 'vue';
import VueProgressBar from 'vue-progressbar';
import Notifications from 'vue-notification';

import App from './App.vue';
import Argon from './plugins/argon-kit';
import store from './plugins/store';
import router from './router';

Vue.config.productionTip = false;

Vue.use(VueProgressBar, {
  color: '#ba3fd9',
  failedColor: 'red',
  thickness: '3px',
  location: 'top',
  position: 'fixed',
  transition: {speed: '0.2s', opacity: '0.6s', termination: 300},
});

Vue.use(Notifications);
Vue.use(Argon);

new Vue({store, router, render: h => h(App)}).$mount('#app');
