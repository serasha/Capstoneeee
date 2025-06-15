<template>
  <footer class="footer bg-dark text-white">
      <div class="footer-wrapper">
      <div class="container py-5">
        <!-- Main Footer Content -->
        <div class="row gy-4">
          <!-- Logo and Description Column -->
          <div class="col-lg-4 col-md-12">
            <div class="text-center text-lg-start">
              <h2 class="logo mb-3">{{ footerData.title }}</h2>
              <p class="description text-muted">
                {{ footerData.description }}
              </p>
              <!-- Social Media Icons -->
              <div class="social-icons mt-3">
                <a 
                  v-for="social in footerData.socialMedia" 
                  :key="social.name"
                  :href="social.url" 
                  class="social-link me-3"
                  :title="social.name"
                >
                  <i :class="social.icon"></i>
                </a>
              </div>
            </div>
          </div>

          <!-- Learn Section -->
          <div class="col-lg-4 col-md-6">
            <div class="text-center text-lg-start">
              <h5 class="section-title mb-3">{{ footerData.learnSection.title }}</h5>
              <ul class="list-unstyled footer-links">
                <li v-for="link in footerData.learnSection.links" :key="link.id" class="mb-2">
                  <a 
                    href="#" 
                    class="footer-link text-decoration-none"
                    @click.prevent="handleLinkClick(link)"
                  >
                    {{ link.text }}
                  </a>
                </li>
              </ul>
            </div>
          </div>

          <!-- Services Section -->
          <div class="col-lg-4 col-md-6">
            <div class="text-center text-lg-start">
              <h5 class="section-title mb-3">{{ footerData.servicesSection.title }}</h5>
              <ul class="list-unstyled footer-links">
                <li v-for="service in footerData.servicesSection.services" :key="service.id" class="mb-2">
                  <a 
                    href="#" 
                    class="footer-link text-decoration-none d-flex align-items-center justify-content-center justify-content-lg-start"
                    @click.prevent="handleServiceClick(service)"
                  >
                    <i :class="service.icon" class="me-2"></i>
                    {{ service.text }}
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>

        <!-- Divider -->
        <hr class="footer-divider my-4">

        <!-- Copyright and Additional Info -->
        <div class="row align-items-center">
          <div class="col-12">
            <p class="copyright mb-0 text-center">
              {{ footerData.copyright }}
            </p>
          </div>
          <div class="col-12 mt-2">
            <div class="text-center">
              <small class="text-muted">
                Version {{ appVersion }} | 
                <a href="#" class="text-muted text-decoration-none" @click.prevent="showInfo">
                  Info Sistem
                </a>
              </small>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Back to Top Button -->
    <button 
      v-show="showBackToTop"
      @click="scrollToTop"
      class="btn btn-primary position-fixed bottom-0 end-0 m-3 rounded-circle back-to-top"
      style="z-index: 1000;"
    >
      <i class="bi bi-arrow-up"></i>
    </button>
  </footer>
</template>

<script>
export default {
  name: 'SPBEFooter',
  data() {
    return {
      showBackToTop: false,
      appVersion: '2.1.0',
      footerData: {
        title: 'SPBE',
        description: 'Layanan Transmigrasi Pemkot Yogyakarta kini bisa diakses melalui SPBE. Daftar, cek info, dan pantau pengajuan lebih mudah langsung dari genggaman Anda.',
        socialMedia: [
          { name: 'Facebook', icon: 'bi bi-facebook', url: '#' },
          { name: 'Twitter', icon: 'bi bi-twitter-x', url: '#' },
          { name: 'Instagram', icon: 'bi bi-instagram', url: '#' },
          { name: 'YouTube', icon: 'bi bi-youtube', url: '#' }
        ],
        
        learnSection: {
          title: 'Pelajari',
          links: [
            { id: 1, text: 'Syarat dan Ketentuan', action: 'terms' },
            { id: 2, text: 'Kebijakan Privasi Data', action: 'privacy' },
            { id: 3, text: 'FAQ', action: 'faq' },
            { id: 4, text: 'Panduan Pengguna', action: 'guide' }
          ]
        },
        servicesSection: {
          title: 'Layanan',
          services: [
            { id: 1, text: 'Pengajuan Formulir Transmigrasi', icon: 'bi bi-file-earmark-text', action: 'form' },
            { id: 2, text: 'Pantau Status Pengajuan', icon: 'bi bi-search', action: 'status' },
            { id: 3, text: 'Bantuan & Konsultasi', icon: 'bi bi-headset', action: 'help' },
            { id: 4, text: 'Download Dokumen', icon: 'bi bi-download', action: 'download' }
          ]
        },
        copyright: 'Copyright ©2025, SPBE by Dinas Komunikasi Informatika Yogyakarta'
      }
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    handleLinkClick(link) {
      console.log('Link clicked:', link)
      this.$emit('link-clicked', link)
      
      switch(link.action) {
        case 'terms':
          break
        case 'privacy':
          break
        case 'faq':
          break
        case 'guide':
          break
      }
    },
    
    handleServiceClick(service) {
      console.log('Service clicked:', service)
      this.$emit('service-clicked', service)
      
      switch(service.action) {
        case 'form':
          break
        case 'status':
          break
        case 'help':
          break
        case 'download':
          break
      }
    },
    
    handleScroll() {
      this.showBackToTop = window.scrollY > 300
    },
    
    scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      })
    },
    
    showInfo() {
      alert('Sistem Informasi SPBE v' + this.appVersion)
    }
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,300;0,400;0,500;0,600;0,700;1,400;1,600;1,700&display=swap');

/* PERBAIKAN UTAMA: Pastikan footer mengisi lebar penuh layar */
.footer {
  background: linear-gradient(135deg, #1a1a1a 0%, #2d2d2d 100%);
  font-family: 'Poppins', sans-serif;
  position: relative;
  width: 100%;
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  min-height: 100%;
}

.footer-wrapper {
  width: 100%;
  display: flex;
  flex-direction: column;
  min-height: 100%;
}

.container {
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
  padding: 2rem 15px;
  flex: 1 0 auto;
}

.footer::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #ff6b6b, #ee5a52, #dc143c, #ff6b6b);
  background-size: 200% 100%;
  animation: gradientMove 3s ease-in-out infinite;
}

@keyframes gradientMove {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}

.logo {
  background: linear-gradient(135deg, #ff6b6b, #ee5a52, #dc143c);
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  font-style: italic;
  font-weight: 700;
  font-size: 2.5rem;
  letter-spacing: 2px;
  text-shadow: 0 2px 4px rgba(0,0,0,0.3);
}

.description {
  font-size: 0.95rem;
  line-height: 1.6;
  color: #cccccc !important;
}

.section-title {
  color: white;
  font-weight: 600;
  font-size: 1.25rem;
  position: relative;
  display: inline-block;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -5px;
  left: 50%;
  transform: translateX(-50%);
  width: 30px;
  height: 2px;
  background: linear-gradient(90deg, #ff6b6b, #dc143c);
}

@media (min-width: 992px) {
  .section-title::after {
    left: 0;
    transform: none;
  }
}

.footer-links {
  margin: 0;
  padding: 0;
}

.footer-link {
  color: #cccccc;
  font-size: 0.95rem;
  transition: all 0.3s ease;
  position: relative;
  padding: 0.25rem 0;
}

.footer-link:hover {
  color: #ff6b6b !important;
  transform: translateX(5px);
}

.footer-link i {
  width: 20px;
  color: #ff6b6b;
}

.social-icons {
  display: flex;
  justify-content: center;
  gap: 1rem;
}

@media (min-width: 992px) {
  .social-icons {
    justify-content: flex-start;
  }
}

.social-link {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: rgba(255, 107, 107, 0.1);
  border: 2px solid rgba(255, 107, 107, 0.3);
  border-radius: 50%;
  color: #ff6b6b;
  text-decoration: none;
  transition: all 0.3s ease;
  font-size: 1.1rem;
}

.social-link:hover {
  background: #ff6b6b;
  color: white;
  transform: translateY(-3px);
  box-shadow: 0 5px 15px rgba(255, 107, 107, 0.4);
}

.footer-divider {
  border: none;
  height: 1px;
  background: linear-gradient(90deg, transparent, #444444, transparent);
  margin: 2rem 0 1.5rem 0;
}

.copyright {
  color: #999999;
  font-size: 0.9rem;
  font-weight: 400;
  text-align: center;
}

.back-to-top {
  width: 50px;
  height: 50px;
  border-radius: 50% !important;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #ff6b6b, #dc143c);
  border: none;
  box-shadow: 0 4px 15px rgba(255, 107, 107, 0.3);
  transition: all 0.3s ease;
}

.back-to-top:hover {
  transform: translateY(-3px);
  box-shadow: 0 6px 20px rgba(255, 107, 107, 0.4);
}

.back-to-top i {
  font-size: 1.3rem;
}

/* Responsive adjustments */
@media (max-width: 767.98px) {
  .container {
    padding: 1.5rem 15px;
  }
  
  .logo {
    font-size: 2rem;
  }
  
  .description {
    font-size: 0.9rem;
  }
  
  .section-title {
    font-size: 1.1rem;
  }
  
  .footer-link {
    font-size: 0.9rem;
  }
  
  .copyright {
    font-size: 0.8rem;
  }
}

@media (max-width: 575.98px) {
  .logo {
    font-size: 1.8rem;
  }
  
  .description {
    font-size: 0.85rem;
  }
  
  .back-to-top {
    width: 45px;
    height: 45px;
  }
}
</style>