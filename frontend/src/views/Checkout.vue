
<template>
  <div class="p-6 max-w-xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">Checkout</h1>
    <div v-if="cart.isEmpty">
      <p class="text-gray-500">Your cart is empty.</p>
    </div>
    <form v-else @submit.prevent="handleCheckout" class="space-y-6 bg-white rounded shadow p-6">
      <div>
        <h2 class="text-lg font-semibold mb-2">Order Summary</h2>
        <ul class="mb-2">
          <li v-for="item in cart.cartItems" :key="item.id" class="flex justify-between border-b py-1">
            <span>{{ item.name }} <span class="text-xs text-gray-400">x{{ item.quantity }}</span></span>
            <span>KES {{ (item.price * item.quantity).toLocaleString() }}</span>
          </li>
        </ul>
        <div class="flex justify-between font-bold text-lg">
          <span>Total</span>
          <span>KES {{ cart.totalAmount.toLocaleString() }}</span>
        </div>
      </div>
      <div>
        <label class="block mb-1 font-medium">Shipping Address</label>
        <input v-model="shippingAddress" type="text" required placeholder="Enter your shipping address" class="border rounded px-3 py-2 w-full mb-3" />
        <label class="block mb-1 font-medium">Phone Number (M-Pesa)</label>
        <input v-model="phone" type="tel" required pattern="^07\d{8}$" maxlength="10" placeholder="07XXXXXXXX" class="border rounded px-3 py-2 w-full" />
      </div>
      <div v-if="error" class="text-red-600">{{ error }}</div>
      <div v-if="success" class="text-green-600">{{ success }}</div>
      <button type="submit" :disabled="loading" class="bg-green-600 text-white px-6 py-2 rounded w-full flex items-center justify-center">
        <span v-if="loading" class="loader mr-2"></span>
        Pay with M-Pesa
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useCartStore } from '@/stores/cart'
import { mpesaAPI, ordersAPI } from '@/services/api'

const auth = useAuthStore()
const cart = useCartStore()
const router = useRouter()

const phone = ref('')
const shippingAddress = ref('')
const loading = ref(false)
const error = ref('')
const success = ref('')

// Redirect if not authenticated (extra guard)
if (!auth.isAuthenticated) {
  router.replace('/login')
}

const handleCheckout = async () => {
  error.value = ''
  success.value = ''
  if (!/^07\d{8}$/.test(phone.value)) {
    error.value = 'Please enter a valid Safaricom number (07XXXXXXXX)'
    return
  }
  if (!shippingAddress.value) {
    error.value = 'Shipping address is required.'
    return
  }
  if (cart.isEmpty) {
    error.value = 'Your cart is empty.'
    return
  }
  loading.value = true
  try {
    // 1. Create order in backend
    const orderPayload = {
      shipping_address: shippingAddress.value,
      phone_number: phone.value,
      items: cart.cartItems.map(item => ({
        product_id: item.id,
        quantity: item.quantity,
        unit_price: item.price
      }))
    }
    await ordersAPI.create(orderPayload)
    // 2. Initiate M-Pesa payment (optional, can be after order creation)
    // const paymentRes = await mpesaAPI.initiateSTK({
    //   order_id: undefined, // If you want to create order first, set this
    //   amount: cart.totalAmount,
    //   phone_number: phone.value
    // })
    // 3. On success, clear cart and show confirmation
    cart.items = []
    cart.saveToStorage()
    success.value = 'Order placed! Please complete the payment on your phone.'
    setTimeout(() => router.push('/orders'), 3000)
  } catch (e: any) {
    error.value = e.message || 'Failed to place order.'
  }
  loading.value = false
}
</script>

<style scoped>
.loader {
  border: 2px solid #f3f3f3;
  border-top: 2px solid #16a34a;
  border-radius: 50%;
  width: 18px;
  height: 18px;
  animation: spin 1s linear infinite;
  display: inline-block;
}
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
