<template>
  <div class="login-container">
    <div class="container-fluid vh-100 position-relative">
      <!-- Back Button (pojok kiri atas) -->
      <button class="btn btn-back position-absolute top-0 start-0 m-4" @click="goToHome">
        <i class="fas fa-arrow-left me-2"></i> Kembali ke Beranda
      </button>
      <div class="row h-100">
        <!-- Left Side - Image -->
        <div class="col-lg-6 d-none d-lg-flex align-items-center justify-content-center bg-light">
          <div class="image-wrapper">
            <img src="/regis.png" alt="Login Illustration" class="img-fluid login-image">
          </div>
        </div>

        <!-- Right Side - Login Form -->
        <div class="col-lg-6 d-flex align-items-center justify-content-center">
          <div class="login-form-wrapper w-100">
            <!-- Hapus tombol back di sini, sudah dipindah ke pojok kiri atas -->
            <div class="login-form-container">
              <div class="text-left mb-4">
                <h2 class="login-title">Login</h2>
                <p class="welcome-text">Selamat datang di layanan kami!</p>
              </div>

              <form @submit.prevent="handleLogin">
                <div class="form-group mb-3">
                  <label for="username" class="form-label">Username</label>
                  <input
                    type="text"
                    class="form-control custom-input"
                    id="username"
                    v-model="loginData.username"
                    placeholder="Masukkan username Anda"
                    :class="{ 'is-invalid': errors.username }"
                  />
                  <div v-if="errors.username" class="invalid-feedback">
                    {{ errors.username }}
                  </div>
                </div>

                <div class="form-group mb-3">
                  <label for="password" class="form-label">Kata Sandi</label>
                  <div class="password-input-wrapper position-relative">
                    <input
                      :type="showPassword ? 'text' : 'password'"
                      class="form-control custom-input"
                      id="password"
                      v-model="loginData.password"
                      placeholder="Masukkan kata sandi Anda"
                      :class="{ 'is-invalid': errors.password }"
                    />
                    <button
                      type="button"
                      class="password-toggle-btn"
                      @click="togglePassword"
                    >
                      <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'" style="color: #dc3545;"></i>
                    </button>
                  </div>
                  <div v-if="errors.password" class="invalid-feedback">
                    {{ errors.password }}
                  </div>
                </div>

                <div class="text-end mb-4">
                  <a href="#" class="forgot-password-link" @click.prevent="forgotPassword">
                    Lupa Kata Sandi?
                  </a>
                </div>

                <button
                  type="submit"
                  class="btn btn-login w-100 mb-4"
                  :disabled="isLoading"
                >
                  <span v-if="isLoading" class="spinner-border spinner-border-sm me-2" role="status"></span>
                  {{ isLoading ? 'Memproses...' : 'LOGIN' }}
                </button>

                <div class="text-center">
                  <span class="register-text">Belum punya akun? </span>
                  <a href="#" class="register-link" @click.prevent="goToRegister">
                    <strong>Daftar</strong>
                  </a>
                  <span class="register-text"> di sini</span>
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
  name: 'LoginComponent',
  data() {
    return {
      loginData: {
        username: '',
        password: ''
      },
      errors: {},
      showPassword: false,
      isLoading: false
    }
  },
  methods: {
    validateForm() {
      this.errors = {};
      
      if (!this.loginData.username.trim()) {
        this.errors.username = 'Username harus diisi';
      }
      
      if (!this.loginData.password.trim()) {
        this.errors.password = 'Kata sandi harus diisi';
      } else if (this.loginData.password.length < 6) {
        this.errors.password = 'Kata sandi minimal 6 karakter';
      }
      
      return Object.keys(this.errors).length === 0;
    },
    
    async handleLogin() {
      if (!this.validateForm()) {
        return;
      }
      this.isLoading = true;
      try {
        const response = await fetch('/api/user/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            namaLengkap: this.loginData.username,
            kataSandi: this.loginData.password
          }),
          credentials: 'include'
        });
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error || 'Gagal login');
        }
        const data = await response.json();
        this.$emit('login-success', data.user);
        this.loginData = { username: '', password: '' };
        // Redirect ke dashboard/home
        this.$router.push('/');
        // Trigger reload Navbar
        this.$root.$emit && this.$root.$emit('auth-changed');
      } catch (error) {
        alert(error.message || 'Login gagal');
      } finally {
        this.isLoading = false;
      }
    },
    
    togglePassword() {
      this.showPassword = !this.showPassword;
    },
    
    forgotPassword() {
      // Emit event atau navigate ke halaman forgot password
      this.$emit('forgot-password');
      console.log('Forgot password clicked');
    },
    
    goToRegister() {
      this.$router.push('/daftar');
    },
    
    goToHome() {
      this.$router.push('/');
    }
  }
}
</script>

<style scoped>
/* Container Styles */
.login-container {
  background: #ffffff;
  min-height: 100vh;
}

/* Image Wrapper */
.image-wrapper {
  width: 100%;
  max-width: 500px;
  padding: 2rem;
}

.login-image {
  width: 100%;
  height: auto;
  animation: fadeInUp 0.8s ease-out;
}

/* Right Side Form */
.login-form-wrapper {
  max-width: 450px;
  padding: 20px;
}

.login-form-container {
  background:  #FEBFBF;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 192, 203, 0.3);
  border-radius: 20px;
  padding: 40px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.login-title {
  color: #333;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 8px;
}

.welcome-text {
  color: #dc3545;
  font-size: 0.9rem;
  margin-bottom: 0;
}

/* Form Styles */
.form-label {
  color: #333;
  font-weight: 500;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

.custom-input {
  background: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.1);
  border-radius: 12px;
  padding: 12px 16px;
  color: #333;
  font-size: 0.9rem;
  transition: all 0.3s ease;
}

.custom-input::placeholder {
  color: rgba(0, 0, 0, 0.4);
}

.custom-input:focus {
  background: rgba(255, 255, 255, 0.9);
  border-color: rgba(220, 53, 69, 0.5);
  box-shadow: 0 0 0 0.2rem rgba(220, 53, 69, 0.1);
  color: #333;
}

.password-input-wrapper {
  position: relative;
}

.password-toggle-btn {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #dc3545;
  cursor: pointer;
  padding: 4px;
  transition: color 0.3s ease;
}

.password-toggle-btn:hover {
  color: #c82333;
}

.forgot-password-link {
  color: #dc3545;
  text-decoration: none;
  font-size: 0.85rem;
  transition: color 0.3s ease;
}

.forgot-password-link:hover {
  color: #c82333;
  text-decoration: underline;
}

.btn-login {
  background: #d1d5db;
  border: none;
  border-radius: 12px;
  padding: 12px;
  color: #666;
  font-weight: 600;
  font-size: 0.95rem;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  transition: all 0.3s ease;
  box-shadow: none;
}

.btn-login:hover:not(:disabled) {
  background: #9ca3af;
  color: #333;
  transform: translateY(-1px);
}

.btn-login:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.register-text {
  color: #666;
  font-size: 0.9rem;
}

.register-link {
  color: #333;
  text-decoration: none;
  transition: color 0.3s ease;
}

.register-link:hover {
  color: #000;
  text-decoration: underline;
}

.btn-back {
  background: #fff;
  color: #dc3545;
  border: 1px solid #dc3545;
  border-radius: 8px;
  font-weight: 600;
  padding: 8px 18px;
  font-size: 1rem;
  transition: background 0.2s, color 0.2s;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  z-index: 10;
}

.btn-back:hover {
  background: #dc3545;
  color: #fff;
}

/* Animations */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

/* Responsive Design */
@media (max-width: 991.98px) {
  .login-form-container {
    background:  #FEBFBF;
    color: #fff;
  }
  
  .login-title {
    color: #333;
  }
  
  .welcome-text {
    color: #666;
  }
  
  .form-label {
    color: #333;
  }
  
  .custom-input {
    background: #fff;
    border: 1px solid #ddd;
    color: #333;
  }
  
  .custom-input::placeholder {
    color: #999;
  }
  
  .custom-input:focus {
    background: #fff;
    border-color: #e74c3c;
    box-shadow: 0 0 0 0.2rem rgba(231, 76, 60, 0.1);
    color: #333;
  }
  
  .password-toggle-btn {
    color: #666;
  }
  
  .password-toggle-btn:hover {
    color: #333;
  }
  
  .forgot-password-link {
    color: #666;
  }
  
  .forgot-password-link:hover {
    color: #333;
  }
  
  .register-text {
    color: #666;
  }
  
  .register-link {
    color: #e74c3c;
  }
  
  .register-link:hover {
    color: #c0392b;
  }
}

@media (max-width: 576px) {
  .login-form-container {
    padding: 20px;
    margin: 20px;
  }
  
  .login-title {
    font-size: 1.75rem;
  }
}
</style>