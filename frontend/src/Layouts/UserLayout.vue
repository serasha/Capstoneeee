<script setup>
import Navbar from '@/components/Navbar.vue'
import Footer from '@/components/Footer.vue'
import HeroSection from '@/components/HeroSection.vue';
</script>

<template>
  <div class="layout-container">
    <Navbar />
    <HeroSection/>
    <main class="content-wrapper">
      <slot />
    </main>
    <Footer />
  </div>
</template>

<style scoped>
.layout-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  width: 100%;
  margin: 0;
  padding: 0;
}

.content-wrapper {
  flex: 1 0 auto;
  width: 100%;
  min-height: calc(100vh - 60px); /* 60px adalah tinggi navbar */
  padding-top: 60px; /* Kompensasi untuk fixed navbar */
}
</style>