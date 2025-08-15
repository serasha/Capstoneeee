<template>
  <div class="formulir-pengajuan">
    <!-- Header -->
    <div class="header-section mb-4">
      <button class="btn btn-link p-0 mb-3 text-decoration-none" @click="goBack">
        <i class="fas fa-arrow-left me-2"></i>
        <span class="fw-medium">
          {{ isEditMode ? 'Edit Formulir Pengajuan' : 'Formulir Pengajuan' }}
        </span>
      </button>
      <!-- Alert untuk mode edit -->
      <div v-if="isEditMode" class="alert alert-info">
        <div class="d-flex align-items-center">
          <i class="fas fa-edit me-2"></i>
          <div>
            <strong>Mode Edit</strong><br>
            <small>Anda sedang mengedit formulir yang telah dikembalikan. Silakan perbaiki data yang diperlukan.</small>
          </div>
        </div>
      </div>
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
            <select
              class="form-select custom-input"
              v-model="formData.jenisKelamin"
              required
            >
              <option disabled value="">Pilih jenis kelamin</option>
              <option value="laki-laki">Laki-laki</option>
              <option value="perempuan">Perempuan</option>
            </select>
          </div>
          <div class="col-md-6">
            <label class="form-label required">Nomor Induk Kependudukan (NIK)</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.nik"
              required
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
            >
          </div>
          <div class="col-md-4">
            <label class="form-label required">Kode Pos</label>
            <input 
              type="text" 
              class="form-control custom-input" 
              v-model="formData.kodePos"
              required
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
            <select
              class="form-select custom-input"
              v-model="formData.status"
              required
            >
              <option disabled value="">Pilih status</option>
              <option value="Kawin">Kawin</option>
              <option value="Belum Kawin">Belum Kawin</option>
            </select>
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
            
            <!-- Upload KTP -->
            <div class="document-upload-section mb-3">
              <label class="document-label">KTP (Kartu Tanda Penduduk)</label>
              <div class="upload-area" @click="triggerFileUpload('ktp')" @dragover.prevent @drop.prevent="handleFileDrop($event, 'ktp')">
                <input 
                  type="file" 
                  ref="ktpInput" 
                  class="d-none" 
                  accept=".pdf,.jpg,.jpeg,.png"
                  @change="handleFileUpload($event, 'ktp')"
                >
                <div class="upload-content">
                  <i class="fas fa-id-card upload-icon"></i>
                  <p class="upload-text mb-0">
                    {{ uploadedFiles.ktp ? uploadedFiles.ktp.name : 'Upload KTP' }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Upload KK -->
            <div class="document-upload-section mb-3">
              <label class="document-label">KK (Kartu Keluarga)</label>
              <div class="upload-area" @click="triggerFileUpload('kk')" @dragover.prevent @drop.prevent="handleFileDrop($event, 'kk')">
                <input 
                  type="file" 
                  ref="kkInput" 
                  class="d-none" 
                  accept=".pdf,.jpg,.jpeg,.png"
                  @change="handleFileUpload($event, 'kk')"
                >
                <div class="upload-content">
                  <i class="fas fa-users upload-icon"></i>
                  <p class="upload-text mb-0">
                    {{ uploadedFiles.kk ? uploadedFiles.kk.name : 'Upload KK' }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Upload Surat Nikah -->
            <div class="document-upload-section mb-3">
              <label class="document-label">Surat Nikah</label>
              <div class="upload-area" @click="triggerFileUpload('surat_nikah')" @dragover.prevent @drop.prevent="handleFileDrop($event, 'surat_nikah')">
                <input 
                  type="file" 
                  ref="suratNikahInput" 
                  class="d-none" 
                  accept=".pdf,.jpg,.jpeg,.png"
                  @change="handleFileUpload($event, 'surat_nikah')"
                >
                <div class="upload-content">
                  <i class="fas fa-heart upload-icon"></i>
                  <p class="upload-text mb-0">
                    {{ uploadedFiles.surat_nikah ? uploadedFiles.surat_nikah.name : 'Upload Surat Nikah' }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Upload Ijazah -->
            <div class="document-upload-section mb-3">
              <label class="document-label">Ijazah</label>
              <div class="upload-area" @click="triggerFileUpload('ijazah')" @dragover.prevent @drop.prevent="handleFileDrop($event, 'ijazah')">
                <input 
                  type="file" 
                  ref="ijazahInput" 
                  class="d-none" 
                  accept=".pdf,.jpg,.jpeg,.png"
                  @change="handleFileUpload($event, 'ijazah')"
                >
                <div class="upload-content">
                  <i class="fas fa-graduation-cap upload-icon"></i>
                  <p class="upload-text mb-0">
                    {{ uploadedFiles.ijazah ? uploadedFiles.ijazah.name : 'Upload Ijazah' }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Upload Pas Foto -->
            <div class="document-upload-section mb-3">
              <label class="document-label">Pas Foto</label>
              <div class="upload-area" @click="triggerFileUpload('pas_foto')" @dragover.prevent @drop.prevent="handleFileDrop($event, 'pas_foto')">
                <input 
                  type="file" 
                  ref="pasFotoInput" 
                  class="d-none" 
                  accept=".jpg,.jpeg,.png"
                  @change="handleFileUpload($event, 'pas_foto')"
                >
                <div class="upload-content">
                  <i class="fas fa-camera upload-icon"></i>
                  <p class="upload-text mb-0">
                    {{ uploadedFiles.pas_foto ? uploadedFiles.pas_foto.name : 'Upload Pas Foto' }}
                  </p>
                </div>
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
            <select 
              class="form-control custom-input" 
              v-model="formData.namaLokasi"
              required
            >
              <option value="">Pilih Lokasi Tujuan</option>
              <option v-for="kota in kotaList" :key="kota.id" :value="kota.nama_kota">
                {{ kota.nama_kota }}, {{ kota.provinsi }}
              </option>
            </select>
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
          {{ isSubmitting ? 'Memproses...' : (isEditMode ? 'PERBARUI DATA' : 'SUBMIT PENDAFTARAN') }}
        </button>
      </div>
    </form>

    <SuccessModal
      :visible="showSuccessModal"
      :userData="{ fullName: formData.namaKepalaKeluarga }"
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
      isEditMode: false,
      editData: null,
      uploadedFiles: {
        ktp: null,
        kk: null,
        surat_nikah: null,
        ijazah: null,
        pas_foto: null
      },
      showSuccessModal: false,
      user: null,
      errorMsg: '',
      fieldErrors: {},
      kotaList: [],
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
  async created() {
    // Cek login
    try {
      const res = await fetch('/api/user/me', { credentials: 'include' });
      if (res.ok) {
        this.user = await res.json();
      } else {
        this.$router.push('/login');
      }
    } catch {
      this.$router.push('/login');
    }
    
    // Load kota list
    this.loadKotaList();
    
    // Cek jika ada data untuk edit
    this.checkEditMode();
  },
  methods: {
    async loadKotaList() {
      try {
        const res = await fetch('/api/kota', { credentials: 'include' });
        if (res.ok) {
          this.kotaList = await res.json();
        }
      } catch (e) {
        console.error('Failed to load kota list:', e);
      }
    },
    goBack() {
      this.$router.go(-1);
    },
    checkEditMode() {
      // Cek apakah ada data edit di localStorage
      const editDataStr = localStorage.getItem('editFormData');
      if (editDataStr) {
        try {
          this.editData = JSON.parse(editDataStr);
          if (this.editData.isEditing) {
            this.isEditMode = true;
            this.loadFormDataForEdit();
            // Hapus data edit dari localStorage setelah dimuat
            localStorage.removeItem('editFormData');
          }
        } catch (e) {
          console.error('Error parsing edit data:', e);
        }
      }
    },
    async loadFormDataForEdit() {
      if (!this.editData || !this.editData.nomorRegistrasi) return;
      
      try {
        // Load data pendaftaran berdasarkan nomor registrasi
        const res = await fetch(`/api/pendaftaran/${this.editData.nomorRegistrasi}`, { 
          credentials: 'include' 
        });
        
        if (res.ok) {
          const data = await res.json();
          
          // Isi form dengan data yang ada
          this.formData = {
            namaKepalaKeluarga: data.nama_pendaftar || this.editData.nama || '',
            jenisKelamin: data.jenis_kelamin || '',
            nik: data.nik || '',
            pilihan: data.jenis_layanan || this.editData.tujuanTransmigrasi || '',
            nomorKK: data.nomor_kk || '',
            pendidikanTerakhir: data.pendidikan_terakhir || '',
            tanggalLahir: data.tanggal_lahir || '',
            pekerjaan: data.pekerjaan || '',
            namaJalan: data.alamat_pendaftar || this.editData.kotaAsal || '',
            nomorTelepon: data.nomor_telepon || '',
            kelurahan: data.kelurahan || '',
            jumlahAnggota: data.jumlah_anggota || '',
            kodePos: data.kode_pos || '',
            provinsi: data.provinsi || '',
            status: data.status || '',
            kemantren: data.kemantren || '',
            statusKawin: data.status_kawin || '',
            kota: data.kota || '',
            namaAnggotaKeluarga: data.nama_anggota_keluarga || '',
            namaLokasi: data.nama_lokasi || '',
            polaUsaha: data.pola_usaha || '',
            nomorPendaftaran: this.editData.nomorRegistrasi || ''
          };
          
          // Set error message untuk mode edit
          this.errorMsg = `Mode Edit: Anda sedang mengedit formulir dengan nomor registrasi ${this.editData.nomorRegistrasi}. Status saat ini: ${this.editData.status}`;
          
        } else {
          // Jika tidak bisa load dari API, gunakan data dari History
          this.formData.namaKepalaKeluarga = this.editData.nama || '';
          this.formData.namaJalan = this.editData.kotaAsal || '';
          this.formData.pilihan = this.editData.tujuanTransmigrasi || '';
          this.formData.nomorPendaftaran = this.editData.nomorRegistrasi || '';
          
          this.errorMsg = `Mode Edit: Menggunakan data dasar. Silakan lengkapi data yang diperlukan untuk formulir ${this.editData.nomorRegistrasi}`;
        }
      } catch (e) {
        console.error('Error loading form data for edit:', e);
        // Fallback ke data dari History
        this.formData.namaKepalaKeluarga = this.editData.nama || '';
        this.formData.namaJalan = this.editData.kotaAsal || '';
        this.formData.pilihan = this.editData.tujuanTransmigrasi || '';
        this.formData.nomorPendaftaran = this.editData.nomorRegistrasi || '';
        
        this.errorMsg = `Mode Edit: Gagal memuat data lengkap. Silakan lengkapi formulir untuk ${this.editData.nomorRegistrasi}`;
      }
    },
    triggerFileUpload(type) {
      const refName = type + 'Input';
      if (this.$refs[refName]) {
        this.$refs[refName].click();
      } else {
        console.error('File input ref not found:', refName, this.$refs);
      }
    },
    handleFileUpload(event, type) {
      const file = event.target.files[0];
      if (file) {
        this.uploadedFiles[type] = file;
      }
    },
    handleFileDrop(event, type) {
      const file = event.dataTransfer.files[0];
      if (file) {
        this.uploadedFiles[type] = file;
      }
    },
    removeFile(type) {
      this.uploadedFiles[type] = null;
    },
    async submitForm() {
      this.isSubmitting = true;
      this.errorMsg = '';
      this.fieldErrors = {};
      try {
        // Tentukan URL dan method berdasarkan mode
        const url = this.isEditMode ? 
          `/api/pendaftaran/${this.editData.nomorRegistrasi}` : 
          '/api/pendaftaran';
        const method = this.isEditMode ? 'PUT' : 'POST';
        
        // Kirim SEMUA data ke backend
        const response = await fetch(url, {
          method: method,
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
          body: JSON.stringify({
            nama_pendaftar: this.formData.namaKepalaKeluarga,
            nik: this.formData.nik,
            tanggal_lahir: this.formData.tanggalLahir,
            jenis_kelamin: this.formData.jenisKelamin,
            nomor_kk: this.formData.nomorKK,
            pendidikan_terakhir: this.formData.pendidikanTerakhir,
            pekerjaan: this.formData.pekerjaan,
            alamat_pendaftar: this.formData.namaJalan,
            nomor_telepon: this.formData.nomorTelepon,
            kelurahan: this.formData.kelurahan,
            jumlah_anggota: this.formData.jumlahAnggota,
            kode_pos: this.formData.kodePos,
            provinsi: this.formData.provinsi,
            status: this.formData.status,
            kemantren: this.formData.kemantren,
            status_kawin: this.formData.statusKawin,
            kota: this.formData.kota,
            nama_anggota_keluarga: this.formData.namaAnggotaKeluarga,
            nama_lokasi: this.formData.namaLokasi,
            pola_usaha: this.formData.polaUsaha,
            nomor_pendaftaran: this.formData.nomorPendaftaran,
            jenis_layanan: this.formData.pilihan,
            cara_pendaftar: 'online',
            dokumen_administrasi_pendaftar: '', // handle upload terpisah jika perlu
            status_pendaftar: this.isEditMode ? 'pending' : 'pending', // Reset status ke pending setelah edit
          }),
        });
        if (!response.ok) {
          let errorData = {};
          try {
            errorData = await response.json();
          } catch (e) {
            errorData = { error: this.isEditMode ? 'Gagal mengupdate data' : 'Gagal mendaftar' };
          }
          if (errorData.errors) {
            this.fieldErrors = errorData.errors;
            this.errorMsg = '';
          } else {
            this.errorMsg = errorData.error || (this.isEditMode ? 'Gagal mengupdate data' : 'Gagal mendaftar');
          }
          throw new Error(this.errorMsg || 'Validasi error');
        }
        // Jika sukses
        const successMessage = this.isEditMode ? 
          'Data berhasil diperbarui! Status telah direset ke pending untuk review ulang.' : 
          'Pendaftaran berhasil! Status: pending.';
        alert(successMessage);
        this.showSuccessModal = true;
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
  padding: 20px;
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  background-color: #fafafa;
  min-height: 80px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.upload-area:hover {
  border-color: #c82333;
  background-color: #fff5f5;
}

.upload-icon {
  font-size: 24px;
  color: #dc3545;
  margin-right: 8px;
}

.upload-text {
  color: #666;
  font-size: 16px;
  font-weight: 500;
}

.document-upload-section {
  border: 1px solid #e9ecef;
  border-radius: 8px;
  padding: 1rem;
  background: #f8f9fa;
}

.document-label {
  font-weight: 600;
  color: #333;
  margin-bottom: 0.5rem;
  display: block;
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