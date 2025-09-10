<template>
  <transition name="cart-sidebar-fade">
    <div v-if="open">
      <!-- Overlay for click outside -->
      <div class="fixed inset-0 bg-black/30 z-40" @click="$emit('close')"></div>
      <aside
        ref="sidebarRef"
        class="fixed inset-y-0 right-0 w-96 max-w-full bg-white shadow-2xl z-50 flex flex-col"
        @keydown.esc="emitClose"
        tabindex="0"
      >
  <div class="flex items-center justify-between p-6 border-b">
        <h2 class="text-xl font-bold">Shopping Cart</h2>
        <button @click="$emit('close')" class="text-gray-500 hover:text-primary-600 text-2xl">&times;</button>
      </div>
  <div class="flex-1 overflow-y-auto p-6">
        <div v-if="cart.isEmpty" class="text-gray-500 text-center mt-12">Your cart is empty.</div>
        <div v-else>
          <div class="space-y-4">
            <div v-for="item in cartItemsWithImages" :key="item.id" class="flex items-center gap-4 border-b pb-4">
              <img :src="item.image_url || '/placeholder-product.jpg'" :alt="item.name" class="w-16 h-16 object-cover rounded border" />
              <div class="flex-1">
                <div class="font-medium">{{ item.name }}</div>
                <div class="text-sm text-gray-500">${{ item.price.toFixed(2) }}</div>
                <div class="flex items-center mt-2">
                  <button @click="cart.decrementItem(item.id)" class="btn btn-xs">-</button>
                  <span class="mx-2">{{ item.quantity }}</span>
                  <button @click="cart.incrementItem(item.id)" class="btn btn-xs">+</button>
                </div>
              </div>
              <div class="flex flex-col items-end">
                <span class="font-bold text-primary-600">${{ (item.price * item.quantity).toFixed(2) }}</span>
                <button @click="cart.removeItem(item.id)" class="btn btn-xs btn-danger mt-2">Remove</button>
              </div>
            </div>
          </div>
        </div>
      </div>
  <div class="border-t p-6">
        <div class="flex justify-between mb-2">
          <span class="font-medium">Total Items:</span>
          <span>{{ cart.itemCount }}</span>
        </div>
        <div class="flex justify-between mb-2">
          <span class="font-medium">Total Amount:</span>
          <span class="font-bold text-lg text-primary-600">${{ cart.totalAmount.toFixed(2) }}</span>
        </div>
  <button class="btn btn-primary w-full mt-4" :disabled="cart.isEmpty" @click="goToCheckout">Checkout</button>
      </div>
      </aside>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { useCartStore } from '../stores/cart'
import { useRouter } from 'vue-router'
import { ref, watch, onMounted, onBeforeUnmount, computed } from 'vue'

const router = useRouter()

function goToCheckout() {
  if (!cart.isEmpty) {
    router.push('/checkout')
    emit('close')
  }
}
const cart = useCartStore()
const props = defineProps<{ open: boolean }>()
const emit = defineEmits(['close'])

const sidebarRef = ref<HTMLElement | null>(null)

function emitClose() {
  emit('close')
}

// Focus sidebar for keydown events when open
watch(() => props.open, (open) => {
  if (open && sidebarRef.value) {
    sidebarRef.value.focus()
  }
})

// Escape key fallback for when sidebar is not focused
function handleEscape(e: KeyboardEvent) {
  if (props.open && e.key === 'Escape') {
    emitClose()
  }
}

onMounted(() => {
  window.addEventListener('keydown', handleEscape)
})
onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleEscape)
})

// Map product names to Unsplash images
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

const cartItemsWithImages = computed(() => {
  const baseURL = import.meta.env.VITE_API_BASE_URL
  return cart.cartItems.map(item => ({
    ...item,
    image_url: imageMap[item.name] || (baseURL + item.image_url)
  }))
})
</script>

<style scoped>
.cart-sidebar-fade-enter-active,
.cart-sidebar-fade-leave-active {
  transition: opacity 0.2s, transform 0.2s;
}
.cart-sidebar-fade-enter-from,
.cart-sidebar-fade-leave-to {
  opacity: 0;
  transform: translateX(100%);
}
.cart-sidebar-fade-enter-to,
.cart-sidebar-fade-leave-from {
  opacity: 1;
  transform: translateX(0);
}
</style>
