<template>
  <nav class="navbar navbar-expand-lg navbar-dark w-100 py-2 main-navbar">
    <div class="container-fluid px-4 d-flex justify-content-between align-items-center w-100">
      <!-- Brand/Logo -->
      <a class="navbar-brand fw-bold ms-3" href="#">
        SPBE
      </a>

      <!-- Navbar items (centered) -->
      <div class="navbar-center d-flex flex-row justify-content-center mx-auto">
        <ul class="navbar-nav flex-row mb-0">
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/beranda">BERANDA</a>
          </li>
          <li class="nav-item">
            <a 
              class="nav-link text-white px-3" 
              href="https://jss.jogjakota.go.id/" 
              target="_blank" 
              rel="noopener"
            >YOGYAKARTA</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/bantuan">BANTUAN</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/panduan">PANDUAN</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/faq">FAQ</a>
          </li>
        </ul>
      </div>

      <!-- Right side buttons -->
      <div class="d-flex gap-2 mb-0 me-4 daftar-masuk-group">
        <template v-if="!loading">
          <template v-if="isAuthenticated && user">
            <span class="text-white me-2">👤 {{ user.username }} <span v-if="user.role">({{ user.role }})</span></span>
            <button class="btn btn-outline-light" @click="logout">Logout</button>
          </template>
          <template v-else>
            <router-link to="/daftar" class="btn btn-outline-warning daftar-btn">
              Daftar
            </router-link>
            <router-link to="/login" class="btn btn-warning masuk-btn">
              Masuk
            </router-link>
          </template>
        </template>

      <div v-else class="d-flex align-items-center gap-3 me-4">
        <div class="d-flex align-items-center position-relative">
          <div class="profile-dropdown" @click="toggleDropdown" @blur="closeDropdown" tabindex="0">
            <img 
              v-if="user.avatar"
              :src="user.avatar" 
              alt="avatar" 
              class="rounded-circle avatar-img-navbar"
              style="width:36px;height:36px;object-fit:cover;border:2px solid #ffca28;"
              @error="showDefaultAvatar = true"
              v-show="!showDefaultAvatar"
            >
            <span v-if="showDefaultAvatar || !user.avatar" class="rounded-circle avatar-img-navbar avatar-default-navbar" style="width:36px;height:36px;display:inline-flex;align-items:center;justify-content:center;background:#ffe082;border:2px solid #ffca28;">
              <!-- SVG icon profile -->
              <svg width="32" height="32" viewBox="0 0 60 60" fill="none">
                <circle cx="30" cy="30" r="30" fill="#ffe082"/>
                <circle cx="30" cy="24" r="12" fill="#ffca28"/>
                <ellipse cx="30" cy="44" rx="16" ry="10" fill="#ffca28"/>
              </svg>
            </span>
            <span class="ms-2 text-white fw-semibold">{{ user.name }}</span>
            <i class="fas fa-caret-down ms-2 text-white"></i>
            <div v-if="dropdownOpen" class="dropdown-menu show">
              <button class="dropdown-item" @mousedown.prevent="logout">
                <i class="fas fa-sign-out-alt me-2"></i> Logout
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
      </div>
  </nav>
</template>

<script>
import { useUserStore } from '@/store/userStore'
import { storeToRefs } from 'pinia'
export default {
  name: 'Navbar',
  setup() {
    const userStore = useUserStore()
    const { user, isAuthenticated, loading } = storeToRefs(userStore)
    const logout = async () => {
      await fetch('/api/user/logout', { method: 'POST', credentials: 'include' })
      userStore.logout()
      window.location.href = '/login'
    }
    userStore.fetchUser()
    return { user, isAuthenticated, loading, logout }
  }
}
</script>

<style scoped>
/* Navbar positioning */
.main-navbar {
  background-color: #222;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1030;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  width: 100vw !important;
  height: 60px !important;
  padding-top: 8px !important;
  padding-bottom: 8px !important;
  border-radius: 0;
}

/* Ensure full width */
.container-fluid {
  max-width: 100% !important;
  width: 100% !important;
  height: 100%;
  display: flex;
  align-items: center;
}

/* Body padding to compensate for fixed navbar */
body {
  padding-top: 60px !important;
}

.navbar-brand {
  font-weight: bold !important;
  border: none !important;
  color: #ff4d4d;
  font-size: 1.5rem;
  font-style: italic;
}

.nav-link {
  font-weight: 500;
  transition: none;
}

/* Lebih spesifik agar tidak tertimpa style lain */
.navbar-nav .nav-link:hover,
.navbar-nav .nav-link:focus {
  color: #ffca28 !important;
  background: transparent !important; /* pastikan tidak ada background saat hover */
}

.btn {
  font-weight: 500;
  transition: all 0.3s ease;
}

.btn-outline-warning:hover {
  background-color: #ffca28;
  color: #000;
}

.btn-warning:hover {
  background-color: #ffb300;
  border-color: #ffb300;
}

.navbar-center {
  flex: 1 1 auto;
  justify-content: center !important;
  display: flex !important;
}

.navbar-nav.flex-row {
  flex-direction: row !important;
  display: flex !important;
  gap: 1rem;
  justify-content: center;
  align-items: center;
}

.navbar-nav .nav-item {
  display: flex;
  align-items: center;
}

.navbar-nav .nav-link {
  padding-left: 1rem !important;
  padding-right: 1rem !important;
  color: #fff !important;
}

/* Responsive: stack menu vertically on mobile */
@media (max-width: 991px) {
  .container-fluid {
    flex-direction: column;
    padding-top: 1rem;
  }
  .navbar-nav.flex-row {
    flex-direction: column !important;
    width: 100%;
    gap: 0;
  }
  .navbar-nav .nav-item {
    justify-content: center;
  }
  .navbar-nav .nav-link {
    padding-left: 0.5rem !important;
    padding-right: 0.5rem !important;
    text-align: center;
  }
  .d-flex {
    flex-direction: column;
    width: 100%;
    margin-top: 1rem;
  }
  .d-flex .btn {
    width: 100%;
    margin-bottom: 0.5rem;
  }
}

.d-flex.gap-2.mb-0.me-4 {
  margin-right: 2rem !important;
}

.daftar-btn {
  border-radius: 4px !important;
  font-size: 16px !important;
  border: 2px solid #ffca28 !important;
  background: transparent !important;
  color: #ffca28 !important;
}

.masuk-btn {
  border-radius: 4px !important;
  font-size: 16px !important;
  background-color: #ffca28;
  border-color: #ffca28;
  color: #000;
}

.daftar-masuk-group > .daftar-btn {
  margin-right: 12px;
}

.avatar-img-navbar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ffca28;
  background: #ffe082;
  display: inline-block;
  vertical-align: middle;
}

.avatar-default-navbar {
  background: #ffe082;
  border: 2px solid #ffca28;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0;
}

.avatar-default-navbar svg {
  width: 32px;
  height: 32px;
  display: block;
}

/* Profile dropdown */
.profile-dropdown {
  display: flex;
  align-items: center;
  cursor: pointer;
  position: relative;
  outline: none;
  user-select: none;
}

.profile-dropdown:focus {
  outline: none;
}

.dropdown-menu {
  position: absolute;
  top: 110%;
  right: 0;
  min-width: 160px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 8px 24px rgba(0,0,0,0.12);
  padding: 0.5rem 0;
  z-index: 100;
  border: 1px solid #eee;
  animation: fadeInDropdown 0.2s;
}

@keyframes fadeInDropdown {
  from { opacity: 0; transform: translateY(-10px);}
  to { opacity: 1; transform: translateY(0);}
}

.dropdown-item {
  width: 100%;
  padding: 10px 18px;
  background: none;
  border: none;
  text-align: left;
  color: #333;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.2s;
  display: flex;
  align-items: center;
}

.dropdown-item:hover {
  background: #f8f9fa;
  color: #dc3545;
}

.fas.fa-caret-down {
  font-size: 1rem;
  margin-left: 4px;
}

@media (max-width: 991px) {
  .avatar-img-navbar,
  .avatar-default-navbar {
    width: 32px !important;
    height: 32px !important;
  }
  .avatar-default-navbar svg {
    width: 28px;
    height: 28px;
  }
}
</style>

