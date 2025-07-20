<template>
  <div class="status-pengajuan-container">
    <!-- Header -->
    <div class="header-section d-flex align-items-center mb-4">
      <button class="btn btn-link p-0 me-3 back-btn" @click="goBack">
        <i class="fas fa-chevron-left"></i>
      </button>
      <h4 class="mb-0 fw-bold text-dark">Pantau Status Pengajuan</h4>
    </div>

    <!-- Status Cards -->
    <div class="status-cards">
      <div 
        v-for="(item, index) in statusData" 
        :key="index"
        class="status-card mb-4"
      >
        <div class="card border-0 shadow-sm">
          <!-- Hapus gambar history-card-image-wrapper di sini -->
          <div class="card-body p-4">
            <div class="row align-items-center">
              <!-- Icon Section -->
              <div class="col-auto">
                <div class="status-icon">
                  <div class="icon-wrapper">
                    <img src="/history.jpg" alt="History Icon" class="history-icon-img" />
                    <div class="check-mark">
                      <i class="fas fa-check text-white"></i>
                    </div>
                  </div>
                </div>
              </div>

              <!-- Content Section -->
              <div class="col">
                <div class="row">
                  <div class="col-md-8">
                    <div class="mb-2">
                      <small class="text-muted">Tanggal Pengajuan: {{ item.tanggalPengajuan }}</small>
                    </div>
                    <h5 class="fw-bold mb-3 text-dark">{{ item.nama }}</h5>
                    
                    <div class="row">
                      <div class="col-sm-6 mb-2">
                        <small class="text-danger fw-medium">Kota Asal</small>
                        <div class="fw-bold text-dark">{{ item.kotaAsal }}</div>
                      </div>
                      <div class="col-sm-6 mb-2">
                        <small class="text-danger fw-medium">Tujuan Transmigrasi</small>
                        <div class="fw-bold text-dark">{{ item.tujuanTransmigrasi }}</div>
                      </div>
                    </div>
                  </div>

                  <!-- Status Section -->
                  <div class="col-md-4 text-md-end">
                    <div class="mb-2">
                      <small class="text-muted">{{ item.nomorRegistrasi }}</small>
                    </div>
                    <div class="mb-3">
                      <small class="text-muted fw-medium">Status</small>
                    </div>
                    <span 
                      :class="getStatusClass(item.status)"
                      class="status-badge px-3 py-2 rounded-pill fw-medium"
                    >
                      {{ item.status }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Tombol Beranda -->
    <div class="text-center my-4">
      <button class="btn btn-danger btn-beranda" @click="goToHome">
        <i class="fas fa-home me-2"></i>Beranda
      </button>
    </div>

    <!-- Empty State (jika tidak ada data) -->
    <div v-if="!loading && statusData.length === 0" class="empty-state text-center py-5">
      <div class="mb-4">
        <i class="fas fa-file-alt fa-3x text-muted"></i>
      </div>
      <h5 class="text-muted">Belum ada pengajuan</h5>
      <p class="text-muted">Data pengajuan transmigrasi akan muncul di sini</p>
    </div>
    <div v-if="errorMsg" class="alert alert-danger mt-3">{{ errorMsg }}</div>
  </div>
</template>

<script>
export default {
  name: 'StatusPengajuan',
  data() {
    return {
      statusData: [],
      loading: true,
      errorMsg: ''
    }
  },
  async created() {
    try {
      const res = await fetch('/api/pendaftaran/user', { credentials: 'include' });
      if (res.ok) {
        const data = await res.json();
        this.statusData = data.map(item => ({
          tanggalPengajuan: item.created_at ? new Date(item.created_at).toLocaleDateString() : '-',
          nama: item.nama_pendaftar || '-',
          kotaAsal: item.alamat_pendaftar || '-',
          tujuanTransmigrasi: item.jenis_layanan || '-',
          nomorRegistrasi: item.id_pendaftaran || '-',
          status: item.status_pendaftar || '-'
        }));
      } else {
        this.errorMsg = 'Gagal mengambil data status pendaftaran';
      }
    } catch (e) {
      this.errorMsg = 'Gagal mengambil data status pendaftaran';
    } finally {
      this.loading = false;
      statusData: []
    }
  },
  mounted() {
    // Ambil data history dari localStorage
    const saved = localStorage.getItem('pendaftaranHistory');
    if (saved) {
      try {
        this.statusData = JSON.parse(saved);
      } catch (e) {
        this.statusData = [];
      }
    }
    // Jika tidak ada data, tampilkan dummy
    if (!this.statusData || this.statusData.length === 0) {
      this.statusData = [
        {
          tanggalPengajuan: '12/05/2025',
          nama: 'HENDRAWAN SUJATMIKO, S.T.',
          kotaAsal: 'Banyuwangi',
          tujuanTransmigrasi: 'Kota Palu',
          nomorRegistrasi: '2025/1/03/CATRANS/YK',
          status: 'Proses'
        },
        {
          tanggalPengajuan: '12/05/2025',
          nama: 'HENDRAWAN SUJATMIKO, S.T.',
          kotaAsal: 'Banyuwangi',
          tujuanTransmigrasi: 'Kota Palu',
          nomorRegistrasi: '2025/1/03/CATRANS/YK',
          status: 'Selesai'
        }
      ]
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1);
    },
    getStatusClass(status) {
      switch((status||'').toLowerCase()) {
        case 'proses':
          return 'bg-warning text-dark';
        case 'selesai':
          return 'bg-success text-white';
        case 'ditolak':
          return 'bg-danger text-white';
        case 'pending':
          return 'bg-secondary text-white';
        default:
          return 'bg-primary text-white';
      }
    },
    goToHome() {
      this.$router.push('/');
    }
  }
}
</script>

<style scoped>
.status-pengajuan-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  min-height: 100vh;
}

.header-section {
  background: transparent;
  padding-bottom: 10px;
}

.back-btn {
  color: #6c757d;
  font-size: 1.2rem;
  text-decoration: none;
  transition: color 0.2s ease;
}

.back-btn:hover {
  color: #495057;
}

.status-cards {
  max-width: 100%;
}

.status-card {
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.status-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 15px rgba(0,0,0,0.1) !important;
}

.card {
  border-radius: 12px;
  background: #FFF8F6;
}

.status-icon {
  position: relative;
  width: 56px;
  height: 56px;
}

.icon-wrapper {
  position: relative;
  width: 56px;
  height: 56px;
  /* background: #696868; */ /* Hapus background */
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}
.icon-wrapper .fa-mobile-alt {
  display: none;
}
.history-icon-img {
  width: 90px;
  height: 90px;
  object-fit: contain;
  display: block;
}

.check-mark {
  position: absolute;
  bottom: -5px;
  right: -5px;
  width: 24px;
  height: 24px;
  background: #00d600;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid white;
}

.check-mark .fa-check {
  font-size: 0.7rem;
}

.status-badge {
  font-size: 0.875rem;
  min-width: 80px;
  display: inline-block;
  text-align: center;
}

.text-danger {
  color: #dc3545 !important;
}

.fw-medium {
  font-weight: 500 !important;
}

.empty-state {
  border-radius: 12px;
  margin: 2rem 0;
}

.history-card-image-wrapper {
  display: none;
}
.history-card-image {
  display: none;
}
@media (max-width: 576px) {
  .history-card-image {
    max-width: 80px;
  }
}

/* Responsive Design */
@media (max-width: 992px) {
  .status-pengajuan-container {
    padding: 10px;
  }
  .card-body {
    padding: 1.2rem !important;
  }
  .status-icon {
    width: 45px;
    height: 45px;
  }
  .icon-wrapper .fa-mobile-alt {
    font-size: 1rem;
  }
  .check-mark {
    width: 16px;
    height: 16px;
  }
  .check-mark .fa-check {
    font-size: 0.5rem;
  }
}

@media (max-width: 768px) {
  .status-icon,
  .icon-wrapper {
    width: 38px;
    height: 38px;
  }
  .history-icon-img {
    width: 28px;
    height: 28px;
  }
}

@media (max-width: 576px) {
  .status-icon,
  .icon-wrapper {
    width: 28px;
    height: 28px;
  }
  .history-icon-img {
    width: 18px;
    height: 18px;
  }
  .header-section h4 {
    font-size: 1rem;
  }
  .status-card {
    margin-bottom: 0.7rem;
  }
  .card-body {
    padding: 0.7rem !important;
  }
  .status-icon {
    width: 32px;
    height: 32px;
  }
  .icon-wrapper .fa-mobile-alt {
    font-size: 0.8rem;
  }
  .check-mark {
    width: 10px;
    height: 10px;
  }
  .check-mark .fa-check {
    font-size: 0.35rem;
  }
  .status-badge {
    font-size: 0.7rem;
    min-width: 48px;
    padding: 3px 6px;
  }
  .empty-state {
    padding: 1.5rem 0;
    margin: 1rem 0;
  }
  .row.align-items-center {
    flex-direction: column;
    align-items: flex-start !important;
  }
  .col-md-8,
  .col-md-4 {
    width: 100% !important;
    max-width: 100% !important;
    text-align: left !important;
    margin-top: 0 !important;
  }
  .col-md-4.text-md-end {
    text-align: left !important;
    margin-top: 1rem;
  }
}

.btn-beranda {
  background: linear-gradient(135deg, #dc3545, #e74c3c);
  color: #fff;
  font-weight: 600;
  border-radius: 10px;
  padding: 12px 32px;
  font-size: 1rem;
  box-shadow: 0 4px 16px rgba(220,53,69,0.12);
  border: none;
  transition: all 0.2s;
}
.btn-beranda:hover {
  background: linear-gradient(135deg, #c82333, #dc3545);
  color: #fff;
  transform: translateY(-2px);
}

/* Custom scrollbar */
.status-pengajuan-container::-webkit-scrollbar {
  width: 6px;
}

.status-pengajuan-container::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.status-pengajuan-container::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

.status-pengajuan-container::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style>