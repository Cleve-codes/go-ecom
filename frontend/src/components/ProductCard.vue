<template>
  <div class="card group cursor-pointer relative overflow-hidden" @click="navigateToProduct">
    <!-- Product Image -->
    <div class="relative overflow-hidden bg-gray-50 aspect-square">

      <img 
        :src="imageSrc"
        :alt="product.name"
        class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-700"
        @error="handleImageError"
        loading="lazy"
      />
      
      <!-- Stock Badge -->
      <div 
        v-if="product.stock <= 5 && product.stock > 0"
        class="absolute top-3 left-3 bg-gradient-to-r from-orange-500 to-orange-600 text-white text-xs px-3 py-1 rounded-full font-medium shadow-lg animate-pulse"
      >
        Only {{ product.stock }} left!
      </div>
      
      <div 
        v-else-if="product.stock === 0"
        class="absolute top-3 left-3 bg-gradient-to-r from-red-500 to-red-600 text-white text-xs px-3 py-1 rounded-full font-medium shadow-lg"
      >
        Out of Stock
      </div>

      <!-- Category Badge -->
      <div class="absolute top-3 right-3 bg-white/90 backdrop-blur-sm text-gray-700 text-xs px-2 py-1 rounded-full font-medium">
        {{ product.category }}
      </div>

      <!-- Overlay with Actions -->
      <div class="absolute inset-0 bg-gradient-to-t from-black/20 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <div class="absolute bottom-4 left-4 right-4 flex justify-between items-end">
          <!-- Quick Add Button -->
          <button 
            v-if="product.stock > 0"
            @click.stop="addToCart"
            class="btn btn-primary btn-sm transform translate-y-2 opacity-0 group-hover:translate-y-0 group-hover:opacity-100 transition-all duration-300 delay-75 shadow-xl"
            :disabled="addingToCart"
          >
            <PlusIcon class="w-4 h-4 mr-1" />
            {{ addingToCart ? 'Adding...' : 'Quick Add' }}
          </button>
          
          <!-- Favorite Button -->
          <button 
            @click.stop="toggleFavorite"
            class="p-2 bg-white/90 backdrop-blur-sm rounded-full text-gray-600 hover:text-red-500 transform translate-y-2 opacity-0 group-hover:translate-y-0 group-hover:opacity-100 transition-all duration-300 delay-100 shadow-lg"
          >
            <HeartIcon 
              class="w-5 h-5 transition-all duration-200" 
              :class="{ 
                'fill-red-500 text-red-500 scale-110': isFavorite,
                'hover:scale-110': !isFavorite 
              }"
            />
          </button>
        </div>
      </div>
    </div>

    <!-- Product Info -->
    <div class="p-4 bg-white">
      <!-- Product Name & Price -->
      <div class="flex items-start justify-between mb-2">
        <h3 class="font-semibold text-gray-900 text-lg leading-tight group-hover:text-primary-600 transition-colors duration-200 line-clamp-1">
          {{ product.name }}
        </h3>
        <div class="text-right ml-2 flex-shrink-0">
          <span class="text-2xl font-bold text-gray-900">
            ${{ formatPrice(product.price) }}
          </span>
        </div>
      </div>

      <!-- Description -->
      <p class="text-gray-600 text-sm mb-3 line-clamp-2 leading-relaxed">
        {{ product.description || 'No description available' }}
      </p>

      <!-- Stock & Rating -->
      <div class="flex items-center justify-between mb-3">
        <!-- Stock Indicator -->
        <div class="flex items-center space-x-2">
          <div 
            class="w-2 h-2 rounded-full"
            :class="stockIndicatorClass"
          ></div>
          <span class="text-xs text-gray-500 font-medium">
            {{ stockText }}
          </span>
        </div>

        <!-- Rating (placeholder for future implementation) -->
        <div class="flex items-center space-x-1">
          <div class="flex text-yellow-400">
            <StarIcon class="w-4 h-4 fill-current" />
            <StarIcon class="w-4 h-4 fill-current" />
            <StarIcon class="w-4 h-4 fill-current" />
            <StarIcon class="w-4 h-4 fill-current" />
            <StarIcon class="w-4 h-4 text-gray-300" />
          </div>
          <span class="text-xs text-gray-500">(4.2)</span>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center space-x-2">
        <button 
          v-if="product.stock > 0"
          @click.stop="addToCart"
          class="flex-1 btn btn-primary btn-sm font-medium"
          :disabled="addingToCart"
        >
          <ShoppingCartIcon class="w-4 h-4 mr-2" />
          {{ addingToCart ? 'Adding...' : 'Add to Cart' }}
        </button>
        
        <button 
          v-else
          class="flex-1 btn btn-sm bg-gray-100 text-gray-400 cursor-not-allowed"
          disabled
        >
          <XCircleIcon class="w-4 h-4 mr-2" />
          Out of Stock
        </button>

        <button 
          @click.stop="navigateToProduct"
          class="btn btn-secondary btn-sm px-3"
          title="View Details"
        >
          <EyeIcon class="w-4 h-4" />
        </button>
      </div>
    </div>

    <!-- Sale Badge (if on sale) -->
    <div 
      v-if="isOnSale"
      class="absolute top-0 left-0 bg-gradient-to-r from-green-500 to-green-600 text-white text-xs px-3 py-1 rounded-br-lg font-bold shadow-lg"
    >
      SALE
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import type { Product } from '../types'
import {
  PlusIcon,
  HeartIcon,
  ShoppingCartIcon,
  XCircleIcon,
  EyeIcon,
  StarIcon,
} from '@heroicons/vue/24/outline'

// Props
interface Props {
  product: Product
}

const props = defineProps<Props>()

// Composables
const router = useRouter()
const cartStore = useCartStore()

// State

const addingToCart = ref(false)
const isFavorite = ref(false) // This would come from a favorites store in a real app
const imageSrc = ref(props.product.image_url || '/placeholder-product.jpg')

// Computed
const stockIndicatorClass = computed(() => {
  if (props.product.stock > 10) return 'bg-green-500'
  if (props.product.stock > 0) return 'bg-orange-500'
  return 'bg-red-500'
})

const stockText = computed(() => {
  if (props.product.stock === 0) return 'Out of stock'
  if (props.product.stock <= 5) return `Only ${props.product.stock} left`
  return 'In stock'
})

const isOnSale = computed(() => {
  // This would be based on actual sale logic
  return Math.random() > 0.8 // Placeholder logic
})

// Methods
const formatPrice = (price: number): string => {
  return price.toFixed(2)
}

const addToCart = async (): Promise<void> => {
  if (addingToCart.value || props.product.stock === 0) return

  addingToCart.value = true
  
  try {
    const success = cartStore.addItem(props.product)
    if (!success) {
      // Could show a toast notification here
      console.warn('Could not add item to cart - insufficient stock')
    }
    // Optional: Show success feedback
  } catch (error) {
    console.error('Error adding to cart:', error)
    // Could show error toast here
  } finally {
    // Small delay for better UX
    setTimeout(() => {
      addingToCart.value = false
    }, 500)
  }
}

const toggleFavorite = (): void => {
  isFavorite.value = !isFavorite.value
  // In a real app, this would sync with a favorites store/API
}

const navigateToProduct = (): void => {
  router.push(`/products/${props.product.id}`)
}

const handleImageError = (): void => {
  imageSrc.value = '/placeholder-product.jpg'
}
</script>

<style scoped>
.line-clamp-1 {
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 1;
  -webkit-box-orient: vertical;
}

.line-clamp-2 {
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}
</style>