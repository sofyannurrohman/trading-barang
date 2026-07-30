<template>
  <div class="min-h-screen bg-slate-50 flex font-sans">
    <!-- Sidebar -->
    <aside class="w-64 bg-slate-900 text-slate-300 flex flex-col transition-all duration-300" :class="{ '-ml-64': !isSidebarOpen }">
      <div class="h-16 flex items-center justify-center border-b border-slate-800 px-4">
        <h1 class="text-xl font-bold text-white flex items-center gap-2">
          <div class="w-8 h-8 bg-indigo-600 rounded flex items-center justify-center text-white">E</div>
          ERP System
        </h1>
      </div>
      
      <div class="flex-1 overflow-y-auto py-4">
        <nav class="space-y-1 px-2">
          <!-- Dashboard -->
          <router-link to="/dashboard" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">
            <span class="mr-3 text-slate-400 group-hover:text-slate-300">📊</span>
            Dashboard
          </router-link>

          <!-- Master Data -->
          <div class="mt-4">
            <button @click="toggleMenu('master')" class="w-full group flex items-center justify-between px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white">
              <div class="flex items-center">
                <span class="mr-3 text-slate-400 group-hover:text-slate-300">📦</span>
                Master Data
              </div>
              <span>{{ openMenus.master ? '▼' : '▶' }}</span>
            </button>
            <div v-show="openMenus.master" class="pl-8 space-y-1 mt-1">
              <router-link to="/dashboard/master-data/products" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Data Produk</router-link>
              <router-link to="/dashboard/master-data/partners" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Data Mitra</router-link>
            </div>
          </div>

          <!-- Gudang & Inventori -->
          <div class="mt-1">
            <button @click="toggleMenu('inventory')" class="w-full group flex items-center justify-between px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white">
              <div class="flex items-center">
                <span class="mr-3 text-slate-400 group-hover:text-slate-300">🏭</span>
                Gudang & Inventori
              </div>
              <span>{{ openMenus.inventory ? '▼' : '▶' }}</span>
            </button>
            <div v-show="openMenus.inventory" class="pl-8 space-y-1 mt-1">
              <router-link to="/dashboard/inventory/inbound" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Barang Masuk</router-link>
              <router-link to="/dashboard/inventory/adjustments" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Penyesuaian</router-link>
              <router-link to="/dashboard/inventory/stock-card" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Kartu Stok</router-link>
            </div>
          </div>

          <!-- Penjualan & Faktur -->
          <div class="mt-1">
            <button @click="toggleMenu('sales')" class="w-full group flex items-center justify-between px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white">
              <div class="flex items-center">
                <span class="mr-3 text-slate-400 group-hover:text-slate-300">🛒</span>
                Penjualan & Faktur
              </div>
              <span>{{ openMenus.sales ? '▼' : '▶' }}</span>
            </button>
            <div v-show="openMenus.sales" class="pl-8 space-y-1 mt-1">
              <router-link to="/dashboard/sales/create" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Buat Faktur</router-link>
              <router-link to="/dashboard/sales/list" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Daftar Faktur</router-link>
            </div>
          </div>

          <!-- Laporan -->
          <div class="mt-1">
            <button @click="toggleMenu('reports')" class="w-full group flex items-center justify-between px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white">
              <div class="flex items-center">
                <span class="mr-3 text-slate-400 group-hover:text-slate-300">📈</span>
                Laporan
              </div>
              <span>{{ openMenus.reports ? '▼' : '▶' }}</span>
            </button>
            <div v-show="openMenus.reports" class="pl-8 space-y-1 mt-1">
              <router-link to="/dashboard/reports/profit" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Laba Kotor</router-link>
              <router-link to="/dashboard/reports/tax" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Pajak Keluaran</router-link>
            </div>
          </div>

          <!-- Pengaturan -->
          <div class="mt-1">
            <button @click="toggleMenu('settings')" class="w-full group flex items-center justify-between px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white">
              <div class="flex items-center">
                <span class="mr-3 text-slate-400 group-hover:text-slate-300">⚙️</span>
                Pengaturan
              </div>
              <span>{{ openMenus.settings ? '▼' : '▶' }}</span>
            </button>
            <div v-show="openMenus.settings" class="pl-8 space-y-1 mt-1">
              <router-link to="/dashboard/settings/company" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Profil Perusahaan</router-link>
              <router-link to="/dashboard/settings/users" class="group flex items-center px-2 py-2 text-sm font-medium rounded-md hover:bg-slate-800 hover:text-white" active-class="bg-slate-800 text-white">Pengguna</router-link>
            </div>
          </div>
        </nav>
      </div>
    </aside>

    <!-- Main Content wrapper -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Header -->
      <header class="bg-white shadow-sm border-b border-slate-200 h-16 flex items-center justify-between px-4 sm:px-6 lg:px-8 z-10">
        <div class="flex items-center">
          <button @click="isSidebarOpen = !isSidebarOpen" class="text-slate-500 hover:text-slate-700 focus:outline-none">
            <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-slate-600 font-medium">Hello, {{ authStore.user?.username }} ({{ authStore.user?.role }})</span>
          <button @click="handleLogout" class="px-4 py-2 border border-slate-300 rounded-md text-sm font-medium text-slate-700 bg-white hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Logout
          </button>
        </div>
      </header>

      <!-- Main viewport -->
      <main class="flex-1 relative overflow-y-auto focus:outline-none">
        <div class="py-6">
          <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <router-view />
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const isSidebarOpen = ref(true)

const openMenus = reactive({
  master: false,
  inventory: false,
  sales: false,
  reports: false,
  settings: false,
})

const toggleMenu = (menu: keyof typeof openMenus) => {
  openMenus[menu] = !openMenus[menu]
}

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>
