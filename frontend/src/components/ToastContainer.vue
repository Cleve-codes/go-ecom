<template>
  <Teleport to="body">
    <div 
      v-if="toasts.length > 0"
      class="fixed top-4 right-4 z-50 space-y-2"
      role="region" 
      aria-label="Notifications"
    >
      <TransitionGroup
        name="toast"
        enter-active-class="transition duration-300 ease-out"
        enter-from-class="transform translate-x-full opacity-0"
        enter-to-class="transform translate-x-0 opacity-100"
        leave-active-class="transition duration-200 ease-in"
        leave-from-class="transform translate-x-0 opacity-100"
        leave-to-class="transform translate-x-full opacity-0"
        move-class="transition duration-200"
      >
        <div
          v-for="toast in toasts"
          :key="toast.id"
          class="max-w-sm w-full bg-white shadow-lg rounded-lg pointer-events-auto overflow-hidden border-l-4"
          :class="toastClasses[toast.type]"
          role="alert"
        >
          <div class="p-4">
            <div class="flex items-start">
              <div class="flex-shrink-0">
                <component :is="toastIcons[toast.type]" class="w-5 h-5" />
              </div>
              <div class="ml-3 w-0 flex-1">
                <p class="text-sm font-medium text-gray-900">
                  {{ toast.title }}
                </p>
                <p v-if="toast.message" class="mt-1 text-sm text-gray-500">
                  {{ toast.message }}
                </p>
              </div>
              <div class="ml-4 flex-shrink-0 flex">
                <button
                  @click="removeToast(toast.id)"
                  class="bg-white rounded-md inline-flex text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
                >
                  <XMarkIcon class="w-5 h-5" />
                </button>
              </div>
            </div>
          </div>
          
          <!-- Progress bar -->
          <div 
            v-if="toast.duration && toast.duration > 0"
            class="h-1 bg-gray-200 overflow-hidden"
          >
            <div 
              class="h-full transition-all duration-linear"
              :class="progressBarClasses[toast.type]"
              :style="{ 
                width: '100%',
                animation: `shrink ${toast.duration}ms linear forwards`
              }"
            />
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import type { Toast } from '../types'
import {
  CheckCircleIcon,
  ExclamationCircleIcon,
  ExclamationTriangleIcon,
  InformationCircleIcon,
  XMarkIcon,
} from '@heroicons/vue/24/outline'

// State
const toasts = ref<Toast[]>([])

// Toast styling classes
const toastClasses = reactive({
  success: 'border-green-500',
  error: 'border-red-500',
  warning: 'border-yellow-500',
  info: 'border-blue-500',
})

const toastIcons = reactive({
  success: CheckCircleIcon,
  error: ExclamationCircleIcon,
  warning: ExclamationTriangleIcon,
  info: InformationCircleIcon,
})

const progressBarClasses = reactive({
  success: 'bg-green-500',
  error: 'bg-red-500',
  warning: 'bg-yellow-500',
  info: 'bg-blue-500',
})

// Methods
const addToast = (toast: Omit<Toast, 'id'>): string => {
  const id = Date.now().toString() + Math.random().toString(36).substr(2, 9)
  const newToast: Toast = {
    id,
    duration: 5000, // Default 5 seconds
    ...toast,
  }
  
  toasts.value.push(newToast)
  
  // Auto remove toast after duration
  if (newToast.duration && newToast.duration > 0) {
    setTimeout(() => {
      removeToast(id)
    }, newToast.duration)
  }
  
  return id
}

const removeToast = (id: string): void => {
  const index = toasts.value.findIndex(toast => toast.id === id)
  if (index > -1) {
    toasts.value.splice(index, 1)
  }
}

const removeAllToasts = (): void => {
  toasts.value = []
}

// Helper methods for different toast types
const showSuccess = (title: string, message?: string, duration = 5000): string => {
  return addToast({ type: 'success', title, message, duration })
}

const showError = (title: string, message?: string, duration = 7000): string => {
  return addToast({ type: 'error', title, message, duration })
}

const showWarning = (title: string, message?: string, duration = 6000): string => {
  return addToast({ type: 'warning', title, message, duration })
}

const showInfo = (title: string, message?: string, duration = 5000): string => {
    return addToast({ type: 'info', title, message, duration })
}

// Expose methods for use in app
</script>

<style scoped>
@keyframes shrink {
  from { width: 100%; }
  to { width: 0%; }
}
</style>