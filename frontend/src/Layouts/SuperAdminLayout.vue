<script setup>
import { defineProps } from 'vue';
import Sidebar from '@/components/Sidebar.vue';
import Dashboard from '@/components/Beranda.vue';

const props = defineProps({
  title: {
    type: String,
    required: true
  }
});
</script>
<template>
  <div class="super-admin-layout d-flex">
    <!-- Sidebar -->
    <Sidebar />
    <!-- Main Content -->
    <div class="main-content flex-grow-1">
      <router-view />
    </div>
  </div>
</template>

<style scoped>
.super-admin-layout {
  display: flex;
}

.header {
  background-color: #f8f9fa;
  padding: 1rem;
  border-bottom: 1px solid #dee2e6;
}

.main-content {
  flex-grow: 1;
  padding: 1rem;
}
</style>