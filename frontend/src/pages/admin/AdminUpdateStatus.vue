<template>
  <div class="admin-update-status">
    <!-- Header with Download and Search -->
    <div class="page-header">
      <h1 class="page-title">Update Status</h1>
      <div class="header-actions">
        <button class="btn-download" @click="downloadData">
          <i class="fas fa-download"></i>
          Download Data
        </button>
        <div class="search-container">
          <label for="search" class="search-label">Cari Data</label>
          <input
            type="text"
            id="search"
            v-model="searchQuery"
            class="search-input"
            placeholder="Cari..."
          />
        </div>
      </div>
    </div>

    <!-- Table Controls -->
    <div class="table-controls">
      <div class="entries-control">
        <label>Show</label>
        <select v-model="entriesPerPage" class="entries-select">
          <option value="5">5</option>
          <option value="10">10</option>
          <option value="25">25</option>
          <option value="50">50</option>
        </select>
        <span>entries</span>
      </div>
    </div>

    <!-- Data Table -->
    <div class="table-section">
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>No</th>
              <th>No. Pendaftaran</th>
              <th>Hari/Tanggal</th>
              <th>Nama Pendaftar</th>
              <th>Kota Asal</th>
              <th>Tujuan Transmigrasi</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in paginatedData" :key="item.id">
              <td>{{ getRowNumber(index) }}</td>
              <td class="registration-number">{{ item.noPendaftaran }}</td>
              <td>{{ item.hariTanggal }}</td>
              <td>{{ item.namaPendaftar }}</td>
              <td>{{ item.kotaAsal }}</td>
              <td>{{ item.tujuanTransmigrasi }}</td>
              <td>
                <button 
                  class="btn-action"
                  :class="getActionButtonClass(item.status)"
                  @click="updateStatus(item)"
                >
                  {{ item.status }}
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
            &lt;&lt;
          </button>
          <button 
            class="pagination-btn" 
            @click="previousPage" 
            :disabled="currentPage === 1"
          >
            &lt;
          </button>
          <span class="page-number">{{ currentPage }}</span>
          <button 
            class="pagination-btn" 
            @click="nextPage" 
            :disabled="currentPage === totalPages"
          >
            &gt;
          </button>
          <button 
            class="pagination-btn" 
            @click="nextPage" 
            :disabled="currentPage === totalPages"
          >
            &gt;&gt;
          </button>
        </div>
      </div>
    </div>

    <!-- Update Status Modal -->
    <div v-if="showModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>Update Status</h3>
          <button class="modal-close" @click="closeModal">&times;</button>
        </div>
        <div class="modal-body">
          <p><strong>{{ selectedItem?.namaPendaftar }}</strong></p>
          <p>No. Pendaftaran: {{ selectedItem?.noPendaftaran }}</p>
          <div class="form-group">
            <label>Status Baru:</label>
            <select v-model="newStatus" class="form-control">
              <option value="Update">Update</option>
              <option value="Selesai">Selesai</option>
              <option value="Pending">Pending</option>
            </select>
          </div>
        </div>
        <div class="modal-footer">
          <button class="btn-cancel" @click="closeModal">Batal</button>
          <button class="btn-save" @click="saveStatus">Simpan</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminUpdateStatus',
  data() {
    return {
      searchQuery: '',
      entriesPerPage: 5,
      currentPage: 1,
      showModal: false,
      selectedItem: null,
      newStatus: '',
      statusData: [
        {
          id: 1,
          noPendaftaran: 'LPR202502130005',
          hariTanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Firda Ro',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Update'
        },
        {
          id: 2,
          noPendaftaran: 'LPR202502130005',
          hariTanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        },
        {
          id: 3,
          noPendaftaran: 'LPR202502130005',
          hariTanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        },
        {
          id: 4,
          noPendaftaran: 'LPR202502130005',
          hariTanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        },
        {
          id: 5,
          noPendaftaran: 'LPR202502130005',
          hariTanggal: 'Kamis 13/Feb/2025',
          namaPendaftar: 'Wahyudi Nurul',
          kotaAsal: 'Yogyakarta',
          tujuanTransmigrasi: 'Sulawesi',
          status: 'Selesai'
        }
      ]
    }
  },
  computed: {
    filteredData() {
      if (!this.searchQuery) {
        return this.statusData;
      }
      return this.statusData.filter(item =>
        item.noPendaftaran.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
        item.namaPendaftar.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
        item.kotaAsal.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
        item.tujuanTransmigrasi.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
    paginatedData() {
      const start = (this.currentPage - 1) * this.entriesPerPage;
      const end = start + parseInt(this.entriesPerPage);
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
    getRowNumber(index) {
      return (this.currentPage - 1) * this.entriesPerPage + index + 1;
    },
    getActionButtonClass(status) {
      return {
        'btn-update': status === 'Update',
        'btn-selesai': status === 'Selesai',
        'btn-pending': status === 'Pending'
      };
    },
    updateStatus(item) {
      this.selectedItem = item;
      this.newStatus = item.status;
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
      this.selectedItem = null;
      this.newStatus = '';
    },
    saveStatus() {
      if (this.selectedItem && this.newStatus) {
        // Update the status in the data
        const index = this.statusData.findIndex(item => item.id === this.selectedItem.id);
        if (index !== -1) {
          this.statusData[index].status = this.newStatus;
        }
        alert('Status berhasil diupdate!');
        this.closeModal();
      }
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
    },
    downloadData() {
      // Simulate download
      const csvContent = this.generateCSV();
      const blob = new Blob([csvContent], { type: 'text/csv' });
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'update-status.csv';
      a.click();
      window.URL.revokeObjectURL(url);
    },
    generateCSV() {
      const headers = ['No', 'No. Pendaftaran', 'Hari/Tanggal', 'Nama Pendaftar', 'Kota Asal', 'Tujuan Transmigrasi', 'Status'];
      const rows = this.filteredData.map((item, index) => [
        index + 1,
        item.noPendaftaran,
        item.hariTanggal,
        item.namaPendaftar,
        item.kotaAsal,
        item.tujuanTransmigrasi,
        item.status
      ]);
      return [headers, ...rows].map(row => row.join(',')).join('\n');
    }
  },
  watch: {
    searchQuery() {
      this.currentPage = 1;
    },
    entriesPerPage() {
      this.currentPage = 1;
    }
  }
}
</script>

<style scoped>
.admin-update-status {
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 2rem;
  gap: 2rem;
}

.page-title {
  font-size: 2rem;
  font-weight: 700;
  color: #333;
  margin: 0;
}

.header-actions {
  display: flex;
  align-items: flex-end;
  gap: 2rem;
}

.btn-download {
  background: #dc3545;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
  display: flex;
  align-items: center;
  gap: 0.5rem;
  text-decoration: underline;
}

.btn-download:hover {
  background: #c82333;
}

.search-container {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.search-label {
  font-weight: 600;
  color: #333;
  font-size: 0.9rem;
}

.search-input {
  padding: 0.75rem 1rem;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  width: 250px;
  font-size: 1rem;
}

.search-input:focus {
  outline: none;
  border-color: #dc3545;
}

.table-controls {
  background: white;
  padding: 1.5rem;
  border-radius: 12px 12px 0 0;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.entries-control {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  color: #333;
}

.entries-select {
  padding: 0.25rem 0.5rem;
  border: 1px solid #ced4da;
  border-radius: 4px;
  background: white;
}

.table-section {
  background: white;
  border-radius: 0 0 12px 12px;
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

.data-table tbody tr:nth-child(even) {
  background: #fff5f5;
}

.data-table tbody tr:hover {
  background: #f8f9fa;
}

.registration-number {
  color: #dc3545;
  font-weight: 600;
  font-family: monospace;
}

.btn-action {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-update {
  background: #dc3545;
  color: white;
}

.btn-update:hover {
  background: #c82333;
}

.btn-selesai {
  background: #28a745;
  color: white;
}

.btn-selesai:hover {
  background: #218838;
}

.btn-pending {
  background: #ffc107;
  color: #000;
}

.btn-pending:hover {
  background: #e0a800;
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
  width: 90%;
  max-width: 500px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 1px solid #e9ecef;
}

.modal-header h3 {
  margin: 0;
  color: #333;
}

.modal-close {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #6c757d;
}

.modal-close:hover {
  color: #333;
}

.modal-body {
  padding: 1.5rem;
}

.form-group {
  margin-top: 1rem;
}

.form-group label {
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

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  padding: 1.5rem;
  border-top: 1px solid #e9ecef;
}

.btn-cancel {
  padding: 0.75rem 1.5rem;
  background: #6c757d;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.btn-cancel:hover {
  background: #5a6268;
}

.btn-save {
  padding: 0.75rem 1.5rem;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.btn-save:hover {
  background: #c82333;
}

/* Responsive */
@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: stretch;
  }
  
  .header-actions {
    flex-direction: column;
    gap: 1rem;
  }
  
  .search-input {
    width: 100%;
  }
  
  .pagination-section {
    flex-direction: column;
    gap: 1rem;
  }
  
  .data-table {
    font-size: 0.875rem;
  }
  
  .data-table th,
  .data-table td {
    padding: 0.75rem 0.5rem;
  }
}
</style>