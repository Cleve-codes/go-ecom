<template>
  <Teleport to="body">
    <div
      class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50"
      role="dialog"
      aria-label="Loading"
    >
      <div class="bg-white rounded-2xl p-8 shadow-2xl flex flex-col items-center space-y-4 max-w-sm mx-4">
        <!-- Loading Spinner -->
        <div class="relative">
          <div class="w-16 h-16 border-4 border-primary-200 border-t-primary-600 rounded-full animate-spin"></div>
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="w-6 h-6 bg-primary-600 rounded-full animate-pulse"></div>
          </div>
        </div>
        
        <!-- Loading Text -->
        <div class="text-center">
          <h3 class="text-lg font-semibold text-gray-900 mb-1">
            {{ title }}
          </h3>
          <p class="text-sm text-gray-600">
            {{ message }}
          </p>
        </div>
        
        <!-- Progress Bar (optional) -->
        <div v-if="showProgress" class="w-full">
          <div class="w-full bg-gray-200 rounded-full h-2">
            <div 
              class="bg-gradient-to-r from-primary-500 to-primary-600 h-2 rounded-full transition-all duration-300"
              :style="{ width: `${progress}%` }"
            ></div>
          </div>
          <p class="text-xs text-gray-500 mt-1 text-center">{{ progress }}% complete</p>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
interface Props {
  title?: string
  message?: string
  showProgress?: boolean
  progress?: number
}

withDefaults(defineProps<Props>(), {
  title: 'Loading...',
  message: 'Please wait while we process your request',
  showProgress: false,
  progress: 0,
})
</script>

<style scoped>
/* Additional loading animations */
@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}
</style>