<template>
  <div class="p-6 max-w-md mx-auto">
    <h1 class="text-2xl font-bold mb-4">Register</h1>
    <form @submit.prevent="onRegister" class="space-y-4">
      <div>
        <label class="block text-sm font-medium mb-1" for="email">Email</label>
        <input v-model="email" id="email" type="email" required class="input input-bordered w-full" />
      </div>
      <div>
        <label class="block text-sm font-medium mb-1" for="fullName">Full Name</label>
        <input v-model="fullName" id="fullName" type="text" required class="input input-bordered w-full" />
      </div>
      <div>
        <label class="block text-sm font-medium mb-1" for="password">Password</label>
        <input v-model="password" id="password" type="password" required class="input input-bordered w-full" />
      </div>
      <button type="submit" class="btn btn-primary w-full" :disabled="loading">
        {{ loading ? 'Registering...' : 'Register' }}
      </button>
      <p v-if="error" class="text-red-500 text-sm mt-2">{{ error }}</p>
      <p class="text-sm mt-4">Already have an account? <router-link to="/login" class="text-primary-600 hover:underline">Login</router-link></p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '../services/api'

const email = ref('')
const fullName = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const router = useRouter()

const onRegister = async () => {
  loading.value = true
  error.value = ''
  try {
    const res = await authAPI.register({ email: email.value, password: password.value, full_name: fullName.value })
    localStorage.setItem('token', res.token)
    localStorage.setItem('user', JSON.stringify(res.user))
    router.push('/')
  } catch (err: any) {
    error.value = err.message || 'Registration failed'
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
