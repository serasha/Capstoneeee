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
        path: 'verifikasi',
        name: 'AdminVerifikasi',
        component: AdminVerifikasi
      },
      {
        path: 'master-kota',
        name: 'AdminMasterKota',
        component: () => import('../pages/admin/AdminMasterKota.vue')
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

// Ubah route guard agar:
// - Guest: hanya akses /, /login, /daftar
// - User: akses / (Home) dan fitur user
// - Admin: akses /admin/*

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore()
  const { user, isAuthenticated, guest, role } = storeToRefs(userStore)
  if (userStore.user === null && !userStore.loading) {
    await userStore.fetchUser()
  }

  // Guest: hanya boleh akses Home, Login, Daftar
  if (guest.value) {
    if (to.path !== '/' && to.path !== '/login' && to.path !== '/daftar') {
      return next('/')
    }
    return next()
  }

  // Admin: boleh akses semua, tapi fitur admin hanya di /admin/*
  if (role.value === 'admin') {
    // Tidak perlu redirect ke dashboard, biarkan akses /
    return next()
  }

  // User: tidak boleh akses /admin/*
  if (role.value === 'user') {
    if (to.path.startsWith('/admin')) {
      return next('/')
    }
    // Jika user login dan akses /login atau /daftar, redirect ke home
    if (to.path === '/login' || to.path === '/daftar') {
      return next('/')
    }
    return next()
  }

  next()
})

export default router
