
<template>
  <div class="p-6 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold mb-4">My Orders</h1>
    <div v-if="loading" class="text-blue-600 mb-4">Loading your orders...</div>
    <div v-if="error" class="text-red-600 mb-4">{{ error }}</div>
    <div v-if="!loading && !error && orders.length === 0" class="text-gray-500">You have no orders yet.</div>
    <div v-if="orders.length > 0" class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-100">
          <tr>
            <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Order ID</th>
            <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Date</th>
            <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
            <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Total</th>
            <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Items</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="order in orders" :key="order.id">
            <td class="px-4 py-2 whitespace-nowrap">#{{ order.id.slice(0, 8) }}</td>
            <td class="px-4 py-2 whitespace-nowrap">{{ order.created_at.split('T')[0] }}</td>
            <td class="px-4 py-2 whitespace-nowrap">
              <span :class="statusClass(order.status)">{{ order.status }}</span>
            </td>
            <td class="px-4 py-2 whitespace-nowrap">KES {{ (order.total_amount ?? 0).toLocaleString() }}</td>
            <td class="px-4 py-2 whitespace-nowrap">
              <ul>
                <li v-for="item in (order.items || [])" :key="item.id">
                  {{ item.product?.name || item.product_id }} x{{ item.quantity }}
                </li>
              </ul>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ordersAPI } from '../services/api';
import { useAuthStore } from '../stores/auth';

const orders = ref<any[]>([])
const loading = ref(false)
const error = ref('')

function statusClass(status: string) {
  switch (status) {
    case 'paid':
    case 'completed':
    case 'delivered':
      return 'px-2 py-1 rounded bg-green-100 text-green-700 text-xs font-semibold'
    case 'pending':
    case 'processing':
      return 'px-2 py-1 rounded bg-yellow-100 text-yellow-700 text-xs font-semibold'
    case 'cancelled':
      return 'px-2 py-1 rounded bg-red-100 text-red-700 text-xs font-semibold'
    default:
      return 'px-2 py-1 rounded bg-gray-100 text-gray-700 text-xs font-semibold'
  }
}

const authStore = useAuthStore();
const router = useRouter();

async function fetchOrders() {
  loading.value = true;
  error.value = '';
  try {
    const result = await ordersAPI.getAll();
    orders.value = Array.isArray(result) ? result : [];
  } catch (e: any) {
    error.value = e.message || 'Failed to fetch orders.';
    orders.value = [];
  }
  loading.value = false;
}

onMounted(async () => {
  await authStore.initializeAuth();
  if (!authStore.isAuthenticated) {
    router.push('/login');
    return;
  }
  await fetchOrders();
});
</script>

<style scoped>
</style>
