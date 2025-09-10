// User Types
export interface User {
  id: string
  email: string
  full_name: string
  role: 'admin' | 'customer'
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  full_name: string
}

export interface LoginResponse {
  message: string
  token: string
  user: User
}

// Product Types
export interface Product {
  id: string
  name: string
  description: string
  price: number
  stock: number
  category: string
  image_url?: string
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface ProductRequest {
  name: string
  description: string
  price: number
  stock: number
  category: string
  image?: File
}

export interface ProductsResponse {
  products: Product[]
  total: number
  page: number
  limit: number
}

// Cart Types
export interface CartItem {
  id: string
  name: string
  price: number
  image_url?: string
  stock: number
  quantity: number
}

// Order Types
export interface Order {
  id: string
  user_id: string
  status: 'pending' | 'paid' | 'processing' | 'shipped' | 'delivered' | 'cancelled'
  total_amount: number
  shipping_address: string
  phone_number?: string
  created_at: string
  updated_at: string
  items?: OrderItem[]
}

export interface OrderItem {
  id: string
  order_id: string
  product_id: string
  quantity: number
  unit_price: number
  product?: Product
}

export interface OrderRequest {
  shipping_address: string
  phone_number: string
  items: Array<{
    product_id: string
    quantity: number
    unit_price: number
  }>
}

// Transaction Types
export interface Transaction {
  id: string
  order_id: string
  mpesa_ref?: string
  status: 'pending' | 'success' | 'failed'
  amount: number
  phone_number: string
  created_at: string
  updated_at: string
}

export interface STKPushRequest {
  order_id: string
  phone_number: string
  amount: number
}

export interface STKPushResponse {
  transaction_id: string
  status: string
  message: string
}

// API Response Types
export interface ApiResponse<T = any> {
  data: T
  message?: string
  error?: string
}

export interface PaginationParams {
  page?: number
  limit?: number
  search?: string
  category?: string
  min_price?: number
  max_price?: number
}

// UI Types
export interface Toast {
  id: string
  type: 'success' | 'error' | 'warning' | 'info'
  title: string
  message?: string
  duration?: number
}

export interface ModalProps {
  isOpen: boolean
  title: string
  description?: string
}

// Form Types
export interface FormField {
  name: string
  label: string
  type: 'text' | 'email' | 'password' | 'number' | 'textarea' | 'select' | 'file'
  placeholder?: string
  required?: boolean
  validation?: {
    min?: number
    max?: number
    pattern?: RegExp
    message?: string
  }
  options?: Array<{ label: string; value: string | number }>
}

// Error Types
export interface ApiError {
  message: string
  status?: number
  code?: string
  details?: Record<string, any>
}

// Store State Types
export interface AuthState {
  user: User | null
  token: string | null
  loading: boolean
  error: string | null
}

export interface CartState {
  items: CartItem[]
}

export interface ProductState {
  products: Product[]
  currentProduct: Product | null
  loading: boolean
  error: string | null
  pagination: {
    page: number
    limit: number
    total: number
  }
}