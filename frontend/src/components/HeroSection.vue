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

        <!-- Kanan: Gambar Hero -->
        <div class="col-md-6 text-center mt-4 mt-md-0">
          <div class="hero-image-container">
            <img src="/Hr.png" alt="Hero Ilustrasi" class="img-fluid hero-image" />
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'HeroSection'
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

.hero-image-container {
  position: relative;
  max-width: 100%;
  min-height: 320px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.hero-image {
  width: 100%;
  height: auto;
  max-width: 650px;
  max-height: 500px;
  margin: 0 auto;
  filter: drop-shadow(0 15px 40px rgba(0,0,0,0.25));
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
  
  .hero-image {
    max-width: 400px;
    max-height: 320px;
  }
  
  .col-md-6.px-5 {
    padding-left: 2rem !important;
    padding-right: 2rem !important;
  }
}
</style>