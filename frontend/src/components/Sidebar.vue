<template>
  <div class="sidebar-wrapper">
    <!-- Sidebar -->
    <nav class="sidebar" :class="{ 'sidebar-collapsed': isCollapsed }">
      <div class="sidebar-content">
        <!-- Logo Section -->
        <div class="sidebar-header">
          <div class="logo-section">
            <h2 class="logo-text">SPBE</h2>
          </div>
          <button 
            class="btn btn-link sidebar-toggle d-lg-none" 
            @click="toggleSidebar"
            aria-label="Toggle Sidebar"
          >
            <i class="fas fa-bars"></i>
          </button>
        </div>

        <!-- Navigation Menu -->
        <div class="sidebar-nav">
          <ul class="nav flex-column">
            <!-- Home/Beranda -->
            <li class="nav-item">
              <router-link 
                to="/dashboard" 
                class="nav-link"
                :class="{ active: isActiveRoute('/dashboard') }"
              >
                <i class="fas fa-home nav-icon"></i>
                <span class="nav-text">Beranda</span>
              </router-link>
            </li>

            <!-- Verifikasi -->
            <li class="nav-item">
              <router-link 
                to="/verifikasi" 
                class="nav-link"
                :class="{ active: isActiveRoute('/verifikasi') }"
              >
                <i class="fas fa-check-square nav-icon"></i>
                <span class="nav-text">Verifikasi</span>
              </router-link>
            </li>

            <!-- Update Status -->
            <li class="nav-item">
              <router-link 
                to="/update-status" 
                class="nav-link"
                :class="{ active: isActiveRoute('/update-status') }"
              >
                <i class="fas fa-sync-alt nav-icon"></i>
                <span class="nav-text">Update Status</span>
              </router-link>
            </li>

            <!-- Log Aktivitas -->
            <li class="nav-item">
              <router-link 
                to="/log-aktivitas" 
                class="nav-link"
                :class="{ active: isActiveRoute('/log-aktivitas') }"
              >
                <i class="fas fa-list-alt nav-icon"></i>
                <span class="nav-text">Log Aktivitas</span>
              </router-link>
            </li>

            <!-- Hak Akses -->
            <li class="nav-item">
              <router-link 
                to="/hak-akses" 
                class="nav-link"
                :class="{ active: isActiveRoute('/hak-akses') }"
              >
                <i class="fas fa-users nav-icon"></i>
                <span class="nav-text">Hak Akses</span>
              </router-link>
            </li>
          </ul>
        </div>

        <!-- Logout Section -->
        <div class="sidebar-footer">
          <button 
            class="btn btn-logout w-100" 
            @click="handleLogout"
          >
            <i class="fas fa-sign-out-alt nav-icon"></i>
            <span class="nav-text">Logout</span>
          </button>
        </div>
      </div>
    </nav>

    <!-- Overlay for mobile -->
    <div 
      class="sidebar-overlay d-lg-none" 
      :class="{ show: !isCollapsed }"
      @click="closeSidebar"
    ></div>
  </div>
</template>

<script>
export default {
  name: 'SPBESidebar',
  data() {
    return {
      isCollapsed: true // Default collapsed on mobile
    }
  },
  methods: {
    toggleSidebar() {
      this.isCollapsed = !this.isCollapsed;
    },
    closeSidebar() {
      this.isCollapsed = true;
    },
    isActiveRoute(path) {
      return this.$route.path === path;
    },
    handleLogout() {
      // Emit event to parent component
      this.$emit('logout');
      // Hapus user dari localStorage dan redirect ke login
      localStorage.removeItem('user');
      this.$router.push('/login');
    }
  },
  mounted() {
    // Auto-collapse on mobile, expand on desktop
    const handleResize = () => {
      if (window.innerWidth >= 992) {
        this.isCollapsed = false;
      } else {
        this.isCollapsed = true;
      }
    };

    handleResize();
    window.addEventListener('resize', handleResize);

    // Cleanup
    this.$once('hook:beforeDestroy', () => {
      window.removeEventListener('resize', handleResize);
    });
  }
}
</script>

<style scoped>
.sidebar-wrapper {
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
  box-shadow: 2px 0 10px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease-in-out;
  z-index: 1050;
  overflow-y: auto;
}

.sidebar-content {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 0;
}

.sidebar-header {
  padding: 1.5rem 1rem;
  border-bottom: 1px solid #dee2e6;
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  color: white;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo-section {
  flex-grow: 1;
}

.logo-text {
  font-size: 2rem;
  font-weight: bold;
  margin: 0;
  letter-spacing: 2px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.sidebar-toggle {
  color: white;
  padding: 0.25rem;
  font-size: 1.2rem;
}

.sidebar-toggle:hover {
  color: #f8f9fa;
}

.sidebar-nav {
  flex-grow: 1;
  padding: 1rem 0;
}

.nav-item {
  margin-bottom: 0.25rem;
}

.nav-link {
  display: flex;
  align-items: center;
  padding: 0.875rem 1.5rem;
  color: #495057;
  text-decoration: none;
  border-radius: 0;
  transition: all 0.2s ease-in-out;
  border-left: 3px solid transparent;
}

.nav-link:hover {
  background-color: #e9ecef;
  color: #dc3545;
  text-decoration: none;
  border-left-color: #dc3545;
}

.nav-link.active {
  background-color: #dc3545;
  color: white;
  border-left-color: #a71e2a;
}

.nav-icon {
  width: 20px;
  text-align: center;
  margin-right: 0.75rem;
  font-size: 1.1rem;
}

.nav-text {
  font-weight: 500;
  font-size: 0.95rem;
}

.sidebar-footer {
  padding: 1rem;
  border-top: 1px solid #dee2e6;
  background-color: #f8f9fa;
}

.btn-logout {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  border: none;
  color: white;
  padding: 0.75rem 1rem;
  border-radius: 0.375rem;
  font-weight: 500;
  transition: all 0.2s ease-in-out;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-logout:hover {
  background: linear-gradient(135deg, #c82333 0%, #a71e2a 100%);
  color: white;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(220, 53, 69, 0.3);
}

.sidebar-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1040;
  opacity: 0;
  visibility: hidden;
  transition: all 0.3s ease-in-out;
}

.sidebar-overlay.show {
  opacity: 1;
  visibility: visible;
}

/* Mobile Responsive */
@media (max-width: 991.98px) {
  .sidebar-collapsed {
    transform: translateX(-100%);
  }
  
  .sidebar:not(.sidebar-collapsed) {
    transform: translateX(0);
  }
}

/* Desktop Responsive */
@media (min-width: 992px) {
  .sidebar {
    position: relative;
    transform: translateX(0) !important;
  }
  
  .sidebar-overlay {
    display: none !important;
  }
  
  .sidebar-toggle {
    display: none !important;
  }
}

/* Custom Scrollbar */
.sidebar::-webkit-scrollbar {
  width: 6px;
}

.sidebar::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.sidebar::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.sidebar::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* Animation for nav items */
.nav-item {
  animation: slideInLeft 0.3s ease-out;
}

.nav-item:nth-child(1) { animation-delay: 0.1s; }
.nav-item:nth-child(2) { animation-delay: 0.15s; }
.nav-item:nth-child(3) { animation-delay: 0.2s; }
.nav-item:nth-child(4) { animation-delay: 0.25s; }
.nav-item:nth-child(5) { animation-delay: 0.3s; }

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

/* Focus states for accessibility */
.nav-link:focus,
.btn-logout:focus,
.sidebar-toggle:focus {
  outline: 2px solid #0d6efd;
  outline-offset: 2px;
}
</style>