<script setup>
import HeroSection from '../components/HeroSection.vue'
import Section from '../components/Section.vue'
import { useUserStore } from '@/store/userStore'
import { storeToRefs } from 'pinia'

const userStore = useUserStore()
const { role } = storeToRefs(userStore)
</script>

<template>
  <div class="home">
    <HeroSection />
    <Section :role="role" />
  </div>
</template>

<style scoped>
.home {
  min-height: 100vh;
}
</style>
