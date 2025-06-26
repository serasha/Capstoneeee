<template>
  <div v-if="visible" class="modal-overlay" @click.self="closeModal">
    <div class="success-container modal-popup">
      <div class="row min-vh-100">
        <div class="col-12 d-flex align-items-center justify-content-center">
          <div class="success-wrapper">
            <!-- Close Button -->
            <button class="modal-close-btn" @click="closeModal" aria-label="Close">&times;</button>
            
            <!-- Success Card -->
            <div class="success-card">
              <!-- Animated Check Icon -->
              <div class="check-icon-wrapper">
                <div class="check-circle">
                  <div class="check-icon">
                    <svg class="checkmark" viewBox="0 0 52 52">
                      <circle class="checkmark__circle" cx="26" cy="26" r="25" fill="none"/>
                      <path class="checkmark__check" fill="none" d="M14.1 27.2l7.1 7.2 16.7-16.8"/>
                    </svg>
                  </div>
                </div>
              </div>

              <!-- Success Message -->
              <div class="success-content">
                <h2 class="success-title" v-if="!showAnimation">
                  {{ successTitle }}
                </h2>
                <h2 class="success-title animate__animated animate__fadeInUp" v-else>
                  {{ successTitle }}
                </h2>
                
                <p class="success-subtitle" v-if="userData">
                  Selamat datang, <strong>{{ userData.fullName }}</strong>!
                </p>
                <p class="success-subtitle" v-else>
                  Akun Anda telah berhasil dibuat dan siap digunakan.
                </p>
              </div>

              <!-- Action Buttons -->
              <div class="success-actions">
                <button 
                  class="btn btn-status"
                  @click="viewApplicationStatus"
                  :disabled="isLoading"
                >
                  <span v-if="!isLoading">Lihat Status Pengajuan</span>
                  <span v-else>
                    <span class="spinner-border spinner-border-sm me-2" role="status"></span>
                    Loading...
                  </span>
                </button>
                
                <div class="secondary-actions mt-3">
                  <button 
                    class="btn btn-outline-secondary"
                    @click="goToHome"
                  >
                    <i class="fas fa-home me-2"></i>
                    Beranda
                  </button>
                </div>
              </div>

              <!-- Additional Info -->
              <div class="additional-info">
                <div class="info-item">
                  <i class="fas fa-envelope text-primary"></i>
                  <span>Email konfirmasi telah dikirim</span>
                </div>
                <div class="info-item">
                  <i class="fas fa-shield-alt text-success"></i>
                  <span>Akun Anda aman dan terverifikasi</span>
                </div>
              </div>
            </div>

            <!-- Floating Elements -->
            <div class="floating-elements">
              <div class="float-element element-1">
                <i class="fas fa-star"></i>
              </div>
              <div class="float-element element-2">
                <i class="fas fa-heart"></i>
              </div>
              <div class="float-element element-3">
                <i class="fas fa-thumbs-up"></i>
              </div>
              <div class="float-element element-4">
                <i class="fas fa-trophy"></i>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Background Particles -->
      <div class="particles">
        <div class="particle" v-for="i in 20" :key="i"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RegistrationSuccess',
  props: {
    userData: {
      type: Object,
      default: null
    },
    successTitle: {
      type: String,
      default: 'Pendaftaran Anda Berhasil!'
    },
    visible: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      showAnimation: false,
      isLoading: false
    }
  },
  methods: {
    viewApplicationStatus() {
      this.isLoading = true;
      
      // Simulate API call
      setTimeout(() => {
        this.isLoading = false;
        this.$emit('view-status');
        // You can redirect or show status modal here
        console.log('Viewing application status...');
      }, 2000);
    },
    
    goToHome() {
      this.$emit('go-to-home');
      // Handle navigation to home
      console.log('Navigating to home...');
    },
    
    addConfetti() {
      // Simple confetti effect
      setTimeout(() => {
        const particles = document.querySelectorAll('.particle');
        particles.forEach((particle, index) => {
          setTimeout(() => {
            particle.style.animation = `confetti 3s ease-out forwards`;
            particle.style.animationDelay = `${index * 0.1}s`;
          }, index * 100);
        });
      }, 1000);
    },
    
    closeModal() {
      this.$emit('close');
    }
  },
  
  mounted() {
    // Trigger animations after component mounts
    setTimeout(() => {
      this.showAnimation = true;
    }, 500);
    
    // Add some celebratory effects
    this.addConfetti();
  }
}
</script>

<style scoped>
/* Modal Overlay */
.modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(44, 62, 80, 0.55);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow-y: auto;
}

/* Modal Popup Container */
.modal-popup {
  min-height: unset;
  
  border-radius: 24px;
  margin: 40px auto;
  position: relative;
  /* background: transparent; */ /* opsional, default transparan */
}

/* Main Container */
.success-container {
  /* background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); */

  min-height: 100vh;
  position: relative;
  overflow: hidden;
}

.success-wrapper {
  position: relative;
  z-index: 10;
  max-width: 500px;
  margin: 0 auto;
  padding: 20px;
}

/* Success Card */
.success-card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 24px;
  padding: 50px 40px;
  text-align: center;
  box-shadow: 
    0 25px 50px rgba(0, 0, 0, 0.1),
    0 10px 25px rgba(0, 0, 0, 0.05);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  animation: slideInUp 0.8s ease-out;
  position: relative;
}

/* Check Icon Animation */
.check-icon-wrapper {
  margin-bottom: 30px;
}

.check-circle {
  width: 120px;
  height: 120px;
  margin: 0 auto;
  background: linear-gradient(135deg, #4CAF50, #81C784);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 15px 30px rgba(76, 175, 80, 0.3);
  animation: bounceIn 1s ease-out 0.5s both;
}

.checkmark {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  display: block;
  stroke-width: 3;
  stroke: #fff;
  stroke-miterlimit: 10;
  box-shadow: inset 0px 0px 0px #4CAF50;
  animation: fill .4s ease-in-out .4s forwards, scale .3s ease-in-out .9s both;
}

.checkmark__circle {
  stroke-dasharray: 166;
  stroke-dashoffset: 166;
  stroke-width: 3;
  stroke-miterlimit: 10;
  stroke: #4CAF50;
  fill: none;
  animation: stroke 0.6s cubic-bezier(0.65, 0, 0.45, 1) forwards;
}

.checkmark__check {
  transform-origin: 50% 50%;
  stroke-dasharray: 48;
  stroke-dashoffset: 48;
  animation: stroke 0.3s cubic-bezier(0.65, 0, 0.45, 1) 0.8s forwards;
}

/* Success Content */
.success-title {
  color: #2c3e50;
  font-size: 2.2rem;
  font-weight: 700;
  margin-bottom: 20px;
  line-height: 1.2;
}

.success-subtitle {
  color: #6c757d;
  font-size: 1.1rem;
  margin-bottom: 30px;
  line-height: 1.5;
}

.success-subtitle strong {
  color: #4CAF50;
  font-weight: 600;
}

/* Action Buttons */
.success-actions {
  margin-bottom: 30px;
}

.btn-status {
  background: linear-gradient(135deg, #dc3545, #e74c3c);
  border: none;
  color: white;
  padding: 15px 30px;
  border-radius: 12px;
  font-weight: 600;
  font-size: 1.1rem;
  width: 100%;
  transition: all 0.3s ease;
  box-shadow: 0 8px 16px rgba(220, 53, 69, 0.3);
}

.btn-status:hover:not(:disabled) {
  background: linear-gradient(135deg, #c82333, #dc3545);
  transform: translateY(-2px);
  box-shadow: 0 12px 24px rgba(220, 53, 69, 0.4);
}

.btn-status:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.secondary-actions .btn {
  border-radius: 10px;
  padding: 10px 20px;
  font-weight: 500;
  transition: all 0.3s ease;
}

.secondary-actions .btn:hover {
  transform: translateY(-1px);
}

/* Additional Info */
.additional-info {
  border-top: 1px solid #e9ecef;
  padding-top: 25px;
  margin-top: 25px;
}

.info-item {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
  color: #6c757d;
  font-size: 0.95rem;
}

.info-item i {
  margin-right: 10px;
  width: 20px;
  text-align: center;
}

/* Floating Elements */
.floating-elements {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 5;
}

.float-element {
  position: absolute;
  color: rgba(255, 255, 255, 0.6);
  font-size: 1.5rem;
  animation: float 6s ease-in-out infinite;
}

.element-1 {
  top: 15%;
  left: 10%;
  animation-delay: 0s;
  color: #FFD700;
}

.element-2 {
  top: 25%;
  right: 15%;
  animation-delay: 1s;
  color: #FF69B4;
}

.element-3 {
  bottom: 30%;
  left: 20%;
  animation-delay: 2s;
  color: #00CED1;
}

.element-4 {
  bottom: 20%;
  right: 10%;
  animation-delay: 3s;
  color: #FFA500;
}

/* Background Particles */
.particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 1;
}

.particle {
  position: absolute;
  width: 8px;
  height: 8px;
  background: rgba(255, 255, 255, 0.5);
  border-radius: 50%;
}

.particle:nth-child(odd) {
  background: rgba(76, 175, 80, 0.6);
}

.particle:nth-child(even) {
  background: rgba(255, 193, 7, 0.6);
}

/* Animations */
@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(50px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounceIn {
  0% {
    opacity: 0;
    transform: scale(0.3);
  }
  50% {
    opacity: 1;
    transform: scale(1.05);
  }
  70% {
    transform: scale(0.9);
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}

@keyframes stroke {
  100% {
    stroke-dashoffset: 0;
  }
}

@keyframes scale {
  0%, 100% {
    transform: none;
  }
  50% {
    transform: scale3d(1.1, 1.1, 1);
  }
}

@keyframes fill {
  100% {
    box-shadow: inset 0px 0px 0px 30px #4CAF50;
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.7;
  }
  50% {
    transform: translateY(-20px) rotate(180deg);
    opacity: 1;
  }
}

@keyframes confetti {
  0% {
    transform: translateY(-100vh) rotate(0deg);
    opacity: 1;
  }
  100% {
    transform: translateY(100vh) rotate(720deg);
    opacity: 0;
  }
}

/* Responsive Design */
@media (max-width: 768px) {
  .success-card {
    margin: 20px;
    padding: 40px 25px;
  }
  
  .success-title {
    font-size: 1.8rem;
  }
  
  .check-circle {
    width: 100px;
    height: 100px;
  }
  
  .checkmark {
    width: 50px;
    height: 50px;
  }
  
  .secondary-actions {
    flex-direction: column;
  }
  
  .secondary-actions .btn {
    width: 100%;
    margin-bottom: 10px;
  }
  
  .secondary-actions .btn:last-child {
    margin-bottom: 0;
  }
}

@media (max-width: 480px) {
  .success-card {
    margin: 15px;
    padding: 30px 20px;
  }
  
  .success-title {
    font-size: 1.6rem;
  }
  
  .success-subtitle {
    font-size: 1rem;
  }
  
  .btn-status {
    padding: 12px 20px;
    font-size: 1rem;
  }
  
  .float-element {
    font-size: 1.2rem;
  }
}

/* Animation delays for staggered effect */
.info-item:nth-child(1) {
  animation: fadeInUp 0.6s ease-out 1.2s both;
}

.info-item:nth-child(2) {
  animation: fadeInUp 0.6s ease-out 1.4s both;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Close Button */
.modal-close-btn {
  position: absolute;
  top: 18px;
  right: 18px;
  background: transparent;
  border: none;
  font-size: 2rem;
  color: #888;
  cursor: pointer;
  z-index: 20;
  transition: color 0.2s;
}
.modal-close-btn:hover {
  color: #dc3545;
}
</style>