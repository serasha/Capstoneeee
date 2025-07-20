<template>
  <div class="register-container">
    <div class="container-fluid p-0 position-relative">
      <!-- Back Button (pojok kiri atas) -->
      <button class="btn btn-back position-absolute top-0 start-0 m-4" @click="goToHome">
        <i class="fas fa-arrow-left me-2"></i> Kembali ke Beranda
      </button>

      <div class="row min-vh-100 g-0">
        <!-- Left side - Image -->
        <div class="col-lg-6 d-flex align-items-center justify-content-center bg-light-pink">
          <div class="image-wrapper">
            <img src="/regis.png" alt="Registration Illustration" class="img-fluid registration-image">
          </div>
        </div>

        <!-- Right side - Registration Form -->
        <div class="col-lg-6 d-flex align-items-center justify-content-center" style="background: #ffffff;">
          <div class="form-wrapper w-100">
            <div class="form-container">
              <div class="text-start mb-4">
                <h1 class="form-title mb-2">Daftar</h1>
                <p class="form-subtitle mb-0">Selamat datang di layanan kami!</p>
              </div>

              <form @submit.prevent="handleSubmit">
                <!-- Username -->
                <div class="mb-4">
                  <label class="form-label-custom mb-3">Username</label>
                  <input
                    type="text"
                    class="form-control custom-input"
                    placeholder="Masukkan username Anda"
                    v-model="formData.namaLengkap"
                    :class="{ 'is-invalid': errors.namaLengkap }"
                  />
                  <div v-if="errors.namaLengkap" class="invalid-feedback">
                    {{ errors.namaLengkap ? errors.namaLengkap.replace('Nama lengkap', 'Username') : '' }}
                  </div>
                </div>
                <!-- Pilihan Role -->
                <div class="mb-4">
                  <label class="form-label-custom mb-3">Role</label>
                  <select v-model="formData.role" class="form-control custom-input">
                    <option value="user">User</option>
                    <option value="admin">Admin</option>
                  </select>
                </div>

                <!-- Kata Sandi -->
                <div class="mb-4 position-relative">
                  <label class="form-label-custom mb-3">Kata Sandi</label>
                  <input
                    :type="showPassword ? 'text' : 'password'"
                    class="form-control custom-input pe-5"
                    placeholder="Masukkan kata sandi Anda"
                    v-model="formData.kataSandi"
                    :class="{ 'is-invalid': errors.kataSandi }"
                  />
                  <button
                    type="button"
                    class="password-toggle"
                    @click="togglePassword"
                  >
                    <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'" class="toggle-icon"></i>
                  </button>
                  <div v-if="errors.kataSandi" class="invalid-feedback">
                    {{ errors.kataSandi }}
                  </div>
                </div>

                <!-- Ulangi Kata Sandi -->
                <div class="mb-4 position-relative">
                  <label class="form-label-custom mb-3">Ulangi Kata Sandi</label>
                  <input
                    :type="showConfirmPassword ? 'text' : 'password'"
                    class="form-control custom-input pe-5"
                    placeholder="Masukkan kata sandi Anda"
                    v-model="formData.ulangiKataSandi"
                    :class="{ 'is-invalid': errors.ulangiKataSandi }"
                  />
                  <button
                    type="button"
                    class="password-toggle"
                    @click="toggleConfirmPassword"
                  >
                    <i :class="showConfirmPassword ? 'fas fa-eye-slash' : 'fas fa-eye'" class="toggle-icon"></i>
                  </button>
                  <div v-if="errors.ulangiKataSandi" class="invalid-feedback">
                    {{ errors.ulangiKataSandi }}
                  </div>
                </div>

                <!-- Lupa Kata Sandi Link -->
                <div class="text-end mb-5">
                  <a href="#" class="forgot-password-link">
                    Lupa Kata Sandi?
                  </a>
                </div>

                <!-- Daftar Button -->
                <button
                  type="submit"
                  class="btn btn-register w-100 mb-4"
                  :disabled="isLoading"
                >
                  <span v-if="isLoading" class="spinner-border spinner-border-sm me-2"></span>
                  DAFTAR
                </button>

                <!-- Login Link -->
                <div class="text-center">
                  <span class="login-text">Sudah punya akun? </span>
                  <a href="#" @click.prevent="goToLogin" class="login-link">
                    Masuk di sini
                  </a>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RegisterComponent',
  data() {
    return {
      formData: {
        namaLengkap: '',
        kataSandi: '',
        ulangiKataSandi: '',
        role: 'user',
      },
      errors: {},
      showPassword: false,
      showConfirmPassword: false,
      isLoading: false
    }
  },
  methods: {
    togglePassword() {
      this.showPassword = !this.showPassword;
    },
    toggleConfirmPassword() {
      this.showConfirmPassword = !this.showConfirmPassword;
    },
    validateForm() {
      this.errors = {};
      
      if (!this.formData.namaLengkap.trim()) {
        this.errors.namaLengkap = 'Username wajib diisi';
      }
      
      if (!this.formData.kataSandi) {
        this.errors.kataSandi = 'Kata sandi wajib diisi';
      } else if (this.formData.kataSandi.length < 6) {
        this.errors.kataSandi = 'Kata sandi minimal 6 karakter';
      }
      
      if (!this.formData.ulangiKataSandi) {
        this.errors.ulangiKataSandi = 'Konfirmasi kata sandi wajib diisi';
      } else if (this.formData.kataSandi !== this.formData.ulangiKataSandi) {
        this.errors.ulangiKataSandi = 'Kata sandi tidak cocok';
      }
      
      return Object.keys(this.errors).length === 0;
    },
    async handleSubmit() {
      if (!this.validateForm()) {
        return;
      }
      this.isLoading = true;
      try {
        const response = await fetch('/api/user/register', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            namaLengkap: this.formData.namaLengkap,
            kataSandi: this.formData.kataSandi,
            role: this.formData.role,
          }),
        });
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error || 'Gagal registrasi');
        }
        alert('Registrasi berhasil!');
        this.formData = {
          namaLengkap: '',
          kataSandi: '',
          ulangiKataSandi: '',
          role: 'user',
        };
      } catch (error) {
        alert(error.message || 'Terjadi kesalahan saat registrasi');
      } finally {
        this.isLoading = false;
      }
    },
    goToLogin() {
      this.$router.push('/login');
    },
    goToHome() {
      this.$router.push('/');
    }
  }
}
</script>

<style scoped>
/* Container Styles */
.register-container {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif;
}

.image-wrapper {
  width: 100%;
  max-width: 500px;
  padding: 2rem;
}

.registration-image {
  width: 100%;
  height: auto;
  animation: fadeInUp 0.8s ease-out;
}

/* Form Container */
.form-wrapper {
  padding: 2rem;
  max-width: 500px;
  margin: 0 auto;
}

.form-container {
  background: #FEBFBF;
  padding: 40px 30px;
  border-radius: 25px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 450px;
  margin: 0 auto;
}

/* Typography */
.form-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #000000;
  margin-bottom: 8px;
}

.form-subtitle {
  font-size: 1rem;
  color: #dc2626;
  font-weight: 400;
}

.form-label-custom {
  font-size: 1rem;
  font-weight: 600;
  color: #000000;
  display: block;
}

/* Input Styles */
.custom-input {
  border: none;
  background: #f5f5f5;
  border-radius: 15px;
  padding: 18px 20px;
  font-size: 16px;
  transition: all 0.3s ease;
  box-shadow: none;
  color: #666666;
  width: 100%;
}

.custom-input::placeholder {
  color: #999999;
  font-weight: 400;
}

.custom-input:focus {
  background: #f0f0f0;
  outline: none;
  box-shadow: 0 0 0 3px rgba(220, 38, 38, 0.1);
  border: 2px solid #dc2626;
}

.custom-input.is-invalid {
  border: 2px solid #dc2626;
  background: #fef2f2;
}

/* Password Toggle */
.password-toggle {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
  z-index: 5;
}

.toggle-icon {
  color: #dc2626;
  font-size: 18px;
}

/* Forgot Password Link */
.forgot-password-link {
  color: #dc2626;
  text-decoration: none;
  font-size: 0.9rem;
  font-weight: 500;
}

.forgot-password-link:hover {
  color: #b91c1c;
  text-decoration: underline;
}

/* Register Button */
.btn-register {
  background: #9ca3af;
  border: none;
  border-radius: 15px;
  padding: 18px 24px;
  font-weight: 700;
  font-size: 16px;
  letter-spacing: 0.5px;
  transition: all 0.3s ease;
  color: white;
  text-transform: uppercase;
}

.btn-register:hover:not(:disabled) {
  background: #6b7280;
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(156, 163, 175, 0.3);
  color: white;
}

.btn-register:disabled {
  background: #d1d5db;
  transform: none;
  box-shadow: none;
}

/* Login Link */
.login-text {
  color: #666666;
  font-size: 0.9rem;
}

.login-link {
  color: #000000;
  font-weight: 700;
  text-decoration: none;
  font-size: 0.9rem;
}

.login-link:hover {
  color: #dc2626;
  text-decoration: underline;
}

/* Back Button */
.btn-back {
  background: #fff;
  color: #dc2626;
  border: 1px solid #dc2626;
  border-radius: 8px;
  font-weight: 600;
  padding: 8px 18px;
  font-size: 1rem;
  transition: all 0.2s ease;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  z-index: 10;
}

.btn-back:hover {
  background: #dc2626;
  color: #fff;
}

/* Error Messages */
.invalid-feedback {
  color: #dc2626;
  font-size: 0.875rem;
  margin-top: 8px;
  font-weight: 500;
}

/* Responsive Design */
@media (max-width: 991.98px) {
  .col-lg-6:first-child {
    display: none;
  }
  
  .form-wrapper {
    padding: 1.5rem;
  }
  
  .form-container {
    padding: 30px 25px;
  }
}

@media (max-width: 576px) {
  .form-wrapper {
    padding: 1rem;
  }
  
  .form-container {
    padding: 25px 20px;
    border-radius: 20px;
    margin: 1rem;
  }
  
  .form-title {
    font-size: 2rem;
  }
  
  .custom-input {
    padding: 16px 18px;
    font-size: 15px;
  }
  
  .btn-register {
    padding: 16px 20px;
    font-size: 15px;
  }
}

/* Animations */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.form-container {
  animation: fadeInUp 0.6s ease-out;
}

.registration-image {
  animation: fadeInUp 0.8s ease-out;
}
</style>