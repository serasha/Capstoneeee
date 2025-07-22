<template>
  <div class="admin-verifikasi">
    <!-- Header with Download and Search -->
    <div class="page-header">
      <h1 class="page-title">Verifikasi</h1>
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
            <tr v-for="(item, index) in tableData" :key="item.id">
              <td>{{ getRowNumber(index) }}</td>
              <td class="registration-number">{{ item.noPendaftaran }}</td>
              <td>{{ item.hariTanggal }}</td>
              <td>{{ item.namaPendaftar }}</td>
              <td>{{ item.kotaAsal }}</td>
              <td>{{ item.tujuanTransmigrasi }}</td>
              <td>
                <button 
                  class="btn-verifikasi"
                  @click="verifikasiData(item)"
                >
                  <i class="fas fa-check me-1"></i>Verifikasi
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="pagination-section">
        <div class="pagination-info">
          Showing {{ startEntry }} to {{ endEntry }} of {{ totalEntries }} entries
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
    <VerifikasiModal
      v-if="modalOpen"
      :visible="modalOpen"
      :pendaftaran="selectedPendaftaran"
      @close="modalOpen = false"
      @success="onVerifikasiSuccess"
    />
  </div>
</template>

<script>
import VerifikasiModal from '@/components/VerifikasiModal.vue'

export default {
  name: 'AdminVerifikasi',
  components: { VerifikasiModal },
  data() {
    return {
      searchQuery: '',
      entriesPerPage: 5,
      currentPage: 1,
      totalEntries: 0,
      tableData: [],
      loading: false,
      modalOpen: false,
      selectedPendaftaran: null
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.totalEntries / this.entriesPerPage) || 1;
    },
    startEntry() {
      return (this.currentPage - 1) * this.entriesPerPage + 1;
    },
    endEntry() {
      return Math.min(this.currentPage * this.entriesPerPage, this.totalEntries);
    }
  },
  methods: {
    getRowNumber(index) {
      return (this.currentPage - 1) * this.entriesPerPage + index + 1;
    },
    async fetchData() {
      this.loading = true;
      const params = new URLSearchParams({
        page: this.currentPage,
        page_size: this.entriesPerPage,
        nama: this.searchQuery
      });
      const res = await fetch(`/api/pendaftaran/list?${params.toString()}`, { credentials: 'include' });
      const json = await res.json();
      this.tableData = (json.data || []).map(item => ({
        id: item.id_pendaftaran,
        noPendaftaran: item.id_pendaftaran,
        hariTanggal: item.tanggal_pendaftar ? new Date(item.tanggal_pendaftar).toLocaleDateString() : '-',
        namaPendaftar: item.nama_pendaftar,
        kotaAsal: item.alamat_pendaftar,
        tujuanTransmigrasi: item.jenis_layanan,
        isVerified: item.status_pendaftar === 'verifikasi'
      }));
      this.totalEntries = json.total || 0;
      this.loading = false;
    },
    verifikasiData(item) {
      this.selectedPendaftaran = item
      this.modalOpen = true
    },
    onVerifikasiSuccess() {
      this.fetchData()
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
        this.fetchData();
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
        this.fetchData();
      }
    },
    downloadData() {
      // Simulate download
      const csvContent = this.generateCSV();
      const blob = new Blob([csvContent], { type: 'text/csv' });
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'verifikasi-data.csv';
      a.click();
      window.URL.revokeObjectURL(url);
    },
    generateCSV() {
      const headers = ['No', 'No. Pendaftaran', 'Hari/Tanggal', 'Nama Pendaftar', 'Kota Asal', 'Tujuan Transmigrasi', 'Status'];
      const rows = this.tableData.map((item, index) => [
        index + 1,
        item.noPendaftaran,
        item.hariTanggal,
        item.namaPendaftar,
        item.kotaAsal,
        item.tujuanTransmigrasi,
        item.isVerified ? 'Terverifikasi' : 'Belum Verifikasi'
      ]);
      return [headers, ...rows].map(row => row.join(',')).join('\n');
    }
  },
  watch: {
    searchQuery() {
      this.currentPage = 1;
      this.fetchData();
    },
    entriesPerPage() {
      this.currentPage = 1;
      this.fetchData();
    }
  },
  mounted() {
    this.fetchData();
  }
}
</script>

<style scoped>
.admin-verifikasi {
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

.btn-verifikasi {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #dc3545;
  color: white;
}

.btn-verifikasi:hover:not(:disabled) {
  background: #c82333;
}

.btn-verifikasi:disabled {
  background: #28a745;
  cursor: not-allowed;
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

.modal-backdrop {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
}
.modal-content {
  background: #fff;
  border-radius: 12px;
  padding: 2rem 2.5rem;
  min-width: 320px;
  max-width: 90vw;
  box-shadow: 0 8px 32px rgba(0,0,0,0.18);
}
.modal-actions {
  margin-top: 1.5rem;
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}
</style>