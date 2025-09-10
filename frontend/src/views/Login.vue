<template>
  <div class="p-6 max-w-md mx-auto">
    <h1 class="text-2xl font-bold mb-4">Login</h1>
    <form @submit.prevent="onLogin" class="space-y-4">
      <div>
        <label class="block text-sm font-medium mb-1" for="email">Email</label>
        <input v-model="email" id="email" type="email" required class="input input-bordered w-full" />
      </div>
      <div>
        <label class="block text-sm font-medium mb-1" for="password">Password</label>
        <input v-model="password" id="password" type="password" required class="input input-bordered w-full" />
      </div>
      <button type="submit" class="btn btn-primary w-full" :disabled="loading">
        {{ loading ? 'Logging in...' : 'Login' }}
      </button>
      <p v-if="error" class="text-red-500 text-sm mt-2">{{ error }}</p>
      <p class="text-sm mt-4">Don't have an account? <router-link to="/register" class="text-primary-600 hover:underline">Sign up</router-link></p>
    </form>
  </div>
</template>

<script setup lang="ts">

import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { authAPI } from '../services/api'


const email = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const onLogin = async () => {
  loading.value = true
  error.value = ''
  try {
    const res = await authAPI.login({ email: email.value, password: password.value })
    // Set auth state in Pinia store
    authStore.setAuthData(res.token, res.user)
    // Redirect to checkout if coming from there, else home
    if (route.query.redirect === 'checkout') {
      router.push('/checkout')
    } else {
      router.push('/')
    }
  } catch (err: any) {
    error.value = err.message || 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.input {
  @apply border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary-500;
}
.btn {
  @apply bg-primary-600 text-white font-semibold py-2 px-4 rounded hover:bg-primary-700 transition;
}
</style>
