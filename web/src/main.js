import Vue from 'vue';

import App from './App.vue';
import Argon from './plugins/argon-kit';
import store from './plugins/store';
import router from './router';

Vue.config.productionTip = false;
Vue.use(Argon);

new Vue({store, router, render: h => h(App)}).$mount('#app');
