<template>
  <nav class="bg-white/95 backdrop-blur-lg border-b border-gray-200/50 sticky top-0 z-50 transition-all duration-300">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        
        <!-- Logo -->
        <div class="flex items-center space-x-4">
          <RouterLink 
            to="/" 
            class="flex items-center space-x-2 text-xl font-bold text-gray-900 hover:text-primary-600 transition-colors"
          >
            <div class="w-8 h-8 bg-gradient-to-br from-primary-500 to-primary-600 rounded-lg flex items-center justify-center shadow-lg">
              <ShoppingBagIcon class="w-5 h-5 text-white" />
            </div>
            <span class="hidden sm:block bg-gradient-to-r from-primary-600 to-primary-700 bg-clip-text text-transparent">
              ShopVue
            </span>
          </RouterLink>
        </div>

        <!-- Desktop Navigation -->
        <div class="hidden md:flex items-center space-x-8">
          <RouterLink 
            to="/products" 
            class="text-gray-700 hover:text-primary-600 font-medium transition-colors relative group"
          >
            Products
            <span class="absolute inset-x-0 -bottom-1 h-0.5 bg-primary-500 scale-x-0 group-hover:scale-x-100 transition-transform origin-left"></span>
          </RouterLink>
          
          <RouterLink 
            v-if="isAdmin" 
            to="/admin" 
            class="text-gray-700 hover:text-primary-600 font-medium transition-colors relative group"
          >
            Admin
            <span class="absolute inset-x-0 -bottom-1 h-0.5 bg-primary-500 scale-x-0 group-hover:scale-x-100 transition-transform origin-left"></span>
          </RouterLink>
        </div>

        <!-- Right Side -->
        <div class="flex items-center space-x-4">
          
          <!-- Search Button (Mobile) -->
          <button 
            class="md:hidden p-2 text-gray-700 hover:text-primary-600 transition-colors"
            @click="showSearch = !showSearch"
          >
            <MagnifyingGlassIcon class="w-5 h-5" />
          </button>

          <!-- Cart Button -->
          <button 
            @click="$emit('open-cart')"
            class="relative p-2 text-gray-700 hover:text-primary-600 transition-all duration-200 hover:scale-110"
          >
            <ShoppingCartIcon class="w-6 h-6" />
            <Transition name="scale">
              <span 
                v-if="cartItemCount > 0"
                class="absolute -top-1 -right-1 bg-gradient-to-r from-primary-500 to-primary-600 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center font-medium shadow-lg"
              >
                {{ cartItemCount > 99 ? '99+' : cartItemCount }}
              </span>
            </Transition>
          </button>

          <!-- User Menu -->
          <div v-if="isAuthenticated" class="relative" ref="userMenuRef">
            <button 
              @click="showUserMenu = !showUserMenu"
              class="flex items-center space-x-2 p-2 text-gray-700 hover:text-primary-600 transition-colors rounded-lg hover:bg-gray-50"
            >
              <div class="w-8 h-8 bg-gradient-to-br from-primary-400 to-primary-600 rounded-full flex items-center justify-center text-white font-medium text-sm shadow-lg">
                {{ userInitials }}
              </div>
              <span class="hidden sm:block font-medium">{{ user?.full_name }}</span>
              <ChevronDownIcon 
                class="w-4 h-4 transition-transform duration-200" 
                :class="{ 'rotate-180': showUserMenu }"
              />
            </button>

            <!-- Dropdown Menu -->
            <Transition
              enter-active-class="transition duration-200 ease-out"
              enter-from-class="transform scale-95 opacity-0"
              enter-to-class="transform scale-100 opacity-100"
              leave-active-class="transition duration-75 ease-in"
              leave-from-class="transform scale-100 opacity-100"
              leave-to-class="transform scale-95 opacity-0"
            >
              <div 
                v-if="showUserMenu"
                class="absolute right-0 mt-2 w-56 bg-white rounded-xl shadow-xl border border-gray-100 py-2 z-50 overflow-hidden"
              >
                <div class="px-4 py-3 border-b border-gray-100">
                  <p class="text-sm font-medium text-gray-900">{{ user?.full_name }}</p>
                  <p class="text-sm text-gray-500 truncate">{{ user?.email }}</p>
                </div>
                
                <div class="py-1">
                  <RouterLink 
                    to="/profile" 
                    @click="showUserMenu = false"
                    class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors"
                  >
                    <UserIcon class="w-4 h-4 mr-3 text-gray-400" />
                    Profile
                  </RouterLink>
                  <RouterLink 
                    to="/orders" 
                    @click="showUserMenu = false"
                    class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors"
                  >
                    <ClipboardDocumentListIcon class="w-4 h-4 mr-3 text-gray-400" />
                    My Orders
                  </RouterLink>
                  <RouterLink 
                    v-if="isAdmin"
                    to="/admin" 
                    @click="showUserMenu = false"
                    class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors"
                  >
                    <Cog6ToothIcon class="w-4 h-4 mr-3 text-gray-400" />
                    Admin Dashboard
                  </RouterLink>
                </div>
                
                <hr class="my-1">
                
                <button 
                  @click="handleLogout"
                  class="flex items-center w-full px-4 py-2 text-sm text-red-600 hover:bg-red-50 transition-colors"
                >
                  <ArrowLeftOnRectangleIcon class="w-4 h-4 mr-3" />
                  Sign Out
                </button>
              </div>
            </Transition>
          </div>

          <!-- Auth Buttons -->
          <div v-else class="flex items-center space-x-2">
            <RouterLink 
              to="/login"
              class="btn btn-secondary btn-sm font-medium"
            >
              Sign In
            </RouterLink>
            <RouterLink 
              to="/register"
              class="btn btn-primary btn-sm font-medium"
            >
              Sign Up
            </RouterLink>
          </div>

          <!-- Mobile Menu Button -->
          <button 
            @click="showMobileMenu = !showMobileMenu"
            class="md:hidden p-2 text-gray-700 hover:text-primary-600 transition-colors"
          >
            <Bars3Icon v-if="!showMobileMenu" class="w-6 h-6" />
            <XMarkIcon v-else class="w-6 h-6" />
          </button>
        </div>
      </div>
    </div>

    <!-- Mobile Search Bar -->
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform -translate-y-2 opacity-0"
      enter-to-class="transform translate-y-0 opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="transform translate-y-0 opacity-100"
      leave-to-class="transform -translate-y-2 opacity-0"
    >
      <div v-if="showSearch" class="md:hidden px-4 py-3 border-t border-gray-200">
        <div class="relative">
          <MagnifyingGlassIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
          <input
            type="text"
            placeholder="Search products..."
            class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
            v-model="searchQuery"
            @keyup.enter="handleSearch"
          />
        </div>
      </div>
    </Transition>

    <!-- Mobile Menu -->
    <Transition
      enter-active-class="transition duration-200 ease-out"
      enter-from-class="transform -translate-y-2 opacity-0"
      enter-to-class="transform translate-y-0 opacity-100"
      leave-active-class="transition duration-150 ease-in"
      leave-from-class="transform translate-y-0 opacity-100"
      leave-to-class="transform -translate-y-2 opacity-0"
    >
      <div v-if="showMobileMenu" class="md:hidden border-t border-gray-200 bg-white/95 backdrop-blur-lg">
        <div class="px-4 py-2 space-y-1">
          <RouterLink 
            to="/products" 
            @click="showMobileMenu = false"
            class="block px-3 py-2 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors font-medium"
          >
            Products
          </RouterLink>
          
          <RouterLink 
            v-if="isAdmin" 
            to="/admin" 
            @click="showMobileMenu = false"
            class="block px-3 py-2 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors font-medium"
          >
            Admin Dashboard
          </RouterLink>

          <div v-if="isAuthenticated" class="pt-2 border-t border-gray-200 mt-2">
            <div class="px-3 py-2 text-sm text-gray-500">
              Signed in as {{ user?.full_name }}
            </div>
            <RouterLink 
              to="/profile" 
              @click="showMobileMenu = false"
              class="block px-3 py-2 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors"
            >
              Profile
            </RouterLink>
            <RouterLink 
              to="/orders" 
              @click="showMobileMenu = false"
              class="block px-3 py-2 text-gray-700 hover:bg-gray-50 rounded-lg transition-colors"
            >
              My Orders
            </RouterLink>
            <button 
              @click="handleLogout"
              class="w-full text-left px-3 py-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors mt-1"
            >
              Sign Out
            </button>
          </div>

          <div v-else class="pt-2 border-t border-gray-200 mt-2 space-y-2">
            <RouterLink 
              to="/login" 
              @click="showMobileMenu = false"
              class="block w-full text-center btn btn-secondary btn-sm"
            >
              Sign In
            </RouterLink>
            <RouterLink 
              to="/register" 
              @click="showMobileMenu = false"
              class="block w-full text-center btn btn-primary btn-sm"
            >
              Sign Up
            </RouterLink>
          </div>
        </div>
      </div>
    </Transition>
  </nav>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useCartStore } from '../stores/cart'
import {
  ShoppingBagIcon,
  ShoppingCartIcon,
  UserIcon,
  ChevronDownIcon,
  Bars3Icon,
  XMarkIcon,
  MagnifyingGlassIcon,
  ClipboardDocumentListIcon,
  Cog6ToothIcon,
  ArrowLeftOnRectangleIcon,
} from '@heroicons/vue/24/outline'

// Stores
const authStore = useAuthStore()
const cartStore = useCartStore()
const router = useRouter()

// Reactive state
const showUserMenu = ref(false)
const showMobileMenu = ref(false)
const showSearch = ref(false)
const searchQuery = ref('')
const userMenuRef = ref<HTMLElement | null>(null)

// Computed properties
const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const user = computed(() => authStore.user)
const userInitials = computed(() => authStore.userInitials)
const cartItemCount = computed(() => cartStore.itemCount)

// Methods
const handleLogout = (): void => {
  authStore.logout()
  showUserMenu.value = false
  showMobileMenu.value = false
  router.push('/')
}

const handleSearch = (): void => {
  if (searchQuery.value.trim()) {
    router.push({ path: '/products', query: { search: searchQuery.value.trim() } })
    showSearch.value = false
    showMobileMenu.value = false
  }
}

// Close dropdown when clicking outside
const handleClickOutside = (event: Event): void => {
  if (userMenuRef.value && !userMenuRef.value.contains(event.target as Node)) {
    showUserMenu.value = false
  }
}

// Close mobile menu on route change
router.afterEach(() => {
  showMobileMenu.value = false
  showSearch.value = false
})

// Lifecycle hooks
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  if (authStore.initializeAuth) authStore.initializeAuth()
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.router-link-active {
  @apply text-primary-600;
}

.scale-enter-active,
.scale-leave-active {
  transition: transform 0.2s ease-in-out;
}

.scale-enter-from,
.scale-leave-to {
  transform: scale(0);
}
</style>