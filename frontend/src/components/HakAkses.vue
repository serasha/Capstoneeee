<template>
  <div class="hak-akses-container">
    <!-- Header Section -->
    <div class="header-section d-flex justify-content-between align-items-center mb-4">
      <h4 class="page-title mb-0">Manajemen Hak Akses</h4>
      <div class="user-info d-flex align-items-center">
        <div class="me-3 text-end">
          <h6 class="mb-0 user-name">{{ userInfo.name }}</h6>
          <small class="text-muted">{{ userInfo.id }}</small>
        </div>
        <img 
          :src="userInfo.avatar" 
          :alt="userInfo.name" 
          class="user-avatar"
        >
      </div>
    </div>

    <!-- Access Form Section -->
    <div class="access-form-section">
      <div class="row justify-content-center">
        <div class="col-lg-6 col-md-8 col-sm-10">
          <div class="form-card">
            <div class="card-body p-4">
              <form @submit.prevent="submitAccess">
                <!-- ID User Field -->
                <div class="mb-4">
                  <label for="idUser" class="form-label">ID User</label>
                  <input 
                    type="text" 
                    id="idUser"
                    v-model="formData.idUser"
                    class="form-control form-control-lg"
                    :class="{ 'is-invalid': errors.idUser }"
                    placeholder="Masukkan ID User"
                    @blur="validateIdUser"
                    @input="clearError('idUser')"
                  >
                  <div v-if="errors.idUser" class="invalid-feedback">
                    {{ errors.idUser }}
                  </div>
                </div>

                <!-- Nama User Field -->
                <div class="mb-4">
                  <label for="namaUser" class="form-label">Nama User</label>
                  <input 
                    type="text" 
                    id="namaUser"
                    v-model="formData.namaUser"
                    class="form-control form-control-lg"
                    :class="{ 'is-invalid': errors.namaUser }"
                    placeholder="Masukkan Nama User"
                    @blur="validateNamaUser"
                    @input="clearError('namaUser')"
                  >
                  <div v-if="errors.namaUser" class="invalid-feedback">
                    {{ errors.namaUser }}
                  </div>
                </div>

                <!-- Pilih Akses Field -->
                <div class="mb-4">
                  <label for="pilihAkses" class="form-label">Pilih Akses</label>
                  <div class="dropdown-container">
                    <select 
                      id="pilihAkses"
                      v-model="formData.pilihAkses"
                      class="form-select form-select-lg"
                      :class="{ 'is-invalid': errors.pilihAkses }"
                      @change="validatePilihAkses"
                    >
                      <option value="">Pilih Level Akses</option>
                      <option 
                        v-for="akses in aksesOptions" 
                        :key="akses.value" 
                        :value="akses.value"
                      >
                        {{ akses.label }}
                      </option>
                    </select>
                    <div v-if="errors.pilihAkses" class="invalid-feedback">
                      {{ errors.pilihAkses }}
                    </div>
                  </div>
                </div>

                <!-- Submit Button -->
                <div class="text-center">
                  <button 
                    type="submit" 
                    class="btn btn-danger btn-lg px-5"
                    :disabled="isSubmitting"
                  >
                    <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2"></span>
                    {{ isSubmitting ? 'Menyimpan...' : 'SIMPAN AKSES' }}
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent Access List -->
    <div class="recent-access-section mt-5">
      <div class="row">
        <div class="col-12">
          <div class="access-list-card">
            <div class="card-header bg-light">
              <h5 class="mb-0">Daftar Hak Akses Terbaru</h5>
            </div>
            <div class="card-body p-0">
              <div class="table-responsive">
                <table class="table table-hover mb-0">
                  <thead class="table-light">
                    <tr>
                      <th>No</th>
                      <th>ID User</th>
                      <th>Nama User</th>
                      <th>Level Akses</th>
                      <th>Tanggal Dibuat</th>
                      <th>Status</th>
                      <th>Aksi</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr 
                      v-for="(access, index) in recentAccess" 
                      :key="access.id"
                    >
                      <td>{{ index + 1 }}</td>
                      <td class="user-id">{{ access.idUser }}</td>
                      <td>{{ access.namaUser }}</td>
                      <td>
                        <span 
                          class="badge"
                          :class="getAccessBadgeClass(access.levelAkses)"
                        >
                          {{ access.levelAkses }}
                        </span>
                      </td>
                      <td>{{ formatDate(access.tanggalDibuat) }}</td>
                      <td>
                        <span 
                          class="badge"
                          :class="access.status === 'Aktif' ? 'bg-success' : 'bg-secondary'"
                        >
                          {{ access.status }}
                        </span>
                      </td>
                      <td>
                        <div class="btn-group btn-group-sm">
                          <button 
                            class="btn btn-outline-primary btn-sm"
                            @click="editAccess(access)"
                            title="Edit"
                          >
                            <i class="fas fa-edit"></i>
                          </button>
                          <button 
                            class="btn btn-outline-danger btn-sm"
                            @click="deleteAccess(access.id)"
                            title="Hapus"
                          >
                            <i class="fas fa-trash"></i>
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Success/Error Messages -->
    <div 
      v-if="message.show" 
      class="alert alert-dismissible fade show mt-3"
      :class="message.type === 'success' ? 'alert-success' : 'alert-danger'"
    >
      {{ message.text }}
      <button 
        type="button" 
        class="btn-close" 
        @click="message.show = false"
      ></button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HakAkses',
  data() {
    return {
      userInfo: {
        name: 'Hendra Hermawan',
        id: '190001290',
        avatar: 'https://via.placeholder.com/50x50/007bff/ffffff?text=HH'
      },
      formData: {
        idUser: '',
        namaUser: '',
        pilihAkses: ''
      },
      errors: {},
      isSubmitting: false,
      message: {
        show: false,
        type: 'success',
        text: ''
      },
      aksesOptions: [
        { value: 'super_admin', label: 'Super Admin' },
        { value: 'admin', label: 'Admin' },
        { value: 'operator', label: 'Operator' },
        { value: 'user', label: 'User' },
        { value: 'guest', label: 'Guest' }
      ],
      recentAccess: [
        {
          id: 1,
          idUser: '190001290',
          namaUser: 'Hendra Hermawan',
          levelAkses: 'Super Admin',
          tanggalDibuat: new Date('2025-01-15'),
          status: 'Aktif'
        },
        {
          id: 2,
          idUser: '190001291',
          namaUser: 'Ahmad Rizki',
          levelAkses: 'Admin',
          tanggalDibuat: new Date('2025-01-14'),
          status: 'Aktif'
        },
        {
          id: 3,
          idUser: '190001292',
          namaUser: 'Siti Nurhaliza',
          levelAkses: 'Operator',
          tanggalDibuat: new Date('2025-01-13'),
          status: 'Aktif'
        },
        {
          id: 4,
          idUser: '190001293',
          namaUser: 'Budi Santoso',
          levelAkses: 'User',
          tanggalDibuat: new Date('2025-01-12'),
          status: 'Nonaktif'
        }
      ]
    }
  },
  methods: {
    validateIdUser() {
      if (!this.formData.idUser.trim()) {
        this.errors.idUser = 'ID User harus diisi'
      } else if (this.formData.idUser.length < 5) {
        this.errors.idUser = 'ID User minimal 5 karakter'
      }
    },
    validateNamaUser() {
      if (!this.formData.namaUser.trim()) {
        this.errors.namaUser = 'Nama User harus diisi'
      } else if (this.formData.namaUser.length < 3) {
        this.errors.namaUser = 'Nama User minimal 3 karakter'
      }
    },
    validatePilihAkses() {
      if (!this.formData.pilihAkses) {
        this.errors.pilihAkses = 'Level akses harus dipilih'
      }
    },
    clearError(field) {
      if (this.errors[field]) {
        delete this.errors[field]
      }
    },
    validateForm() {
      this.errors = {}
      this.validateIdUser()
      this.validateNamaUser()
      this.validatePilihAkses()
      return Object.keys(this.errors).length === 0
    },
    async submitAccess() {
      if (!this.validateForm()) {
        return
      }

      this.isSubmitting = true
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1500))
        
        // Add to recent access
        const newAccess = {
          id: this.recentAccess.length + 1,
          idUser: this.formData.idUser,
          namaUser: this.formData.namaUser,
          levelAkses: this.aksesOptions.find(opt => opt.value === this.formData.pilihAkses)?.label,
          tanggalDibuat: new Date(),
          status: 'Aktif'
        }
        
        this.recentAccess.unshift(newAccess)
        
        // Reset form
        this.formData = {
          idUser: '',
          namaUser: '',
          pilihAkses: ''
        }
        
        this.showMessage('success', 'Hak akses berhasil disimpan!')
        
      } catch (error) {
        this.showMessage('error', 'Terjadi kesalahan saat menyimpan data')
      } finally {
        this.isSubmitting = false
      }
    },
    editAccess(access) {
      this.formData = {
        idUser: access.idUser,
        namaUser: access.namaUser,
        pilihAkses: this.aksesOptions.find(opt => opt.label === access.levelAkses)?.value || ''
      }
      
      // Scroll to form
      document.querySelector('.form-card').scrollIntoView({ 
        behavior: 'smooth' 
      })
    },
    async deleteAccess(id) {
      if (confirm('Apakah Anda yakin ingin menghapus hak akses ini?')) {
        this.recentAccess = this.recentAccess.filter(access => access.id !== id)
        this.showMessage('success', 'Hak akses berhasil dihapus!')
      }
    },
    getAccessBadgeClass(levelAkses) {
      const badgeClasses = {
        'Super Admin': 'bg-danger',
        'Admin': 'bg-warning text-dark',
        'Operator': 'bg-info',
        'User': 'bg-success',
        'Guest': 'bg-secondary'
      }
      return badgeClasses[levelAkses] || 'bg-secondary'
    },
    formatDate(date) {
      return new Intl.DateTimeFormat('id-ID', {
        day: '2-digit',
        month: 'short',
        year: 'numeric'
      }).format(date)
    },
    showMessage(type, text) {
      this.message = {
        show: true,
        type,
        text
      }
      
      setTimeout(() => {
        this.message.show = false
      }, 5000)
    }
  }
}
</script>

<style scoped>
.hak-akses-container {
  background: #f8f9fa;
  min-height: 100vh;
  padding: 24px;
}

.header-section {
  background: #fff;
  padding: 20px 24px;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.08);
  border: 1px solid #e9ecef;
}

.page-title {
  color: #333;
  font-weight: 600;
}

.user-info {
  background: rgba(0,123,255,0.1);
  padding: 8px 16px;
  border-radius: 25px;
  border: 1px solid rgba(0,123,255,0.2);
}

.user-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #007bff;
}

.user-name {
  color: #333;
  font-weight: 600;
}

.access-form-section {
  margin-top: 32px;
}

.form-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  border: none;
  overflow: hidden;
}

.form-label {
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.form-control-lg, .form-select-lg {
  border: 2px solid #e9ecef;
  border-radius: 8px;
  padding: 12px 16px;
  font-size: 1rem;
  transition: all 0.3s ease;
}

.form-control-lg:focus, .form-select-lg:focus {
  border-color: #007bff;
  box-shadow: 0 0 0 0.2rem rgba(0,123,255,0.25);
}

.form-control-lg.is-invalid, .form-select-lg.is-invalid {
  border-color: #dc3545;
}

.dropdown-container {
  position: relative;
}

.btn-danger {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  border: none;
  border-radius: 8px;
  padding: 12px 32px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(220,53,69,0.3);
}

.btn-danger:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(220,53,69,0.4);
}

.btn-danger:disabled {
  opacity: 0.6;
  transform: none;
  box-shadow: none;
}

.recent-access-section {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  overflow: hidden;
}

.access-list-card {
  border: none;
  border-radius: 12px;
  overflow: hidden;
}

.card-header {
  border-bottom: 2px solid #e9ecef;
  padding: 16px 24px;
}

.table th {
  background-color: #f8f9fa;
  border-color: #dee2e6;
  font-weight: 600;
  color: #495057;
  padding: 12px 16px;
  border-top: none;
}

.table td {
  padding: 12px 16px;
  vertical-align: middle;
  border-color: #dee2e6;
}

.table tbody tr:hover {
  background-color: rgba(0,123,255,0.05);
}

.user-id {
  font-family: 'Courier New', monospace;
  font-weight: 500;
  color: #007bff;
}

.badge {
  font-size: 0.8em;
  padding: 6px 12px;
  border-radius: 20px;
  font-weight: 500;
}

.btn-group-sm .btn {
  padding: 4px 8px;
  font-size: 0.8em;
  border-radius: 4px;
}

.btn-outline-primary:hover {
  transform: scale(1.05);
}

.btn-outline-danger:hover {
  transform: scale(1.05);
}

.alert {
  border-radius: 8px;
  border: none;
  padding: 16px 20px;
  font-weight: 500;
}

.alert-success {
  background: linear-gradient(135deg, #d4edda 0%, #c3e6cb 100%);
  color: #155724;
}

.alert-danger {
  background: linear-gradient(135deg, #f8d7da 0%, #f5c6cb 100%);
  color: #721c24;
}

/* Loading Animation */
.spinner-border-sm {
  width: 1rem;
  height: 1rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .hak-akses-container {
    padding: 16px;
  }
  
  .header-section {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 16px;
  }
  
  .user-info {
    align-self: flex-end;
  }
  
  .form-card .card-body {
    padding: 24px 16px;
  }
  
  .btn-danger {
    width: 100%;
    padding: 14px;
  }
  
  .table-responsive {
    font-size: 0.9em;
  }
  
  .table th, .table td {
    padding: 8px 12px;
  }
}

@media (max-width: 576px) {
  .page-title {
    font-size: 1.2rem;
  }
  
  .user-info {
    flex-direction: column;
    text-align: center;
    padding: 12px;
  }
  
  .user-avatar {
    width: 40px;
    height: 40px;
  }
  
  .form-control-lg, .form-select-lg {
    padding: 10px 12px;
  }
  
  .table {
    font-size: 0.8em;
  }
  
  .btn-group-sm .btn {
    padding: 2px 6px;
    font-size: 0.7em;
  }
}
</style>