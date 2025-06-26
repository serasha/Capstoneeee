<template>
  <div class="sidebar-container">
    <!-- Sidebar -->
    <div class="sidebar" :class="{ 'sidebar-mobile-open': isMobileMenuOpen }">
      <!-- Logo Section -->
      <div class="sidebar-header">
        <h2 class="logo-text">SPBE</h2>
      </div>

      <!-- Navigation Menu -->
      <nav class="sidebar-nav">
        <ul class="nav-list">
          <!-- Beranda -->
          <li class="nav-item">
            <router-link to="/beranda" class="nav-link active" @click="closeMobileMenu">
              <i class="fas fa-home nav-icon"></i>
              <span class="nav-text">Beranda</span>
            </router-link>
          </li>

          <!-- Verifikasi -->
          <li class="nav-item">
            <router-link to="/verifikasi" class="nav-link" @click="closeMobileMenu">
              <i class="fas fa-check-circle nav-icon"></i>
              <span class="nav-text">Verifikasi</span>
            </router-link>
          </li>

          <!-- Update Status -->
          <li class="nav-item">
            <router-link to="/update-status" class="nav-link" @click="closeMobileMenu">
              <i class="fas fa-sync-alt nav-icon"></i>
              <span class="nav-text">Update Status</span>
            </router-link>
          </li>

          <!-- Log Aktivitas -->
          <li class="nav-item">
            <router-link to="/log-aktivitas" class="nav-link" @click="closeMobileMenu">
              <i class="fas fa-list-alt nav-icon"></i>
              <span class="nav-text">Log Aktivitas</span>
            </router-link>
          </li>

          <!-- Hak Akses -->
          <li class="nav-item">
            <router-link to="/hak-akses" class="nav-link" @click="closeMobileMenu">
              <i class="fas fa-user-shield nav-icon"></i>
              <span class="nav-text">Hak Akses</span>
            </router-link>
          </li>
        </ul>
      </nav>

      <!-- Logout Button -->
      <div class="sidebar-footer">
        <button class="logout-btn" @click="handleLogout">
          <i class="fas fa-sign-out-alt logout-icon"></i>
          <span class="logout-text">Logout</span>
        </button>
      </div>
    </div>

    <!-- Mobile Menu Toggle Button -->
    <button 
      class="mobile-menu-toggle d-lg-none" 
      @click="toggleMobileMenu"
      :class="{ 'menu-open': isMobileMenuOpen }"
    >
      <span class="hamburger-line"></span>
      <span class="hamburger-line"></span>
      <span class="hamburger-line"></span>
    </button>

    <!-- Mobile Overlay -->
    <div 
      class="mobile-overlay d-lg-none" 
      :class="{ 'show': isMobileMenuOpen }"
      @click="closeMobileMenu"
    ></div>
  </div>
</template>

<script>
export default {
  name: 'SidebarComponent',
  data() {
    return {
      isMobileMenuOpen: false
    }
  },
  methods: {
    toggleMobileMenu() {
      this.isMobileMenuOpen = !this.isMobileMenuOpen;
    },
    closeMobileMenu() {
      this.isMobileMenuOpen = false;
    },
    handleLogout() {
      // Handle logout logic here
      if (confirm('Apakah Anda yakin ingin keluar?')) {
        // Add your logout logic
        this.$router.push('/login');
        // or emit event to parent component
        this.$emit('logout');
      }
    }
  },
  mounted() {
    // Close mobile menu when clicking outside
    document.addEventListener('click', (e) => {
      if (!e.target.closest('.sidebar-container') && this.isMobileMenuOpen) {
        this.closeMobileMenu();
      }
    });
  }
}
</script>

<style scoped>
.sidebar-container {
  position: relative;
}

.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 280px;
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  border-right: 1px solid #dee2e6;
  z-index: 1000;
  transition: transform 0.3s ease;
  overflow-y: auto;
  box-shadow: 2px 0 10px rgba(0,0,0,0.1);
}

/* Mobile responsiveness */
@media (max-width: 991.98px) {
  .sidebar {
    transform: translateX(-100%);
  }
  
  .sidebar-mobile-open {
    transform: translateX(0);
  }
}

.sidebar-header {
  padding: 2rem 1.5rem;
  text-align: center;
  border-bottom: 1px solid #dee2e6;
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  margin-bottom: 1rem;
}

.logo-text {
  font-size: 2.5rem;
  font-weight: bold;
  color: white;
  margin: 0;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
  letter-spacing: 2px;
}

.sidebar-nav {
  flex: 1;
  padding: 0 1rem;
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
  border-radius: 10px;
  transition: all 0.3s ease;
  font-weight: 500;
  position: relative;
  overflow: hidden;
}

.nav-link:hover {
  background: linear-gradient(135deg, #e9ecef 0%, #f8f9fa 100%);
  color: #dc3545;
  text-decoration: none;
  transform: translateX(5px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

.nav-link.active {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(220, 53, 69, 0.3);
}

.nav-link.active::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  width: 4px;
  background: white;
  border-radius: 0 2px 2px 0;
}

.nav-icon {
  font-size: 1.2rem;
  margin-right: 1rem;
  width: 20px;
  text-align: center;
}

.nav-text {
  font-size: 1rem;
}

.sidebar-footer {
  padding: 1.5rem;
  border-top: 1px solid #dee2e6;
  margin-top: auto;
}

.logout-btn {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 1rem 1.5rem;
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-weight: 500;
  transition: all 0.3s ease;
  cursor: pointer;
}

.logout-btn:hover {
  background: linear-gradient(135deg, #c82333 0%, #a71e2a 100%);
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(220, 53, 69, 0.4);
}

.logout-icon {
  font-size: 1.2rem;
  margin-right: 1rem;
}

.logout-text {
  font-size: 1rem;
}

/* Mobile Menu Toggle Button */
.mobile-menu-toggle {
  position: fixed;
  top: 1rem;
  left: 1rem;
  z-index: 1001;
  background: #dc3545;
  border: none;
  border-radius: 8px;
  width: 50px;
  height: 50px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 12px rgba(220, 53, 69, 0.3);
  transition: all 0.3s ease;
}

.mobile-menu-toggle:hover {
  background: #c82333;
  transform: scale(1.05);
}

.hamburger-line {
  width: 25px;
  height: 3px;
  background: white;
  margin: 2px 0;
  transition: all 0.3s ease;
  border-radius: 2px;
}

.menu-open .hamburger-line:nth-child(1) {
  transform: rotate(45deg) translate(6px, 6px);
}

.menu-open .hamburger-line:nth-child(2) {
  opacity: 0;
}

.menu-open .hamburger-line:nth-child(3) {
  transform: rotate(-45deg) translate(6px, -6px);
}

/* Mobile Overlay */
.mobile-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 999;
  opacity: 0;
  visibility: hidden;
  transition: all 0.3s ease;
}

.mobile-overlay.show {
  opacity: 1;
  visibility: visible;
}

/* Desktop specific styles */
@media (min-width: 992px) {
  .sidebar {
    position: relative;
    transform: none;
    box-shadow: none;
    border-right: 1px solid #dee2e6;
  }
  
  .mobile-menu-toggle {
    display: none;
  }
}

/* Scrollbar styling */
.sidebar::-webkit-scrollbar {
  width: 6px;
}

.sidebar::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.sidebar::-webkit-scrollbar-thumb {
  background: #dc3545;
  border-radius: 3px;
}

.sidebar::-webkit-scrollbar-thumb:hover {
  background: #c82333;
}

/* Animation for nav items */
.nav-item {
  animation: slideInLeft 0.3s ease forwards;
  opacity: 0;
}

.nav-item:nth-child(1) { animation-delay: 0.1s; }
.nav-item:nth-child(2) { animation-delay: 0.2s; }
.nav-item:nth-child(3) { animation-delay: 0.3s; }
.nav-item:nth-child(4) { animation-delay: 0.4s; }
.nav-item:nth-child(5) { animation-delay: 0.5s; }

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}
</style>