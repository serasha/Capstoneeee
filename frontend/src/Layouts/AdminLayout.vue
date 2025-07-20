<template>
  <div class="admin-layout">
    <!-- Sidebar -->
    <AdminSidebar />
    
    <!-- Main Content -->
    <div class="main-content">
      <!-- User Profile Header -->
      <div class="user-header">
        <div class="user-info">
          <span class="user-name">{{ userData.name }}</span>
          <span class="user-id">{{ userData.id }}</span>
          <img :src="userData.avatar" :alt="userData.name" class="user-avatar">
        </div>
      </div>
      
      <!-- Page Content -->
      <div class="page-content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import AdminSidebar from '@/components/AdminSidebar.vue'

export default {
  name: 'AdminLayout',
  components: {
    AdminSidebar
  },
  data() {
    return {
      userData: {
        name: 'Heltrizulfikar. java',
        id: '190002290',
        avatar: 'https://via.placeholder.com/50x50/6c757d/ffffff?text=HJ'
      }
    }
  }
}
</script>

<style scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
  background: #f8f9fa;
}

.main-content {
  flex: 1;
  margin-left: 280px;
  display: flex;
  flex-direction: column;
}

.user-header {
  background: white;
  padding: 1rem 2rem;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  justify-content: flex-end;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-name {
  font-weight: 600;
  color: #333;
}

.user-id {
  color: #6c757d;
  font-size: 0.9rem;
}

.user-avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #e9ecef;
}

.page-content {
  flex: 1;
  padding: 2rem;
}

/* Responsive */
@media (max-width: 768px) {
  .main-content {
    margin-left: 0;
  }
  
  .user-header {
    padding: 1rem;
  }
  
  .page-content {
    padding: 1rem;
  }
}
</style>