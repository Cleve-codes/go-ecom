import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { CartItem, Product } from '../types'
import { readonly } from 'vue'

export const useCartStore = defineStore('cart', () => {
  // State
  const items = ref<CartItem[]>([])

  // Initialize from localStorage
  const initializeCart = () => {
    const storedCart = localStorage.getItem('cart')
    if (storedCart) {
      try {
        items.value = JSON.parse(storedCart)
      } catch (e) {
        console.error('Failed to parse stored cart data:', e)
        items.value = []
      }
    }
  }

  // Getters
  const itemCount = computed(() => 
    items.value.reduce((total, item) => total + item.quantity, 0)
  )
  
  const totalAmount = computed(() => 
    items.value.reduce((total, item) => total + (item.price * item.quantity), 0)
  )
  
  const cartItems = computed(() => items.value)
  
  const getItemById = computed(() => (id: string) => 
    items.value.find(item => item.id === id)
  )

  const isEmpty = computed(() => items.value.length === 0)

  const uniqueItemCount = computed(() => items.value.length)

  // Actions
  const saveToStorage = (): void => {
    localStorage.setItem('cart', JSON.stringify(items.value))
  }

  const addItem = (product: Product, quantity: number = 1): boolean => {
    const existingItemIndex = items.value.findIndex(item => item.id === product.id)
    
    if (existingItemIndex > -1) {
      const existingItem = items.value[existingItemIndex]
      const newQuantity = existingItem.quantity + quantity
      if (newQuantity <= product.stock) {
        existingItem.quantity = newQuantity
        saveToStorage()
        return true
      } else {
        // Can't add more than available stock
        return false
      }
    } else {
      if (quantity <= product.stock) {
        const cartItem: CartItem = {
          id: product.id,
          name: product.name,
          price: product.price,
          image_url: product.image_url,
          stock: product.stock,
          quantity,
        }
        items.value.push(cartItem)
        saveToStorage()
        return true
      } else {
        return false
      }
    }
  }

  const removeItem = (productId: string): void => {
    const index = items.value.findIndex(item => item.id === productId)
    if (index > -1) {
      items.value.splice(index, 1)
      saveToStorage()
    }
  }

  const updateQuantity = (productId: string, quantity: number): boolean => {
    const item = items.value.find(item => item.id === productId)
    
    if (!item) return false
    
    if (quantity <= 0) {
      removeItem(productId)
      return true
    }
    
    if (quantity <= item.stock) {
      item.quantity = quantity
      saveToStorage()
      return true
    }
    
    return false
  }

  const incrementItem = (productId: string): boolean => {
    const item = items.value.find(item => item.id === productId)
    if (item && item.quantity < item.stock) {
      item.quantity++
      saveToStorage()
      return true
    }
    return false
  }

  const decrementItem = (productId: string): boolean => {
    const item = items.value.find(item => item.id === productId)
    if (item) {
      if (item.quantity > 1) {
        item.quantity--
        saveToStorage()
        return true
      } else {
        removeItem(productId)
        return true
      }
    }
    return false
  }

  const clearCart = (): void => {
    items.value = []
    saveToStorage()
  }

  const updateItemStock = (productId: string, newStock: number): void => {
    const item = items.value.find(item => item.id === productId)
    if (item) {
      item.stock = newStock
      // If quantity exceeds new stock, adjust it
      if (item.quantity > newStock) {
        if (newStock === 0) {
          removeItem(productId)
        } else {
          item.quantity = newStock
        }
      }
      saveToStorage()
    }
  }

  const getCartSummary = computed(() => ({
    itemCount: itemCount.value,
    uniqueItemCount: uniqueItemCount.value,
    totalAmount: totalAmount.value,
    isEmpty: isEmpty.value,
    items: items.value,
  }))

  // Initialize cart on store creation
  initializeCart()

  return {
    // State
    items: readonly(items),
    
    // Getters
    itemCount,
    totalAmount,
    cartItems,
    getItemById,
    isEmpty,
    uniqueItemCount,
    getCartSummary,
    
    // Actions
    addItem,
    removeItem,
    updateQuantity,
    incrementItem,
    decrementItem,
    clearCart,
    updateItemStock,
    initializeCart,
    saveToStorage,
  }
})