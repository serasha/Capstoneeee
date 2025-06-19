import { createRouter, createWebHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import Daftar from '../pages/Daftar.vue'
import LandingPage from '../pages/LandingPage.vue'

const routes = [
  { path: '/', component: LandingPage },
  { path: '/login', component: Login },
  { path: '/daftar', component: Daftar }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
