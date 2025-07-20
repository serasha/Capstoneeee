import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
    isAuthenticated: false,
    loading: false,
    error: '',
    guest: true, // true jika belum login
    role: 'guest', // 'user', 'admin', atau 'guest'
  }),
  actions: {
    async fetchUser() {
      this.loading = true
      try {
        const res = await fetch('/api/user/me', { credentials: 'include' })
        if (res.ok) {
          this.user = await res.json()
          this.isAuthenticated = true
          this.guest = false
          this.role = this.user.role || 'user'
        } else {
          this.user = null
          this.isAuthenticated = false
          this.guest = true
          this.role = 'guest'
        }
      } catch {
        this.user = null
        this.isAuthenticated = false
        this.guest = true
        this.role = 'guest'
      } finally {
        this.loading = false
      }
    },
    setUser(user) {
      this.user = user
      this.isAuthenticated = !!user
      this.guest = !user
      this.role = user && user.role ? user.role : 'guest'
    },
    logout() {
      this.user = null
      this.isAuthenticated = false
      this.guest = true
      this.role = 'guest'
    }
  },
  persist: true
}) 