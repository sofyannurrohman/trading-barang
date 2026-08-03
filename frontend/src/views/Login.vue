<template>
  <div class="w-full lg:grid lg:min-h-screen lg:grid-cols-2 xl:min-h-screen">
    <!-- Left Branding Side -->
    <div class="hidden bg-primary lg:flex flex-col justify-center items-center text-primary-foreground p-12">
      <div class="max-w-md space-y-6">
        <h1 class="text-4xl font-extrabold tracking-tight">Sistem Manajemen Perdagangan & ERP</h1>
        <p class="text-lg opacity-90 text-balance leading-relaxed">
          Platform terintegrasi untuk mengelola inventaris, perdagangan, dan administrasi ERP secara efisien dan profesional.
        </p>
      </div>
    </div>
    
    <!-- Right Form Side -->
    <div class="flex items-center justify-center py-12 bg-background min-h-screen lg:min-h-full">
      <div class="mx-auto w-full max-w-md space-y-6 px-4">
        <div class="text-center space-y-2">
          <h2 class="text-3xl font-bold tracking-tight">Login</h2>
          <p class="text-muted-foreground text-sm">
            Masukkan username dan password Anda untuk masuk ke sistem.
          </p>
        </div>

        <form class="space-y-4" @submit.prevent="handleLogin">
          <div class="space-y-2">
            <label for="username" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">Username</label>
            <Input 
              id="username" 
              v-model="username" 
              type="text" 
              required 
              placeholder="Masukkan username" 
            />
          </div>
          <div class="space-y-2">
            <label for="password" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">Password</label>
            <Input 
              id="password" 
              v-model="password" 
              type="password" 
              required 
              placeholder="Masukkan password" 
            />
          </div>

          <div v-if="error" class="text-destructive text-sm font-medium text-center bg-destructive/10 p-3 rounded-md">
            {{ error }}
          </div>

          <Button type="submit" :disabled="loading" class="w-full mt-2">
            {{ loading ? 'Signing in...' : 'Sign in' }}
          </Button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import api from '../plugins/axios'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { toast } from 'vue-sonner'

const username = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

const router = useRouter()
const authStore = useAuthStore()

const handleLogin = async () => {
  loading.value = true
  error.value = ''
  try {
    const response = await api.post('/auth/login', {
      username: username.value,
      password: password.value
    })
    
    authStore.login(response.data.token, response.data.user)
    toast.success('Login berhasil')
    router.push('/dashboard')
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>
