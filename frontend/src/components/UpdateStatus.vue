<template>
  <div class="update-status-dashboard">
    <!-- Header Section -->
    <div class="dashboard-header d-flex justify-content-between align-items-center mb-4">
      <div class="user-info d-flex align-items-center">
        <img 
          src="https://via.placeholder.com/40x40/6c757d/ffffff?text=HH" 
          alt="User Avatar" 
          class="user-avatar me-3"
        >
        <div>
          <h5 class="mb-0">Hendra Hermawan</h5>
          <small class="text-muted">190002290</small>
        </div>
      </div>
      <button class="btn btn-outline-danger btn-sm">
        <i class="fas fa-download me-2"></i>Download Data
      </button>
    </div>

    <!-- Search and Filter Section -->
    <div class="filter-section mb-4">
      <div class="row">
        <div class="col-md-6 col-lg-4 mb-3">
          <label class="form-label">Show entries</label>
          <select v-model="entriesPerPage" @change="updatePagination" class="form-select">
            <option value="5">5</option>
            <option value="10">10</option>
            <option value="25">25</option>
            <option value="50">50</option>
          </select>
        </div>
        <div class="col-md-6 col-lg-4 mb-3">
          <label class="form-label">Cari Data</label>
          <input 
            v-model="searchQuery" 
            type="text" 
            class="form-control" 
            placeholder="Cari berdasarkan nama, kota, dll..."
          >
        </div>
        <div class="col-md-6 col-lg-4 mb-3">
          <label class="form-label">Filter Status</label>
          <select v-model="statusFilter" class="form-select">
            <option value="">Semua Status</option>
            <option value="Selesai">Selesai</option>
            <option value="Update">Update</option>
            <option value="Pending">Pending</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Data Table -->
    <div class="table-responsive">
      <table class="table table-striped table-hover">
        <thead class="table-light">
          <tr>
            <th scope="col">No</th>
            <th scope="col">No. Pendaftaran</th>
            <th scope="col">Hari/Tanggal</th>
            <th scope="col">Nama Pendaftar</th>
            <th scope="col">Kota Asal</th>
            <th scope="col">Tujuan Transmigrasi</th>
            <th scope="col">Status</th>
            <th scope="col">Aksi</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in paginatedData" :key="item.id">
            <td>{{ (currentPage - 1) * entriesPerPage + index + 1 }}</td>
            <td class="text-primary fw-semibold">{{ item.noPendaftaran }}</td>
            <td>
              <small class="text-muted">{{ item.tanggal }}</small>
            </td>
            <td class="fw-semibold">{{ item.namaPendaftar }}</td>
            <td>{{ item.kotaAsal }}</td>
            <td>{{ item.tujuanTransmigrasi }}</td>
            <td>
              <span 
                :class="getStatusBadgeClass(item.status)"
                class="badge"
              >
                {{ item.status }}
              </span>
            </td>
            <td>
              <div class="btn-group" role="group">
                <button 
                  @click="updateStatus(item)"
                  class="btn btn-sm btn-outline-primary"
                  data-bs-toggle="modal" 
                  data-bs-target="#updateStatusModal"
                >
                  <i class="fas fa-edit me-1"></i>Update
                </button>
                <button 
                  @click="viewDetail(item)"
                  class="btn btn-sm btn-outline-info"
                >
                  <i class="fas fa-eye me-1"></i>Detail
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination -->
    <div class="d-flex justify-content-between align-items-center mt-4">
      <div class="pagination-info">
        <small class="text-muted">
          Showing {{ startEntry }} to {{ endEntry }} of {{ filteredData.length }} entries
        </small>
      </div>
      <nav aria-label="Page navigation">
        <ul class="pagination pagination-sm mb-0">
          <li class="page-item" :class="{ disabled: currentPage === 1 }">
            <button 
              @click="changePage(currentPage - 1)" 
              class="page-link"
              :disabled="currentPage === 1"
            >
              <i class="fas fa-chevron-left"></i>
            </button>
          </li>
          <li 
            v-for="page in visiblePages" 
            :key="page"
            class="page-item" 
            :class="{ active: page === currentPage }"
          >
            <button @click="changePage(page)" class="page-link">{{ page }}</button>
          </li>
          <li class="page-item" :class="{ disabled: currentPage === totalPages }">
            <button 
              @click="changePage(currentPage + 1)" 
              class="page-link"
              :disabled="currentPage === totalPages"
            >
              <i class="fas fa-chevron-right"></i>
            </button>
          </li>
        </ul>
      </nav>
    </div>

    <!-- Update Status Modal -->
    <div class="modal fade" id="updateStatusModal" tabindex="-1" aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Update Status Pendaftaran</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal"></button>
          </div>
          <div class="modal-body">
            <div v-if="selectedItem" class="mb-3">
              <h6 class="text-muted">Detail Pendaftar:</h6>
              <p class="mb-1"><strong>No. Pendaftaran:</strong> {{ selectedItem.noPendaftaran }}</p>
              <p class="mb-1"><strong>Nama:</strong> {{ selectedItem.namaPendaftar }}</p>
              <p class="mb-3"><strong>Status Saat Ini:</strong> 
                <span :class="getStatusBadgeClass(selectedItem.status)" class="badge ms-1">
                  {{ selectedItem.status }}
                </span>
              </p>
            </div>
            <div class="mb-3">
              <label for="newStatus" class="form-label">Status Baru</label>
              <select v-model="newStatus" id="newStatus" class="form-select">
                <option value="">Pilih Status</option>
                <option value="Pending">Pending</option>
                <option value="Update">Perlu Update</option>
                <option value="Selesai">Selesai</option>
                <option value="Ditolak">Ditolak</option>
              </select>
            </div>
            <div class="mb-3">
              <label for="statusNote" class="form-label">Catatan (Opsional)</label>
              <textarea 
                v-model="statusNote" 
                id="statusNote" 
                class="form-control" 
                rows="3"
                placeholder="Tambahkan catatan untuk perubahan status..."
              ></textarea>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Batal</button>
            <button 
              @click="confirmUpdateStatus" 
              type="button" 
              class="btn btn-primary"
              :disabled="!newStatus"
            >
              <i class="fas fa-save me-2"></i>Simpan Perubahan
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading Overlay -->
    <div v-if="isLoading" class="loading-overlay">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UpdateStatusDashboard',
  data() {
    return {
      // Sample data - replace with actual API data
      registrationData: [
        {
          id: 1,
          noPendaftaran: 'LPR202502130005',
          tanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Firda Ro',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Update'
        },
        {
          id: 2,
          noPendaftaran: 'LPR202502130005',
          tanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        },
        {
          id: 3,
          noPendaftaran: 'LPR202502130005',
          tanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        },
        {
          id: 4,
          noPendaftaran: 'LPR202502130005',
          tanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        },
        {
          id: 5,
          noPendaftaran: 'LPR202502130005',
          tanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        }
      ],
      entriesPerPage: 5,
      currentPage: 1,
      searchQuery: '',
      statusFilter: '',
      selectedItem: null,
      newStatus: '',
      statusNote: '',
      isLoading: false
    }
  },
  computed: {
    filteredData() {
      let filtered = this.registrationData;
      
      // Filter by search query
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase();
        filtered = filtered.filter(item => 
          item.namaPendaftar.toLowerCase().includes(query) ||
          item.noPendaftaran.toLowerCase().includes(query) ||
          item.kotaAsal.toLowerCase().includes(query) ||
          item.tujuanTransmigrasi.toLowerCase().includes(query)
        );
      }
      
      // Filter by status
      if (this.statusFilter) {
        filtered = filtered.filter(item => item.status === this.statusFilter);
      }
      
      return filtered;
    },
    totalPages() {
      return Math.ceil(this.filteredData.length / this.entriesPerPage);
    },
    paginatedData() {
      const start = (this.currentPage - 1) * this.entriesPerPage;
      const end = start + this.entriesPerPage;
      return this.filteredData.slice(start, end);
    },
    startEntry() {
      return (this.currentPage - 1) * this.entriesPerPage + 1;
    },
    endEntry() {
      const end = this.currentPage * this.entriesPerPage;
      return end > this.filteredData.length ? this.filteredData.length : end;
    },
    visiblePages() {
      const pages = [];
      const start = Math.max(1, this.currentPage - 2);
      const end = Math.min(this.totalPages, this.currentPage + 2);
      
      for (let i = start; i <= end; i++) {
        pages.push(i);
      }
      
      return pages;
    }
  },
  methods: {
    getStatusBadgeClass(status) {
      const statusClasses = {
        'Selesai': 'bg-success',
        'Update': 'bg-danger',
        'Pending': 'bg-warning text-dark',
        'Ditolak': 'bg-secondary'
      };
      return statusClasses[status] || 'bg-secondary';
    },
    updatePagination() {
      this.currentPage = 1;
    },
    changePage(page) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
      }
    },
    updateStatus(item) {
      this.selectedItem = { ...item };
      this.newStatus = '';
      this.statusNote = '';
    },
    viewDetail(item) {
      // Implement view detail functionality
      console.log('View detail for:', item);
      // You can emit an event or use router to navigate to detail page
      this.$emit('view-detail', item);
    },
    async confirmUpdateStatus() {
      if (!this.newStatus) return;
      
      this.isLoading = true;
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1000));
        
        // Update the status in local data
        const index = this.registrationData.findIndex(item => item.id === this.selectedItem.id);
        if (index !== -1) {
          this.registrationData[index].status = this.newStatus;
        }
        
        // Close modal
        const modal = bootstrap.Modal.getInstance(document.getElementById('updateStatusModal'));
        modal.hide();
        
        // Show success message
        this.showNotification('Status berhasil diperbarui!', 'success');
        
        // Emit event to parent component
        this.$emit('status-updated', {
          item: this.selectedItem,
          newStatus: this.newStatus,
          note: this.statusNote
        });
        
      } catch (error) {
        console.error('Error updating status:', error);
        this.showNotification('Gagal memperbarui status!', 'error');
      } finally {
        this.isLoading = false;
      }
    },
    showNotification(message, type) {
      // Simple notification - you can replace with a toast library
      const alertClass = type === 'success' ? 'alert-success' : 'alert-danger';
      const alertHtml = `
        <div class="alert ${alertClass} alert-dismissible fade show position-fixed" 
             style="top: 20px; right: 20px; z-index: 9999;" role="alert">
          ${message}
          <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
        </div>
      `;
      document.body.insertAdjacentHTML('beforeend', alertHtml);
      
      // Auto remove after 3 seconds
      setTimeout(() => {
        const alert = document.querySelector('.alert');
        if (alert) alert.remove();
      }, 3000);
    }
  },
  watch: {
    searchQuery() {
      this.currentPage = 1;
    },
    statusFilter() {
      this.currentPage = 1;
    }
  }
}
</script>

<style scoped>
.update-status-dashboard {
  background: #f8f9fa;
  padding: 1.5rem;
  border-radius: 10px;
  min-height: 100vh;
}

.dashboard-header {
  background: white;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.filter-section {
  background: white;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.table-responsive {
  background: white;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  overflow: hidden;
}

.table {
  margin-bottom: 0;
}

.table th {
  background: #f1f3f4;
  font-weight: 600;
  border: none;
  padding: 1rem 0.75rem;
}

.table td {
  padding: 1rem 0.75rem;
  border: none;
  border-bottom: 1px solid #dee2e6;
  vertical-align: middle;
}

.table tbody tr:hover {
  background-color: rgba(0,123,255,0.05);
}

.badge {
  font-size: 0.75rem;
  padding: 0.375rem 0.75rem;
}

.btn-group .btn {
  margin-right: 0.25rem;
}

.btn-group .btn:last-child {
  margin-right: 0;
}

.pagination {
  background: white;
  padding: 1rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.pagination-info {
  background: white;
  padding: 1rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}

.page-link {
  border: none;
  color: #495057;
  padding: 0.5rem 0.75rem;
  margin: 0 0.125rem;
  border-radius: 5px;
}

.page-link:hover {
  background-color: #e9ecef;
  color: #495057;
}

.page-item.active .page-link {
  background-color: #007bff;
  color: white;
}

.modal-content {
  border: none;
  border-radius: 10px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.2);
}

.modal-header {
  border-bottom: 1px solid #dee2e6;
  padding: 1.5rem;
}

.modal-body {
  padding: 1.5rem;
}

.modal-footer {
  border-top: 1px solid #dee2e6;
  padding: 1.5rem;
}

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255,255,255,0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

/* Responsive Design */
@media (max-width: 768px) {
  .update-status-dashboard {
    padding: 1rem;
  }
  
  .dashboard-header {
    flex-direction: column;
    gap: 1rem;
  }
  
  .filter-section .row {
    gap: 0.5rem;
  }
  
  .table-responsive {
    font-size: 0.875rem;
  }
  
  .btn-group {
    flex-direction: column;
    gap: 0.25rem;
  }
  
  .btn-group .btn {
    margin-right: 0;
    width: 100%;
  }
}

@media (max-width: 576px) {
  .table th,
  .table td {
    padding: 0.5rem 0.25rem;
  }
  
  .badge {
    font-size: 0.7rem;
    padding: 0.25rem 0.5rem;
  }
  
  .pagination {
    padding: 0.5rem;
  }
  
  .pagination-info {
    padding: 0.5rem;
  }
}
</style>