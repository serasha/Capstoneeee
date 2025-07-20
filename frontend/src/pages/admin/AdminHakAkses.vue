<template>
  <div class="admin-hak-akses">
    <div class="page-header">
      <h1 class="page-title">Hak Akses</h1>
    </div>

    <!-- Form Section -->
    <div class="form-section">
      <div class="form-container">
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="idUser" class="form-label">ID User</label>
            <input
              type="text"
              id="idUser"
              v-model="formData.idUser"
              class="form-control"
              placeholder="Masukkan ID User"
              required
            />
          </div>

          <div class="form-group">
            <label for="namaUser" class="form-label">Nama User</label>
            <input
              type="text"
              id="namaUser"
              v-model="formData.namaUser"
              class="form-control"
              placeholder="Masukkan Nama User"
              required
            />
          </div>

          <div class="form-group">
            <label for="pilihAkses" class="form-label">Pilih Akses</label>
            <select
              id="pilihAkses"
              v-model="formData.pilihAkses"
              class="form-control"
              required
            >
              <option value="">Pilih Level Akses</option>
              <option value="admin">Admin</option>
              <option value="operator">Operator</option>
              <option value="user">User</option>
            </select>
          </div>

          <button type="submit" class="btn-submit" :disabled="isSubmitting">
            <span v-if="isSubmitting" class="spinner"></span>
            SIMPAN AKSES
          </button>
        </form>
      </div>
    </div>

    <!-- Users List -->
    <div class="users-list">
      <h2 class="section-title">Daftar Pengguna</h2>
      <div class="table-container">
        <table class="data-table">
          <thead>
            <tr>
              <th>No</th>
              <th>ID User</th>
              <th>Nama User</th>
              <th>Level Akses</th>
              <th>Status</th>
              <th>Aksi</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(user, index) in users" :key="user.id">
              <td>{{ index + 1 }}</td>
              <td>{{ user.idUser }}</td>
              <td>{{ user.namaUser }}</td>
              <td>
                <span class="badge" :class="getBadgeClass(user.levelAkses)">
                  {{ user.levelAkses }}
                </span>
              </td>
              <td>
                <span class="badge" :class="user.status === 'Aktif' ? 'badge-success' : 'badge-secondary'">
                  {{ user.status }}
                </span>
              </td>
              <td>
                <button class="btn-edit" @click="editUser(user)">
                  <i class="fas fa-edit"></i>
                </button>
                <button class="btn-delete" @click="deleteUser(user.id)">
                  <i class="fas fa-trash"></i>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminHakAkses',
  data() {
    return {
      formData: {
        idUser: '',
        namaUser: '',
        pilihAkses: ''
      },
      isSubmitting: false,
      users: [
        {
          id: 1,
          idUser: 'USR001',
          namaUser: 'Ahmad Rizki',
          levelAkses: 'Admin',
          status: 'Aktif'
        },
        {
          id: 2,
          idUser: 'USR002',
          namaUser: 'Siti Nurhaliza',
          levelAkses: 'Operator',
          status: 'Aktif'
        },
        {
          id: 3,
          idUser: 'USR003',
          namaUser: 'Budi Santoso',
          levelAkses: 'User',
          status: 'Nonaktif'
        }
      ]
    }
  },
  methods: {
    async submitForm() {
      this.isSubmitting = true;
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1000));
        
        // Add new user to list
        const newUser = {
          id: this.users.length + 1,
          idUser: this.formData.idUser,
          namaUser: this.formData.namaUser,
          levelAkses: this.formData.pilihAkses,
          status: 'Aktif'
        };
        
        this.users.push(newUser);
        
        // Reset form
        this.formData = {
          idUser: '',
          namaUser: '',
          pilihAkses: ''
        };
        
        alert('Hak akses berhasil disimpan!');
      } catch (error) {
        alert('Terjadi kesalahan saat menyimpan data');
      } finally {
        this.isSubmitting = false;
      }
    },
    
    editUser(user) {
      this.formData = {
        idUser: user.idUser,
        namaUser: user.namaUser,
        pilihAkses: user.levelAkses.toLowerCase()
      };
    },
    
    deleteUser(userId) {
      if (confirm('Apakah Anda yakin ingin menghapus user ini?')) {
        this.users = this.users.filter(user => user.id !== userId);
        alert('User berhasil dihapus!');
      }
    },
    
    getBadgeClass(level) {
      const classes = {
        'Admin': 'badge-danger',
        'Operator': 'badge-warning',
        'User': 'badge-info'
      };
      return classes[level] || 'badge-secondary';
    }
  }
}
</script>

<style scoped>
.admin-hak-akses {
  max-width: 1200px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 2rem;
}

.page-title {
  font-size: 2rem;
  font-weight: 700;
  color: #333;
}

.form-section {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
}

.form-container {
  max-width: 600px;
  margin: 0 auto;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #333;
}

.form-control {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-control:focus {
  outline: none;
  border-color: #dc3545;
  box-shadow: 0 0 0 0.2rem rgba(220, 53, 69, 0.1);
}

.btn-submit {
  width: 100%;
  padding: 1rem;
  background: #dc3545;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn-submit:hover:not(:disabled) {
  background: #c82333;
}

.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.spinner {
  width: 20px;
  height: 20px;
  border: 2px solid transparent;
  border-top: 2px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.users-list {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.section-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #333;
  margin-bottom: 1.5rem;
}

.table-container {
  overflow-x: auto;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
}

.data-table th,
.data-table td {
  padding: 1rem;
  text-align: left;
  border-bottom: 1px solid #e9ecef;
}

.data-table th {
  background: #f8f9fa;
  font-weight: 600;
  color: #333;
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

.badge-danger {
  background: #dc3545;
  color: white;
}

.badge-warning {
  background: #ffc107;
  color: #000;
}

.badge-info {
  background: #17a2b8;
  color: white;
}

.badge-success {
  background: #28a745;
  color: white;
}

.badge-secondary {
  background: #6c757d;
  color: white;
}

.btn-edit,
.btn-delete {
  padding: 0.5rem;
  margin: 0 0.25rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-edit {
  background: #17a2b8;
  color: white;
}

.btn-edit:hover {
  background: #138496;
}

.btn-delete {
  background: #dc3545;
  color: white;
}

.btn-delete:hover {
  background: #c82333;
}

/* Responsive */
@media (max-width: 768px) {
  .form-section,
  .users-list {
    padding: 1rem;
  }
  
  .data-table {
    font-size: 0.875rem;
  }
  
  .data-table th,
  .data-table td {
    padding: 0.5rem;
  }
}
</style>