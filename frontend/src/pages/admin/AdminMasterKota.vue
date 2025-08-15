<template>
  <div class="admin-master-kota">
    <div class="page-header">
      <h1 class="page-title">Master Kota Tujuan</h1>
      <button class="btn btn-primary" @click="showAddModal = true">
        <PlusIcon style="width:20px;height:20px;" class="me-2" />Tambah Kota
      </button>
    </div>

    <!-- Search and Filter -->
    <div class="filter-section">
      <div class="row">
        <div class="col-md-6">
          <input
            type="text"
            v-model="searchQuery"
            class="form-control"
            placeholder="Cari kota atau provinsi..."
          />
        </div>
        <div class="col-md-6">
          <select v-model="statusFilter" class="form-control">
            <option value="">Semua Status</option>
            <option value="true">Aktif</option>
            <option value="false">Nonaktif</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Data Table -->
    <div class="table-section">
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>No</th>
              <th>Nama Kota</th>
              <th>Provinsi</th>
              <th>Status</th>
              <th>Tanggal Dibuat</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(kota, index) in paginatedData" :key="kota.id">
              <td>{{ getRowNumber(index) }}</td>
              <td>{{ kota.nama_kota }}</td>
              <td>{{ kota.provinsi }}</td>
              <td>
                <span 
                  class="badge"
                  :class="kota.status_aktif ? 'badge-success' : 'badge-secondary'"
                >
                  {{ kota.status_aktif ? 'Aktif' : 'Nonaktif' }}
                </span>
              </td>
              <td>{{ formatDate(kota.created_at) }}</td>
              <td>
                <button 
                  class="btn btn-sm btn-outline-primary me-2"
                  @click="editKota(kota)"
                >
                  <PencilSquareIcon style="width:20px;height:20px;" />
                </button>
                <button 
                  class="btn btn-sm btn-outline-danger"
                  @click="deleteKota(kota.id)"
                >
                  <TrashIcon style="width:20px;height:20px;" />
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="pagination-section">
        <div class="pagination-info">
          Showing {{ startEntry }} to {{ endEntry }} of {{ filteredData.length }} entries
        </div>
        <div class="pagination-controls">
          <button 
            class="pagination-btn" 
            @click="previousPage" 
            :disabled="currentPage === 1"
          >
            <
          </button>
          <span class="page-number">{{ currentPage }}</span>
          <button 
            class="pagination-btn" 
            @click="nextPage" 
            :disabled="currentPage === totalPages"
          >
            >
          </button>
        </div>
      </div>
    </div>

    <!-- Add/Edit Modal -->
    <div v-if="showAddModal || showEditModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h5 class="modal-title">
            {{ showAddModal ? 'Tambah Kota Baru' : 'Edit Kota' }}
          </h5>
          <button class="btn-close" @click="closeModal">&times;</button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="submitForm">
            <div class="form-group mb-3">
              <label class="form-label">Nama Kota</label>
              <input
                type="text"
                v-model="formData.nama_kota"
                class="form-control"
                required
              />
            </div>
            <div class="form-group mb-3">
              <label class="form-label">Provinsi</label>
              <input
                type="text"
                v-model="formData.provinsi"
                class="form-control"
                required
              />
            </div>
            <div class="form-group mb-3">
              <label class="form-label">Status</label>
              <select v-model="formData.status_aktif" class="form-control">
                <option :value="true">Aktif</option>
                <option :value="false">Nonaktif</option>
              </select>
            </div>
          </form>
        </div>
        <div class="modal-footer">
          <button class="btn btn-secondary" @click="closeModal">Batal</button>
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
  </div>
</template>

<script>
import { PlusIcon, PencilSquareIcon, TrashIcon } from '@heroicons/vue/24/solid'
export default {
  name: 'AdminMasterKota',
  components: { PlusIcon, PencilSquareIcon, TrashIcon },
  data() {
    return {
      searchQuery: '',
      statusFilter: '',
      currentPage: 1,
      entriesPerPage: 10,
      kotaData: [],
      showAddModal: false,
      showEditModal: false,
      isSubmitting: false,
      formData: {
        nama_kota: '',
        provinsi: '',
        status_aktif: true
      },
      editingId: null
    }
  },
  computed: {
    filteredData() {
      let filtered = this.kotaData;
      
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase();
        filtered = filtered.filter(kota => 
          kota.nama_kota.toLowerCase().includes(query) ||
          kota.provinsi.toLowerCase().includes(query)
        );
      }
      
      if (this.statusFilter !== '') {
        filtered = filtered.filter(kota => 
          kota.status_aktif.toString() === this.statusFilter
        );
      }
      
      return filtered;
    },
    paginatedData() {
      const start = (this.currentPage - 1) * this.entriesPerPage;
      const end = start + this.entriesPerPage;
      return this.filteredData.slice(start, end);
    },
    totalPages() {
      return Math.ceil(this.filteredData.length / this.entriesPerPage);
    },
    startEntry() {
      return (this.currentPage - 1) * this.entriesPerPage + 1;
    },
    endEntry() {
      const end = this.currentPage * this.entriesPerPage;
      return Math.min(end, this.filteredData.length);
    }
  },
  methods: {
    async fetchData() {
      try {
        const res = await fetch('/api/kota', { credentials: 'include' });
        if (res.ok) {
          this.kotaData = await res.json();
        }
      } catch (e) {
        console.error('Failed to fetch kota data:', e);
      }
    },
    getRowNumber(index) {
      return (this.currentPage - 1) * this.entriesPerPage + index + 1;
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleDateString('id-ID');
    },
    editKota(kota) {
      this.formData = { ...kota };
      this.editingId = kota.id;
      this.showEditModal = true;
    },
    async deleteKota(id) {
      if (confirm('Apakah Anda yakin ingin menghapus kota ini?')) {
        try {
          const res = await fetch(`/api/kota/${id}`, {
            method: 'DELETE',
            credentials: 'include'
          });
          if (res.ok) {
            this.fetchData();
            alert('Kota berhasil dihapus!');
          }
        } catch (e) {
          alert('Gagal menghapus kota');
        }
      }
    },
    async submitForm() {
      this.isSubmitting = true;
      try {
        const url = this.showAddModal ? '/api/kota' : `/api/kota/${this.editingId}`;
        const method = this.showAddModal ? 'POST' : 'PUT';
        
        const res = await fetch(url, {
          method,
          headers: { 'Content-Type': 'application/json' },
          credentials: 'include',
          body: JSON.stringify(this.formData)
        });
        
        if (res.ok) {
          this.fetchData();
          this.closeModal();
          alert(this.showAddModal ? 'Kota berhasil ditambahkan!' : 'Kota berhasil diupdate!');
        }
      } catch (e) {
        alert('Gagal menyimpan data kota');
      } finally {
        this.isSubmitting = false;
      }
    },
    closeModal() {
      this.showAddModal = false;
      this.showEditModal = false;
      this.formData = {
        nama_kota: '',
        provinsi: '',
        status_aktif: true
      };
      this.editingId = null;
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }
    }
  },
  mounted() {
    this.fetchData();
  }
}
</script>

<style scoped>
.admin-master-kota {
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.page-title {
  font-size: 2rem;
  font-weight: 700;
  color: #333;
  margin: 0;
}

.filter-section {
  background: white;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  margin-bottom: 1rem;
}

.table-section {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.table-container {
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table th {
  background: #f8f9fa;
  padding: 1rem;
  text-align: left;
  font-weight: 600;
  color: #333;
  border-bottom: 2px solid #e9ecef;
}

.data-table td {
  padding: 1rem;
  border-bottom: 1px solid #e9ecef;
}

.data-table tbody tr:hover {
  background: #f8f9fa;
}

.badge {
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 500;
}

.badge-success {
  background: #28a745;
  color: white;
}

.badge-secondary {
  background: #6c757d;
  color: white;
}

.pagination-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-top: 1px solid #e9ecef;
}

.pagination-info {
  color: #6c757d;
  font-size: 0.9rem;
}

.pagination-controls {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.pagination-btn {
  padding: 0.5rem 0.75rem;
  border: 1px solid #dee2e6;
  background: white;
  color: #333;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.pagination-btn:hover:not(:disabled) {
  background: #e9ecef;
}

.pagination-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-number {
  padding: 0.5rem 0.75rem;
  background: #dc3545;
  color: white;
  border-radius: 4px;
  font-weight: 600;
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
  max-width: 500px;
  width: 90%;
}

.modal-header {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-body {
  padding: 1.5rem;
}

.modal-footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid #e9ecef;
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
}

.btn-close {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
}

.form-group {
  margin-bottom: 1rem;
}

.form-label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #333;
}

.form-control {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  font-size: 1rem;
}

.form-control:focus {
  outline: none;
  border-color: #dc3545;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-primary {
  background: #dc3545;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #c82333;
}

.btn-secondary {
  background: #6c757d;
  color: white;
}

.btn-secondary:hover {
  background: #5a6268;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.spinner-border-sm {
  width: 1rem;
  height: 1rem;
  border-width: 0.1em;
}
</style>