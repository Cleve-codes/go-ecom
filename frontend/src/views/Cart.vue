
<template>
  <div class="p-6">
    <h1 class="text-2xl font-bold mb-4">Shopping Cart</h1>
    <div v-if="cart.isEmpty" class="text-gray-500">Your cart is empty.</div>
    <div v-else>
      <div class="overflow-x-auto">
        <table class="min-w-full bg-white rounded-lg shadow">
          <thead>
            <tr class="bg-gray-100 text-gray-700">
              <th class="py-2 px-4 text-left">Product</th>
              <th class="py-2 px-4 text-left">Price</th>
              <th class="py-2 px-4 text-left">Quantity</th>
              <th class="py-2 px-4 text-left">Total</th>
              <th class="py-2 px-4"></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in cart.cartItems" :key="item.id" class="border-b">
              <td class="py-2 px-4 flex items-center gap-3">
                <img :src="item.image_url || '/placeholder-product.jpg'" :alt="item.name" class="w-14 h-14 object-cover rounded border" />
                <span>{{ item.name }}</span>
              </td>
              <td class="py-2 px-4">${{ item.price.toFixed(2) }}</td>
              <td class="py-2 px-4">
                <button @click="cart.decrementItem(item.id)" class="btn btn-xs">-</button>
                <span class="mx-2">{{ item.quantity }}</span>
                <button @click="cart.incrementItem(item.id)" class="btn btn-xs">+</button>
              </td>
              <td class="py-2 px-4">${{ (item.price * item.quantity).toFixed(2) }}</td>
              <td class="py-2 px-4">
                <button @click="cart.removeItem(item.id)" class="btn btn-xs btn-danger">Remove</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="mt-6 flex justify-end">
        <div class="bg-gray-50 rounded-lg p-6 w-full max-w-xs">
          <div class="flex justify-between mb-2">
            <span class="font-medium">Total Items:</span>
            <span>{{ cart.itemCount }}</span>
          </div>
          <div class="flex justify-between mb-2">
            <span class="font-medium">Total Amount:</span>
            <span class="font-bold text-lg text-primary-600">${{ cart.totalAmount.toFixed(2) }}</span>
          </div>
          <button class="btn btn-primary w-full mt-4" :disabled="cart.isEmpty">Checkout</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useCartStore } from '../stores/cart'
const cart = useCartStore()
</script>

<style scoped>
</style>
