<template>
  <div class="login-container">
    <div class="container-fluid vh-100">
      <div class="row h-100 align-items-center justify-content-center">
        <div class="col-12 col-md-6 col-lg-4">
          <div class="login-card">
            <div class="card border-0 shadow-lg">
              <div class="card-body p-5">
                <div class="text-center mb-4">
                  <h2 class="login-title">Login</h2>
                  <p class="welcome-text">Selamat datang di layanan kami!</p>
                </div>
                
                <form @submit.prevent="handleLogin">
                  <div class="mb-4">
                    <label for="username" class="form-label fw-semibold">Username</label>
                    <input 
                      type="text" 
                      class="form-control form-control-lg custom-input" 
                      id="username"
                      v-model="loginForm.username"
                      placeholder="Masukkan username Anda"
                      :class="{ 'is-invalid': errors.username }"
                      required
                    >
                    <div v-if="errors.username" class="invalid-feedback">
                      {{ errors.username }}
                    </div>
                  </div>
                  
                  <div class="mb-4">
                    <label for="password" class="form-label fw-semibold">Kata Sandi</label>
                    <div class="input-group">
                      <input 
                        :type="showPassword ? 'text' : 'password'" 
                        class="form-control form-control-lg custom-input" 
                        id="password"
                        v-model="loginForm.password"
                        placeholder="Masukkan kata sandi Anda"
                        :class="{ 'is-invalid': errors.password }"
                        required
                      >
                      <button 
                        class="btn btn-outline-secondary password-toggle" 
                        type="button"
                        @click="togglePassword"
                      >
                        <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                      </button>
                      <div v-if="errors.password" class="invalid-feedback">
                        {{ errors.password }}
                      </div>
                    </div>
                    <div class="text-end mt-2">
                      <a href="#" class="forgot-link" @click.prevent="forgotPassword">
                        Lupa Kata Sandi?
                      </a>
                    </div>
                  </div>
                  
                  <div class="d-grid mb-4">
                    <button 
                      type="submit" 
                      class="btn btn-primary btn-lg login-btn"
                      :disabled="isLoading"
                    >
                      <span v-if="isLoading" class="spinner-border spinner-border-sm me-2"></span>
                      {{ isLoading ? 'Memproses...' : 'LOGIN' }}
                    </button>
                  </div>
                  
                  <div class="text-center">
                    <span class="text-muted">Belum punya akun? </span>
                    <a href="#" class="register-link fw-semibold" @click.prevent="goToRegister">
                      Daftar di sini
                    </a>
                  </div>
                </form>
              </div>
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
      loginForm: {
        username: '',
        password: ''
      },
      showPassword: false,
      isLoading: false,
      errors: {}
    }
  },
  methods: {
    togglePassword() {
      this.showPassword = !this.showPassword;
    },
    
    validateForm() {
      this.errors = {};
      
      if (!this.loginForm.username.trim()) {
        this.errors.username = 'Username wajib diisi';
      } else if (this.loginForm.username.length < 3) {
        this.errors.username = 'Username minimal 3 karakter';
      }
      
      if (!this.loginForm.password.trim()) {
        this.errors.password = 'Kata sandi wajib diisi';
      } else if (this.loginForm.password.length < 6) {
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
        // Simulasi API call
        await new Promise(resolve => setTimeout(resolve, 1500));
        
        // Emit event ke parent component atau redirect
        this.$emit('login-success', {
          username: this.loginForm.username,
          timestamp: new Date()
        });
        
        // Reset form
        this.loginForm = {
          username: '',
          password: ''
        };
        
        // Show success message
        this.$emit('show-message', {
          type: 'success',
          message: 'Login berhasil! Selamat datang.'
        });
        
      } catch (error) {
        this.$emit('show-message', {
          type: 'error',
          message: 'Login gagal. Silakan periksa kembali username dan kata sandi Anda.'
        });
      } finally {
        this.isLoading = false;
      }
    },
    
    forgotPassword() {
      this.$emit('forgot-password');
    },
    
    goToRegister() {
      this.$emit('go-to-register');
    }
  }
}
</script>

<style scoped>
.login-container {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

.login-card {
  max-width: 450px;
  margin: 0 auto;
}

.card {
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}

.login-title {
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 0.5rem;
}

.welcome-text {
  color: #e74c3c;
  font-size: 1.1rem;
  margin-bottom: 0;
}

.form-label {
  color: #2c3e50;
  margin-bottom: 0.75rem;
}

.custom-input {
  border: 2px solid #e9ecef;
  border-radius: 12px;
  padding: 0.75rem 1rem;
  font-size: 1rem;
  transition: all 0.3s ease;
  background: rgba(255, 255, 255, 0.9);
}

.custom-input:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 0.2rem rgba(102, 126, 234, 0.25);
  background: #fff;
}

.custom-input.is-invalid {
  border-color: #dc3545;
}

.password-toggle {
  border: 2px solid #e9ecef;
  border-left: none;
  border-radius: 0 12px 12px 0;
  background: rgba(255, 255, 255, 0.9);
  transition: all 0.3s ease;
}

.password-toggle:hover {
  background: #f8f9fa;
  border-color: #667eea;
}

.login-btn {
  background: linear-gradient(45deg, #667eea, #764ba2);
  border: none;
  border-radius: 12px;
  font-weight: 600;
  letter-spacing: 1px;
  padding: 0.75rem 2rem;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
  background: linear-gradient(45deg, #5a6fd8, #6a42a6);
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.forgot-link {
  color: #e74c3c;
  text-decoration: none;
  font-size: 0.9rem;
  transition: color 0.3s ease;
}

.forgot-link:hover {
  color: #c0392b;
  text-decoration: underline;
}

.register-link {
  color: #667eea;
  text-decoration: none;
  transition: color 0.3s ease;
}

.register-link:hover {
  color: #5a6fd8;
  text-decoration: underline;
}

.invalid-feedback {
  font-size: 0.875rem;
  margin-top: 0.25rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .login-container {
    padding: 1rem;
  }
  
  .card-body {
    padding: 2rem !important;
  }
  
  .login-title {
    font-size: 2rem;
  }
  
  .welcome-text {
    font-size: 1rem;
  }
}

@media (max-width: 576px) {
  .card-body {
    padding: 1.5rem !important;
  }
  
  .login-title {
    font-size: 1.75rem;
  }
  
  .custom-input {
    font-size: 0.9rem;
  }
}

/* Animation */
.login-card {
  animation: slideUp 0.6s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>