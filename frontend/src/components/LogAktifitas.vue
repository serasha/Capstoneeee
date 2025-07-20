<template>
  <div class="log-aktivitas-container">
    <!-- Header Section -->
    <div class="header-section d-flex justify-content-between align-items-center mb-4">
      <div class="user-info d-flex align-items-center">
        <img 
          :src="userInfo.avatar" 
          :alt="userInfo.name" 
          class="user-avatar me-3"
        >
        <div>
          <h5 class="mb-0 user-name">{{ userInfo.name }}</h5>
          <small class="text-muted">{{ userInfo.id }}</small>
        </div>
      </div>
      <button 
        class="btn btn-outline-danger btn-sm"
        @click="downloadData"
      >
        <i class="fas fa-download me-2"></i>
        Download Data
      </button>
    </div>

    <!-- Search and Filter Section -->
    <div class="filter-section mb-4">
      <div class="row">
        <div class="col-md-6">
          <div class="d-flex align-items-center">
            <label class="me-2 text-nowrap">Show</label>
            <select 
              v-model="entriesPerPage" 
              class="form-select form-select-sm me-2"
              style="width: auto;"
              @change="updatePagination"
            >
              <option value="5">5</option>
              <option value="10">10</option>
              <option value="25">25</option>
              <option value="50">50</option>
            </select>
            <span class="text-nowrap">entries</span>
          </div>
        </div>
        <div class="col-md-6">
          <div class="d-flex justify-content-end">
            <label class="me-2 align-self-center text-nowrap">Cari Data</label>
            <input 
              type="text" 
              v-model="searchQuery"
              class="form-control form-control-sm"
              placeholder="Cari..."
              style="width: 200px;"
              @input="filterData"
            >
          </div>
        </div>
      </div>
    </div>

    <!-- Table Section -->
    <div class="table-responsive">
      <table class="table table-hover">
        <thead class="table-light">
          <tr>
            <th style="width: 5%;">No</th>
            <th style="width: 20%;">No. Pendaftaran</th>
            <th style="width: 15%;">Hari/Tanggal</th>
            <th style="width: 60%;">Aktivitas</th>
          </tr>
        </thead>
        <tbody>
          <tr 
            v-for="(item, index) in paginatedData" 
            :key="item.id"
            class="activity-row"
          >
            <td>{{ getRowNumber(index) }}</td>
            <td class="registration-number">{{ item.registrationNumber }}</td>
            <td class="date-column">
              <div class="date-info">
                <div class="day">{{ item.day }}</div>
                <div class="date">{{ item.date }}</div>
              </div>
            </td>
            <td>
              <span 
                class="badge activity-badge"
                :class="getActivityBadgeClass(item.activity)"
              >
                {{ item.activity }}
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- No Data Message -->
    <div v-if="filteredData.length === 0" class="text-center py-4">
      <p class="text-muted">Tidak ada data yang ditemukan</p>
    </div>

    <!-- Pagination Section -->
    <div class="pagination-section d-flex justify-content-between align-items-center mt-4">
      <div class="pagination-info">
        <span class="text-muted">
          Showing {{ getStartRecord() }} to {{ getEndRecord() }} of {{ filteredData.length }} entries
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
  </div>
</template>

<script>
export default {
  name: 'LogAktivitas',
  data() {
    return {
      userInfo: {
        name: 'Hendra Hermawan',
        id: '190001290',
        avatar: 'https://via.placeholder.com/40x40/007bff/ffffff?text=HH'
      },
      searchQuery: '',
      entriesPerPage: 10,
      currentPage: 1,
      activities: [
        {
          id: 1,
          registrationNumber: 'LPR202502130005',
          day: 'Kamis',
          date: '13/Feb/2025',
          activity: 'Verifikasi Berkas'
        },
        {
          id: 2,
          registrationNumber: 'LPR202502130005',
          day: 'Kamis',
          date: '13/Feb/2025',
          activity: 'Verifikasi Berkas'
        },
        {
          id: 3,
          registrationNumber: 'LPR202502130005',
          day: 'Kamis',
          date: '13/Feb/2025',
          activity: 'Verifikasi Berkas'
        },
        {
          id: 4,
          registrationNumber: 'LPR202502130005',
          day: 'Kamis',
          date: '13/Feb/2025',
          activity: 'Verifikasi Berkas'
        },
        {
          id: 5,
          registrationNumber: 'LPR202502130005',
          day: 'Kamis',
          date: '13/Feb/2025',
          activity: 'Verifikasi Berkas'
        },
        {
          id: 6,
          registrationNumber: 'LPR202502140006',
          day: 'Jumat',
          date: '14/Feb/2025',
          activity: 'Pendaftaran Baru'
        },
        {
          id: 7,
          registrationNumber: 'LPR202502140007',
          day: 'Jumat',
          date: '14/Feb/2025',
          activity: 'Review Dokumen'
        },
        {
          id: 8,
          registrationNumber: 'LPR202502150008',
          day: 'Sabtu',
          date: '15/Feb/2025',
          activity: 'Approval Final'
        }
      ],
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
  mounted() {
    this.filteredData = [...this.activities]
  },
  methods: {
    filterData() {
      if (!this.searchQuery.trim()) {
        this.filteredData = [...this.activities]
      } else {
        const query = this.searchQuery.toLowerCase()
        this.filteredData = this.activities.filter(item => 
          item.registrationNumber.toLowerCase().includes(query) ||
          item.day.toLowerCase().includes(query) ||
          item.date.toLowerCase().includes(query) ||
          item.activity.toLowerCase().includes(query)
        )
      }
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
    getActivityBadgeClass(activity) {
      const activityClasses = {
        'Verifikasi Berkas': 'bg-danger',
        'Pendaftaran Baru': 'bg-success',
        'Review Dokumen': 'bg-warning',
        'Approval Final': 'bg-info'
      }
      return activityClasses[activity] || 'bg-secondary'
    },
    downloadData() {
      // Simulate data download
      const csvContent = this.generateCSV()
      const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' })
      const link = document.createElement('a')
      const url = URL.createObjectURL(blob)
      link.setAttribute('href', url)
      link.setAttribute('download', `log_aktivitas_${new Date().toISOString().split('T')[0]}.csv`)
      link.style.visibility = 'hidden'
      document.body.appendChild(link)
      link.click()
      document.body.removeChild(link)
    },
    generateCSV() {
      const headers = ['No', 'No. Pendaftaran', 'Hari/Tanggal', 'Aktivitas']
      const csvRows = [headers.join(',')]
      
      this.activities.forEach((item, index) => {
        const row = [
          index + 1,
          item.registrationNumber,
          `${item.day} ${item.date}`,
          item.activity
        ]
        csvRows.push(row.join(','))
      })
      
      return csvRows.join('\n')
    }
  }
}
</script>

<style scoped>
.log-aktivitas-container {
  background: #fff;
  padding: 24px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  margin: 20px;
}

.header-section {
  border-bottom: 1px solid #e9ecef;
  padding-bottom: 20px;
}

.user-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e9ecef;
}

.user-name {
  color: #333;
  font-weight: 600;
}

.filter-section {
  background: #f8f9fa;
  padding: 16px;
  border-radius: 6px;
  border: 1px solid #e9ecef;
}

.table {
  margin-bottom: 0;
}

.table th {
  background-color: #f8f9fa;
  border-color: #dee2e6;
  font-weight: 600;
  color: #495057;
  padding: 12px 8px;
}

.table td {
  padding: 12px 8px;
  vertical-align: middle;
  border-color: #dee2e6;
}

.activity-row:hover {
  background-color: #f8f9fa;
}

.registration-number {
  font-family: 'Courier New', monospace;
  font-weight: 500;
  color: #dc3545;
}

.date-info {
  text-align: center;
}

.date-info .day {
  font-weight: 500;
  color: #333;
  font-size: 0.9em;
}

.date-info .date {
  color: #666;
  font-size: 0.85em;
}

.activity-badge {
  font-size: 0.85em;
  padding: 6px 12px;
  border-radius: 20px;
  font-weight: 500;
}

.pagination-section {
  border-top: 1px solid #e9ecef;
  padding-top: 20px;
}

.pagination-info {
  color: #6c757d;
  font-size: 0.9em;
}

.pagination .page-link {
  color: #007bff;
  border-color: #dee2e6;
  padding: 6px 12px;
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

.btn-outline-danger {
  border-color: #dc3545;
  color: #dc3545;
}

.btn-outline-danger:hover {
  background-color: #dc3545;
  border-color: #dc3545;
  color: #fff;
}

/* Responsive Design */
@media (max-width: 768px) {
  .log-aktivitas-container {
    margin: 10px;
    padding: 16px;
  }
  
  .header-section {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 16px;
  }
  
  .filter-section .row > div {
    margin-bottom: 16px;
  }
  
  .filter-section .d-flex {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 8px;
  }
  
  .table-responsive {
    font-size: 0.9em;
  }
  
  .pagination-section {
    flex-direction: column;
    gap: 16px;
    text-align: center;
  }
}

@media (max-width: 576px) {
  .user-info {
    flex-direction: column;
    align-items: flex-start !important;
    gap: 8px;
  }
  
  .user-avatar {
    width: 40px;
    height: 40px;
  }
  
  .table td, .table th {
    padding: 8px 4px;
    font-size: 0.85em;
  }
  
  .activity-badge {
    font-size: 0.75em;
    padding: 4px 8px;
  }
}
</style>