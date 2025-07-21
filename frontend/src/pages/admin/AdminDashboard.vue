<template>
  <div class="admin-dashboard">
    <!-- Admin Landing Section -->
    <section class="admin-landing-section">
      <div class="admin-landing-bg">
        <h1 class="admin-landing-title">Modul Layanan Pendaftaran<br />dan Pendataan Transmigrasi</h1>
        <p class="admin-landing-desc">
          Layanan Transmigrasi dari Pemerintah Kota Yogyakarta kini hadir di genggaman Anda melalui Sistem Pemerintahan Berbasis Elektronik (SPBE). Proses pendaftaran, verifikasi data, update status, hingga bantuan keberangkatan bisa diakses dengan mudah, cepat, dan transparan.
        </p>
        <div class="admin-landing-services">
          <router-link to="/admin/dashboard" class="admin-service-card">
            <span class="admin-service-icon">📝</span>
            <span class="admin-service-label">Pendaftaran Formulir</span>
          </router-link>
          <router-link to="/admin/verifikasi" class="admin-service-card">
            <span class="admin-service-icon">📑</span>
            <span class="admin-service-label">Verifikasi Data Pendaftar</span>
          </router-link>
          <router-link to="/admin/bantuan" class="admin-service-card">
            <span class="admin-service-icon">💬</span>
            <span class="admin-service-label">Bantuan & Konsultasi</span>
          </router-link>
        </div>
      </div>
    </section>

    <!-- Stats Cards -->
    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon">
          <i class="fas fa-users"></i>
        </div>
        <div class="stat-content">
          <h3 class="stat-number">{{ stats.totalUsers }}</h3>
          <p class="stat-label">Total Pengguna</p>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">
          <i class="fas fa-file-alt"></i>
        </div>
        <div class="stat-content">
          <h3 class="stat-number">{{ stats.totalApplications }}</h3>
          <p class="stat-label">Total Pendaftaran</p>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">
          <i class="fas fa-check-circle"></i>
        </div>
        <div class="stat-content">
          <h3 class="stat-number">{{ stats.verifiedApplications }}</h3>
          <p class="stat-label">Terverifikasi</p>
        </div>
      </div>

      <div class="stat-card">
        <div class="stat-icon">
          <i class="fas fa-clock"></i>
        </div>
        <div class="stat-content">
          <h3 class="stat-number">{{ stats.pendingApplications }}</h3>
          <p class="stat-label">Menunggu</p>
        </div>
      </div>
    </div>

    <!-- Quick Actions -->
    <div class="quick-actions">
      <h2 class="section-title">Aksi Cepat</h2>
      <div class="actions-grid">
        <router-link to="/admin/verifikasi" class="action-card">
          <i class="fas fa-check-square"></i>
          <span>Verifikasi Data</span>
        </router-link>
        
        <router-link to="/admin/log-aktivitas" class="action-card">
          <i class="fas fa-list-alt"></i>
          <span>Log Aktivitas</span>
        </router-link>
        
        <router-link to="/admin/hak-akses" class="action-card">
          <i class="fas fa-users"></i>
          <span>Hak Akses</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminDashboard',
  data() {
    return {
      stats: {
        totalUsers: 0,
        totalApplications: 0,
        verifiedApplications: 0,
        pendingApplications: 0
      }
    }
  },
  async created() {
    try {
      const res = await fetch('/api/admin/dashboard-stat', { credentials: 'include' })
      if (res.ok) {
        const data = await res.json()
        this.stats.totalUsers = data.total_users
        this.stats.totalApplications = data.total_pendaftaran
        this.stats.verifiedApplications = data.terverifikasi
        this.stats.pendingApplications = data.menunggu
      }
    } catch (e) {
      // Optional: tampilkan error
    }
  }
}
</script>

<style scoped>
.admin-dashboard {
  max-width: 1200px;
  margin: 0 auto;
}

.welcome-section {
  margin-bottom: 2rem;
}

.page-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #333;
  margin-bottom: 0.5rem;
}

.page-subtitle {
  color: #6c757d;
  font-size: 1.1rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 3rem;
}

.stat-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  gap: 1.5rem;
  transition: transform 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-2px);
}

.stat-icon {
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, #dc3545, #c82333);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 1.5rem;
}

.stat-number {
  font-size: 2rem;
  font-weight: 700;
  color: #333;
  margin: 0;
}

.stat-label {
  color: #6c757d;
  margin: 0;
  font-size: 0.9rem;
}

.quick-actions {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #333;
  margin-bottom: 1.5rem;
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.action-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 12px;
  text-decoration: none;
  color: #333;
  transition: all 0.3s ease;
  border: 2px solid transparent;
}

.action-card:hover {
  background: #dc3545;
  color: white;
  transform: translateY(-2px);
  text-decoration: none;
}

.action-card i {
  font-size: 2rem;
}

.action-card span {
  font-weight: 600;
}

/* Responsive */
@media (max-width: 768px) {
  .page-title {
    font-size: 2rem;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
  }
  
  .stat-card {
    padding: 1.5rem;
  }
  
  .actions-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .action-card {
    padding: 1.5rem;
  }
}

@media (max-width: 480px) {
  .actions-grid {
    grid-template-columns: 1fr;
  }
}

/* Admin Landing Section Styles */
.admin-landing-section {
  background: linear-gradient(rgba(220,53,69,0.85), rgba(220,53,69,0.85)), url('/hero-ilustrasi.png');
  background-size: cover;
  background-position: center;
  padding: 3rem 0 2rem 0;
  color: #fff;
  border-radius: 16px;
  margin-bottom: 2.5rem;
}
.admin-landing-bg {
  max-width: 1100px;
  margin: 0 auto;
  text-align: left;
  padding: 2rem 2rem 2rem 2rem;
}
.admin-landing-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 1.2rem;
  line-height: 1.2;
}
.admin-landing-desc {
  font-size: 1.1rem;
  margin-bottom: 2.2rem;
  max-width: 700px;
}
.admin-landing-services {
  display: flex;
  flex-wrap: wrap;
  gap: 1.5rem;
}
.admin-service-card {
  background: #fff;
  color: #dc3545;
  border: 2px solid #dc3545;
  border-radius: 14px;
  padding: 1.5rem 1.2rem;
  min-width: 180px;
  min-height: 120px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-decoration: none;
  font-weight: 600;
  font-size: 1.1rem;
  box-shadow: 0 2px 10px rgba(220,53,69,0.08);
  transition: all 0.2s;
}
.admin-service-card:hover {
  background: #dc3545;
  color: #fff;
  border-color: #fff;
  transform: translateY(-2px) scale(1.04);
}
.admin-service-icon {
  font-size: 2.2rem;
  margin-bottom: 0.5rem;
}
.admin-service-label {
  text-align: center;
  font-size: 1rem;
}
@media (max-width: 768px) {
  .admin-landing-title {
    font-size: 1.5rem;
  }
  .admin-landing-bg {
    padding: 1rem;
  }
  .admin-landing-services {
    flex-direction: column;
    gap: 1rem;
  }
  .admin-service-card {
    min-width: 100%;
    min-height: 90px;
    font-size: 1rem;
    padding: 1rem 0.5rem;
  }
}
</style>