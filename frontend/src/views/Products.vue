<template>
  <div class="container mx-auto py-8 px-4">
    <h1 class="text-2xl font-bold mb-6">Products</h1>
    <div v-if="loading" class="text-center py-12 text-gray-500">Loading products...</div>
    <div v-else-if="error" class="text-center py-12 text-red-500">{{ error }}</div>
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
      <ProductCard
        v-for="product in products"
        :key="product.id"
        :product="product"
        @click="goToProduct(product.id)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { productsAPI } from '../services/api'
import ProductCard from '../components/ProductCard.vue'

const products = ref([])
const loading = ref(true)
const error = ref('')
const router = useRouter()

const imageMap: Record<string, string> = {
  'Laptop Pro': 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=600&q=80',
  'Wireless Headphones': 'https://images.unsplash.com/photo-1511367461989-f85a21fda167?auto=format&fit=crop&w=600&q=80',
  'Coffee Mug': 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=600&q=80',
  'Running Shoes': 'https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=600&q=80',
}

const goToProduct = (id: string) => {
  router.push(`/products/${id}`)
}

onMounted(async () => {
  loading.value = true
  error.value = ''
  try {
    const res = await productsAPI.getAll()
    products.value = (res.products || []).map((p: any) => ({
      ...p,
      image_url: imageMap[p.name] || p.image_url
    }))
  } catch (err: any) {
    error.value = err.message || 'Failed to load products.'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
</style>
