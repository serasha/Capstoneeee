<template>
  <div class="hak-akses-container">
    <!-- Header Section -->
    <div class="header-section d-flex justify-content-between align-items-center mb-4">
      <h4 class="page-title mb-0">Manajemen Hak Akses</h4>
      <div class="header-actions d-flex gap-2">
        <button 
          class="btn btn-success btn-sm"
          @click="showAddModal = true"
        >
          <i class="fas fa-plus me-2"></i>Tambah User
        </button>
        <button 
          class="btn btn-outline-primary btn-sm"
          @click="refreshData"
          :disabled="loading"
        >
          <i class="fas fa-sync-alt me-2" :class="{ 'fa-spin': loading }"></i>Refresh
        </button>
      </div>
    </div>

    <!-- Search and Filter Section -->
    <div class="filter-section mb-4">
      <div class="row">
        <div class="col-md-4">
          <label class="form-label">Cari Data</label>
          <input 
            type="text" 
            v-model="searchQuery"
            class="form-control"
            placeholder="Cari berdasarkan nama, username, email..."
            @input="filterData"
          >
        </div>
        <div class="col-md-3">
          <label class="form-label">Filter Level Akses</label>
          <select v-model="levelFilter" @change="filterData" class="form-select">
            <option value="">Semua Level</option>
            <option value="Admin">Admin</option>
            <option value="user">User</option>
            <option value="operator">Operator</option>
          </select>
        </div>
        <div class="col-md-3">
          <label class="form-label">Filter Status</label>
          <select v-model="statusFilter" @change="filterData" class="form-select">
            <option value="">Semua Status</option>
            <option value="true">Aktif</option>
            <option value="false">Nonaktif</option>
          </select>
        </div>
        <div class="col-md-2">
          <label class="form-label">Show</label>
          <select v-model="entriesPerPage" @change="updatePagination" class="form-select">
            <option value="5">5</option>
            <option value="10">10</option>
            <option value="25">25</option>
            <option value="50">50</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Data Table -->
    <div class="table-section">
      <div class="table-responsive">
        <table class="table table-hover">
          <thead class="table-light">
            <tr>
              <th style="width: 5%;">No</th>
              <th style="width: 12%;">ID User</th>
              <th style="width: 20%;">Nama User</th>
              <th style="width: 15%;">Username</th>
              <th style="width: 18%;">Email</th>
              <th style="width: 12%;">Level Akses</th>
              <th style="width: 8%;">Status</th>
              <th style="width: 10%;">Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr 
              v-for="(item, index) in paginatedData" 
              :key="`${item.type}-${item.id}`"
              class="table-row"
            >
              <td>{{ getRowNumber(index) }}</td>
              <td>
                <span class="user-id-badge" :class="getUserIdBadgeClass(item.type)">
                  {{ item.id_user }}
                </span>
              </td>
              <td>
                <div class="user-name-cell">
                  <strong>{{ item.nama_user }}</strong>
                  <small class="text-muted d-block">{{ item.type === 'admin' ? 'Administrator' : 'Regular User' }}</small>
                </div>
              </td>
              <td>{{ item.username }}</td>
              <td>{{ item.email || '-' }}</td>
              <td>
                <span 
                  class="badge level-badge"
                  :class="getLevelBadgeClass(item.level_akses)"
                >
                  {{ item.level_akses }}
                </span>
              </td>
              <td>
                <div class="status-cell">
                  <span 
                    class="badge status-badge"
                    :class="item.status_aktif ? 'bg-success' : 'bg-secondary'"
                  >
                    {{ item.status_aktif ? 'Aktif' : 'Nonaktif' }}
                  </span>
                  <button 
                    v-if="item.type === 'admin'"
                    class="btn btn-sm btn-outline-secondary ms-1"
                    @click="toggleStatus(item)"
                    :disabled="loading"
                    title="Toggle Status"
                  >
                    <i class="fas fa-toggle-on" v-if="item.status_aktif"></i>
                    <i class="fas fa-toggle-off" v-else></i>
                  </button>
                </div>
              </td>
              <td>
                <div class="btn-group btn-group-sm">
                  <button 
                    class="btn btn-outline-primary"
                    @click="editItem(item)"
                    title="Edit"
                    :disabled="loading"
                  >
                    <i class="fas fa-edit"></i>
                  </button>
                  <button 
                    class="btn btn-outline-danger"
                    @click="deleteItem(item)"
                    title="Hapus"
                    :disabled="loading"
                  >
                    <i class="fas fa-trash"></i>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Empty State -->
      <div v-if="!loading && filteredData.length === 0" class="empty-state text-center py-5">
        <i class="fas fa-users fa-3x text-muted mb-3"></i>
        <h5 class="text-muted">Tidak ada data ditemukan</h5>
        <p class="text-muted">Coba ubah filter pencarian atau tambah user baru</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="loading-state text-center py-5">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
        <p class="text-muted mt-2">Memuat data...</p>
      </div>
    </div>

    <!-- Pagination -->
    <div class="pagination-section d-flex justify-content-between align-items-center mt-4">
      <div class="pagination-info">
        <span class="text-muted">
          Menampilkan {{ getStartRecord() }} sampai {{ getEndRecord() }} dari {{ filteredData.length }} data
        </span>
      </div>
      <nav aria-label="Page navigation">
        <ul class="pagination pagination-sm mb-0">
          <li class="page-item" :class="{ disabled: currentPage === 1 }">
            <button 
              class="page-link" 
              @click="previousPage"
              :disabled="currentPage === 1"
            >
              &laquo;
            </button>
          </li>
          <li class="page-item" :class="{ disabled: currentPage === 1 }">
            <button 
              class="page-link" 
              @click="previousPage"
              :disabled="currentPage === 1"
            >
              &lt;
            </button>
          </li>
          <li 
            v-for="page in visiblePages" 
            :key="page"
            class="page-item"
            :class="{ active: page === currentPage }"
          >
            <button 
              class="page-link" 
              @click="goToPage(page)"
              v-if="page !== '...'"
            >
              {{ page }}
            </button>
            <span class="page-link" v-else>...</span>
          </li>
          <li class="page-item" :class="{ disabled: currentPage === totalPages }">
            <button 
              class="page-link" 
              @click="nextPage"
              :disabled="currentPage === totalPages"
            >
              &gt;
            </button>
          </li>
          <li class="page-item" :class="{ disabled: currentPage === totalPages }">
            <button 
              class="page-link" 
              @click="nextPage"
              :disabled="currentPage === totalPages"
            >
              &raquo;
            </button>
          </li>
        </ul>
      </nav>
    </div>

    <!-- Add/Edit Modal -->
    <div v-if="showAddModal || showEditModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h5 class="modal-title">
            {{ showAddModal ? 'Tambah User Baru' : 'Edit User' }}
          </h5>
          <button class="btn-close" @click="closeModal">&times;</button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="submitForm">
            <div class="row">
              <div class="col-md-6">
                <div class="form-group mb-3">
                  <label class="form-label required">Nama User</label>
                  <input
                    type="text"
                    v-model="formData.nama_user"
                    class="form-control"
                    :class="{ 'is-invalid': errors.nama_user }"
                    placeholder="Masukkan nama lengkap"
                    required
                  />
                  <div v-if="errors.nama_user" class="invalid-feedback">
                    {{ errors.nama_user }}
                  </div>
                </div>
              </div>
              <div class="col-md-6">
                <div class="form-group mb-3">
                  <label class="form-label required">Username</label>
                  <input
                    type="text"
                    v-model="formData.username"
                    class="form-control"
                    :class="{ 'is-invalid': errors.username }"
                    placeholder="Masukkan username"
                    required
                  />
                  <div v-if="errors.username" class="invalid-feedback">
                    {{ errors.username }}
                  </div>
                </div>
              </div>
            </div>

            <div class="row">
              <div class="col-md-6">
                <div class="form-group mb-3">
                  <label class="form-label">Email</label>
                  <input
                    type="email"
                    v-model="formData.email"
                    class="form-control"
                    :class="{ 'is-invalid': errors.email }"
                    placeholder="Masukkan email (opsional untuk user)"
                  />
                  <div v-if="errors.email" class="invalid-feedback">
                    {{ errors.email }}
                  </div>
                </div>
              </div>
              <div class="col-md-6">
                <div class="form-group mb-3">
                  <label class="form-label required">Level Akses</label>
                  <select
                    v-model="formData.level_akses"
                    class="form-control"
                    :class="{ 'is-invalid': errors.level_akses }"
                    required
                  >
                    <option value="">Pilih Level Akses</option>
                    <option value="admin">Admin</option>
                    <option value="user">User</option>
                    <option value="operator">Operator</option>
                  </select>
                  <div v-if="errors.level_akses" class="invalid-feedback">
                    {{ errors.level_akses }}
                  </div>
                </div>
              </div>
            </div>

            <div class="row" v-if="showAddModal">
              <div class="col-md-6">
                <div class="form-group mb-3">
                  <label class="form-label required">Password</label>
                  <input
                    type="password"
                    v-model="formData.password"
                    class="form-control"
                    :class="{ 'is-invalid': errors.password }"
                    placeholder="Masukkan password"
                    required
                  />
                  <div v-if="errors.password" class="invalid-feedback">
                    {{ errors.password }}
                  </div>
                </div>
              </div>
              <div class="col-md-6">
                <div class="form-group mb-3">
                  <label class="form-label required">Konfirmasi Password</label>
                  <input
                    type="password"
                    v-model="formData.confirm_password"
                    class="form-control"
                    :class="{ 'is-invalid': errors.confirm_password }"
                    placeholder="Konfirmasi password"
                    required
                  />
                  <div v-if="errors.confirm_password" class="invalid-feedback">
                    {{ errors.confirm_password }}
                  </div>
                </div>
              </div>
            </div>

            <div class="row" v-if="formData.level_akses === 'admin'">
              <div class="col-12">
                <div class="form-group mb-3">
                  <div class="form-check">
                    <input
                      type="checkbox"
                      id="statusAktif"
                      v-model="formData.status_aktif"
                      class="form-check-input"
                    />
                    <label class="form-check-label" for="statusAktif">
                      Status Aktif
                    </label>
                  </div>
                  <small class="text-muted">Admin yang tidak aktif tidak dapat login ke sistem</small>
                </div>
              </div>
            </div>
          </form>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="closeModal" :disabled="isSubmitting">
            Batal
          </button>
          <button 
            class="btn btn-primary" 
            @click="submitForm"
            :disabled="isSubmitting"
          >
            <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-2"></span>
            {{ showAddModal ? 'Tambah' : 'Update' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Success/Error Messages -->
    <div 
      v-if="message.show" 
      class="alert alert-dismissible fade show position-fixed"
      :class="message.type === 'success' ? 'alert-success' : 'alert-danger'"
      style="top: 20px; right: 20px; z-index: 9999; min-width: 300px;"
    >
      <i :class="message.type === 'success' ? 'fas fa-check-circle' : 'fas fa-exclamation-circle'" class="me-2"></i>
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
      loading: false,
      searchQuery: '',
      levelFilter: '',
      statusFilter: '',
      entriesPerPage: 10,
      currentPage: 1,
      showAddModal: false,
      showEditModal: false,
      isSubmitting: false,
      editingItem: null,
      
      formData: {
        nama_user: '',
        username: '',
        email: '',
        password: '',
        confirm_password: '',
        level_akses: '',
        status_aktif: true
      },
      
      errors: {},
      
      message: {
        show: false,
        type: 'success',
        text: ''
      },
      
      hakAksesData: [],
      filteredData: []
    }
  },
  
  computed: {
    totalPages() {
      return Math.ceil(this.filteredData.length / this.entriesPerPage)
    },
    
    paginatedData() {
      const start = (this.currentPage - 1) * this.entriesPerPage
      const end = start + this.entriesPerPage
      return this.filteredData.slice(start, end)
    },
    
    visiblePages() {
      const pages = []
      const total = this.totalPages
      const current = this.currentPage
      
      if (total <= 7) {
        for (let i = 1; i <= total; i++) {
          pages.push(i)
        }
      } else {
        if (current <= 4) {
          for (let i = 1; i <= 5; i++) {
            pages.push(i)
          }
          pages.push('...')
          pages.push(total)
        } else if (current >= total - 3) {
          pages.push(1)
          pages.push('...')
          for (let i = total - 4; i <= total; i++) {
            pages.push(i)
          }
        } else {
          pages.push(1)
          pages.push('...')
          for (let i = current - 1; i <= current + 1; i++) {
            pages.push(i)
          }
          pages.push('...')
          pages.push(total)
        }
      }
      
      return pages
    }
  },
  
  async created() {
    await this.fetchData()
  },
  
  methods: {
    async fetchData() {
      this.loading = true
      try {
        const response = await fetch('/api/hak-akses', {
          credentials: 'include'
        })
        
        if (!response.ok) {
          throw new Error('Gagal mengambil data')
        }
        
        this.hakAksesData = await response.json()
        this.filterData()
      } catch (error) {
        this.showMessage('error', 'Gagal mengambil data: ' + error.message)
      } finally {
        this.loading = false
      }
    },
    
    filterData() {
      let filtered = [...this.hakAksesData]
      
      // Filter berdasarkan search query
      if (this.searchQuery.trim()) {
        const query = this.searchQuery.toLowerCase()
        filtered = filtered.filter(item => 
          item.nama_user.toLowerCase().includes(query) ||
          item.username.toLowerCase().includes(query) ||
          (item.email && item.email.toLowerCase().includes(query)) ||
          item.id_user.toLowerCase().includes(query)
        )
      }
      
      // Filter berdasarkan level akses
      if (this.levelFilter) {
        filtered = filtered.filter(item => 
          item.level_akses.toLowerCase() === this.levelFilter.toLowerCase()
        )
      }
      
      // Filter berdasarkan status
      if (this.statusFilter !== '') {
        const isActive = this.statusFilter === 'true'
        filtered = filtered.filter(item => item.status_aktif === isActive)
      }
      
      this.filteredData = filtered
      this.currentPage = 1
    },
    
    updatePagination() {
      this.currentPage = 1
    },
    
    getRowNumber(index) {
      return (this.currentPage - 1) * this.entriesPerPage + index + 1
    },
    
    getStartRecord() {
      return this.filteredData.length === 0 ? 0 : (this.currentPage - 1) * this.entriesPerPage + 1
    },
    
    getEndRecord() {
      const end = this.currentPage * this.entriesPerPage
      return Math.min(end, this.filteredData.length)
    },
    
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--
      }
    },
    
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++
      }
    },
    
    goToPage(page) {
      if (page !== '...' && page >= 1 && page <= this.totalPages) {
        this.currentPage = page
      }
    },
    
    getUserIdBadgeClass(type) {
      return {
        'badge-admin': type === 'admin',
        'badge-user': type === 'user'
      }
    },
    
    getLevelBadgeClass(level) {
      const classes = {
        'Admin': 'bg-danger',
        'admin': 'bg-danger',
        'user': 'bg-primary',
        'operator': 'bg-warning text-dark'
      }
      return classes[level] || 'bg-secondary'
    },
    
    editItem(item) {
      this.editingItem = item
      this.formData = {
        nama_user: item.nama_user,
        username: item.username,
        email: item.email || '',
        password: '',
        confirm_password: '',
        level_akses: item.level_akses.toLowerCase(),
        status_aktif: item.status_aktif
      }
      this.showEditModal = true
      this.errors = {}
    },
    
    async deleteItem(item) {
      if (!confirm(`Apakah Anda yakin ingin menghapus ${item.nama_user}?`)) {
        return
      }
      
      this.loading = true
      try {
        const response = await fetch(`/api/hak-akses/${item.id}?type=${item.type}`, {
          method: 'DELETE',
          credentials: 'include'
        })
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Gagal menghapus data')
        }
        
        this.showMessage('success', `${item.nama_user} berhasil dihapus`)
        await this.fetchData()
      } catch (error) {
        this.showMessage('error', error.message)
      } finally {
        this.loading = false
      }
    },
    
    async toggleStatus(item) {
      if (item.type !== 'admin') {
        this.showMessage('error', 'Hanya admin yang memiliki status aktif/nonaktif')
        return
      }
      
      this.loading = true
      try {
        const response = await fetch(`/api/hak-akses/${item.id}/toggle-status?type=${item.type}`, {
          method: 'PATCH',
          credentials: 'include'
        })
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Gagal mengubah status')
        }
        
        const result = await response.json()
        this.showMessage('success', result.message)
        await this.fetchData()
      } catch (error) {
        this.showMessage('error', error.message)
      } finally {
        this.loading = false
      }
    },
    
    validateForm() {
      this.errors = {}
      
      if (!this.formData.nama_user.trim()) {
        this.errors.nama_user = 'Nama user wajib diisi'
      }
      
      if (!this.formData.username.trim()) {
        this.errors.username = 'Username wajib diisi'
      } else if (this.formData.username.length < 3) {
        this.errors.username = 'Username minimal 3 karakter'
      }
      
      if (this.formData.level_akses === 'admin' && !this.formData.email.trim()) {
        this.errors.email = 'Email wajib diisi untuk admin'
      }
      
      if (!this.formData.level_akses) {
        this.errors.level_akses = 'Level akses wajib dipilih'
      }
      
      if (this.showAddModal) {
        if (!this.formData.password) {
          this.errors.password = 'Password wajib diisi'
        } else if (this.formData.password.length < 6) {
          this.errors.password = 'Password minimal 6 karakter'
        }
        
        if (!this.formData.confirm_password) {
          this.errors.confirm_password = 'Konfirmasi password wajib diisi'
        } else if (this.formData.password !== this.formData.confirm_password) {
          this.errors.confirm_password = 'Password tidak cocok'
        }
      }
      
      return Object.keys(this.errors).length === 0
    },
    
    async submitForm() {
      if (!this.validateForm()) {
        return
      }
      
      this.isSubmitting = true
      
      try {
        const url = this.showAddModal 
          ? '/api/hak-akses' 
          : `/api/hak-akses/${this.editingItem.id}?type=${this.editingItem.type}`
        
        const method = this.showAddModal ? 'POST' : 'PUT'
        
        const response = await fetch(url, {
          method,
          headers: {
            'Content-Type': 'application/json'
          },
          credentials: 'include',
          body: JSON.stringify(this.formData)
        })
        
        if (!response.ok) {
          const errorData = await response.json()
          throw new Error(errorData.error || 'Gagal menyimpan data')
        }
        
        const result = await response.json()
        this.showMessage('success', result.message)
        this.closeModal()
        await this.fetchData()
      } catch (error) {
        this.showMessage('error', error.message)
      } finally {
        this.isSubmitting = false
      }
    },
    
    closeModal() {
      this.showAddModal = false
      this.showEditModal = false
      this.editingItem = null
      this.formData = {
        nama_user: '',
        username: '',
        email: '',
        password: '',
        confirm_password: '',
        level_akses: '',
        status_aktif: true
      }
      this.errors = {}
    },
    
    async refreshData() {
      await this.fetchData()
      this.showMessage('success', 'Data berhasil diperbarui')
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
  font-size: 1.5rem;
}

.header-actions .btn {
  font-weight: 500;
}

.filter-section {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.08);
  border: 1px solid #e9ecef;
}

.form-label {
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

.form-control, .form-select {
  border: 2px solid #e9ecef;
  border-radius: 8px;
  padding: 10px 12px;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.form-control:focus, .form-select:focus {
  border-color: #007bff;
  box-shadow: 0 0 0 0.2rem rgba(0,123,255,0.25);
}

.table-section {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.08);
  border: 1px solid #e9ecef;
  overflow: hidden;
}

.table {
  margin-bottom: 0;
}

.table th {
  background-color: #f8f9fa;
  border-color: #dee2e6;
  font-weight: 600;
  color: #495057;
  padding: 16px 12px;
  border-top: none;
  font-size: 0.9rem;
}

.table td {
  padding: 16px 12px;
  vertical-align: middle;
  border-color: #dee2e6;
  font-size: 0.9rem;
}

.table-row:hover {
  background-color: rgba(0,123,255,0.05);
}

.user-id-badge {
  font-family: 'Courier New', monospace;
  font-weight: 600;
  font-size: 0.8rem;
  padding: 4px 8px;
  border-radius: 4px;
  color: white;
}

.badge-admin {
  background-color: #dc3545;
}

.badge-user {
  background-color: #007bff;
}

.user-name-cell strong {
  color: #333;
  font-size: 0.95rem;
}

.user-name-cell small {
  font-size: 0.75rem;
}

.level-badge {
  font-size: 0.8em;
  padding: 6px 12px;
  border-radius: 20px;
  font-weight: 500;
}

.status-cell {
  display: flex;
  align-items: center;
}

.status-badge {
  font-size: 0.8em;
  padding: 6px 12px;
  border-radius: 20px;
  font-weight: 500;
}

.btn-group-sm .btn {
  padding: 6px 10px;
  font-size: 0.8rem;
  border-radius: 4px;
}

.empty-state, .loading-state {
  padding: 60px 20px;
}

.pagination-section {
  background: white;
  padding: 20px 24px;
  border-top: 1px solid #e9ecef;
}

.pagination-info {
  font-size: 0.9rem;
  color: #6c757d;
}

.pagination .page-link {
  color: #007bff;
  border-color: #dee2e6;
  padding: 8px 12px;
  font-size: 0.9rem;
}

.pagination .page-item.active .page-link {
  background-color: #007bff;
  border-color: #007bff;
}

.pagination .page-item.disabled .page-link {
  color: #6c757d;
  background-color: #fff;
  border-color: #dee2e6;
}

/* Modal Styles */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 12px;
  max-width: 600px;
  width: 90%;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
}

.modal-header {
  padding: 20px 24px;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.btn-close {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #6c757d;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-close:hover {
  color: #333;
}

.modal-body {
  padding: 24px;
}

.modal-footer {
  padding: 20px 24px;
  border-top: 1px solid #e9ecef;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

.form-group {
  margin-bottom: 1rem;
}

.form-label.required::after {
  content: ' *';
  color: #dc3545;
}

.form-control.is-invalid, .form-select.is-invalid {
  border-color: #dc3545;
}

.invalid-feedback {
  color: #dc3545;
  font-size: 0.875rem;
  margin-top: 4px;
}

.form-check-input:checked {
  background-color: #007bff;
  border-color: #007bff;
}

.form-check-label {
  font-weight: 500;
  color: #333;
}

/* Alert Styles */
.alert {
  border: none;
  border-radius: 8px;
  font-weight: 500;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.alert-success {
  background: linear-gradient(135deg, #d4edda 0%, #c3e6cb 100%);
  color: #155724;
}

.alert-danger {
  background: linear-gradient(135deg, #f8d7da 0%, #f5c6cb 100%);
  color: #721c24;
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
  
  .header-actions {
    width: 100%;
    justify-content: flex-end;
  }
  
  .filter-section .row > div {
    margin-bottom: 16px;
  }
  
  .table-responsive {
    font-size: 0.8rem;
  }
  
  .table th, .table td {
    padding: 12px 8px;
  }
  
  .pagination-section {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
  
  .modal-content {
    width: 95%;
    margin: 20px;
  }
  
  .modal-body .row > div {
    margin-bottom: 12px;
  }
}

@media (max-width: 576px) {
  .btn-group-sm {
    flex-direction: column;
    gap: 4px;
  }
  
  .btn-group-sm .btn {
    width: 100%;
  }
  
  .status-cell {
    flex-direction: column;
    gap: 4px;
  }
  
  .user-id-badge {
    font-size: 0.7rem;
    padding: 2px 6px;
  }
  
  .level-badge, .status-badge {
    font-size: 0.7rem;
    padding: 4px 8px;
  }
}

/* Loading Animation */
.spinner-border-sm {
  width: 1rem;
  height: 1rem;
}

.fa-spin {
  animation: fa-spin 1s infinite linear;
}

@keyframes fa-spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Smooth Transitions */
.table-row {
  transition: background-color 0.2s ease;
}

.btn {
  transition: all 0.2s ease;
}

.form-control, .form-select {
  transition: border-color 0.2s ease, box-shadow 0.2s ease;
}
</style>