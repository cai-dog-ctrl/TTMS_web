import Vue from 'vue'
import VueRouter from 'vue-router'
import HomePage from '../views/HomePage.vue'
import FirstPage from '../views/FirstPage.vue'
// import { ElLoadingComponent } from 'element-ui/types/loading'
import Login from '../views/Login.vue'
import SignIn from '../views/SignIn'

Vue.use(VueRouter)

const routes = [
  {
    path:"/",
    redirect:"/home",
  },
  {
    path:"/login",
    component:Login
  },
  {
    path:"/signin",
    component:SignIn
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
