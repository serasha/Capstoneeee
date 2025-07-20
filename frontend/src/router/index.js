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
  // Cek role user dari backend
  let user = null
  try {
    const res = await fetch('/api/user/me', { credentials: 'include' })
    if (res.ok) {
      user = await res.json()
    }
  } catch {}

  // Jika route /admin, hanya admin yang boleh akses
  if (to.path.startsWith('/admin')) {
    if (!user || user.role !== 'admin') {
      return next('/login')
    }
    return next()
  }

  // Jika user login dan role admin, redirect ke /admin/dashboard
  if (user && user.role === 'admin' && (to.path === '/' || to.path === '/login' || to.path === '/daftar')) {
    return next('/admin/dashboard')
  }

  // Jika user login dan role user, dan akses /login atau /daftar, redirect ke home
  if (user && user.role === 'user' && (to.path === '/login' || to.path === '/daftar')) {
    return next('/')
  }

  next()
})

export default router
