import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
import './assets/css/golabl.css'
import axios from 'axios'

Vue.config.productionTip = false

Vue.prototype.$http = axios
 axios.defaults.baseURL = 'http://127.0.0.1:8080/api/'
//axios.defaults.baseURL = 'http://124.70.199.139:8080/api/'
axios.interceptors.request.use(config=>{  //请求拦截器（在请求之前进行一些配置）
  config.headers.Authorization=window.sessionStorage.getItem('token')
  return config
})
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
