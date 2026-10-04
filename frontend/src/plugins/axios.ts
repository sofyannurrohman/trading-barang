import axios from 'axios'
import { useAuthStore } from '../stores/auth'
import { notify } from '../lib/notify'
import router from '../router'

declare module 'axios' {
  export interface AxiosRequestConfig {
    /** Jika true, axios interceptor tidak akan menampilkan toast error otomatis untuk request ini */
    silent?: boolean
    /** Jika diisi, axios interceptor otomatis menampilkan toast sukses dengan teks ini */
    successMessage?: string
    /** @deprecated Diganti dengan silent / penanganan eksplisit */
    skipToast?: boolean
  }
}

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
  timeout: 30000,
})

// Request Interceptor
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

// Response Interceptor
api.interceptors.response.use(
  (response) => {
    // Tampilkan custom success toast jika didefinisikan di config request
    if (response.config.successMessage) {
      notify.success(response.config.successMessage)
    }
    return response
  },
  (error) => {
    const isSilent = error.config?.silent || error.config?.skipToast
    const status = error.response?.status

    // Tangani 401 Unauthorized (Sesi habis / token invalid)
    if (status === 401) {
      const authStore = useAuthStore()
      authStore.logout()
      if (router.currentRoute.value.path !== '/login') {
        notify.warning('Sesi Anda telah berakhir', 'Silakan masuk kembali untuk melanjutkan.')
        router.push('/login')
      }
      return Promise.reject(error)
    }

    // Tangani 403 Forbidden (Tidak ada hak akses)
    if (status === 403 && !isSilent) {
      notify.error('Akses Ditolak', 'Anda tidak memiliki hak akses untuk tindakan ini.')
      return Promise.reject(error)
    }

    // Tangani Network Error (Server down / tidak terhubung)
    if (!error.response && !isSilent) {
      notify.error('Koneksi Gagal', 'Tidak dapat terhubung ke server. Periksa jaringan Anda.')
      return Promise.reject(error)
    }

    // Catatan: Error 400, 422, 500 diteruskan ke Promise.reject agar komponen pemanggil
    // dapat menangani feedback kontekstual secara spesifik via notify.error(err).
    return Promise.reject(error)
  }
)

export default api
