<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <section class="relative bg-gradient-to-br from-primary-50 via-white to-primary-100 overflow-hidden">
      <div class="absolute inset-0 bg-hero-pattern opacity-5"></div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 sm:py-28">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
          
          <!-- Hero Content -->
          <div class="text-center lg:text-left animate-fade-in">
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold text-gray-900 leading-tight">
              Modern
              <span class="bg-gradient-to-r from-primary-600 to-primary-700 bg-clip-text text-transparent">
                Shopping
              </span>
              <br />Experience
            </h1>
            <p class="mt-6 text-lg sm:text-xl text-gray-600 leading-relaxed max-w-2xl">
              Discover amazing products with seamless shopping experience. 
              Fast delivery, secure payments, and exceptional customer service.
            </p>
            
            <div class="mt-8 flex flex-col sm:flex-row gap-4 justify-center lg:justify-start">
              <RouterLink 
                to="/products"
                class="btn btn-primary btn-lg group"
              >
                <ShoppingBagIcon class="w-5 h-5 mr-2 group-hover:animate-bounce" />
                Shop Now
              </RouterLink>
              
              <button 
                @click="scrollToFeatures"
                class="btn btn-secondary btn-lg group"
              >
                <PlayCircleIcon class="w-5 h-5 mr-2" />
                Learn More
              </button>
            </div>

            <!-- Stats -->
            <div class="mt-12 grid grid-cols-3 gap-6 text-center lg:text-left">
              <div class="animate-slide-up" style="animation-delay: 0.2s">
                <div class="text-3xl font-bold text-primary-600">10K+</div>
                <div class="text-sm text-gray-600 font-medium">Happy Customers</div>
              </div>
              <div class="animate-slide-up" style="animation-delay: 0.4s">
                <div class="text-3xl font-bold text-primary-600">500+</div>
                <div class="text-sm text-gray-600 font-medium">Products</div>
              </div>
              <div class="animate-slide-up" style="animation-delay: 0.6s">
                <div class="text-3xl font-bold text-primary-600">99.9%</div>
                <div class="text-sm text-gray-600 font-medium">Uptime</div>
              </div>
            </div>
          </div>

          <!-- Hero Product Banner -->
          <div class="relative animate-fade-in" style="animation-delay: 0.3s">
            <div class="relative mx-auto w-full max-w-lg bg-white rounded-2xl shadow-2xl p-6 flex flex-col items-center">
              <img v-if="heroProduct && heroProduct.image_url" :src="heroProduct.image_url" :alt="heroProduct.name" class="w-60 h-60 object-contain rounded-xl mb-4 border" />
              <div v-else class="w-60 h-60 bg-gray-100 rounded-xl mb-4 flex items-center justify-center">
                <ShoppingCartIcon class="w-20 h-20 text-primary-500" />
              </div>
              <h2 v-if="heroProduct" class="text-2xl font-bold text-gray-900 mb-2">{{ heroProduct.name }}</h2>
              <p v-if="heroProduct" class="text-lg text-primary-600 font-semibold mb-2">${{ heroProduct.price.toFixed(2) }}</p>
              <p v-if="heroProduct" class="text-gray-600 text-center mb-4 line-clamp-2">{{ heroProduct.description }}</p>
              <RouterLink v-if="heroProduct" :to="`/products/${heroProduct.id}`" class="btn btn-primary">View Product</RouterLink>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section ref="featuresSection" class="py-20 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <h2 class="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
            Why Choose ShopVue?
          </h2>
          <p class="text-lg text-gray-600 max-w-2xl mx-auto">
            We're committed to providing the best shopping experience with cutting-edge technology and customer-first approach.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div 
            v-for="(feature, index) in features" 
            :key="feature.title"
            class="text-center group hover:transform hover:scale-105 transition-all duration-300"
            :style="{ animationDelay: `${index * 0.2}s` }"
          >
            <div class="w-16 h-16 bg-gradient-to-br from-primary-500 to-primary-600 rounded-2xl flex items-center justify-center mx-auto mb-6 group-hover:shadow-xl transition-shadow">
              <component :is="feature.icon" class="w-8 h-8 text-white" />
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-3">{{ feature.title }}</h3>
            <p class="text-gray-600 leading-relaxed">{{ feature.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Products Section -->
    <section class="py-20 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between mb-12">
          <div>
            <h2 class="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
              Featured Products
            </h2>
            <p class="text-lg text-gray-600">
              Discover our most popular items
            </p>
          </div>
          <RouterLink 
            to="/products"
            class="btn btn-secondary group hidden sm:flex"
          >
            View All
            <ArrowRightIcon class="w-4 h-4 ml-2 group-hover:translate-x-1 transition-transform" />
          </RouterLink>
        </div>

        <!-- Products Grid -->
        <div v-if="loading" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div v-for="i in 4" :key="i" class="bg-white rounded-xl p-4 animate-pulse">
            <div class="aspect-square bg-gray-200 rounded-lg mb-4"></div>
            <div class="h-4 bg-gray-200 rounded mb-2"></div>
            <div class="h-4 bg-gray-200 rounded w-2/3"></div>
          </div>
        </div>

        <div v-else-if="featuredProducts.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <ProductCard 
            v-for="product in featuredProducts" 
            :key="product.id" 
            :product="product"
            class="animate-fade-in"
          />
        </div>

        <div v-else class="text-center py-12">
          <ShoppingBagIcon class="w-16 h-16 text-gray-400 mx-auto mb-4" />
          <h3 class="text-lg font-medium text-gray-900 mb-2">No products available</h3>
          <p class="text-gray-600">Check back later for amazing products!</p>
        </div>

        <!-- Mobile View All Button -->
        <div class="text-center mt-8 sm:hidden">
          <RouterLink to="/products" class="btn btn-primary">
            View All Products
          </RouterLink>
        </div>
      </div>
    </section>

    <!-- Newsletter Section -->
    <section class="py-20 bg-gradient-to-r from-primary-600 to-primary-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <div class="max-w-2xl mx-auto">
          <h2 class="text-3xl sm:text-4xl font-bold text-white mb-4">
            Stay Updated
          </h2>
          <p class="text-lg text-primary-100 mb-8">
            Get the latest updates on new products, special offers, and exclusive deals.
          </p>

          <div class="flex flex-col sm:flex-row gap-4 max-w-md mx-auto">
            <input 
              v-model="newsletterEmail"
              type="email" 
              placeholder="Enter your email"
              class="flex-1 px-4 py-3 rounded-lg border-0 focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-primary-600"
            />
            <button 
              @click="subscribeNewsletter"
              class="btn bg-white text-primary-600 hover:bg-gray-50 px-8 py-3 font-medium"
              :disabled="!newsletterEmail || subscribing"
            >
              {{ subscribing ? 'Subscribing...' : 'Subscribe' }}
            </button>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'
import { productsAPI } from '../services/api'
import type { Product } from '../types'
import ProductCard from '../components/ProductCard.vue'
import {
  ShoppingBagIcon,
  ShoppingCartIcon,
  HeartIcon,
  PlayCircleIcon,
  ArrowRightIcon,
  TruckIcon,
  ShieldCheckIcon,
  ChatBubbleLeftRightIcon,
} from '@heroicons/vue/24/outline'

// State
const featuredProducts = ref<Product[]>([])
const heroProduct = ref<Product | null>(null)
const loading = ref(true)
const newsletterEmail = ref('')
const subscribing = ref(false)
const featuresSection = ref<HTMLElement | null>(null)

// Features data
const features = [
  {
    title: 'Fast Delivery',
    description: 'Get your orders delivered quickly with our express shipping options.',
    icon: TruckIcon,
  },
  {
    title: 'Secure Payment',
    description: 'Shop with confidence using our secure M-Pesa integration.',
    icon: ShieldCheckIcon,
  },
  {
    title: '24/7 Support',
    description: 'Our customer support team is always here to help you.',
    icon: ChatBubbleLeftRightIcon,
  },
]

// Image mapping (updated with product names from DB)
const imageMap: Record<string, string> = {
  'Apple Watch Series 9': 'https://images.unsplash.com/photo-1434494878577-86c23bcb06b9?auto=format&fit=crop&w=600&q=80',
  'LG C3 55" OLED': 'https://images.unsplash.com/photo-1593359677879-a4bb92f829d1?auto=format&fit=crop&w=600&q=80',
  'Apple iPad Air (2024)': 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?auto=format&fit=crop&w=600&q=80',
  'Samsung Galaxy Tab S9': 'https://images.unsplash.com/photo-1544244015-0df4b3ffc6b0?auto=format&fit=crop&w=600&q=80',
  'Sony WH-1000XM5': 'https://images.unsplash.com/photo-1511367461989-f85a21fda167?auto=format&fit=crop&w=600&q=80',
  'iPhone 15 Pro': 'https://images.unsplash.com/photo-1592750475338-74b7b21085ab?auto=format&fit=crop&w=600&q=80',
  'Samsung Galaxy S24 Ultra': 'https://images.unsplash.com/photo-1592750475338-74b7b21085ab?auto=format&fit=crop&w=600&q=80',
  'MacBook Pro 16" M3': 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=600&q=80',
  'Dell XPS 13': 'https://images.unsplash.com/photo-1517336714731-489689fd1ca8?auto=format&fit=crop&w=600&q=80',
  'Sony Bravia XR 65" OLED': 'https://images.unsplash.com/photo-1593359677879-a4bb92f829d1?auto=format&fit=crop&w=600&q=80',
}

// Methods
const loadFeaturedProducts = async (): Promise<void> => {
  try {
    loading.value = true
    const response = await productsAPI.getAll({ limit: 8 })
    const arr = Array.isArray(response) ? response : response.products
    const baseURL = import.meta.env.VITE_API_BASE_URL
    featuredProducts.value = arr.map((p: any) => ({
      ...p,
      image_url: imageMap[p.name] || (baseURL + p.image_url)
    }))
  } catch (error) {
    console.error('Failed to load featured products:', error)
  } finally {
    loading.value = false
  }
}

const scrollToFeatures = (): void => {
  featuresSection.value?.scrollIntoView({ behavior: 'smooth' })
}

const subscribeNewsletter = async (): Promise<void> => {
  if (!newsletterEmail.value) return
  
  subscribing.value = true
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    console.log('Subscribed:', newsletterEmail.value)
    newsletterEmail.value = ''
  } catch (error) {
    console.error('Failed to subscribe:', error)
  } finally {
    subscribing.value = false
  }
}

// Lifecycle
onMounted(async () => {
  await loadFeaturedProducts()
  // Pick the first featured product for the hero
  const baseURL = import.meta.env.VITE_API_BASE_URL
  heroProduct.value = featuredProducts.value.length > 0
    ? { ...featuredProducts.value[0], image_url: imageMap[featuredProducts.value[0].name] || (baseURL + featuredProducts.value[0].image_url) }
    : null
})
</script>

<style scoped>
@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}
</style>