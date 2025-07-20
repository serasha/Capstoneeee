<template>
  <div class="dashboard-container">
    <!-- User Profile Header -->
    <div class="user-profile-header">
      <div class="container-fluid">
        <div class="row justify-content-end">
          <div class="col-auto">
            <div class="user-profile-card">
              <div class="user-info">
                <div class="user-details">
                  <h5 class="user-name">{{ userData.name }}</h5>
                  <p class="user-id">{{ userData.id }}</p>
                </div>
                <div class="user-avatar">
                  <img 
                    v-if="userData.avatar"
                    :src="userData.avatar" 
                    :alt="userData.name"
                    class="avatar-img"
                    @error="showDefaultAvatar = true"
                    v-show="!showDefaultAvatar"
                  >
                  <span v-if="showDefaultAvatar || !userData.avatar" class="avatar-img avatar-default">
                    <!-- SVG icon profile -->
                    <svg width="60" height="60" viewBox="0 0 60 60" fill="none">
                      <circle cx="30" cy="30" r="30" fill="#e9ecef"/>
                      <circle cx="30" cy="24" r="12" fill="#adb5bd"/>
                      <ellipse cx="30" cy="44" rx="16" ry="10" fill="#adb5bd"/>
                    </svg>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Dashboard Content -->
    <div class="dashboard-content">
      <div class="container-fluid">
        <!-- Welcome Section -->
        <div class="welcome-section">
          <div class="row justify-content-center">
            <div class="col-lg-8 col-md-10">
              <div class="welcome-card">
                <h1 class="welcome-title">
                  Selamat Datang di Sistem Pemerintahan<br>
                  Berbasis Elektronik (SPBE)
                </h1>
                
                <div class="services-section">
                  <h3 class="services-title">Layanan Yang Tersedia</h3>
                  <p class="services-subtitle">Silahkan pilih layanan di bawah ini:</p>
                  
                  <!-- Service Cards -->
                  <div class="row g-4 justify-content-center">
                    <div class="col-lg-4 col-md-6" v-for="service in services" :key="service.id">
                      <div 
                        class="service-card" 
                        @click="navigateToService(service.route)"
                        :class="{ 'service-card-hover': true }"
                      >
                        <div class="service-icon">
                          <i :class="service.icon"></i>
                        </div>
                        <div class="service-content">
                          <h5 class="service-title">{{ service.title }}</h5>
                          <p class="service-description">{{ service.description }}</p>
                        </div>
                        <div class="service-arrow">
                          <i class="fas fa-chevron-right"></i>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SPBEDashboard',
  data() {
    return {
      userData: {
        name: 'Hendra Hermawan',
        id: '190000290',
        avatar: '' // Kosongkan agar default icon muncul
      },
      showDefaultAvatar: false,
      services: [
        {
          id: 1,
          title: 'Verifikasi Data Pendaftar',
          description: 'Verifikasi dan validasi data pendaftar sistem',
          icon: 'fas fa-check-circle',
          route: '/verifikasi',
          color: '#28a745'
        },
        {
          id: 2,
          title: 'Update Status Pengajuan',
          description: 'Update status pengajuan dan permohonan',
          icon: 'fas fa-sync-alt',
          route: '/update-status',
          color: '#17a2b8'
        },
        {
          id: 3,
          title: 'Log Aktivitas',
          description: 'Pantau dan kelola log aktivitas sistem',
          icon: 'fas fa-clipboard-list',
          route: '/log-aktivitas',
          color: '#ffc107'
        }
      ]
    }
  },
  methods: {
    navigateToService(route) {
      this.$router.push(route);
    },
    handleImageError(event) {
      this.showDefaultAvatar = true;
    },
    handleLogout() {
      localStorage.removeItem('user');
      this.$router.push('/login');
    }
  },
  mounted() {
    // You can fetch user data here
    this.loadUserData();
  },
  methods: {
    ...this.methods,
    async loadUserData() {
      try {
        // Simulate API call to get user data
        // const response = await this.$api.getUserProfile();
        // this.userData = response.data;
        // Kosongkan avatar jika tidak ada, biar default icon muncul
        if (!this.userData.avatar) {
          this.showDefaultAvatar = true;
        }
        console.log('User data loaded');
      } catch (error) {
        console.error('Error loading user data:', error);
      }
    }
  }
}
</script>

<style scoped>
.dashboard-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  padding: 0;
}

.user-profile-header {
  background: white;
  border-bottom: 1px solid #dee2e6;
  padding: 1.5rem 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.user-profile-card {
  background: white;
  border-radius: 15px;
  padding: 1rem 1.5rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border: 1px solid #e9ecef;
  transition: all 0.3s ease;
}

.user-profile-card:hover {
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.15);
  transform: translateY(-2px);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-details {
  text-align: right;
}

.user-name {
  font-size: 1.1rem;
  font-weight: 600;
  color: #343a40;
  margin: 0;
  line-height: 1.2;
}

.user-id {
  font-size: 0.9rem;
  color: #6c757d;
  margin: 0;
  font-weight: 500;
}

.user-avatar {
  position: relative;
  width: 60px;
  height: 60px;
}

.avatar-img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #dc3545;
  box-shadow: 0 2px 8px rgba(220, 53, 69, 0.2);
  background: #e9ecef;
  display: inline-block;
  vertical-align: middle;
}

.avatar-default {
  display: flex;
  align-items: center;
  justify-content: center;
  background: #e9ecef;
  border: 3px solid #dc3545;
}

.dashboard-content {
  padding: 3rem 0;
}

.welcome-section {
  padding: 2rem 0;
}

.welcome-card {
  background: white;
  border-radius: 20px;
  padding: 3rem 2rem;
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.1);
  text-align: center;
  border: 1px solid #e9ecef;
}

.welcome-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #343a40;
  margin-bottom: 3rem;
  line-height: 1.3;
}

.services-section {
  margin-top: 2rem;
}

.services-title {
  font-size: 1.8rem;
  font-weight: 600;
  color: #495057;
  margin-bottom: 1rem;
}

.services-subtitle {
  font-size: 1.1rem;
  color: #6c757d;
  margin-bottom: 2.5rem;
}

.service-card {
  background: white;
  border: 2px solid #dc3545;
  border-radius: 15px;
  padding: 2rem 1.5rem;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  min-height: 200px;
}

.service-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(220, 53, 69, 0.05) 0%, rgba(220, 53, 69, 0.1) 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.service-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(220, 53, 69, 0.2);
  border-color: #a71e2a;
}

.service-card:hover::before {
  opacity: 1;
}

.service-icon {
  font-size: 3rem;
  color: #dc3545;
  margin-bottom: 1.5rem;
  position: relative;
  z-index: 2;
}

.service-content {
  position: relative;
  z-index: 2;
}

.service-title {
  font-size: 1.2rem;
  font-weight: 600;
  color: #343a40;
  margin-bottom: 0.75rem;
  line-height: 1.3;
}

.service-description {
  font-size: 0.95rem;
  color: #6c757d;
  margin-bottom: 0;
  line-height: 1.4;
}

.service-arrow {
  position: absolute;
  top: 1rem;
  right: 1rem;
  color: #dc3545;
  font-size: 1rem;
  opacity: 0;
  transition: all 0.3s ease;
  z-index: 2;
}

.service-card:hover .service-arrow {
  opacity: 1;
  transform: translateX(5px);
}

/* Responsive Design */
@media (max-width: 1199.98px) {
  .welcome-title {
    font-size: 2.2rem;
  }
  
  .service-card {
    min-height: 180px;
  }
}

@media (max-width: 991.98px) {
  .welcome-title {
    font-size: 2rem;
  }
  
  .welcome-card {
    padding: 2.5rem 1.5rem;
  }
  
  .services-title {
    font-size: 1.6rem;
  }
}

@media (max-width: 767.98px) {
  .dashboard-content {
    padding: 2rem 0;
  }
  
  .welcome-title {
    font-size: 1.8rem;
    margin-bottom: 2rem;
  }
  
  .welcome-card {
    padding: 2rem 1rem;
  }
  
  .user-profile-header {
    padding: 1rem 0;
  }
  
  .user-profile-card {
    padding: 0.75rem 1rem;
  }
  
  .avatar-img {
    width: 50px;
    height: 50px;
  }
  
  .user-name {
    font-size: 1rem;
  }
  
  .user-id {
    font-size: 0.8rem;
  }
  
  .service-card {
    min-height: 160px;
    padding: 1.5rem 1rem;
  }
  
  .service-icon {
    font-size: 2.5rem;
    margin-bottom: 1rem;
  }
  
  .service-title {
    font-size: 1.1rem;
  }
  
  .service-description {
    font-size: 0.9rem;
  }
}

@media (max-width: 575.98px) {
  .welcome-title {
    font-size: 1.6rem;
  }
  
  .services-title {
    font-size: 1.4rem;
  }
  
  .services-subtitle {
    font-size: 1rem;
  }
}

/* Animation */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.welcome-card {
  animation: fadeInUp 0.6s ease-out;
}

.service-card {
  animation: fadeInUp 0.6s ease-out;
}

.service-card:nth-child(1) { animation-delay: 0.1s; }
.service-card:nth-child(2) { animation-delay: 0.2s; }
.service-card:nth-child(3) { animation-delay: 0.3s; }

/* Focus states for accessibility */
.service-card:focus {
  outline: 2px solid #0d6efd;
  outline-offset: 2px;
}

/* Print styles */
@media print {
  .user-profile-header {
    display: none;
  }
  
  .service-card {
    border: 1px solid #000;
    box-shadow: none;
  }
  
  .service-card:hover {
    transform: none;
  }
}
</style>