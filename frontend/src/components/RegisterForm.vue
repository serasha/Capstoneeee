<template>
  <div class="formulir-pengajuan">
    <!-- Header -->
    <div class="header-section mb-4">
      <button class="btn btn-link p-0 mb-3 text-decoration-none" @click="goBack">
        <i class="fas fa-arrow-left me-2"></i>
        <span class="fw-medium">Formulir Pengajuan</span>
      </button>
    </div>

    <form @submit.prevent="submitForm" class="registration-form">
      <div v-if="errorMsg" class="alert alert-danger mb-3">{{ errorMsg }}</div>
      <!-- Data Pribadi & Keluarga Section -->
      <div class="form-section mb-4">
        <div class="section-header">
          <h5 class="section-title">DATA PRIBADI & KELUARGA</h5>
          <p class="section-subtitle">Silahkan isi data di bawah ini:</p>
        </div>

        <div class="row g-3">
          <div class="col-md-6">
            <label class="form-label required">Nama Lengkap Kepala Keluarga</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.namaKepalaKeluarga"
              ref="namaKepalaKeluarga"
              required
              :class="{ 'is-invalid': fieldErrors.nama_pendaftar }"
            >
            <div v-if="fieldErrors.nama_pendaftar" class="invalid-feedback d-block">{{ fieldErrors.nama_pendaftar }}</div>
          </div>
          <div class="col-md-6">
            <label class="form-label required">Jenis Kelamin</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.jenisKelamin"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Nomor Induk Kependudukan (NIK)</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.nik"
              required
              inputmode="numeric"
              pattern="[0-9]*"
            >
          </div>
          <div class="col-md-6">
            <label class="form-label">Pilih di sini:</label>
            <div class="d-flex gap-2">
              <button 
                type="button" 
                class="btn btn-outline-danger custom-btn"
                :class="{ active: formData.pilihan === 'Laki-laki' }"
                @click="formData.pilihan = 'Laki-laki'"
              >
                Laki-laki
              </button>
              <button 
                type="button" 
                class="btn btn-outline-danger custom-btn"
                :class="{ active: formData.pilihan === 'Perempuan' }"
                @click="formData.pilihan = 'Perempuan'"
              >
                Perempuan
              </button>
            </div>
          </div>
          <div class="col-md-6">
            <label class="form-label required">Nomor Kartu Keluarga</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.nomorKK"
              required
              inputmode="numeric"
              pattern="[0-9]*"
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Pendidikan Terakhir</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.pendidikanTerakhir"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Tanggal Lahir</label>
            <input 
              type="date" 
              class="form-control custom-input" 
              v-model="formData.tanggalLahir"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Pekerjaan</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.pekerjaan"
              required
            >
          </div>
        </div>
      </div>

      <!-- Alamat Rumah Section -->
      <div class="form-section mb-4">
        <div class="section-header">
          <h5 class="section-title">ALAMAT RUMAH</h5>
          <p class="section-subtitle">Silahkan isi data di bawah ini:</p>
        </div>

        <div class="row g-3">
          <div class="col-md-6">
            <label class="form-label required">Nama Jalan/Desa</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.namaJalan"
              ref="namaJalan"
              required
              :class="{ 'is-invalid': fieldErrors.alamat_pendaftar }"
            >
            <div v-if="fieldErrors.alamat_pendaftar" class="invalid-feedback d-block">{{ fieldErrors.alamat_pendaftar }}</div>
          </div>
          <div class="col-md-6">
            <label class="form-label required">Nomor Telepon</label>
            <input 
              type="tel" 
              class="form-control custom-input" 
              v-model="formData.nomorTelepon"
              required
              inputmode="numeric"
              pattern="[0-9]*"
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Kelurahan</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.kelurahan"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Jumlah Anggota Keluarga</label>
            <input 
              type="number" 
              class="form-control custom-input" 
              v-model="formData.jumlahAnggota"
              required
              inputmode="numeric"
              pattern="[0-9]*"
            >
          </div>
          <div class="col-md-4">
            <label class="form-label required">Kode Pos</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.kodePos"
              required
              inputmode="numeric"
              pattern="[0-9]*"
            >
          </div>
          <div class="col-md-4">
            <label class="form-label required">Provinsi</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.provinsi"
              required
            >
          </div>
          <div class="col-md-4">
            <label class="form-label required">Status (Kawin/Belum Kawin)</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.status"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Kemantren</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.kemantren"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label">Pilih di sini:</label>
            <div class="d-flex gap-2">
              <button 
                type="button" 
                class="btn btn-outline-danger custom-btn"
                :class="{ active: formData.statusKawin === 'Kawin' }"
                @click="formData.statusKawin = 'Kawin'"
              >
                Kawin
              </button>
              <button 
                type="button" 
                class="btn btn-outline-danger custom-btn"
                :class="{ active: formData.statusKawin === 'Belum Kawin' }"
                @click="formData.statusKawin = 'Belum Kawin'"
              >
                Belum Kawin
              </button>
            </div>
          </div>
          <div class="col-md-6">
            <label class="form-label required">Kota</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.kota"
              required
            >
          </div>
          <div class="col-md-6">
            <label class="form-label required">Nama Anggota Keluarga</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.namaAnggotaKeluarga"
              required
            >
          </div>
          <div class="col-12">
            <label class="form-label required">Upload Dokumen (Fotocopy: KTP, KK, Buku Nikah)</label>
            <div class="upload-area" @click="triggerFileUpload" @dragover.prevent @drop.prevent="handleFileDrop">
              <input 
                type="file" 
                ref="fileInput" 
                class="d-none" 
                multiple 
                accept=".pdf,.jpg,.jpeg,.png"
                @change="handleFileUpload"
              >
              <div class="upload-content">
                <i class="fas fa-cloud-upload-alt upload-icon"></i>
                <p class="upload-text mb-0">
                  {{ uploadedFiles.length > 0 ? `${uploadedFiles.length} file(s) selected` : 'Click to upload or drag files here' }}
                </p>
              </div>
            </div>
            <div v-if="uploadedFiles.length > 0" class="uploaded-files mt-2">
              <div v-for="(file, index) in uploadedFiles" :key="index" class="uploaded-file">
                <span>{{ file.name }}</span>
                <button type="button" class="btn btn-sm btn-outline-danger" @click="removeFile(index)">
                  <i class="fas fa-times"></i>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Daerah Tujuan Transmigrasi Section -->
      <div class="form-section mb-4">
        <div class="section-header">
          <h5 class="section-title">DAERAH TUJUAN TRANSMIGRASI</h5>
        </div>

        <div class="row g-3">
          <div class="col-12">
            <label class="form-label required">Nama Lokasi Yang Diminati</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.namaLokasi"
              required
            >
          </div>
          <div class="col-12">
            <label class="form-label required">Pola Usaha Yang Diminati</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.polaUsaha"
              required
            >
          </div>
          <div class="col-12">
            <label class="form-label required">Nomor Pendaftaran Saudara</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.nomorPendaftaran"
              required
              inputmode="numeric"
              pattern="[0-9]*"
            >
          </div>
        </div>
      </div>

      <!-- Disclaimer -->
      <div class="disclaimer-section mb-4">
        <p class="disclaimer-text">
          SAYA SECARA SADAR DAN SUKARELA MENGAJUKAN PERMOHONAN MENGIKUTI PROGRAM TRANSMIGRASI YANG DISELENGGARAKAN OLEH PEMERINTAH. SAYA 
          MENYATAKAN BELUM PERNAH MENGIKUTI PROGRAM TRANSMIGRASI DAN DATA YANG DIISI BENAR DAN DAPAT DIPERTANGGUNGJAWABKAN, SERTA SETUJU UNTUK 
          MENAATI SEGALA PERATURAN DALAM RANGKA PENDAFTARAN TRANSMIGRASI INI.
        </p>
      </div>

      <!-- Submit Button -->
      <div class="text-center">
        <button type="submit" class="btn btn-submit" :disabled="isSubmitting">
          <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2"></span>
          {{ isSubmitting ? 'Memproses...' : 'SUBMIT PENDAFTARAN' }}
        </button>
      </div>
    </form>

    <SuccessModal
      :visible="showSuccessModal"
       :userData="{ fullName: loggedInUserName || formData.namaKepalaKeluarga }"
      @close="showSuccessModal = false"
      @view-status="goToStatus"
      @go-to-dashboard="goToDashboard"
      @go-to-home="goToHome"
    />
  </div>
</template>

<script>
import SuccessModal from './SuccessModal.vue';

export default {
  name: 'FormulirPendaftaran',
  components: { SuccessModal },
  data() {
    return {
      isSubmitting: false,
      uploadedFiles: [],
      showSuccessModal: false,
      formData: {
        namaKepalaKeluarga: '',
        jenisKelamin: '',
        nik: '',
        pilihan: '',
        nomorKK: '',
        pendidikanTerakhir: '',
        tanggalLahir: '',
        pekerjaan: '',
        namaJalan: '',
        nomorTelepon: '',
        kelurahan: '',
        jumlahAnggota: '',
        kodePos: '',
        provinsi: '',
        status: '',
        kemantren: '',
        statusKawin: '',
        kota: '',
        namaAnggotaKeluarga: '',
        namaLokasi: '',
        polaUsaha: '',
        nomorPendaftaran: ''
      }
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1);
    },
    triggerFileUpload() {
      this.$refs.fileInput.click();
    },
    handleFileUpload(event) {
      const files = Array.from(event.target.files);
      this.uploadedFiles = [...this.uploadedFiles, ...files];
    },
    handleFileDrop(event) {
      const files = Array.from(event.dataTransfer.files);
      this.uploadedFiles = [...this.uploadedFiles, ...files];
    },
    removeFile(index) {
      this.uploadedFiles.splice(index, 1);
    },
    async submitForm() {
      this.isSubmitting = true;
      this.errorMsg = '';
      this.fieldErrors = {};
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // Handle form submission
        console.log('Form Data:', this.formData);
        console.log('Uploaded Files:', this.uploadedFiles);
        
        this.showSuccessModal = true; // Tampilkan modal setelah submit sukses
        
      } catch (error) {
        // errorMsg dan fieldErrors sudah di-set
      } finally {
        this.isSubmitting = false;
      }
    },
    goToStatus() {
      this.showSuccessModal = false;
      this.$router.push('/status');
    },
    goToDashboard() {
      this.showSuccessModal = false;
      this.$router.push('/dashboard');
    },
    goToHome() {
      this.showSuccessModal = false;
      this.$router.push('/');
    },
  }
}
</script>

<style scoped>
.formulir-pendaftaran {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
}

.header-section .btn-link {
  color: #333;
  font-size: 18px;
  text-decoration: none !important;
}

.header-section .btn-link:hover {
  color: #dc3545;
}

.form-section {
  background: #ffffff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
  overflow: hidden;
  margin-bottom: 24px;
}

.section-header {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  padding: 20px;
  text-align: center;
  color: white;
}

.section-title {
  font-size: 16px;
  font-weight: 700;
  margin: 0;
  letter-spacing: 0.5px;
}

.section-subtitle {
  font-size: 14px;
  margin: 8px 0 0 0;
  opacity: 0.9;
}

.form-section .row {
  padding: 24px;
}

.form-label {
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
  font-size: 14px;
}

.form-label.required::after {
  content: ' *';
  color: #dc3545;
}

.custom-input {
  border: 2px solid #e9ecef;
  border-radius: 8px;
  padding: 12px 16px;
  font-size: 14px;
  transition: all 0.3s ease;
  background-color: #fafafa;
}

.custom-input:focus {
  border-color: #dc3545;
  box-shadow: 0 0 0 0.2rem rgba(220, 53, 69, 0.15);
  background-color: #ffffff;
}

.custom-btn {
  border: 2px solid #dc3545;
  color: #dc3545;
  font-weight: 600;
  padding: 8px 20px;
  border-radius: 6px;
  transition: all 0.3s ease;
  font-size: 14px;
}

.custom-btn:hover {
  background-color: #dc3545;
  color: white;
  border-color: #dc3545;
}

.custom-btn.active {
  background-color: #dc3545;
  color: white;
  border-color: #dc3545;
}

.upload-area {
  border: 2px dashed #dc3545;
  border-radius: 12px;
  padding: 40px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background-color: #fafafa;
}

.upload-area:hover {
  border-color: #c82333;
  background-color: #fff5f5;
}

.upload-icon {
  font-size: 48px;
  color: #dc3545;
  margin-bottom: 16px;
}

.upload-text {
  color: #666;
  font-size: 16px;
  font-weight: 500;
}

.uploaded-files {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.uploaded-file {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: #f8f9fa;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 14px;
  border: 1px solid #dee2e6;
}

.disclaimer-section {
  background-color: #fff3cd;
  border: 1px solid #ffeaa7;
  border-radius: 8px;
  padding: 20px;
}

.disclaimer-text {
  font-size: 13px;
  line-height: 1.6;
  color: #856404;
  margin: 0;
  text-align: justify;
}

.btn-submit {
  background: linear-gradient(135deg, #28a745 0%, #20c997 100%);
  border: none;
  color: white;
  padding: 15px 40px;
  font-size: 16px;
  font-weight: 700;
  border-radius: 8px;
  transition: all 0.3s ease;
  min-width: 200px;
  letter-spacing: 0.5px;
}

.btn-submit:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(40, 167, 69, 0.3);
}

.btn-submit:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

/* Responsive Design */
@media (max-width: 768px) {
  .formulir-pendaftaran {
    padding: 16px;
  }
  
  .form-section .row {
    padding: 20px 16px;
  }
  
  .section-header {
    padding: 16px;
  }
  
  .section-title {
    font-size: 14px;
  }
  
  .upload-area {
    padding: 30px 20px;
  }
  
  .upload-icon {
    font-size: 36px;
  }
  
  .custom-btn {
    padding: 10px 16px;
    font-size: 13px;
  }
  
  .btn-submit {
    width: 100%;
    padding: 14px;
  }
}

@media (max-width: 576px) {
  .d-flex.gap-2 {
    flex-direction: column;
  }
  
  .custom-btn {
    width: 100%;
  }
  
  .uploaded-file {
    width: 100%;
    justify-content: space-between;
  }
}
</style>