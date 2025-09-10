import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authAPI } from '../services/api'
import { readonly } from 'vue'
import type { User, LoginRequest, RegisterRequest, AuthState } from '../types'

export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref<User | null>(null)
  const token = ref<string | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Initialize from localStorage
  const initializeAuth = () => {
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')
    
    if (storedToken && storedUser) {
      token.value = storedToken
      try {
        user.value = JSON.parse(storedUser)
      } catch (e) {
        console.error('Failed to parse stored user data:', e)
        clearAuthData()
      }
    }
  }

  // Getters
  const isAuthenticated = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 'admin')
  const isCustomer = computed(() => user.value?.role === 'customer')
  const userInitials = computed(() => {
    if (!user.value?.full_name) return 'U'
    return user.value.full_name
      .split(' ')
      .map(name => name.charAt(0).toUpperCase())
      .join('')
      .slice(0, 2)
  })

  // Actions
  const clearAuthData = () => {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  const setAuthData = (authToken: string, userData: User) => {
    token.value = authToken
    user.value = userData
    localStorage.setItem('token', authToken)
    localStorage.setItem('user', JSON.stringify(userData))
  }

  const login = async (credentials: LoginRequest): Promise<{ success: boolean; error?: string }> => {
    loading.value = true
    error.value = null

    try {
      const response = await authAPI.login(credentials)
      setAuthData(response.token, response.user)
      return { success: true }
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Login failed'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      loading.value = false
    }
  }

  const register = async (userData: RegisterRequest): Promise<{ success: boolean; error?: string }> => {
    loading.value = true
    error.value = null

    try {
      const response = await authAPI.register(userData)
      setAuthData(response.token, response.user)
      return { success: true }
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Registration failed'
      error.value = errorMessage
      return { success: false, error: errorMessage }
    } finally {
      loading.value = false
    }
  }

  const fetchProfile = async (): Promise<void> => {
    if (!token.value) return

    try {
      const response = await authAPI.getProfile()
      user.value = response.user
      localStorage.setItem('user', JSON.stringify(response.user))
    } catch (err) {
      console.error('Failed to fetch profile:', err)
      logout()
    }
  }

  const logout = (): void => {
    clearAuthData()
    error.value = null
    // Optional: Redirect to login page
    // router.push('/login')
  }

  const clearError = (): void => {
    error.value = null
  }

  const updateProfile = (updatedUser: Partial<User>): void => {
    if (user.value) {
      user.value = { ...user.value, ...updatedUser }
      localStorage.setItem('user', JSON.stringify(user.value))
    }
  }

  // Initialize auth state
  initializeAuth()

  return {
    // State
    user,
    token,
    loading,
    error,
    // Getters
    isAuthenticated,
    isAdmin,
    isCustomer,
    userInitials,
    // Actions
    login,
    register,
    logout,
    fetchProfile,
    clearError,
    updateProfile,
    initializeAuth,
    clearAuthData,
    setAuthData,
  }
})