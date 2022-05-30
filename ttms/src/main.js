import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import './assets/css/golabl.css'
import axios from 'axios'

Vue.prototype.$http = axios
//axios.defaults.baseURL = 'http://127.0.0.1:8080/api/'
axios.defaults.baseURL = 'http://124.70.199.139:8080/api/'
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
