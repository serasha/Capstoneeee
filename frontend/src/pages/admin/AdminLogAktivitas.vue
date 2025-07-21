<template>
  <div class="admin-log-aktivitas">
    <!-- Header with Download and Search -->
    <div class="page-header">
      <h1 class="page-title">Log Aktivitas</h1>
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
              <th>Tanggal</th>
              <th>Admin</th>
              <th>Aksi</th>
              <th>Target</th>
              <th>Deskripsi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in tableData" :key="item.id">
              <td>{{ getRowNumber(index) }}</td>
              <td>{{ item.tanggal }}</td>
              <td>{{ item.admin }}</td>
              <td>{{ item.aksi }}</td>
              <td>{{ item.target }}</td>
              <td>{{ item.deskripsi }}</td>
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
  </div>
</template>

<script>
export default {
  name: 'AdminLogAktivitas',
  data() {
    return {
      searchQuery: '',
      entriesPerPage: 5,
      currentPage: 1,
      totalEntries: 0,
      tableData: [],
      loading: false
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
        page_size: this.entriesPerPage
      });
      const res = await fetch(`/api/log-aktivitas?${params.toString()}`, { credentials: 'include' });
      const json = await res.json();
      this.tableData = (json.data || []).map(item => ({
        id: item.id,
        tanggal: item.created_at ? new Date(item.created_at).toLocaleString() : '-',
        admin: item.admin && item.admin.username ? item.admin.username : (item.admin_id || '-'),
        aksi: item.aksi,
        target: item.target,
        deskripsi: item.deskripsi
      }));
      this.totalEntries = json.total || 0;
      this.loading = false;
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
      const csvContent = this.generateCSV();
      const blob = new Blob([csvContent], { type: 'text/csv' });
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'log-aktivitas.csv';
      a.click();
      window.URL.revokeObjectURL(url);
    },
    generateCSV() {
      const headers = ['No', 'Tanggal', 'Admin', 'Aksi', 'Target', 'Deskripsi'];
      const rows = this.tableData.map((item, index) => [
        index + 1,
        item.tanggal,
        item.admin,
        item.aksi,
        item.target,
        item.deskripsi
      ]);
      return [headers, ...rows].map(row => row.join(',')).join('\n');
    }
  },
  watch: {
    searchQuery() {
      this.currentPage = 1;
      // Optional: implement search on backend if needed
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
.admin-log-aktivitas {
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

.activity-badge {
  background: #dc3545;
  color: white;
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 500;
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
</style>