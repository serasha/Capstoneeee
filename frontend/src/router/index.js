import { createRouter, createWebHistory } from 'vue-router'
import Login from '../pages/Login.vue'
import Daftar from '../pages/Daftar.vue'
import Home from '../pages/Home.vue'
import Userlayout from '../Layouts/UserLayout.vue'
import RegisterForm from '../components/RegisterForm.vue'
import SuccessModal from '../components/SuccessModal.vue'
import History from '../components/History.vue' 
import AdminLayout from '../Layouts/AdminLayout.vue'
import AdminDashboard from '../pages/admin/AdminDashboard.vue'
import AdminHakAkses from '../pages/admin/AdminHakAkses.vue'
import AdminLogAktivitas from '../pages/admin/AdminLogAktivitas.vue'
import AdminUpdateStatus from '../pages/admin/AdminUpdateStatus.vue'
import AdminVerifikasi from '../pages/admin/AdminVerifikasi.vue'
import { useUserStore } from '@/store/userStore'
import { storeToRefs } from 'pinia'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/daftar',
    name: 'Daftar',
    component: Daftar
  },
  {
    path: '/RegisterForm',
    name: 'RegisterForm',
    component: RegisterForm
  },
  {
    path: '/success',
    name: 'SuccessModal',
    component: SuccessModal
  },
  {
    path: '/status',
    name: 'StatusPengajuan',
    component: History
  },
  {
    path: '/history',
    name: 'History',
    component: History
  },
  {
    path: '/admin',
    component: AdminLayout,
    children: [
      {
        path: 'dashboard',
        name: 'AdminDashboard',
        component: AdminDashboard
      },
      {
        path: 'hak-akses',
        name: 'AdminHakAkses',
        component: AdminHakAkses
      },
      {
        path: 'log-aktivitas',
        name: 'AdminLogAktivitas',
        component: AdminLogAktivitas
      },
      {
        path: 'update-status',
        name: 'AdminUpdateStatus',
        component: AdminUpdateStatus
      },
      {
        path: 'verifikasi',
        name: 'AdminVerifikasi',
        component: AdminVerifikasi
      }
    ]
  },
  { 
    path: '/',
    component: Userlayout,
    children: [
      {
        path: '',
        name: 'Home',
        component: Home
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Global navigation guard untuk otorisasi dan redirect sesuai role
router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()
  const { user, isAuthenticated, guest, role } = storeToRefs(userStore)
  // Pastikan state sudah di-fetch
  if (userStore.user === null && !userStore.loading) {
    await userStore.fetchUser()
  }

  // Jika guest (belum login), hanya boleh akses /login, /daftar
  if (guest.value) {
    if (to.path !== '/login' && to.path !== '/daftar') {
      return next('/login')
    }
    return next()
  }

  // Jika route /admin, hanya admin yang boleh akses
  if (to.path.startsWith('/admin')) {
    if (role.value !== 'admin') {
      return next('/')
    }
    return next()
  }

  // Jika user login dan role admin, dan akses route user (bukan /admin), redirect ke /admin/dashboard
  if (role.value === 'admin' && !to.path.startsWith('/admin')) {
    return next('/admin/dashboard')
  }

  // Jika user login dan role user, dan akses route /admin, redirect ke /
  if (role.value === 'user' && to.path.startsWith('/admin')) {
    return next('/')
  }

  // Jika user login dan role user, dan akses /login atau /daftar, redirect ke home
  if (role.value === 'user' && (to.path === '/login' || to.path === '/daftar')) {
    return next('/')
  }

  // Jika user login dan role admin, dan akses /login atau /daftar, redirect ke /admin/dashboard
  if (role.value === 'admin' && (to.path === '/login' || to.path === '/daftar')) {
    return next('/admin/dashboard')
  }

  next()
})

export default router
