<template>
  <div class="min-h-screen w-full flex bg-slate-950 text-slate-100 font-sans selection:bg-rose-500 selection:text-white">
    <!-- Left Hero & Branding Section (Visible on Large Screens) -->
    <div class="hidden lg:flex lg:w-1/2 xl:w-7/12 relative flex-col justify-between p-12 overflow-hidden"
         style="background: radial-gradient(circle at 10% 20%, #681A1A 0%, #3B0D0D 50%, #1A0505 100%);">
      
      <!-- Ambient Glow & Pattern Elements -->
      <div class="absolute top-0 right-0 -mr-20 -mt-20 w-96 h-96 rounded-full bg-rose-600/10 blur-3xl pointer-events-none"></div>
      <div class="absolute bottom-0 left-0 -ml-20 -mb-20 w-96 h-96 rounded-full bg-amber-600/10 blur-3xl pointer-events-none"></div>
      <div class="absolute inset-0 bg-[linear-gradient(to_right,#ffffff08_1px,transparent_1px),linear-gradient(to_bottom,#ffffff08_1px,transparent_1px)] bg-[size:32px_32px] pointer-events-none"></div>

      <!-- Top Header Brand -->
      <div class="relative z-10 flex items-center gap-3">
        <div class="w-11 h-11 rounded-xl bg-gradient-to-br from-rose-500 to-red-800 p-0.5 shadow-lg shadow-rose-950/50 flex items-center justify-center border border-rose-400/30">
          <div class="w-full h-full bg-[#3B0D0D]/80 backdrop-blur rounded-[10px] flex items-center justify-center">
            <LayoutDashboard class="w-6 h-6 text-rose-300" />
          </div>
        </div>
        <div>
          <span class="text-xl font-bold tracking-tight text-rose-100">Trading Barang</span>
          <span class="block text-xs font-medium text-rose-300/70 tracking-wider uppercase">Enterprise ERP System</span>
        </div>
      </div>

      <!-- Center Hero Message & Highlights -->
      <div class="relative z-10 max-w-xl space-y-8 my-auto py-12">
        <div class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full bg-rose-500/10 border border-rose-500/20 text-rose-300 text-xs font-semibold backdrop-blur-sm shadow-sm">
          <Sparkles class="w-3.5 h-3.5 text-rose-400" />
          <span>Solusi Manajemen Terpadu &amp; Faktur Pajak</span>
        </div>

        <div class="space-y-4">
          <h1 class="text-4xl xl:text-5xl font-extrabold tracking-tight text-white leading-tight">
            Akurasi Operasional &amp; Kontrol Bisnis <span class="bg-gradient-to-r from-rose-300 via-red-200 to-amber-200 bg-clip-text text-transparent">Real-Time.</span>
          </h1>
          <p class="text-slate-300/85 text-base xl:text-lg leading-relaxed font-normal">
            Kelola persediaan multi-barang, penerbitan invoice dengan PPN otomatis, dan analisa margin laba bersih dalam satu platform yang handal dan presisi.
          </p>
        </div>

        <!-- System Feature Badges -->
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 pt-2">
          <div class="p-4 rounded-xl bg-white/[0.04] border border-white/10 backdrop-blur-md hover:bg-white/[0.07] transition-colors">
            <Boxes class="w-5 h-5 text-rose-400 mb-2" />
            <h4 class="text-sm font-semibold text-white">Inventori Cerdas</h4>
            <p class="text-xs text-slate-400 mt-1">Perhitungan HPP otomatis &amp; peringatan stok.</p>
          </div>
          <div class="p-4 rounded-xl bg-white/[0.04] border border-white/10 backdrop-blur-md hover:bg-white/[0.07] transition-colors">
            <Receipt class="w-5 h-5 text-rose-400 mb-2" />
            <h4 class="text-sm font-semibold text-white">Faktur &amp; Pajak</h4>
            <p class="text-xs text-slate-400 mt-1">Kalkulasi PPN 11% &amp; cetak invoice instan.</p>
          </div>
          <div class="p-4 rounded-xl bg-white/[0.04] border border-white/10 backdrop-blur-md hover:bg-white/[0.07] transition-colors">
            <TrendingUp class="w-5 h-5 text-rose-400 mb-2" />
            <h4 class="text-sm font-semibold text-white">Laba Kotor</h4>
            <p class="text-xs text-slate-400 mt-1">Laporan analitik finansial transparan.</p>
          </div>
        </div>
      </div>

      <!-- Bottom Compliance / Trust Footer -->
      <div class="relative z-10 flex items-center justify-between text-xs text-slate-400/80 border-t border-white/10 pt-6">
        <div class="flex items-center gap-2">
          <ShieldCheck class="w-4 h-4 text-emerald-400" />
          <span>Role-Based Access &amp; Enkripsi Terproteksi</span>
        </div>
        <span>&copy; {{ new Date().getFullYear() }} PT Trading Barang</span>
      </div>
    </div>

    <!-- Right Side Login Form Section -->
    <div class="w-full lg:w-1/2 xl:w-5/12 flex items-center justify-center p-6 sm:p-10 md:p-12 bg-slate-900 border-l border-slate-800 relative">
      <div class="w-full max-w-md space-y-7">
        
        <!-- Mobile Logo Header (Hidden on Desktop) -->
        <div class="lg:hidden flex flex-col items-center text-center mb-6">
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-rose-500 to-red-800 p-0.5 shadow-lg flex items-center justify-center mb-3">
            <div class="w-full h-full bg-slate-900 rounded-[10px] flex items-center justify-center">
              <LayoutDashboard class="w-6 h-6 text-rose-400" />
            </div>
          </div>
          <h2 class="text-xl font-bold text-white tracking-tight">Trading Barang ERP</h2>
          <span class="text-xs text-slate-400">Enterprise Trading System</span>
        </div>

        <!-- Form Title Header -->
        <div class="space-y-2">
          <h2 class="text-2xl sm:text-3xl font-bold tracking-tight text-white">Selamat Datang</h2>
          <p class="text-sm text-slate-400">
            Silakan masukkan <span class="text-slate-200 font-medium">Email</span> atau <span class="text-slate-200 font-medium">Username</span> serta kata sandi Anda.
          </p>
        </div>

        <!-- Error Feedback Alert -->
        <div v-if="error" class="flex items-start gap-3 p-3.5 rounded-xl bg-rose-500/10 border border-rose-500/30 text-rose-300 text-sm animate-shake">
          <AlertCircle class="w-5 h-5 flex-shrink-0 text-rose-400 mt-0.5" />
          <span class="flex-1 font-medium">{{ error }}</span>
        </div>

        <!-- Login Form -->
        <form class="space-y-4" @submit.prevent="handleLogin">
          
          <!-- Email / Username Input -->
          <div class="space-y-1.5">
            <label for="identifier" class="text-xs font-semibold text-slate-300 uppercase tracking-wider block">
              Email atau Username
            </label>
            <div class="relative flex items-center">
              <div class="absolute left-3.5 text-slate-400 pointer-events-none flex items-center justify-center">
                <Mail class="w-4 h-4" />
              </div>
              <input
                id="identifier"
                v-model="identifier"
                type="text"
                required
                autocomplete="username"
                placeholder="nama@email.com atau username"
                class="w-full pl-10 pr-4 py-2.5 bg-slate-950/70 border border-slate-700/80 rounded-xl text-slate-100 placeholder-slate-500 text-sm focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 transition-all shadow-inner"
              />
            </div>
          </div>

          <!-- Password Input -->
          <div class="space-y-1.5">
            <div class="flex items-center justify-between">
              <label for="password" class="text-xs font-semibold text-slate-300 uppercase tracking-wider block">
                Kata Sandi
              </label>
            </div>
            <div class="relative flex items-center">
              <div class="absolute left-3.5 text-slate-400 pointer-events-none flex items-center justify-center">
                <Lock class="w-4 h-4" />
              </div>
              <input
                id="password"
                v-model="password"
                :type="showPassword ? 'text' : 'password'"
                required
                autocomplete="current-password"
                placeholder="Masukkan kata sandi Anda"
                class="w-full pl-10 pr-11 py-2.5 bg-slate-950/70 border border-slate-700/80 rounded-xl text-slate-100 placeholder-slate-500 text-sm focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 transition-all shadow-inner"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3.5 text-slate-400 hover:text-slate-200 focus:outline-none transition-colors"
                tabindex="-1"
                :title="showPassword ? 'Sembunyikan sandi' : 'Tampilkan sandi'"
              >
                <EyeOff v-if="showPassword" class="w-4 h-4" />
                <Eye v-else class="w-4 h-4" />
              </button>
            </div>
          </div>

          <!-- Submit Button -->
          <div class="pt-2">
            <button
              type="submit"
              :disabled="loading"
              class="w-full py-3 px-4 rounded-xl font-semibold text-sm text-white shadow-lg transition-all duration-200 flex items-center justify-center gap-2 cursor-pointer disabled:opacity-60 disabled:cursor-not-allowed group relative overflow-hidden"
              style="background: linear-gradient(135deg, #A82020 0%, #7B1D1D 60%, #5C1A1A 100%); box-shadow: 0 4px 15px rgba(168, 32, 32, 0.35);"
            >
              <span v-if="loading" class="inline-flex items-center gap-2">
                <Loader2 class="w-4 h-4 animate-spin text-rose-200" />
                <span>Memverifikasi akun...</span>
              </span>
              <span v-else class="inline-flex items-center gap-2">
                <span>Masuk ke Dashboard</span>
                <ArrowRight class="w-4 h-4 text-rose-200 group-hover:translate-x-0.5 transition-transform" />
              </span>
            </button>
          </div>
        </form>

        <!-- Demo Account Quick Info Card -->
        <div class="pt-3 border-t border-slate-800/80">
          <div class="p-3.5 rounded-xl bg-slate-950/60 border border-slate-800 text-xs text-slate-400 space-y-2">
            <div class="flex items-center justify-between text-slate-300 font-medium">
              <span class="flex items-center gap-1.5">
                <KeyRound class="w-3.5 h-3.5 text-amber-400" />
                Akun Akses Default
              </span>
              <button 
                type="button" 
                @click="fillDemoCredentials"
                class="text-rose-400 hover:text-rose-300 font-semibold underline underline-offset-2 transition-colors cursor-pointer"
              >
                Gunakan Demo
              </button>
            </div>
            <div class="flex items-center justify-between font-mono text-[11px] bg-slate-900/90 px-2.5 py-1.5 rounded-lg border border-slate-800">
              <span class="text-slate-400">Email / Sandi:</span>
              <span class="text-slate-200">admin@tradingbarang.com / password</span>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import api from '../plugins/axios'
import { toast } from 'vue-sonner'
import {
  LayoutDashboard,
  Sparkles,
  Boxes,
  Receipt,
  TrendingUp,
  ShieldCheck,
  Mail,
  Lock,
  Eye,
  EyeOff,
  ArrowRight,
  Loader2,
  AlertCircle,
  KeyRound
} from '@lucide/vue'

const identifier = ref('')
const password = ref('')
const showPassword = ref(false)
const error = ref('')
const loading = ref(false)

const router = useRouter()
const authStore = useAuthStore()

const fillDemoCredentials = () => {
  identifier.value = 'admin@tradingbarang.com'
  password.value = 'password'
  error.value = ''
}

const handleLogin = async () => {
  if (!identifier.value || !password.value) {
    error.value = 'Silakan isi email/username dan password'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const response = await api.post('/auth/login', {
      email: identifier.value,
      username: identifier.value,
      password: password.value
    }, { skipToast: true } as any)

    authStore.login(response.data.token, response.data.user)
    toast.success(`Selamat datang kembali, ${response.data.user?.username || 'Admin'}!`)
    router.push('/dashboard')
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Gagal masuk. Periksa kembali email/username dan password Anda.'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
@keyframes shake {
  0%, 100% { transform: translateX(0); }
  20%, 60% { transform: translateX(-4px); }
  40%, 80% { transform: translateX(4px); }
}

.animate-shake {
  animation: shake 0.3s ease-in-out;
}
</style>
