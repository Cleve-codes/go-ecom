<template>
  <div class="p-6">
      <h1 class="text-3xl font-bold mb-8 text-gray-800">Admin Dashboard</h1>

      <!-- Summary Cards -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-10">
        <div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
          <span class="text-gray-500">Total Users</span>
          <span class="text-2xl font-bold">{{ summary.users }}</span>
        </div>
        <div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
          <span class="text-gray-500">Total Orders</span>
          <span class="text-2xl font-bold">{{ summary.orders }}</span>
        </div>
        <div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
          <span class="text-gray-500">Total Revenue</span>
          <span class="text-2xl font-bold">${{ (summary.revenue ?? 0).toLocaleString() }}</span>
        </div>
        <div class="bg-white rounded-lg shadow p-6 flex flex-col items-center">
          <span class="text-gray-500">Products</span>
          <span class="text-2xl font-bold">{{ summary.products }}</span>
        </div>
      </div>

      <!-- Charts -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-10">
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-lg font-semibold mb-4">Sales Trend</h2>
          <Bar :data="salesChartData" :options="salesChartOptions" />
        </div>
        <div class="bg-white rounded-lg shadow p-6">
          <h2 class="text-lg font-semibold mb-4">Order Status</h2>
          <Doughnut :data="orderStatusData" :options="orderStatusOptions" />
        </div>
      </div>

      <!-- Recent Orders Table -->
      <div class="bg-white rounded-lg shadow p-6">
        <h2 class="text-lg font-semibold mb-4">Recent Orders</h2>
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Order ID</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">User</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Amount</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Date</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="order in recentOrders" :key="order.id">
                <td class="px-4 py-2 whitespace-nowrap">#{{ order.id }}</td>
                <td class="px-4 py-2 whitespace-nowrap">{{ order.user }}</td>
                <td class="px-4 py-2 whitespace-nowrap">${{ (order.amount ?? 0).toLocaleString() }}</td>
                <td class="px-4 py-2 whitespace-nowrap">
                  <span :class="statusClass(order.status)">{{ order.status }}</span>
                </td>
                <td class="px-4 py-2 whitespace-nowrap">{{ order.date }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Product Management Section -->
      <div class="bg-white rounded-lg shadow p-6 mt-10">
        <h2 class="text-lg font-semibold mb-4">Manage Products</h2>
        <form class="flex flex-wrap gap-4 mb-6" @submit.prevent="addProduct" enctype="multipart/form-data">
          <input v-model="newProduct.name" class="border rounded px-2 py-1" placeholder="Name" required />
          <input v-model="newProduct.description" class="border rounded px-2 py-1" placeholder="Description" />
          <input v-model.number="newProduct.price" type="number" class="border rounded px-2 py-1" placeholder="Price" required min="0" />
          <input v-model.number="newProduct.stock" type="number" class="border rounded px-2 py-1" placeholder="Stock" min="0" />
          <input v-model="newProduct.category" class="border rounded px-2 py-1" placeholder="Category" required />
          <input type="file" accept="image/*" @change="onImageChange" />
          <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded" :disabled="addingProduct">Add Product</button>
        </form>
        <div v-if="addError" class="text-red-600 mb-2">{{ addError }}</div>
        <div v-if="productsLoading" class="text-blue-600 mb-2">Loading products...</div>
        <div v-if="productsError" class="text-red-600 mb-2">{{ productsError }}</div>
        <div class="overflow-x-auto" v-if="!productsLoading && !productsError">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-100">
              <tr>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">ID</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Name</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Price</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Stock</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Category</th>
                <th class="px-4 py-2 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="product in products" :key="product.id">
                <td class="px-4 py-2 whitespace-nowrap">{{ product.id }}</td>
                <td class="px-4 py-2 whitespace-nowrap">{{ product.name }}</td>
                <td class="px-4 py-2 whitespace-nowrap">${{ product.price }}</td>
                <td class="px-4 py-2 whitespace-nowrap">{{ product.stock }}</td>
                <td class="px-4 py-2 whitespace-nowrap">{{ product.category }}</td>
                <td class="px-4 py-2 whitespace-nowrap product-actions">
                  <button @click="deleteProduct(product.id)" class="bg-red-500 text-white px-2 py-1 rounded">Delete</button>
                </td>
              </tr>
              <tr v-if="products.length === 0">
                <td colspan="6" class="text-center text-gray-400 py-4">No products found.</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <!-- End Product Management Section -->
    </div>
  </template>

  <script setup lang="ts">
  import { ref, onMounted } from 'vue'
  import { Bar, Doughnut } from 'vue-chartjs'
  import {
    Chart,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
    ArcElement
  } from 'chart.js'
  import { productsAPI, ordersAPI } from '@/services/api'
  import axios from 'axios'

  Chart.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend, ArcElement)

  const summary = ref({ users: 0, orders: 0, revenue: 0, products: 0 })
  const salesChartData = ref({ labels: [], datasets: [] })
  const salesChartOptions = {
    responsive: true,
    plugins: { legend: { display: false }, title: { display: false } },
    scales: { y: { beginAtZero: true } }
  }
  const orderStatusData = ref({ labels: [], datasets: [] })
  const orderStatusOptions = {
    responsive: true,
    plugins: { legend: { position: 'bottom' }, title: { display: false } }
  }
  const recentOrders = ref<any[]>([])

  // Product management state
  const products = ref<any[]>([])
  const productsLoading = ref(false)
  const productsError = ref('')
  const newProduct = ref({ name: '', description: '', price: 0, stock: 0, category: '', image: null })
  const addingProduct = ref(false)
  const addError = ref('')

  function statusClass(status: string) {
    switch (status) {
      case 'completed':
      case 'Completed':
        return 'px-2 py-1 rounded bg-green-100 text-green-700 text-xs font-semibold'
      case 'pending':
      case 'Pending':
        return 'px-2 py-1 rounded bg-yellow-100 text-yellow-700 text-xs font-semibold'
      case 'cancelled':
      case 'Cancelled':
        return 'px-2 py-1 rounded bg-red-100 text-red-700 text-xs font-semibold'
      default:
        return 'px-2 py-1 rounded bg-gray-100 text-gray-700 text-xs font-semibold'
    }
  }

  // Patch: update products list and summary count
  async function fetchDashboardData() {
    // Fetch products
    productsLoading.value = true
    productsError.value = ''
    try {
      const productsRes = await productsAPI.getAll({ limit: 1000 })
      console.log('Fetched products:', productsRes)
      if (!productsRes || !Array.isArray(productsRes.products)) {
        throw new Error('Invalid products response from API')
      }
      products.value = productsRes.products
      summary.value.products = productsRes.products.length
    } catch (e) {
      productsError.value = (e as any).message || 'Failed to fetch products.'
      products.value = []
      summary.value.products = 0
    }
    productsLoading.value = false

    // Fetch orders
    const orders = await ordersAPI.getAllAdmin()
    summary.value.orders = orders.length
    summary.value.revenue = orders.reduce((sum, o) => sum + (o.total_amount || 0), 0)

    // Sales trend (by month)
    const salesByMonth: Record<string, number> = {}
    orders.forEach(order => {
      const month = new Date(order.created_at).toLocaleString('default', { month: 'short' })
      salesByMonth[month] = (salesByMonth[month] || 0) + (order.total_amount || 0)
    })
    const months = Object.keys(salesByMonth)
    salesChartData.value = {
      labels: months,
      datasets: [
        {
          label: 'Sales',
          backgroundColor: '#3b82f6',
          data: months.map(m => salesByMonth[m]),
          borderRadius: 6
        }
      ]
    }

    // Order status breakdown
    const statusCounts: Record<string, number> = {}
    orders.forEach(order => {
      statusCounts[order.status] = (statusCounts[order.status] || 0) + 1
    })
    orderStatusData.value = {
      labels: Object.keys(statusCounts),
      datasets: [
        {
          label: 'Orders',
          backgroundColor: ['#fbbf24', '#10b981', '#ef4444', '#6366f1', '#f472b6', '#a3e635'],
          data: Object.values(statusCounts)
        }
      ]
    }

    // Recent orders (latest 5)
    recentOrders.value = orders
      .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
      .slice(0, 5)
      .map(order => ({
        id: order.id,
        user: order.user_id,
        amount: order.total_amount,
        status: order.status,
        date: order.created_at.split('T')[0]
      }))

    // Fetch users (admin endpoint)
    try {
  const res = await axios.get(`${import.meta.env.VITE_API_BASE_URL}/admin/users`, {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      })
      summary.value.users = res.data.length
    } catch (e) {
      summary.value.users = 0
    }
  }

  function onImageChange(e: Event) {
    const target = e.target as HTMLInputElement
    if (target.files && target.files[0]) {
      newProduct.value.image = target.files[0]
    } else {
      newProduct.value.image = null
    }
  }

  async function addProduct() {
    addingProduct.value = true
    addError.value = ''
    try {
      const payload = { ...newProduct.value }
      if (
        !payload.name ||
        payload.price === '' || isNaN(Number(payload.price)) ||
        payload.stock === '' || isNaN(Number(payload.stock)) ||
        !payload.category
      ) {
        addError.value = 'Name, price, stock, and category are required.';
        addingProduct.value = false;
        return;
      }
      const formData = new FormData()
      Object.entries(payload).forEach(([key, value]) => {
        if (key === 'image' && value instanceof File) {
          formData.append('image', value)
        } else if (value !== undefined && value !== null) {
          formData.append(key, String(value))
        }
      })
      // Debug: log FormData contents
      for (const pair of formData.entries()) {
        console.log('FormData:', pair[0], pair[1])
      }
      await productsAPI.create(formData)
      await fetchDashboardData()
      newProduct.value = { name: '', description: '', price: 0, stock: 0, category: '', image: null }
    } catch (e) {
      addError.value = (e as any).message || 'Failed to add product.'
    }
    addingProduct.value = false
  }

  async function deleteProduct(id: string) {
    if (!confirm('Are you sure you want to delete this product?')) return
    await productsAPI.delete(id)
    await fetchDashboardData()
  }

  onMounted(fetchDashboardData)
  </script>

  <style scoped>
  /* Custom styles if needed */
  </style>
