import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/services/api'
import router from '@/router'
import type { LoginCredentials, RegisterCredentials } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  const userId = ref<string | null>(localStorage.getItem('userId'))
  const token = ref<string | null>(localStorage.getItem('token'))
  const isGuest = ref(localStorage.getItem('isGuest') === 'true')
  const isAdmin = ref(localStorage.getItem('isAdmin') === 'true')
  const isAuthenticated = computed(() => !!token.value)

  function checkAuth(): void {
    const storedToken = localStorage.getItem('token')
    const storedUserId = localStorage.getItem('userId')
    if (storedToken && storedUserId) {
      token.value = storedToken
      userId.value = storedUserId
      isGuest.value = localStorage.getItem('isGuest') === 'true'
      isAdmin.value = localStorage.getItem('isAdmin') === 'true'
    }
  }

  function setSession(newToken: string, newUserId: string, guest: boolean, admin: boolean): void {
    token.value = newToken
    userId.value = String(newUserId)
    isGuest.value = guest
    isAdmin.value = admin
    localStorage.setItem('token', newToken)
    localStorage.setItem('userId', String(newUserId))
    localStorage.setItem('isGuest', String(guest))
    localStorage.setItem('isAdmin', String(admin))
  }

  async function login(credentials: LoginCredentials): Promise<{ success: boolean; message: string }> {
    try {
      const data = await api.login(credentials)
      setSession(data.token, data.user_id, false, data.is_admin === true)
      return { success: true, message: data.message }
    } catch (error: any) {
      return { success: false, message: error.response?.data?.error || 'Error al iniciar sesión' }
    }
  }

  async function loginAsGuest(): Promise<{ success: boolean; message: string }> {
    try {
      const data = await api.guestLogin()
      setSession(data.token, data.user_id, true, false)
      return { success: true, message: data.message }
    } catch (error: any) {
      return {
        success: false,
        message: error.response?.data?.error || 'Error al entrar como invitado'
      }
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
    isGuest.value = false
    isAdmin.value = false
    localStorage.removeItem('token')
    localStorage.removeItem('userId')
    localStorage.removeItem('isGuest')
    localStorage.removeItem('isAdmin')
    router.push('/login')
  }

  return {
    userId,
    token,
    isGuest,
    isAdmin,
    isAuthenticated,
    checkAuth,
    login,
    loginAsGuest,
    register,
    logout
  }
})
