<script setup>
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/Footer.vue'
</script>

<template>
  <div style="min-height: 100vh; display: flex; flex-direction: column;">
    <Navbar />
    <div style="flex: 1 0 auto;">
      <!-- Tempatkan konten halaman di sini -->
      <slot />
    </div>
    <Footer />
  </div>
</template>