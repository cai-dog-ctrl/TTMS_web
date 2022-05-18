import Vue from 'vue'
import VueRouter from 'vue-router'
import HomePage from '../views/HomePage.vue'
import FirstPage from '../views/FirstPage.vue'
Vue.use(VueRouter)

const routes = [
  {
    path:"/",
    redirect:"/home",
  },
  {
    path:"/home",
    component:HomePage,
    redirect:"/first",
    children:[
      {
        path:"/first",
        component:FirstPage
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
