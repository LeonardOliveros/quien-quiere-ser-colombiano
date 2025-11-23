import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/services/api'
import router from '@/router'
import type { LoginCredentials, RegisterCredentials } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  const userId = ref<string | null>(localStorage.getItem('userId'))
  const token = ref<string | null>(localStorage.getItem('token'))
  const isAuthenticated = computed(() => !!token.value)

  function checkAuth(): void {
    const storedToken = localStorage.getItem('token')
    const storedUserId = localStorage.getItem('userId')
    if (storedToken && storedUserId) {
      token.value = storedToken
      userId.value = storedUserId
    }
  }

  async function login(credentials: LoginCredentials): Promise<{ success: boolean; message: string }> {
    try {
      const data = await api.login(credentials)
      token.value = data.token
      userId.value = data.user_id
      localStorage.setItem('token', data.token)
      localStorage.setItem('userId', data.user_id)
      return { success: true, message: data.message }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al iniciar sesión' }
    }
  }

  async function register(credentials: RegisterCredentials): Promise<{ success: boolean; message: string }> {
    try {
      const data = await api.register(credentials)
      return { success: true, message: data.message }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al registrarse' }
    }
  }

  function logout(): void {
    token.value = null
    userId.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('userId')
    router.push('/login')
  }

  return {
    userId,
    token,
    isAuthenticated,
    checkAuth,
    login,
    register,
    logout
  }
})
