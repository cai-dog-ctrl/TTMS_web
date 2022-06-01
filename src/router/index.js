import Vue from 'vue'
import VueRouter from 'vue-router'
import HomePage from '../views/HomePage.vue'
import FirstPage from '../views/FirstPage.vue'
// import { ElLoadingComponent } from 'element-ui/types/loading'
import Login from '../views/Login.vue'
import SignIn from '../views/SignIn.vue'
import MovieInfo from '../views/MovieInfo.vue'
import BuyTickets from '../views/BuyTickets.vue'
import MangerHome from '../views/MangerHome.vue'
import MangerUser from '../views/MangerUser.vue'
import MangerWelcome from '../views/MangerWelcome.vue'
import MangerMovie from '../views/MangerMovie.vue'
import Borad from '../views/Borad.vue'
import Movie from '../views/Movie.vue'

Vue.use(VueRouter)


const routes = [
  {
    path: "/",
    redirect: "/home",
  },

  {
    path: '/mangerhome',
    component: MangerHome,
    redirect: '/mangerwelcome',
    children: [
      {path: '/mangerwelcome',component: MangerWelcome},
      {path: '/mangeruser',component: MangerUser},
      {path:'/manger_movie',component:MangerMovie}
    ]

  },

  // {
  //   path: "/mangeruser",
  //   component: MangerUser
  // },

  // {
  //   path: "mangerwelcome",
  //   component: MangerWelcome
  // },

  {
    path: "/login",
    component: Login
  },

  {
    path: "/signin",
    component: SignIn
  },
  
  {
    path: "/home",
    component: HomePage,
    redirect: "/first",
    children: [
      {
        path: "/first",
        component: FirstPage
      },
      {
        path: "/movie/:id",
        component: MovieInfo
      },
      {
        path: "/buytickets/:id",
        component: BuyTickets
      },
      {
        path: "/borad",
        component: Borad,
      },
      {
        path: "/movie",
        component: Movie
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
