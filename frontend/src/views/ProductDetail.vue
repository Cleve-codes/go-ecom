<template>
  <div class="container mx-auto py-8 px-4 max-w-5xl">
    <div v-if="loading" class="flex justify-center items-center h-64">
      <span class="text-lg text-gray-500">Loading product...</span>
    </div>
    <div v-else-if="error" class="flex flex-col items-center justify-center h-64">
      <span class="text-red-500 text-lg font-semibold mb-2">{{ error }}</span>
      <RouterLink to="/products" class="btn btn-primary mt-2">Back to Products</RouterLink>
    </div>
    <div v-else-if="product" class="flex flex-col md:flex-row gap-8 bg-white rounded-xl shadow-lg p-6">
      <div class="flex-shrink-0 w-full md:w-1/2 flex items-center justify-center">
        <img
          :src="product.image_url"
          :alt="product.name"
          class="rounded-xl object-contain max-h-96 w-full bg-gray-100 border"
          @error="onImgError"
        />
      </div>
      <div class="flex-1 flex flex-col gap-4">
        <h2 class="text-2xl font-bold text-gray-900">{{ product.name }}</h2>
        <p class="text-lg text-primary-600 font-semibold">Ksh {{ product.price.toLocaleString() }}</p>
        <p class="text-gray-700">{{ product.description }}</p>
        <div class="flex items-center gap-4 mt-4">
          <button
            class="btn btn-primary px-6 py-2"
            :disabled="addingToCart"
            @click="addToCart"
          >
            <span v-if="addingToCart">Adding...</span>
            <span v-else>Add to Cart</span>
          </button>
        </div>
      </div>
    </div>
    <div v-if="relatedProducts.length" class="mt-12">
      <h3 class="text-xl font-bold mb-4 text-gray-900">Related Products</h3>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div v-for="item in relatedProducts" :key="item.id" class="bg-white rounded-xl shadow p-4 flex flex-col items-center hover:shadow-lg transition">
          <img :src="item.image_url" :alt="item.name" class="w-32 h-32 object-contain rounded mb-3 border bg-gray-50" />
          <div class="font-semibold text-gray-900 mb-1">{{ item.name }}</div>
          <div class="text-primary-600 font-bold mb-2">Ksh {{ item.price.toLocaleString() }}</div>
          <RouterLink :to="`/products/${item.id}`" class="btn btn-sm btn-secondary">View</RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch, computed } from 'vue'
import { useRoute, useRouter, RouterLink } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { productsAPI } from '../services/api'

const route = useRoute()
const router = useRouter()
const cartStore = useCartStore()

const product = ref<any>(null)
const loading = ref(true)
const error = ref<string | null>(null)
const addingToCart = ref(false)
const relatedProducts = ref<any[]>([])
const imageMap: Record<string, string> = {
  'Laptop Pro': 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=600&q=80',
  'Wireless Headphones': 'https://images.unsplash.com/photo-1511367461989-f85a21fda167?auto=format&fit=crop&w=600&q=80',
  'Coffee Mug': 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=600&q=80',
  'Running Shoes': 'https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=600&q=80',
}

const fetchProduct = async (id: string | number) => {
  loading.value = true
  error.value = null
  try {
    const res = await productsAPI.getById(id as string)
    product.value = {
      ...res,
      image_url: imageMap[res.name] || res.image_url
    }
    await fetchRelatedProducts(res.id)
  } catch (e: any) {
    error.value = e?.message || 'Product not found.'
  } finally {
    loading.value = false
  }
}

const fetchRelatedProducts = async (currentId: string) => {
  try {
    const res = await productsAPI.getAll()
    // Exclude current product
    relatedProducts.value = res.filter((p: any) => p.id !== currentId).map((p: any) => ({
      ...p,
      image_url: imageMap[p.name] || p.image_url
    }))
  } catch (e) {
    relatedProducts.value = []
  }
}

const addToCart = async () => {
  if (!product.value) return
  addingToCart.value = true
  try {
    cartStore.addToCart(product.value)
  } finally {
    addingToCart.value = false
  }
}

const onImgError = (e: Event) => {
  const target = e.target as HTMLImageElement
  target.src = '/public/favicon.ico'
}

onMounted(() => {
  fetchProduct(route.params.id as string)
})

watch(() => route.params.id, (id) => {
  fetchProduct(id as string)
})
</script>

<style scoped>
</style>
