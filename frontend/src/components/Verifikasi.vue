<template>
  <div class="verifikasi-container">
    <!-- User Profile Header -->
    <div class="user-profile-header">
      <div class="container-fluid">
        <div class="row justify-content-end">
          <div class="col-auto">
            <div class="user-profile-card">
              <div class="user-info">
                <div class="user-details">
                  <h5 class="user-name">{{ userData.name }}</h5>
                  <p class="user-id">{{ userData.id }}</p>
                </div>
                <div class="user-avatar">
                  <img 
                    v-if="userData.avatar"
                    :src="userData.avatar" 
                    :alt="userData.name"
                    class="avatar-img"
                    @error="showDefaultAvatar = true"
                    v-show="!showDefaultAvatar"
                  >
                  <span v-if="showDefaultAvatar || !userData.avatar" class="avatar-img avatar-default">
                    <!-- SVG icon profile -->
                    <svg width="60" height="60" viewBox="0 0 60 60" fill="none">
                      <circle cx="30" cy="30" r="30" fill="#e9ecef"/>
                      <circle cx="30" cy="24" r="12" fill="#adb5bd"/>
                      <ellipse cx="30" cy="44" rx="16" ry="10" fill="#adb5bd"/>
                    </svg>
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="verifikasi-content">
      <div class="container-fluid">
        <div class="content-card">
          <!-- Action Buttons -->
          <div class="action-header">
            <div class="row align-items-center">
              <div class="col-md-6">
                <button 
                  class="btn btn-download"
                  @click="downloadData"
                  :disabled="loading"
                >
                  <ArrowDownTrayIcon style="width:20px;height:20px;" class="me-2" />
                  Download Data
                </button>
              </div>
              <div class="col-md-6">
                <div class="search-section">
                  <label for="search-input" class="search-label">Cari Data</label>
                  <div class="search-input-group">
                    <input
                      id="search-input"
                      type="text"
                      class="form-control search-input"
                      placeholder="Masukkan kata kunci..."
                      v-model="searchQuery"
                      @input="handleSearch"
                    >
                    <button class="btn btn-search" @click="performSearch">
                      <MagnifyingGlassIcon style="width:20px;height:20px;" />
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Table Controls -->
          <div class="table-controls">
            <div class="row align-items-center">
              <div class="col-md-6">
                <div class="entries-control">
                  <label for="entries-select">Show</label>
                  <select 
                    id="entries-select"
                    class="form-select entries-select"
                    v-model="entriesPerPage"
                    @change="changeEntriesPerPage"
                  >
                    <option value="5">5</option>
                    <option value="10">10</option>
                    <option value="25">25</option>
                    <option value="50">50</option>
                  </select>
                  <span>entries</span>
                </div>
              </div>
            </div>
          </div>

          <!-- Data Table -->
          <div class="table-container">
            <div class="table-responsive">
              <table class="table table-striped data-table">
                <thead>
                  <tr>
                    <th scope="col" class="table-header">No</th>
                    <th scope="col" class="table-header">No. Pendaftaran</th>
                    <th scope="col" class="table-header">Hari/Tanggal</th>
                    <th scope="col" class="table-header">Nama Pendaftar</th>
                    <th scope="col" class="table-header">Kota Asal</th>
                    <th scope="col" class="table-header">Tujuan Transmigrasi</th>
                    <th scope="col" class="table-header">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr 
                    v-for="(item, index) in paginatedData" 
                    :key="item.id"
                    class="table-row"
                    :class="{ 'table-row-even': index % 2 === 0 }"
                  >
                    <td class="table-cell">{{ getRowNumber(index) }}</td>
                    <td class="table-cell">
                      <span class="registration-number">{{ item.registrationNumber }}</span>
                    </td>
                    <td class="table-cell">
                      <span class="date-text">{{ formatDate(item.date) }}</span>
                    </td>
                    <td class="table-cell">
                      <span class="name-text" :style="{ color: item.nameColor }">
                        {{ item.name }}
                      </span>
                    </td>
                    <td class="table-cell">
                      <span class="city-text">{{ item.originCity }}</span>
                    </td>
                    <td class="table-cell">
                      <span class="destination-text" :style="{ color: item.destinationColor }">
                        {{ item.destination }}
                      </span>
                    </td>
                    <td class="table-cell">
                      <button 
                        class="btn btn-verify"
                        @click="verifyData(item)"
                        :disabled="item.isVerified"
                      >
                        <CheckIcon style="width:20px;height:20px;" class="me-1" />
                        {{ item.isVerified ? 'Terverifikasi' : 'Verifikasi' }}
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <!-- Pagination -->
          <div class="pagination-container">
            <div class="row align-items-center">
              <div class="col-md-6">
                <div class="pagination-info">
                  Showing {{ startEntry }} to {{ endEntry }} of {{ totalEntries }} entries
                </div>
              </div>
              <div class="col-md-6">
                <nav aria-label="Table pagination">
                  <ul class="pagination pagination-custom">
                    <li class="page-item" :class="{ disabled: currentPage === 1 }">
                      <button 
                        class="page-link"
                        @click="previousPage"
                        :disabled="currentPage === 1"
                      >
                        <ChevronDoubleLeftIcon style="width:20px;height:20px;" />
                      </button>
                    </li>
                    <li class="page-item" :class="{ disabled: currentPage === 1 }">
                      <button 
                        class="page-link"
                        @click="previousPage"
                        :disabled="currentPage === 1"
                      >
                        <ChevronLeftIcon style="width:20px;height:20px;" />
                      </button>
                    </li>
                    <li 
                      class="page-item"
                      v-for="page in visiblePages"
                      :key="page"
                      :class="{ active: page === currentPage }"
                    >
                      <button 
                        class="page-link"
                        @click="goToPage(page)"
                      >
                        {{ page }}
                      </button>
                    </li>
                    <li class="page-item" :class="{ disabled: currentPage === totalPages }">
                      <button 
                        class="page-link"
                        @click="nextPage"
                        :disabled="currentPage === totalPages"
                      >
                        <ChevronRightIcon style="width:20px;height:20px;" />
                      </button>
                    </li>
                    <li class="page-item" :class="{ disabled: currentPage === totalPages }">
                      <button 
                        class="page-link"
                        @click="goToPage(totalPages)"
                        :disabled="currentPage === totalPages"
                      >
                        <ChevronDoubleRightIcon style="width:20px;height:20px;" />
                      </button>
                    </li>
                  </ul>
                </nav>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Loading Overlay -->
    <div class="loading-overlay" v-if="loading">
      <div class="loading-spinner">
        <div class="spinner-border text-primary" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
        <p class="loading-text">Memuat data...</p>
      </div>
    </div>
  </div>
</template>

<script>
import { ArrowDownTrayIcon, MagnifyingGlassIcon, CheckCircleIcon, ChevronDoubleLeftIcon, ChevronLeftIcon, ChevronRightIcon, ChevronDoubleRightIcon, CheckIcon } from '@heroicons/vue/24/solid'
export default {
  name: 'VerifikasiData',
  components: { ArrowDownTrayIcon, MagnifyingGlassIcon, CheckCircleIcon, ChevronDoubleLeftIcon, ChevronLeftIcon, ChevronRightIcon, ChevronDoubleRightIcon, CheckIcon },
  data() {
    return {
      loading: false,
      searchQuery: '',
      entriesPerPage: 5,
      currentPage: 1,
      userData: {
        name: 'Hendra Hermawan',
        id: '190000290',
        avatar: '' // Kosongkan agar default icon muncul
      },
      showDefaultAvatar: false,
      tableData: [
        {
          id: 1,
          registrationNumber: 'LPR20250213005',
          date: '2025-02-13',
          name: 'Firda Ro',
          nameColor: '#dc3545',
          originCity: 'Yogyakarta',
          destination: 'Sulawesi',
          destinationColor: '#dc3545',
          isVerified: true
        },
        {
          id: 2,
          registrationNumber: 'LPR20250213005',
          date: '2025-02-13',
          name: 'Wahyudi Nurul',
          nameColor: '#dc3545',
          originCity: 'Yogyakarta',
          destination: 'Sulawesi',
          destinationColor: '#dc3545',
          isVerified: true
        },
        {
          id: 3,
          registrationNumber: 'LPR20250213005',
          date: '2025-02-13',
          name: 'Wahyudi Nurul',
          nameColor: '#dc3545',
          originCity: 'Yogyakarta',
          destination: 'Sulawesi',
          destinationColor: '#dc3545',
          isVerified: true
        },
        {
          id: 4,
          registrationNumber: 'LPR20250213005',
          date: '2025-02-13',
          name: 'Wahyudi Nurul',
          nameColor: '#dc3545',
          originCity: 'Yogyakarta',
          destination: 'Sulawesi',
          destinationColor: '#dc3545',
          isVerified: true
        },
        {
          id: 5,
          registrationNumber: 'LPR20250213005',
          date: '2025-02-13',
          name: 'Wahyudi Nurul',
          nameColor: '#dc3545',
          originCity: 'Yogyakarta',
          destination: 'Sulawesi',
          destinationColor: '#dc3545',
          isVerified: true
        }
      ],
      filteredData: []
    }
  },
  computed: {
    paginatedData() {
      const start = (this.currentPage - 1) * this.entriesPerPage;
      const end = start + parseInt(this.entriesPerPage);
      return this.filteredData.slice(start, end);
    },
    totalPages() {
      return Math.ceil(this.filteredData.length / this.entriesPerPage);
    },
    totalEntries() {
      return this.filteredData.length;
    },
    startEntry() {
      return (this.currentPage - 1) * this.entriesPerPage + 1;
    },
    endEntry() {
      const end = this.currentPage * this.entriesPerPage;
      return Math.min(end, this.totalEntries);
    },
    visiblePages() {
      const pages = [];
      const total = this.totalPages;
      const current = this.currentPage;
      
      if (total <= 7) {
        for (let i = 1; i <= total; i++) {
          pages.push(i);
        }
      } else {
        if (current <= 4) {
          for (let i = 1; i <= 5; i++) {
            pages.push(i);
          }
          pages.push('...');
          pages.push(total);
        } else if (current >= total - 3) {
          pages.push(1);
          pages.push('...');
          for (let i = total - 4; i <= total; i++) {
            pages.push(i);
          }
        } else {
          pages.push(1);
          pages.push('...');
          for (let i = current - 1; i <= current + 1; i++) {
            pages.push(i);
          }
          pages.push('...');
          pages.push(total);
        }
      }
      
      return pages;
    }
  },
  methods: {
    handleImageError(event) {
      this.showDefaultAvatar = true;
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      const days = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu'];
      const months = ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun', 'Jul', 'Agu', 'Sep', 'Okt', 'Nov', 'Des'];
      
      const dayName = days[date.getDay()];
      const day = date.getDate();
      const month = months[date.getMonth()];
      const year = date.getFullYear();
      
      return `${dayName} ${day}/${month}/${year}`;
    },
    getRowNumber(index) {
      return (this.currentPage - 1) * this.entriesPerPage + index + 1;
    },
    handleSearch() {
      this.performSearch();
    },
    performSearch() {
      const query = this.searchQuery.toLowerCase();
      if (query === '') {
        this.filteredData = [...this.tableData];
      } else {
        this.filteredData = this.tableData.filter(item => 
          item.name.toLowerCase().includes(query) ||
          item.registrationNumber.toLowerCase().includes(query) ||
          item.originCity.toLowerCase().includes(query) ||
          item.destination.toLowerCase().includes(query)
        );
      }
      this.currentPage = 1;
    },
    changeEntriesPerPage() {
      this.currentPage = 1;
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
    goToPage(page) {
      if (page !== '...' && page >= 1 && page <= this.totalPages) {
        this.currentPage = page;
      }
    },
    async verifyData(item) {
      this.loading = true;
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1000));
        
        // Update item status
        item.isVerified = true;
        
        this.$toast.success('Data berhasil diverifikasi');
      } catch (error) {
        this.$toast.error('Gagal memverifikasi data');
      } finally {
        this.loading = false;
      }
    },
    async downloadData() {
      this.loading = true;
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // Create CSV data
        const csvData = this.generateCSV();
        this.downloadCSV(csvData, 'data-verifikasi.csv');
        
        this.$toast.success('Data berhasil didownload');
      } catch (error) {
        this.$toast.error('Gagal mendownload data');
      } finally {
        this.loading = false;
      }
    },
    generateCSV() {
      const headers = ['No', 'No. Pendaftaran', 'Hari/Tanggal', 'Nama Pendaftar', 'Kota Asal', 'Tujuan Transmigrasi', 'Status'];
      const rows = this.filteredData.map((item, index) => [
        index + 1,
        item.registrationNumber,
        this.formatDate(item.date),
        item.name,
        item.originCity,
        item.destination,
        item.isVerified ? 'Terverifikasi' : 'Belum Verifikasi'
      ]);
      
      return [headers, ...rows].map(row => row.join(',')).join('\n');
    },
    downloadCSV(csvData, filename) {
      const blob = new Blob([csvData], { type: 'text/csv;charset=utf-8;' });
      const link = document.createElement('a');
      const url = URL.createObjectURL(blob);
      link.setAttribute('href', url);
      link.setAttribute('download', filename);
      link.style.visibility = 'hidden';
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    }
  },
  mounted() {
    this.filteredData = [...this.tableData];
    if (!this.userData.avatar) {
      this.showDefaultAvatar = true;
    }
  }
}
</script>

<style scoped>
.verifikasi-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
}

.user-profile-header {
  background: white;
  border-bottom: 1px solid #dee2e6;
  padding: 1.5rem 0;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.user-profile-card {
  background: white;
  border-radius: 15px;
  padding: 1rem 1.5rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border: 1px solid #e9ecef;
  transition: all 0.3s ease;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-details {
  text-align: right;
}

.user-name {
  font-size: 1.1rem;
  font-weight: 600;
  color: #343a40;
  margin: 0;
}

.user-id {
  font-size: 0.9rem;
  color: #6c757d;
  margin: 0;
}

.avatar-img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #dc3545;
  background: #e9ecef;
  display: inline-block;
  vertical-align: middle;
}

.avatar-default {
  display: flex;
  align-items: center;
  justify-content: center;
  background: #e9ecef;
  border: 3px solid #dc3545;
}

.user-avatar {
  position: relative;
  width: 60px;
  height: 60px;
}

.user-avatar img,
.user-avatar .avatar-default {
  width: 60px;
  height: 60px;
}

.user-avatar .avatar-default svg {
  width: 60px;
  height: 60px;
  display: block;
}

.user-avatar .avatar-default {
  padding: 0;
}

.verifikasi-content {
  padding: 2rem 0;
}

.content-card {
  background: white;
  border-radius: 15px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.action-header {
  padding: 1.5rem;
  border-bottom: 1px solid #e9ecef;
  background: #f8f9fa;
}

.btn-download {
  background: linear-gradient(135deg, #dc3545 0%, #c82333 100%);
  border: none;
  color: white;
  padding: 0.75rem 1.5rem;
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.btn-download:hover {
  background: linear-gradient(135deg, #c82333 0%, #a71e2a 100%);
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(220, 53, 69, 0.3);
}

.search-section {
  text-align: right;
}

.search-label {
  font-weight: 500;
  color: #495057;
  margin-bottom: 0.5rem;
  display: block;
}

.search-input-group {
  display: flex;
  gap: 0.5rem;
}

.search-input {
  border: 2px solid #e9ecef;
  border-radius: 8px;
  padding: 0.75rem 1rem;
  transition: border-color 0.3s ease;
}

.search-input:focus {
  border-color: #dc3545;
  box-shadow: 0 0 0 0.2rem rgba(220, 53, 69, 0.1);
}

.btn-search {
  background: #dc3545;
  border: none;
  color: white;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.btn-search:hover {
  background: #c82333;
}

.table-controls {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid #e9ecef;
}

.entries-control {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9rem;
  color: #495057;
}

.entries-select {
  width: auto;
  display: inline-block;
  padding: 0.375rem 0.75rem;
  border: 1px solid #ced4da;
  border-radius: 4px;
}

.table-container {
  overflow-x: auto;
}

.data-table {
  margin: 0;
  border-collapse: separate;
  border-spacing: 0;
}

.table-header {
  background: #f8f9fa;
  border-bottom: 2px solid #dee2e6;
  padding: 1rem;
  font-weight: 600;
  color: #495057;
  text-align: center;
  white-space: nowrap;
}

.table-row {
  transition: background-color 0.2s ease;
}

.table-row:hover {
  background-color: #f8f9fa;
}

.table-row-even {
  background-color: rgba(248, 249, 250, 0.5);
}

.table-cell {
  padding: 1rem;
  text-align: center;
  vertical-align: middle;
  border-bottom: 1px solid #e9ecef;
}

.registration-number {
  font-family: 'Courier New', monospace;
  font-weight: 600;
  color: #495057;
}

.date-text {
  font-size: 0.9rem;
  color: #6c757d;
}

.name-text {
  font-weight: 500;
}

.city-text {
  color: #495057;
}

.destination-text {
  font-weight: 500;
}

.btn-verify {
  background: #dc3545;
  border: none;
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  font-size: 0.875rem;
  font-weight: 500;
  transition: all 0.3s ease;
}

.btn-verify:hover:not(:disabled) {
  background: #c82333;
  transform: translateY(-1px);
}

.btn-verify:disabled {
  background: #28a745;
  cursor: not-allowed;
}

.pagination-container {
  padding: 1.5rem;
  border-top: 1px solid #e9ecef;
  background: #f8f9fa;
}

.pagination-info {
  color: #6c757d;
  font-size: 0.9rem;
}

.pagination-custom {
  justify-content: flex-end;
  margin: 0;
}

.pagination-custom .page-link {
  border: 1px solid #dee2e6;
  color: #495057;
  padding: 0.5rem 0.75rem;
  margin: 0 2px;
  border-radius: 6px;
  transition: all 0.3s ease;
}

.pagination-custom .page-item.active .page-link {
  background: #dc3545;
  border-color: #dc3545;
  color: white;
}

.pagination-custom .page-link:hover:not(.disabled) {
  background: #e9ecef;
  border-color: #adb5bd;
}

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.loading-spinner {
  background: white;
  border-radius: 15px;
  padding: 2rem;
  text-align: center;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.loading-text {
  margin-top: 1rem;
  color: #495057;
  font-weight: 500;
}

/* Responsive Design */
@media (max-width: 991.98px) {
  .action-header .row {
    gap: 1rem;
  }
  
  .search-section {
    text-align: left;
    margin-top: 1rem;
  }
  
  .pagination-custom {
    justify-content: center;
    margin-top: 1rem;
  }
}

@media (max-width: 767.98px) {
  .verifikasi-content {
    padding: 1rem 0;
  }
  
  .user-profile-header {
    padding: 1rem 0;
  }
  
  .user-profile-card {
    padding: 0.75rem 1rem;
  }
  
  .avatar-img {
    width: 50px;
    height: 50px;
  }
  
  .table-container {
    font-size: 0.875rem;
  }
  
  .table-header,
  .table-cell {
    padding: 0.75rem 0.5rem;
  }
  
  .btn-verify {
    font-size: 0.75rem;
    padding: 0.375rem 0.75rem;
  }
  
  .search-input-group {
    flex-direction: column;
  }
  
  .btn-search {
    margin-top: 0.5rem;
  }
}

@media (max-width: 575.98px) {
  .table-header,
  .table-cell {
    padding: 0.5rem 0.25rem;
    font-size: 0.8rem;
  }
  
  .registration-number {
    font-size: 0.75rem;
  }
  
  .pagination-custom .page-link {
    padding: 0.375rem 0.5rem;
    font-size: 0.875rem;
  }
}

/* Animation */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.content-card {
  animation: fadeIn 0.6s ease-out;
}

/* Print Styles */
@media print {
  .action-header,
  .pagination-container,
  .user-profile-header {
    display: none;
  }
  
  .data-table {
    font-size: 0.8rem;
  }
  
  .btn-verify {
    display: none;
  }
}
</style>