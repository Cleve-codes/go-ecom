<template>
  <div class="p-6 max-w-xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">My Profile</h1>
    <div v-if="loading" class="text-blue-600 mb-4">Loading profile...</div>
    <div v-if="error" class="text-red-600 mb-4">{{ error }}</div>
    <div v-if="user" class="bg-white rounded shadow p-6">
      <div class="mb-4">
        <span class="font-semibold">Full Name:</span> {{ user.full_name }}
      </div>
      <div class="mb-4">
        <span class="font-semibold">Email:</span> {{ user.email }}
      </div>
      <div class="mb-4">
        <span class="font-semibold">Role:</span> {{ user.role }}
      </div>
      <div class="mb-4">
        <span class="font-semibold">Joined:</span> {{ user.created_at.split('T')[0] }}
      </div>
      <button @click="logout" class="bg-red-600 text-white px-4 py-2 rounded">Logout</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { authAPI } from '@/services/api'

const auth = useAuthStore()
const router = useRouter()
const user = ref(auth.user)
const loading = ref(false)
const error = ref('')

const fetchProfile = async () => {
  loading.value = true
  error.value = ''
  try {
    const res = await authAPI.getProfile()
    user.value = res.user
  } catch (e: any) {
    error.value = e.message || 'Failed to load profile.'
  }
  loading.value = false
}

const logout = () => {
  auth.logout()
  router.replace('/login')
}

onMounted(() => {
  if (!auth.isAuthenticated) {
    router.replace('/login')
    return
  }
  fetchProfile()
})
</script>

<style scoped>
</style>
