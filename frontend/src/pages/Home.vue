<template>
  <div class="home">
    <HeroSection />
    <Section />
  </div>
</template>

<script>
import HeroSection from '../components/HeroSection.vue'
import Section from '../components/Section.vue'

export default {
  name: 'Home',
  components: {
    HeroSection,
    Section
  }
}
</script>

<style scoped>
.home {
  min-height: 100vh;
}
</style>
