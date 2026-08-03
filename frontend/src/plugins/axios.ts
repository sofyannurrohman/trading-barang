import axios from 'axios'
import { useAuthStore } from '../stores/auth'
import { toast } from 'vue-sonner'
import router from '../router'

const api = axios.create({
  baseURL: 'http://localhost:8080',
})

api.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore()
    const token = authStore.token
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

api.interceptors.response.use(
  (response) => {
    const method = response.config.method?.toLowerCase()
    if (method && ['post', 'put', 'delete'].includes(method)) {
      if (!response.config.url?.includes('/login') && !response.config.url?.includes('/register')) {
        const msg = response.data?.message || 'Operasi berhasil disimpan'
        toast.success(msg)
      }
    }
    return response
  },
  (error) => {
    const msg = error.response?.data?.error || error.message || 'Terjadi kesalahan pada sistem'
    toast.error(msg)
    
    if (error.response?.status === 401) {
      const authStore = useAuthStore()
      authStore.logout()
      if (router.currentRoute.value.path !== '/login') {
        router.push('/login')
      }
    }
    return Promise.reject(error)
  }
)

export default api
