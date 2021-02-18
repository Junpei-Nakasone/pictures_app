import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
import LoginPage from '@/views/LoginPage.vue'
import LatestPosts from '@/components/LatestPosts.vue'
import ViewCategory from '@/components/ViewCategory.vue'
import PrefectureCategory from '@/components/PrefectureCategory.vue'
import NewPost from '@/components/NewPost.vue'
import Sample from '@/components/Sample.vue'
import SignUp from '@/views/SignUp.vue'
import Buefy from '@/components/buefy.vue'
import Detail from '@/components/detail.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: Home,
    children: [
      {
        path: '',
        component: LatestPosts
      },
      {
        path: 'view_category',
        component: ViewCategory
      },
      {
        path: 'prefecture_category',
        component: PrefectureCategory
      },
      {
        path: '/newPost',
        component: NewPost
      },
      {
        path: '/sample',
        component: Sample
      },
      {
        path: '/buefy',
        component: Buefy
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/login',
    component: LoginPage
  },
  {
    path: '/signup',
    component: SignUp
  },
  {
    path: '/detail/:picture_id',
    name: 'detail',
    component: Detail
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
