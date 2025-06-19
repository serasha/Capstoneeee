<template>
  <section class="hero-section text-white">
    <div class="container-fluid hero-container">
      <div class="row align-items-center h-100">
        <!-- Kiri: Teks -->
        <div class="col-md-6 px-5 py-4 hero-text">
          <h1 class="fw-bold mb-3">Modul Layanan Pendaftaran<br>dan Pendataan Transmigrasi</h1>
          <p>
            Layanan Transmigrasi dari Pemerintah Kota Yogyakarta kini hadir di genggaman Anda melalui 
            Sistem Pemerintahan Berbasis Elektronik (SPBE). Proses pendaftaran, informasi lokasi tujuan, 
            hingga bantuan keberangkatan bisa diakses dengan mudah, cepat, dan transparan.
          </p>
        </div>

        <!-- Kanan: Lottie Animation -->
        <div class="col-md-6 text-center mt-4 mt-md-0">
          <div class="lottie-container" :class="{ loaded: isLoaded }">
            <!-- Spinner SVG -->
            <div v-if="!isLoaded" class="custom-spinner">
              <svg width="48" height="48" viewBox="0 0 48 48">
                <circle
                  cx="24"
                  cy="24"
                  r="20"
                  fill="none"
                  stroke="#fff"
                  stroke-width="4"
                  stroke-linecap="round"
                  stroke-dasharray="100"
                  stroke-dashoffset="60"
                >
                  <animateTransform
                    attributeName="transform"
                    type="rotate"
                    from="0 24 24"
                    to="360 24 24"
                    dur="1s"
                    repeatCount="indefinite"
                  />
                </circle>
              </svg>
            </div>
            <!-- Lottie animation container -->
            <div ref="lottieContainer" class="lottie-animation"></div>
            
            <!-- Option 2: Jika menggunakan Vue Lottie component -->
            <!-- <lottie-animation
              ref="anim" 
              :animation-data="animationData"
              :auto-play="true"
              :loop="true"
              :speed="1"
              class="lottie-animation"
            /> -->
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import lottie from 'lottie-web'
// Jika menggunakan Vue Lottie component, uncomment line berikut:
// import LottieAnimation from 'lottie-vuejs/src/LottieAnimation.vue'

export default {
  name: 'HeroSection',
  // components: {
  //   LottieAnimation // Jika menggunakan Vue Lottie component
  // },
  data() {
    return {
      animationData: null,
      lottieAnimation: null,
      isLoaded: false
    }
  },
  mounted() {
    this.loadLottieAnimation()
  },
  beforeUnmount() {
    if (this.lottieAnimation) {
      this.lottieAnimation.destroy()
    }
  },
  methods: {
    async loadLottieAnimation() {
      try {
        // Option 1: Load dari CDN LottieFiles (recommended)
        const animationUrl = 'https://assets3.lottiefiles.com/packages/lf20_jcikwtux.json' // Contoh: Digital services animation
        
        // Option 2: Load dari file lokal
        // const animationData = await import('@/assets/animations/hero-animation.json')
        
        this.lottieAnimation = lottie.loadAnimation({
          container: this.$refs.lottieContainer,
          renderer: 'svg',
          loop: true,
          autoplay: true,
          path: animationUrl // Untuk load dari URL
          // animationData: animationData.default // Untuk load dari file lokal
        })

        this.lottieAnimation.addEventListener('DOMLoaded', () => {
          this.isLoaded = true
        })
        this.lottieAnimation.addEventListener('complete', () => {
          console.log('Animation completed')
        })

      } catch (error) {
        console.error('Error loading animation:', error)
        // Fallback ke static illustration jika gagal load
        this.showFallbackIllustration()
        this.isLoaded = true // Hide spinner on error/fallback
      }
    },

    showFallbackIllustration() {
      // Fallback SVG jika Lottie gagal load
      this.$refs.lottieContainer.innerHTML = `
        <svg viewBox="0 0 400 300" class="fallback-illustration">
          <rect width="400" height="300" fill="#7CB342" rx="20"/>
          <circle cx="150" cy="120" r="30" fill="#FFDBCB"/>
          <rect x="130" y="150" width="40" height="60" fill="#3F51B5" rx="10"/>
          <circle cx="250" cy="120" r="30" fill="#FFDBCB"/>
          <rect x="230" y="150" width="40" height="60" fill="#E91E63" rx="10"/>
          <text x="200" y="250" text-anchor="middle" fill="white" font-size="16">Digital Services</text>
        </svg>
      `
    },

    // Method untuk kontrol animasi
    playAnimation() {
      if (this.lottieAnimation) {
        this.lottieAnimation.play()
      }
    },

    pauseAnimation() {
      if (this.lottieAnimation) {
        this.lottieAnimation.pause()
      }
    },

    stopAnimation() {
      if (this.lottieAnimation) {
        this.lottieAnimation.stop()
      }
    }
  }
}
</script>

<style scoped>
.hero-section {
  background: linear-gradient(rgba(38, 37, 37, 0.2), rgba(38, 37, 37, 0.2)), url(hero-ilustrasi.png);
  background-size: cover;
  background-position: center;
  position: relative;
  overflow: hidden;
  min-height: 100vh;
  display: flex;
  align-items: center;
}

.hero-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    radial-gradient(circle at 20% 20%, rgba(255,255,255,0.1) 1px, transparent 1px),
    radial-gradient(circle at 80% 80%, rgba(255,255,255,0.1) 1px, transparent 1px);
  background-size: 50px 50px, 30px 30px;
}

.lottie-container {
  position: relative;
  max-width: 100%;
  height: auto;
  min-height: 320px;
}

.lottie-animation {
  width: 100%;
  height: auto;
  max-width: 650px;
  max-height: 500px;
  margin: 0 auto;
  filter: drop-shadow(0 15px 40px rgba(0,0,0,0.25));
}

.fallback-illustration {
  width: 100%;
  height: auto;
  max-width: 400px;
  filter: drop-shadow(0 10px 30px rgba(0,0,0,0.2));
}

/* Custom SVG Spinner */
.custom-spinner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 2;
  background: none;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  pointer-events: none;
}

.lottie-container.loaded .custom-spinner {
  display: none;
}

h1 {
  font-size: 2.8rem;
  line-height: 1.2;
  margin-bottom: 2rem;
  position: relative;
  z-index: 2;
}

p {
  font-size: 1.1rem;
  line-height: 1.8;
  opacity: 0.9;
  position: relative;
  z-index: 2;
  margin-bottom: 2rem;
}

.hero-container {
  min-height: 100vh;
  padding-top: 3rem;
  padding-bottom: 3rem;
}

.hero-text {
  padding-top: 8rem !important;
}

@media (max-width: 768px) {
  .hero-section {
    min-height: 90vh;
  }
  
  h1 {
    font-size: 2rem;
  }
  
  p {
    font-size: 1rem;
    line-height: 1.6;
  }
  
  .lottie-animation {
    max-width: 400px;
    max-height: 320px;
  }
  
  .col-md-6.px-5 {
    padding-left: 2rem !important;
    padding-right: 2rem !important;
  }
}
</style>