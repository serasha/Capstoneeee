<template>
  <div class="admin-sidebar">
    <!-- Logo Section -->
    <div class="sidebar-header">
      <h1 class="logo">SPBE</h1>
    </div>

    <!-- Navigation Menu -->
    <nav class="sidebar-nav">
      <ul class="nav-list">
        <li class="nav-item">
          <router-link 
            to="/admin/dashboard" 
            class="nav-link"
            :class="{ active: isActiveRoute('/admin/dashboard') }"
          >
            <i class="fas fa-home nav-icon"></i>
            <span class="nav-text">Beranda</span>
          </router-link>
        </li>

        <li class="nav-item">
          <router-link 
            to="/admin/verifikasi" 
            class="nav-link"
            :class="{ active: isActiveRoute('/admin/verifikasi') }"
          >
            <i class="fas fa-check-square nav-icon"></i>
            <span class="nav-text">Verifikasi</span>
          </router-link>
        </li>

        <li class="nav-item">
          <router-link 
            to="/admin/log-aktivitas" 
            class="nav-link"
            :class="{ active: isActiveRoute('/admin/log-aktivitas') }"
          >
            <i class="fas fa-list-alt nav-icon"></i>
            <span class="nav-text">Log Aktivitas</span>
          </router-link>
        </li>

        <li class="nav-item">
          <router-link 
            to="/admin/hak-akses" 
            class="nav-link"
            :class="{ active: isActiveRoute('/admin/hak-akses') }"
          >
            <i class="fas fa-users nav-icon"></i>
            <span class="nav-text">Hak Akses</span>
          </router-link>
        </li>

        <li class="nav-item">
          <router-link 
            to="/admin/master-kota" 
            class="nav-link"
            :class="{ active: isActiveRoute('/admin/master-kota') }"
          >
            <i class="fas fa-map-marker-alt nav-icon"></i>
            <span class="nav-text">Master Kota</span>
          </router-link>
        </li>
      </ul>
    </nav>

    <!-- Logout Section -->
    <div class="sidebar-footer">
      <button class="logout-btn" @click="handleLogout">
        <i class="fas fa-sign-out-alt nav-icon"></i>
        <span class="nav-text">Logout</span>
      </button>
    </div>
  </div>
</template>

<script>
import { useUserStore } from '@/store/userStore'

export default {
  name: 'AdminSidebar',
  methods: {
    isActiveRoute(path) {
      return this.$route.path === path;
    },
    async handleLogout() {
      const userStore = useUserStore()
      try {
        await fetch('/api/user/logout', { method: 'POST', credentials: 'include' })
      } catch {}
      userStore.logout()
      window.location.href = '/login'
    }
  }
}
</script>

<style scoped>
.admin-sidebar {
  width: 280px;
  height: 100vh;
  background: #f8f9fa;
  border-right: 1px solid #e9ecef;
  display: flex;
  flex-direction: column;
  position: fixed;
  left: 0;
  top: 0;
  z-index: 1000;
}

.sidebar-header {
  padding: 2rem 1.5rem;
  border-bottom: 1px solid #e9ecef;
}

.logo {
  font-size: 2.5rem;
  font-weight: bold;
  color: #dc3545;
  margin: 0;
  font-style: italic;
  text-align: left;
}

.sidebar-nav {
  flex: 1;
  padding: 1rem 0;
}

.nav-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.nav-item {
  margin-bottom: 0.5rem;
}

.nav-link {
  display: flex;
  align-items: center;
  padding: 1rem 1.5rem;
  color: #6c757d;
  text-decoration: none;
  transition: all 0.3s ease;
  border-left: 3px solid transparent;
}

.nav-link:hover {
  background-color: #e9ecef;
  color: #dc3545;
  border-left-color: #dc3545;
}

.nav-link.active {
  background-color: rgba(220, 53, 69, 0.1);
  color: #dc3545;
  border-left-color: #dc3545;
  font-weight: 600;
}

.nav-icon {
  width: 20px;
  margin-right: 0.75rem;
  font-size: 1.1rem;
}

.nav-text {
  font-size: 1rem;
}

.sidebar-footer {
  padding: 1.5rem;
  border-top: 1px solid #e9ecef;
}

.logout-btn {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 1rem;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.logout-btn:hover {
  background: #c82333;
  transform: translateY(-1px);
}

/* Responsive */
@media (max-width: 768px) {
  .admin-sidebar {
    width: 100%;
    height: auto;
    position: relative;
  }
  
  .nav-link {
    padding: 0.75rem 1rem;
  }
  
  .logo {
    font-size: 2rem;
  }
}
</style>