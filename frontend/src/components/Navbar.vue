<template>
  <nav class="navbar navbar-expand-lg navbar-dark w-100 py-2 main-navbar">
    <div class="container-fluid px-4 d-flex justify-content-between align-items-center w-100">
      <!-- Brand/Logo -->
      <a class="navbar-brand fw-bold ms-3" href="#">
        SPBE
      </a>

      <!-- Navbar items (centered) -->
      <div class="navbar-center d-flex flex-row justify-content-center mx-auto">
        <ul class="navbar-nav flex-row mb-0">
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/beranda">BERANDA</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/yogyakarta">YOGYAKARTA</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/bantuan">BANTUAN</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/panduan">PANDUAN</a>
          </li>
          <li class="nav-item">
            <a class="nav-link text-white px-3" href="/faq">FAQ</a>
          </li>
        </ul>
      </div>

      <!-- Right side buttons -->
      <div class="d-flex gap-2 mb-0 me-4 daftar-masuk-group">
        <router-link to="/daftar" class="btn btn-outline-warning daftar-btn">
          Daftar
        </router-link>
        <router-link to="/login" class="btn btn-warning masuk-btn">
          Masuk
        </router-link>
       
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      // Anda bisa menambahkan data reactive di sini jika diperlukan
    }
  },
  methods: {
    // Anda bisa menambahkan methods di sini jika diperlukan
  }
}
</script>

<style scoped>
/* Navbar positioning */
.main-navbar {
  background-color: #222;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1030;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  width: 100vw !important;
  height: 60px !important;
  padding-top: 8px !important;
  padding-bottom: 8px !important;
  border-radius: 0;
}

/* Ensure full width */
.container-fluid {
  max-width: 100% !important;
  width: 100% !important;
  height: 100%;
  display: flex;
  align-items: center;
}

/* Body padding to compensate for fixed navbar */
body {
  padding-top: 60px !important;
}

.navbar-brand {
  font-weight: bold !important;
  border: none !important;
  color: #ff4d4d;
  font-size: 1.5rem;
  font-style: italic;
}

.nav-link {
  font-weight: 500;
  transition: color 0.3s ease;
}

.nav-link:hover {
  color: #ffca28 !important;
}

.btn {
  font-weight: 500;
  transition: all 0.3s ease;
}

.btn-outline-warning:hover {
  background-color: #ffca28;
  color: #000;
}

.btn-warning:hover {
  background-color: #ffb300;
  border-color: #ffb300;
}

.navbar-center {
  flex: 1 1 auto;
  justify-content: center !important;
  display: flex !important;
}

.navbar-nav.flex-row {
  flex-direction: row !important;
  display: flex !important;
  gap: 1rem;
  justify-content: center;
  align-items: center;
}

.navbar-nav .nav-item {
  display: flex;
  align-items: center;
}

.navbar-nav .nav-link {
  padding-left: 1rem !important;
  padding-right: 1rem !important;
  color: #fff !important;
}

/* Responsive: stack menu vertically on mobile */
@media (max-width: 991px) {
  .container-fluid {
    flex-direction: column;
    padding-top: 1rem;
  }
  .navbar-nav.flex-row {
    flex-direction: column !important;
    width: 100%;
    gap: 0;
  }
  .navbar-nav .nav-item {
    justify-content: center;
  }
  .navbar-nav .nav-link {
    padding-left: 0.5rem !important;
    padding-right: 0.5rem !important;
    text-align: center;
  }
  .d-flex {
    flex-direction: column;
    width: 100%;
    margin-top: 1rem;
  }
  .d-flex .btn {
    width: 100%;
    margin-bottom: 0.5rem;
  }
}

.d-flex.gap-2.mb-0.me-4 {
  margin-right: 2rem !important;
}

.daftar-btn {
  border-radius: 4px !important;
  font-size: 16px !important;
  border: 2px solid #ffca28 !important;
  background: transparent !important;
  color: #ffca28 !important;
}

.masuk-btn {
  border-radius: 4px !important;
  font-size: 16px !important;
  background-color: #ffca28;
  border-color: #ffca28;
  color: #000;
}

.daftar-masuk-group > .daftar-btn {
  margin-right: 12px;
}
</style>
