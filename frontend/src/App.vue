<template>
  <div id="app" class="min-h-screen bg-gray-50">
  <!-- Navigation -->
  <NavBar @open-cart="cartSidebarOpen = true" />
    
    <!-- Main Content -->
    <main class="flex-1">
      <RouterView v-slot="{ Component }">
        <Transition
          name="page"
          mode="out-in"
          enter-active-class="transition duration-300 ease-out"
          enter-from-class="opacity-0 transform translate-y-4"
          enter-to-class="opacity-100 transform translate-y-0"
          leave-active-class="transition duration-200 ease-in"
          leave-from-class="opacity-100 transform translate-y-0"
          leave-to-class="opacity-0 transform -translate-y-2"
        >
          <component :is="Component" />
        </Transition>
      </RouterView>
    </main>

  <!-- Cart Sidebar -->
  <CartSidebar :open="cartSidebarOpen" @close="cartSidebarOpen = false" />

  <!-- Toast Notifications -->
  <ToastContainer />
    
  <!-- Loading Overlay -->
  <LoadingOverlay v-if="isAppLoading" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { RouterView } from 'vue-router'
import { useAuthStore } from './stores/auth'
import { useCartStore } from './stores/cart'
import NavBar from './components/NavBar.vue'
import ToastContainer from './components/ToastContainer.vue'
import LoadingOverlay from './components/LoadingOverlay.vue'
import CartSidebar from './components/CartSidebar.vue'

// Stores
const authStore = useAuthStore()
const cartStore = useCartStore()

// Cart sidebar state
const cartSidebarOpen = ref(false)

// Computed
const isAppLoading = computed(() => authStore.loading)

// Initialize app
onMounted(async () => {
  // Initialize auth from localStorage
  authStore.initializeAuth()
  
  // Initialize cart from localStorage
  cartStore.initializeCart()
  
  // Fetch user profile if authenticated (to ensure data is fresh)
  if (authStore.isAuthenticated) {
    await authStore.fetchProfile()
  }
})
</script>

<style>
/* Global styles are imported in main.ts */
#app {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* Page transition animations */
.page-enter-active,
.page-leave-active {
  transition: all 0.3s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

/* Custom scrollbar for the entire app */
::-webkit-scrollbar {
  width: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
}

/* Ensure images don't break layout */
img {
  max-width: 100%;
  height: auto;
}

/* Focus styles for better accessibility */
button:focus-visible,
input:focus-visible,
select:focus-visible,
textarea:focus-visible {
  outline: 2px solid #3b82f6;
  outline-offset: 2px;
}
</style>