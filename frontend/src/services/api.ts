import axios, { type AxiosResponse, type AxiosError } from 'axios'
import type {
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  User,
  Product,
  ProductRequest,
  ProductsResponse,
  Order,
  OrderRequest,
  STKPushRequest,
  STKPushResponse,
  Transaction,
  PaginationParams,
  ApiResponse
} from '../types'


// Create axios instance
const api = axios.create({
  baseURL: 'http://localhost:8082/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Attach Authorization header for all requests if token exists
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers = config.headers || {}
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

// Helper function to handle API responses
const handleResponse = <T>(response: AxiosResponse<T>): T => response.data

// Helper function to handle API errors
const handleError = (error: AxiosError): never => {
  const message = (error.response?.data as any)?.error || error.message || 'An error occurred'
  throw new Error(message)
}

// Auth API
export const authAPI = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    try {
      const response = await api.post<LoginResponse>('/auth/login', credentials)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  register: async (userData: RegisterRequest): Promise<LoginResponse> => {
    try {
      const response = await api.post<LoginResponse>('/auth/register', userData)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  getProfile: async (): Promise<{ user: User }> => {
    try {
      const response = await api.get<{ user: User }>('/auth/profile')
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },
}

// Products API
export const productsAPI = {
  getAll: async (params: PaginationParams = {}): Promise<ProductsResponse> => {
    try {
      const response = await api.get<ProductsResponse>('/products', { params })
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  getById: async (id: string): Promise<Product> => {
    try {
      const response = await api.get<Product>(`/products/${id}`)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  create: async (formData: FormData): Promise<Product> => {
    try {
      const response = await api.post<Product>('/products', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      })
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  update: async (id: string, product: Partial<ProductRequest>): Promise<Product> => {
    try {
      const formData = new FormData()
      
      // Append product data
      Object.entries(product).forEach(([key, value]) => {
        if (key === 'image' && value instanceof File) {
          formData.append('image', value)
        } else if (value !== undefined && value !== null) {
          formData.append(key, String(value))
        }
      })

      const response = await api.put<Product>(`/products/${id}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      })
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  delete: async (id: string): Promise<void> => {
    try {
      await api.delete(`/products/${id}`)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },
}

// Orders API
export const ordersAPI = {
  // Fetch orders for the current user
  getAll: async (): Promise<Order[]> => {
    try {
      const response = await api.get<Order[]>('/orders')
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  // Fetch all orders (admin only)
  getAllAdmin: async (): Promise<Order[]> => {
    try {
      const response = await api.get<Order[]>('/admin/orders')
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  getById: async (id: string): Promise<Order> => {
    try {
      const response = await api.get<Order>(`/orders/${id}`)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  create: async (order: OrderRequest): Promise<Order> => {
    try {
      const response = await api.post<Order>('/orders', order)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  updateStatus: async (id: string, status: Order['status']): Promise<Order> => {
    try {
      const response = await api.patch<Order>(`/orders/${id}/status`, { status })
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },
}

// M-Pesa API
export const mpesaAPI = {
  initiateSTK: async (payment: STKPushRequest): Promise<STKPushResponse> => {
    try {
      const response = await api.post<STKPushResponse>('/mpesa/stkpush', payment)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },

  checkStatus: async (transactionId: string): Promise<Transaction> => {
    try {
      const response = await api.get<Transaction>(`/mpesa/status/${transactionId}`)
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },
}

// Health check
export const healthAPI = {
  check: async (): Promise<{ status: string; message: string }> => {
    try {
      const response = await api.get<{ status: string; message: string }>('/health')
      return handleResponse(response)
    } catch (error) {
      return handleError(error as AxiosError)
    }
  },
}

export default api