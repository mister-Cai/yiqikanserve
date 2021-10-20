import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from "./store/index.js"
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
// import { Lazyload } from 'vant'

// Vue.use(Lazyload)

Vue.config.productionTip = false

// Vue.use(VueI18n)
// import i18nMessages from './i18n'
// const i18n = new VueI18n({
//   locale: 'en',
//   messages: i18nMessages
// })

export let app=new Vue({
  router,
  // i18n,
  render: h => h(App),
  store
}).$mount('#app')
