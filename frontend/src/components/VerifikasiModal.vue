<template>
  <div v-if="visible" class="modal-backdrop">
    <div class="modal-content">
      <h3 class="modal-title">Verifikasi Pendaftaran</h3>
      <div v-if="pendaftaran" class="modal-body">
        <div class="modal-row">
          <span class="modal-label">Nama:</span>
          <span class="modal-value">{{ pendaftaran.namaPendaftar }}</span>
        </div>
        <div class="modal-row">
          <span class="modal-label">No. Pendaftaran:</span>
          <span class="modal-value">{{ pendaftaran.noPendaftaran }}</span>
        </div>
        <div class="modal-row">
          <span class="modal-label">Status:</span>
          <select v-model="status" class="modal-select">
            <option value="verifikasi">Verifikasi</option>
            <option value="ditolak">Tolak</option>
          </select>
        </div>
      </div>
      <div class="modal-actions">
        <button class="btn btn-primary" @click="submit" :disabled="loading">
          <span v-if="loading" class="spinner"></span>
          Simpan
        </button>
        <button class="btn btn-secondary" @click="$emit('close')">Batal</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'VerifikasiModal',
  props: {
    visible: Boolean,
    pendaftaran: Object
  },
  data() {
    return {
      status: this.pendaftaran && this.pendaftaran.isVerified ? 'verifikasi' : 'pending',
      loading: false
    }
  },
  watch: {
    pendaftaran(newVal) {
      this.status = newVal && newVal.isVerified ? 'verifikasi' : 'pending'
    }
  },
  methods: {
    async submit() {
      this.loading = true
      await fetch(`/api/pendaftaran/${this.pendaftaran.noPendaftaran}/verifikasi`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({ status_pendaftar: this.status })
      })
      this.loading = false
      this.$emit('success')
      this.$emit('close')
    }
  }
}
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(220,53,69,0.12); /* merah SPBE transparan */
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  transition: background 0.2s;
}
.modal-content {
  background: #fff;
  border-radius: 18px;
  padding: 2.2rem 2.5rem 1.5rem 2.5rem;
  min-width: 340px;
  max-width: 95vw;
  box-shadow: 0 8px 32px rgba(220,53,69,0.18), 0 2px 8px rgba(0,0,0,0.08);
  animation: modalFadeIn 0.18s;
}
@keyframes modalFadeIn {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}
.modal-title {
  font-size: 1.35rem;
  font-weight: 700;
  color: #dc3545;
  margin-bottom: 1.2rem;
  text-align: left;
}
.modal-body {
  margin-bottom: 1.5rem;
}
.modal-row {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}
.modal-label {
  min-width: 120px;
  font-weight: 600;
  color: #333;
  font-size: 1rem;
}
.modal-value {
  color: #6c757d;
  font-size: 1rem;
}
.modal-select {
  border: 2px solid #dc3545;
  border-radius: 8px;
  padding: 0.4rem 1rem;
  font-size: 1rem;
  color: #dc3545;
  background: #fff5f5;
  margin-left: 0.5rem;
  outline: none;
  transition: border 0.2s;
}
.modal-select:focus {
  border-color: #c82333;
}
.modal-actions {
  margin-top: 1.2rem;
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}
.btn {
  font-size: 1rem;
  font-weight: 600;
  border-radius: 8px;
  padding: 0.6rem 1.5rem;
  border: none;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}
.btn-primary {
  background: linear-gradient(90deg, #dc3545 60%, #c82333 100%);
  color: #fff;
  box-shadow: 0 2px 8px rgba(220,53,69,0.08);
}
.btn-primary:hover {
  background: #c82333;
}
.btn-secondary {
  background: #f1f1f1;
  color: #333;
}
.btn-secondary:hover {
  background: #e9ecef;
}
.spinner {
  display: inline-block;
  width: 1.1em;
  height: 1.1em;
  border: 2px solid #fff;
  border-top: 2px solid #dc3545;
  border-radius: 50%;
  margin-right: 0.5em;
  animation: spin 0.7s linear infinite;
  vertical-align: middle;
}
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
@media (max-width: 600px) {
  .modal-content {
    padding: 1.2rem 0.7rem 1rem 0.7rem;
    min-width: 0;
  }
  .modal-label {
    min-width: 80px;
    font-size: 0.95rem;
  }
}
</style> 