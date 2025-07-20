import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
    isAuthenticated: false,
    loading: false,
    error: ''
  }),
  actions: {
    async fetchUser() {
      this.loading = true
      try {
        const res = await fetch('/api/user/me', { credentials: 'include' })
        if (res.ok) {
          this.user = await res.json()
          this.isAuthenticated = true
        } else {
          this.user = null
          this.isAuthenticated = false
        }
      } catch {
        this.user = null
        this.isAuthenticated = false
      } finally {
        this.loading = false
      }
    },
    setUser(user) {
      this.user = user
      this.isAuthenticated = !!user
    },
    logout() {
      this.user = null
      this.isAuthenticated = false
    }
  },
  persist: true
}) 