<template>
  <div class="transmigrasi-section">
    <!-- Hero Section -->
    <section class="hero-section py-5">
      <div class="container">
        <div class="text-center mb-5">
          <h1 class="hero-title mb-3">Layanan Yang Tersedia</h1>
          <p class="hero-subtitle">Silahkan pilih layanan di bawah ini:</p>
        </div>
        
        <!-- Service Cards -->
        <div class="row justify-content-center g-4 mb-5">
          <template v-if="role === 'admin'">
            <div class="col-lg-3 col-md-4 col-sm-6">
              <router-link to="/admin/dashboard" class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow">
                <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                  <span class="service-icon" style="font-size:2.5rem;">📝</span>
                </div>
                <div class="service-title text-center">
                  <div class="title-line">Pendaftaran</div>
                  <div class="title-line">Formulir</div>
                </div>
              </router-link>
            </div>
            <div class="col-lg-3 col-md-4 col-sm-6">
              <router-link to="/admin/verifikasi" class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow">
                <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                  <span class="service-icon" style="font-size:2.5rem;">📑</span>
                </div>
                <div class="service-title text-center">
                  <div class="title-line">Verifikasi Data</div>
                  <div class="title-line">Pendaftar</div>
                </div>
              </router-link>
            </div>
            <div class="col-lg-3 col-md-4 col-sm-6">
              <router-link to="/admin/update-status" class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow">
                <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                  <span class="service-icon" style="font-size:2.5rem;">🔄</span>
                </div>
                <div class="service-title text-center">
                  <div class="title-line">Update Status</div>
                  <div class="title-line">Pengajuan</div>
                </div>
              </router-link>
            </div>
            <div class="col-lg-3 col-md-4 col-sm-6">
              <router-link to="/admin/bantuan" class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow">
                <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                  <span class="service-icon" style="font-size:2.5rem;">💬</span>
                </div>
                <div class="service-title text-center">
                  <div class="title-line">Bantuan &</div>
                  <div class="title-line">Konsultasi</div>
                </div>
              </router-link>
            </div>
          </template>
          <template v-else>
            <template v-if="!loading">
              <template v-if="user">
                <div class="col-lg-3 col-md-4 col-sm-6">
                  <div class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow" @click="goToFormulir">
                    <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                      <span class="service-icon" style="font-size:2.5rem;">📝</span>
                    </div>
                    <div class="service-title text-center">
                      <div class="title-line">Pendaftaran</div>
                      <div class="title-line">Formulir</div>
                    </div>
                  </div>
                </div>
                <div class="col-lg-3 col-md-4 col-sm-6">
                  <div class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow" @click="goToStatus">
                    <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                      <span class="service-icon" style="font-size:2.5rem;">📊</span>
                    </div>
                    <div class="service-title text-center">
                      <div class="title-line">Pantau Status</div>
                      <div class="title-line">Pengajuan</div>
                    </div>
                  </div>
                </div>
              </template>
              <!-- Bantuan & Konsultasi selalu muncul -->
              <div class="col-lg-3 col-md-4 col-sm-6">
                <div class="service-card d-flex flex-column justify-content-center align-items-center p-4 rounded shadow" @click="goToBantuan">
                  <div class="icon-wrapper d-flex justify-content-center align-items-center rounded-circle mb-3">
                    <span class="service-icon" style="font-size:2.5rem;">💬</span>
                  </div>
                  <div class="service-title text-center">
                    <div class="title-line">Bantuan &</div>
                    <div class="title-line">Konsultasi</div>
                  </div>
                </div>
              </div>
            </template>
          </template>
        </div>
      </div>
    </section>

    <!-- SPBE Info Section -->
    <section class="spbe-section py-5">
      <div class="container">
        <div class="text-center mb-4">
          <h2 class="section-title">Apa sih SPBE itu?</h2>
          <p class="section-description mx-auto">
            Layanan Transmigrasi dari Pemerintah Kota Yogyakarta kini bisa diakses dengan mudah 
            melalui Sistem Pemerintahan Berbasis Elektronik (SPBE). Cari informasi, daftar, dan pantau 
            status pengajuan langsung dari genggaman Anda. Wujudkan masa depan baru di 
            wilayah transmigrasi dengan lebih cepat dan praktis.
          </p>
        </div>
      </div>
    </section>

    <!-- Registration Process Section -->
    <section class="process-section py-5">
      <div class="container">
        <div class="row align-items-center">
          <div class="col-lg-6 mb-4 mb-lg-0">
            <h2 class="process-title mb-4">Bagaimana Cara Pendaftaran Formulir Baru?</h2>
            <div class="process-steps">
              <div class="step-item d-flex mb-3">
                <span class="step-number">1.</span>
                <span class="step-text">Baca informasi program transmigrasi.</span>
              </div>
              <div class="step-item d-flex mb-3">
                <span class="step-number">2.</span>
                <span class="step-text">Siapkan dokumen persyaratan.</span>
              </div>
              <div class="step-item d-flex mb-3">
                <span class="step-number">3.</span>
                <span class="step-text">Isi formulir pendaftaran online.</span>
              </div>
              <div class="step-item d-flex mb-3">
                <span class="step-number">4.</span>
                <span class="step-text">Tunggu verifikasi dan jadwal seleksi.</span>
              </div>
              <div class="step-item d-flex">
                <span class="step-number">5.</span>
                <span class="step-text">Ikuti pelatihan dan proses keberangkatan.</span>
              </div>
            </div>
          </div>
          <div class="col-lg-6">
            <!-- Ganti ilustrasi dengan gambar Sec-1.png -->
            <div class="process-illustration text-center">
              <img src="/Sec-1.png" alt="Ilustrasi Pendaftaran" class="img-fluid" style="max-width: 350px;" />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Benefits Section -->
    <section class="benefits-section py-5">
      <div class="container">
        <div class="row align-items-center">
          <!-- Pindahkan gambar Sec-2.png ke kiri -->
          <div class="col-lg-6 mb-4 mb-lg-0 text-center">
            <img src="/Sec-2.png" alt="Manfaat Transmigrasi" class="img-fluid" style="max-width: 350px;" />
          </div>
          <div class="col-lg-6">
            <h2 class="benefits-title mb-4">Manfaat Transmigrasi untuk Masyarakat Yogyakarta?</h2>
            <div class="benefits-list">
              <ul class="mb-0">
                <li class="benefit-item d-flex mb-3">
                  <div class="benefit-icon me-3">
                    <i class="bi bi-house-door"></i>
                  </div>
                  <span class="benefit-text">Lahan tempat tinggal dan pertanian di lokasi baru.</span>
                </li>
                <li class="benefit-item d-flex mb-3">
                  <div class="benefit-icon me-3">
                    <i class="bi bi-mortarboard"></i>
                  </div>
                  <span class="benefit-text">Pembinaan dan pelatihan keterampilan.</span>
                </li>
                <li class="benefit-item d-flex">
                  <div class="benefit-icon me-3">
                    <i class="bi bi-people"></i>
                  </div>
                  <span class="benefit-text">Dukungan sarana dan prasarana dari pemerintah.</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'LayananTransmigrasi',
  props: {
    role: {
      type: String,
      default: 'guest'
    }
  },
  data() {
    return {
      user: null,
      loading: true
    }
  },
  async created() {
    try {
      const res = await fetch('/api/user/me', { credentials: 'include' });
      if (res.ok) {
        this.user = await res.json();
      } else {
        this.user = null;
      }
    } catch {
      this.user = null;
    } finally {
      this.loading = false;
    }
  },
  methods: {
    goToFormulir() {
      this.$router.push('/RegisterForm');
    },
    goToStatus() {
      this.$router.push('/status');
    },
    goToBantuan() {
      this.$router.push('/bantuan');
    }
  }
}
</script>

<style scoped>
/* Variables */
:root {
  --primary-color: #4A90E2;
  --secondary-color: #7B68EE;
  --accent-color: #FF6B9D;
  --success-color: #50E3C2;
  --warning-color: #F5A623;
  --light-bg: #F8F9FA;
  --dark-text: #2C3E50;
  --light-text: #6C757D;
  --card-bg: #FFE6E6;
  --card-border: #C01B1B;
  --icon-bg: #fbdada;
}

/* General Styles */
.transmigrasi-section {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  line-height: 1.6;
}

/* Hero Section */
.hero-section {
  background: #f8f9fa;
  color: #333;
  min-height: 400px;
  display: flex;
  align-items: center;
}

.hero-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
  color: #333;
}

.hero-subtitle {
  font-size: 1.1rem;
  color: #666;
}

/* Service Cards - New Design */
.service-card {
  background-color: #FFE6E6;
  border: 3px solid #C01B1B;
  width: 240px;
  height: 260px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
}

.service-card:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 20px rgba(192, 27, 27, 0.2) !important;
}

.icon-wrapper {
  width: 80px;
  height: 80px;
  background-color: var(--icon-bg);
  border: 2px solid rgba(192, 27, 27, 0.1);
}

.service-icon {
  font-size: 2.5rem;
  color: #C01B1B;
}

.service-title {
  color: #C01B1B;
  font-weight: bold;
  font-size: 1.25rem;
  line-height: 1.3;
}

.title-line {
  margin-bottom: 0.25rem;
}

.title-line:last-child {
  margin-bottom: 0;
}

/* SPBE Section */
.spbe-section {
  background-color: var(--light-bg);
}

.section-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--dark-text);
  margin-bottom: 1.5rem;
}

.section-description {
  font-size: 1.1rem;
  color: var(--light-text);
  max-width: 800px;
  line-height: 1.8;
}

/* Process Section */
.process-section {
  background: white;
}

.process-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--dark-text);
}

.step-item {
  align-items: flex-start;
}

.step-number {
  background: linear-gradient(135deg, var(--primary-color), var(--secondary-color));
  color: black;
  width: 30px;
  height: 30px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  margin-right: 15px;
  flex-shrink: 0;
  font-size: 0.9rem;
}

.step-text {
  color: var(--dark-text);
  font-weight: 500;
  padding-top: 4px;
}

/* Benefits Section */
.benefits-section {
  background: var(--light-bg);
}

.benefits-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--dark-text);
}

.benefit-item {
  align-items: center;
}

.benefit-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, var(--success-color), var(--primary-color));
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  flex-shrink: 0;
}

.benefit-text {
  color: var(--dark-text);
  font-weight: 500;
}

.benefits-list ul {
  padding-left: 0;
  list-style: none;
}

.benefits-list li {
  list-style-type: none;
}

/* Responsive Design */
@media (max-width: 768px) {
  .hero-title {
    font-size: 2rem;
  }
  
  .section-title,
  .process-title,
  .benefits-title {
    font-size: 1.5rem;
  }
  
  .service-card {
    width: 200px;
    height: 220px;
  }
  
  .icon-wrapper {
    width: 60px;
    height: 60px;
  }
  
  .service-icon {
    font-size: 2rem;
  }
  
  .service-title {
    font-size: 1.1rem;
  }
}

@media (max-width: 576px) {
  .hero-section {
    min-height: 300px;
  }
  
  .hero-title {
    font-size: 1.75rem;
  }
  
  .section-title,
  .process-title,
  .benefits-title {
    font-size: 1.25rem;
  }
  
  .service-card {
    width: 180px;
    height: 200px;
    margin-bottom: 1rem;
  }
  
  .icon-wrapper {
    width: 50px;
    height: 50px;
  }
  
  .service-icon {
    font-size: 1.5rem;
  }
  
  .service-title {
    font-size: 1rem;
  }
}
</style>