import { toast } from 'vue-sonner'

/**
 * Ekstraksi pesan error dari berbagai format error (Axios, native Error, string, dll)
 */
export function extractErrorMessage(error: any, defaultMessage = 'Terjadi kesalahan pada sistem'): string {
  if (!error) return defaultMessage
  if (typeof error === 'string') return error

  // Axios Response error
  if (error.response?.data) {
    const data = error.response.data
    if (typeof data === 'string') return data
    if (data.error && typeof data.error === 'string') return data.error
    if (data.message && typeof data.message === 'string') return data.message
    if (data.errors && Array.isArray(data.errors)) return data.errors.join(', ')
    if (data.errors && typeof data.errors === 'object') {
      return Object.values(data.errors).flat().join(', ')
    }
  }

  // Network / general error message
  if (error.message) {
    if (error.message === 'Network Error') {
      return 'Gagal terhubung ke server. Periksa koneksi internet Anda.'
    }
    return error.message
  }

  return defaultMessage
}

interface NotifyOptions {
  description?: string
  duration?: number
  action?: {
    label: string
    onClick: () => void
  }
  cancel?: {
    label: string
    onClick?: () => void
  }
  onDismiss?: () => void
  onAutoClose?: () => void
  id?: string | number
}

/**
 * Service Notifikasi Terpadu menggunakan Sonner
 */
export const notify = {
  /**
   * Menampilkan toast sukses
   */
  success(title: string, description?: string | NotifyOptions, options?: NotifyOptions) {
    if (typeof description === 'object') {
      return toast.success(title, description)
    }
    return toast.success(title, {
      description,
      ...options,
    })
  },

  /**
   * Menampilkan toast error dengan auto-parse error object
   */
  error(titleOrError: any, description?: string | NotifyOptions, options?: NotifyOptions) {
    let title = 'Gagal melakukan aksi'
    let desc = typeof description === 'string' ? description : undefined
    const opt = typeof description === 'object' ? description : options

    if (typeof titleOrError === 'string') {
      title = titleOrError
    } else if (titleOrError) {
      // Jika yang dioper adalah object error, jadikan pesan utamanya sebagai title
      title = extractErrorMessage(titleOrError)
    }

    return toast.error(title, {
      description: desc,
      ...opt,
    })
  },

  /**
   * Menampilkan toast peringatan (warning)
   */
  warning(title: string, description?: string | NotifyOptions, options?: NotifyOptions) {
    if (typeof description === 'object') {
      return toast.warning(title, description)
    }
    return toast.warning(title, {
      description,
      ...options,
    })
  },

  /**
   * Menampilkan toast informasi (info)
   */
  info(title: string, description?: string | NotifyOptions, options?: NotifyOptions) {
    if (typeof description === 'object') {
      return toast.info(title, description)
    }
    return toast.info(title, {
      description,
      ...options,
    })
  },

  /**
   * Menampilkan toast status Promise (loading, success, error)
   */
  promise<T>(
    promise: Promise<T> | (() => Promise<T>),
    data: {
      loading: string
      success: string | ((data: T) => string)
      error: string | ((err: any) => string)
    },
    options?: NotifyOptions
  ) {
    return toast.promise(promise, {
      loading: data.loading,
      success: data.success,
      error: (err: any) => {
        if (typeof data.error === 'function') {
          return data.error(err)
        }
        return data.error || extractErrorMessage(err)
      },
      ...options,
    })
  },

  /**
   * Menutup toast tertentu atau semua toast
   */
  dismiss(toastId?: string | number) {
    toast.dismiss(toastId)
  },
}

export default notify
