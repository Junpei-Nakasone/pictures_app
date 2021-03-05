import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
import LoginPage from '@/views/LoginPage.vue'
import LatestPosts from '@/components/LatestPosts.vue'
import ViewCategory from '@/components/ViewCategory.vue'
import PrefectureCategory from '@/components/PrefectureCategory.vue'
import NewPost from '@/components/NewPost.vue'
import SignUp from '@/views/SignUp.vue'
import Detail from '@/components/detail.vue'
import UserPage from '@/views/UserPage.vue'

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
      }
    ]
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
  },
  {
    path: '/userpage/:user_id',
    name: 'userpage',
    component: UserPage
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
