import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '../stores/auth'

// Lazy-loaded components for better performance
const Home = () => import('../views/Home.vue')
const Products = () => import('../views/Products.vue')
const ProductDetail = () => import('../views/ProductDetail.vue')
const Checkout = () => import('../views/Checkout.vue')
const Login = () => import('../views/Login.vue')
const Register = () => import('../views/Register.vue')
const Profile = () => import('../views/Profile.vue')
const Orders = () => import('../views/Orders.vue')
const OrderDetail = () => import('../views/OrderDetail.vue')

// Admin components
const AdminDashboard = () => import('../views/admin/Dashboard.vue')
const AdminProducts = () => import('../views/admin/Products.vue')
const AdminOrders = () => import('../views/admin/Orders.vue')

// Error pages
const NotFound = () => import('../views/errors/NotFound.vue')
const Unauthorized = () => import('../views/errors/Unauthorized.vue')

const routes: RouteRecordRaw[] = [
  {
    path: '/products',
    name: 'products',
    component: Products,
    meta: {
      title: 'Products - ShopVue',
      description: 'Browse our amazing collection of products'
    }
  },
  {
    path: '/',
    name: 'home',
    component: Home,
    meta: {
      title: 'Home - ShopVue',
      description: 'Welcome to ShopVue - Your modern e-commerce destination'
    }
  },
  {
    path: '/products/:id',
    name: 'product-detail',
    component: ProductDetail,
    props: true,
    meta: {
      title: 'Product Details - ShopVue'
    }
  },
  {
    path: '/checkout',
    name: 'checkout',
    component: Checkout,
    meta: {
      title: 'Checkout - ShopVue',
      requiresAuth: true
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      title: 'Sign In - ShopVue',
      requiresGuest: true
    }
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
    meta: {
      title: 'Sign Up - ShopVue',
      requiresGuest: true
    }
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile,
    meta: {
      title: 'My Profile - ShopVue',
      requiresAuth: true
    }
  },
  {
    path: '/orders',
    name: 'orders',
    component: Orders,
    meta: {
      title: 'My Orders - ShopVue',
      requiresAuth: true
    }
  },
  {
    path: '/orders/:id',
    name: 'order-detail',
    component: OrderDetail,
    props: true,
    meta: {
      title: 'Order Details - ShopVue',
      requiresAuth: true
    }
  },
  
  // Admin routes
  {
    path: '/admin',
    name: 'admin',
    redirect: '/admin/dashboard',
    meta: {
      requiresAuth: true,
      requiresAdmin: true
    },
    children: [
      {
        path: 'dashboard',
        name: 'admin-dashboard',
        component: AdminDashboard,
        meta: {
          title: 'Admin Dashboard - ShopVue',
          requiresAuth: true,
          requiresAdmin: true
        }
      },
      {
        path: 'products',
        name: 'admin-products',
        component: AdminProducts,
        meta: {
          title: 'Manage Products - Admin - ShopVue',
          requiresAuth: true,
          requiresAdmin: true
        }
      },
      {
        path: 'orders',
        name: 'admin-orders',
        component: AdminOrders,
        meta: {
          title: 'Manage Orders - Admin - ShopVue',
          requiresAuth: true,
          requiresAdmin: true
        }
      }
    ]
  },

  // Error pages
  {
    path: '/unauthorized',
    name: 'unauthorized',
    component: Unauthorized,
    meta: {
      title: 'Unauthorized - ShopVue'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: NotFound,
    meta: {
      title: 'Page Not Found - ShopVue'
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    // Always scroll to top when changing pages
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // Set page title
  if (to.meta.title) {
    document.title = to.meta.title as string
  }
  
  // Set meta description
  if (to.meta.description) {
    const metaDescription = document.querySelector('meta[name="description"]')
    if (metaDescription) {
      metaDescription.setAttribute('content', to.meta.description as string)
    }
  }
  
  // Check authentication requirements
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next({ name: 'login', query: { redirect: to.fullPath } })
    return
  }
  
  // Check admin requirements
  if (to.meta.requiresAdmin && !authStore.isAdmin) {
    next({ name: 'unauthorized' })
    return
  }
  
  // Redirect authenticated users away from guest pages
  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next({ name: 'home' })
    return
  }
  
  next()
})

// Global route error handler
router.onError((error) => {
  console.error('Router error:', error)
  // You could show a toast notification here
})

export default router