<template>
  <div class="register-container">
    <div class="container-fluid vh-100">
      <div class="row h-100 align-items-center">
        <!-- Left Side - Illustration -->
        <div class="col-12 col-lg-6 d-none d-lg-flex justify-content-center align-items-center">
          <div class="illustration-wrapper">
            <div class="illustration">
              <!-- Phone mockup -->
              <div class="phone-mockup">
                <div class="phone-screen">
                  <div class="app-header">
                    <div class="app-icon"></div>
                    <div class="app-lines">
                      <div class="line line-1"></div>
                      <div class="line line-2"></div>
                    </div>
                  </div>
                  <div class="form-preview">
                    <div class="input-line"></div>
                    <div class="input-line"></div>
                    <div class="input-line short"></div>
                    <div class="dots">••••••••</div>
                  </div>
                  <div class="signup-btn">SIGN UP</div>
                </div>
              </div>
              
              <!-- Character -->
              <div class="character">
                <div class="character-head"></div>
                <div class="character-body"></div>
                <div class="character-arm"></div>
                <div class="character-legs"></div>
              </div>
              
              <!-- Security elements -->
              <div class="security-lock">
                <div class="lock-body"></div>
                <div class="lock-shackle"></div>
                <div class="lock-keyhole"></div>
              </div>
              
              <!-- Floating elements -->
              <div class="floating-elements">
                <div class="float-element element-1"></div>
                <div class="float-element element-2"></div>
                <div class="float-element element-3"></div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Right Side - Registration Form -->
        <div class="col-12 col-lg-6">
          <div class="register-form-wrapper">
            <div class="card border-0 shadow-lg">
              <div class="card-body p-5">
                <div class="text-center mb-4">
                  <h2 class="register-title">Daftar</h2>
                  <p class="welcome-text">Selamat datang di layanan kami!</p>
                </div>
                
                <form @submit.prevent="handleRegister">
                  <div class="mb-4">
                    <label for="fullName" class="form-label fw-semibold">Nama Lengkap</label>
                    <input 
                      type="text" 
                      class="form-control form-control-lg custom-input" 
                      id="fullName"
                      v-model="registerForm.fullName"
                      placeholder="Masukkan nama lengkap Anda"
                      :class="{ 'is-invalid': errors.fullName }"
                      required
                    >
                    <div v-if="errors.fullName" class="invalid-feedback">
                      {{ errors.fullName }}
                    </div>
                  </div>
                  
                  <div class="mb-4">
                    <label for="username" class="form-label fw-semibold">Username</label>
                    <input 
                      type="text" 
                      class="form-control form-control-lg custom-input" 
                      id="username"
                      v-model="registerForm.username"
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
                        v-model="registerForm.password"
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
                  </div>
                  
                  <div class="mb-4">
                    <label for="confirmPassword" class="form-label fw-semibold">Ulangi Kata Sandi</label>
                    <div class="input-group">
                      <input 
                        :type="showConfirmPassword ? 'text' : 'password'" 
                        class="form-control form-control-lg custom-input" 
                        id="confirmPassword"
                        v-model="registerForm.confirmPassword"
                        placeholder="Masukkan kata sandi Anda"
                        :class="{ 'is-invalid': errors.confirmPassword }"
                        required
                      >
                      <button 
                        class="btn btn-outline-secondary password-toggle" 
                        type="button"
                        @click="toggleConfirmPassword"
                      >
                        <i :class="showConfirmPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                      </button>
                      <div v-if="errors.confirmPassword" class="invalid-feedback">
                        {{ errors.confirmPassword }}
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
                      class="btn btn-primary btn-lg register-btn"
                      :disabled="isLoading"
                    >
                      <span v-if="isLoading" class="spinner-border spinner-border-sm me-2"></span>
                      {{ isLoading ? 'Memproses...' : 'DAFTAR' }}
                    </button>
                  </div>
                  
                  <div class="text-center">
                    <span class="text-muted">Sudah punya akun? </span>
                    <a href="#" class="login-link fw-semibold" @click.prevent="goToLogin">
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
  </div>
</template>

<script>
export default {
  name: 'RegisterComponent',
  data() {
    return {
      registerForm: {
        fullName: '',
        username: '',
        password: '',
        confirmPassword: ''
      },
      showPassword: false,
      showConfirmPassword: false,
      isLoading: false,
      errors: {}
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
      
      // Validate full name
      if (!this.registerForm.fullName.trim()) {
        this.errors.fullName = 'Nama lengkap wajib diisi';
      } else if (this.registerForm.fullName.length < 2) {
        this.errors.fullName = 'Nama lengkap minimal 2 karakter';
      }
      
      // Validate username
      if (!this.registerForm.username.trim()) {
        this.errors.username = 'Username wajib diisi';
      } else if (this.registerForm.username.length < 3) {
        this.errors.username = 'Username minimal 3 karakter';
      } else if (!/^[a-zA-Z0-9_]+$/.test(this.registerForm.username)) {
        this.errors.username = 'Username hanya boleh berisi huruf, angka, dan underscore';
      }
      
      // Validate password
      if (!this.registerForm.password.trim()) {
        this.errors.password = 'Kata sandi wajib diisi';
      } else if (this.registerForm.password.length < 6) {
        this.errors.password = 'Kata sandi minimal 6 karakter';
      } else if (!/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/.test(this.registerForm.password)) {
        this.errors.password = 'Kata sandi harus mengandung huruf besar, huruf kecil, dan angka';
      }
      
      // Validate confirm password
      if (!this.registerForm.confirmPassword.trim()) {
        this.errors.confirmPassword = 'Konfirmasi kata sandi wajib diisi';
      } else if (this.registerForm.password !== this.registerForm.confirmPassword) {
        this.errors.confirmPassword = 'Konfirmasi kata sandi tidak cocok';
      }
      
      return Object.keys(this.errors).length === 0;
    },
    
    async handleRegister() {
      if (!this.validateForm()) {
        return;
      }
      
      this.isLoading = true;
      
      try {
        // Simulasi API call
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // Check if username already exists (simulation)
        const existingUsers = ['admin', 'user', 'test'];
        if (existingUsers.includes(this.registerForm.username.toLowerCase())) {
          throw new Error('Username sudah digunakan');
        }
        
        // Emit event ke parent component
        this.$emit('register-success', {
          fullName: this.registerForm.fullName,
          username: this.registerForm.username,
          timestamp: new Date()
        });
        
        // Reset form
        this.registerForm = {
          fullName: '',
          username: '',
          password: '',
          confirmPassword: ''
        };
        
        // Show success message
        this.$emit('show-message', {
          type: 'success',
          message: 'Pendaftaran berhasil! Silakan login dengan akun Anda.'
        });
        
      } catch (error) {
        this.$emit('show-message', {
          type: 'error',
          message: error.message || 'Pendaftaran gagal. Silakan coba lagi.'
        });
      } finally {
        this.isLoading = false;
      }
    },
    
    forgotPassword() {
      this.$emit('forgot-password');
    },
    
    goToLogin() {
      this.$emit('go-to-login');
    }
  }
}
</script>

<style scoped>
.register-container {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

.illustration-wrapper {
  position: relative;
  width: 100%;
  max-width: 500px;
  height: 400px;
}

.illustration {
  position: relative;
  width: 100%;
  height: 100%;
}

/* Phone Mockup */
.phone-mockup {
  position: absolute;
  top: 20px;
  right: 50px;
  width: 180px;
  height: 300px;
  background: #2c3e50;
  border-radius: 25px;
  padding: 15px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
  z-index: 3;
}

.phone-screen {
  width: 100%;
  height: 100%;
  background: white;
  border-radius: 15px;
  padding: 20px 15px;
  position: relative;
}

.app-header {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.app-icon {
  width: 30px;
  height: 30px;
  background: #e74c3c;
  border-radius: 8px;
  margin-right: 10px;
}

.app-lines {
  flex: 1;
}

.line {
  height: 4px;
  background: #ecf0f1;
  border-radius: 2px;
  margin-bottom: 5px;
}

.line-1 {
  width: 80%;
}

.line-2 {
  width: 60%;
}

.form-preview {
  margin-bottom: 20px;
}

.input-line {
  height: 8px;
  background: #ecf0f1;
  border-radius: 4px;
  margin-bottom: 15px;
}

.input-line.short {
  width: 70%;
}

.dots {
  font-size: 12px;
  color: #bdc3c7;
  margin-bottom: 20px;
}

.signup-btn {
  background: #e74c3c;
  color: white;
  text-align: center;
  padding: 8px;
  border-radius: 6px;
  font-size: 10px;
  font-weight: bold;
}

/* Character */
.character {
  position: absolute;
  bottom: 0;
  left: 50px;
  z-index: 2;
}

.character-head {
  width: 60px;
  height: 60px;
  background: #f4d1ae;
  border-radius: 50%;
  position: relative;
  margin-bottom: -10px;
}

.character-body {
  width: 80px;
  height: 120px;
  background: #e74c3c;
  border-radius: 40px 40px 20px 20px;
}

.character-arm {
  position: absolute;
  width: 15px;
  height: 80px;
  background: #f4d1ae;
  border-radius: 10px;
  right: -5px;
  top: 70px;
  transform: rotate(-20deg);
}

.character-legs {
  position: absolute;
  bottom: -60px;
  left: 10px;
  width: 60px;
  height: 60px;
  background: #2c3e50;
  border-radius: 30px 30px 0 0;
}

/* Security Lock */
.security-lock {
  position: absolute;
  top: 50px;
  left: 100px;
  width: 40px;
  height: 50px;
  z-index: 4;
}

.lock-body {
  width: 40px;
  height: 30px;
  background: #e74c3c;
  border-radius: 5px;
  position: relative;
  margin-top: 15px;
}

.lock-shackle {
  width: 25px;
  height: 20px;
  border: 4px solid #e74c3c;
  border-bottom: none;
  border-radius: 15px 15px 0 0;
  position: absolute;
  top: -5px;
  left: 50%;
  transform: translateX(-50%);
}

.lock-keyhole {
  width: 6px;
  height: 6px;
  background: white;
  border-radius: 50%;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

/* Floating Elements */
.floating-elements {
  position: absolute;
  width: 100%;
  height: 100%;
}

.float-element {
  position: absolute;
  border-radius: 10px;
  opacity: 0.3;
}

.element-1 {
  width: 20px;
  height: 20px;
  background: #3498db;
  top: 30px;
  left: 30px;
  animation: float 3s ease-in-out infinite;
}

.element-2 {
  width: 15px;
  height: 15px;
  background: #f39c12;
  top: 150px;
  right: 80px;
  animation: float 3s ease-in-out infinite 1s;
}

.element-3 {
  width: 25px;
  height: 25px;
  background: #2ecc71;
  bottom: 80px;
  left: 20px;
  animation: float 3s ease-in-out infinite 2s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-20px);
  }
}

/* Form Styles */
.register-form-wrapper {
  max-width: 450px;
  margin: 0 auto;
  padding: 2rem;
}

.card {
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}

.register-title {
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

.register-btn {
  background: linear-gradient(45deg, #667eea, #764ba2);
  border: none;
  border-radius: 12px;
  font-weight: 600;
  letter-spacing: 1px;
  padding: 0.75rem 2rem;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.3);
}

.register-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
  background: linear-gradient(45deg, #5a6fd8, #6a42a6);
}

.register-btn:disabled {
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

.login-link {
  color: #667eea;
  text-decoration: none;
  transition: color 0.3s ease;
}

.login-link:hover {
  color: #5a6fd8;
  text-decoration: underline;
}

.invalid-feedback {
  font-size: 0.875rem;
  margin-top: 0.25rem;
}

/* Responsive Design */
@media (max-width: 991px) {
  .register-form-wrapper {
    padding: 1rem;
  }
}

@media (max-width: 768px) {
  .card-body {
    padding: 2rem !important;
  }
  
  .register-title {
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
  
  .register-title {
    font-size: 1.75rem;
  }
  
  .custom-input {
    font-size: 0.9rem;
  }
}

/* Animation */
.register-form-wrapper {
  animation: slideInRight 0.6s ease-out;
}

.illustration {
  animation: slideInLeft 0.6s ease-out;
}

@keyframes slideInRight {
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}
</style>