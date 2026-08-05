<template>
  <div class="min-h-screen bg-slate-50 flex font-sans">
    <!-- Sidebar -->
    <aside
      class="w-64 flex flex-col transition-all duration-300 flex-shrink-0 shadow-2xl"
      :class="{ '-ml-64': !isSidebarOpen }"
      style="background: linear-gradient(180deg, #4A1010 0%, #5C1A1A 40%, #5C1A1A 100%);"
    >
      <!-- Logo Header -->
      <div class="h-16 flex items-center px-4 flex-shrink-0" style="border-bottom: 1px solid rgba(252,129,129,0.15);">
        <h1 class="text-xl font-bold flex items-center gap-3" style="color: #FECDD3;">
          <div
            class="w-9 h-9 rounded-lg flex items-center justify-center flex-shrink-0 shadow-lg"
            style="background: linear-gradient(135deg, #C53030, #9B2C2C); border: 1px solid rgba(252,129,129,0.3);"
          >
            <LayoutDashboard class="w-5 h-5" style="color: #FECDD3;" />
          </div>
          <span class="text-base font-semibold tracking-wide" style="color: #FECDD3;">ERP System</span>
        </h1>
      </div>

      <!-- Navigation -->
      <div class="flex-1 overflow-y-auto py-4 scrollbar-thin">
        <nav class="space-y-0.5 px-3">

          <!-- Dashboard -->
          <router-link
            to="/dashboard"
            class="sidebar-link group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            active-class="sidebar-link-active"
          >
            <LayoutDashboard class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
            Dashboard
          </router-link>

          <!-- Master Data -->
          <div class="mt-1">
            <button
              @click="toggleMenu('master')"
              class="sidebar-link group w-full flex items-center justify-between px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            >
              <div class="flex items-center">
                <Database class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
                Master Data
              </div>
              <ChevronDown
                class="w-4 h-4 flex-shrink-0 transition-transform duration-200"
                :class="openMenus.master ? 'rotate-0' : '-rotate-90'"
                style="color: rgba(252,129,129,0.6);"
              />
            </button>
            <div v-show="openMenus.master" class="mt-0.5 space-y-0.5 pl-4">
              <div class="pl-3 border-l" style="border-color: rgba(252,129,129,0.2);">
                <router-link
                  to="/dashboard/master-data/products"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Data Produk
                </router-link>
                <router-link
                  to="/dashboard/master-data/partners"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Data Mitra
                </router-link>
              </div>
            </div>
          </div>

          <!-- Gudang & Inventori -->
          <div class="mt-0.5">
            <button
              @click="toggleMenu('inventory')"
              class="sidebar-link group w-full flex items-center justify-between px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            >
              <div class="flex items-center">
                <Warehouse class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
                Gudang &amp; Inventori
              </div>
              <ChevronDown
                class="w-4 h-4 flex-shrink-0 transition-transform duration-200"
                :class="openMenus.inventory ? 'rotate-0' : '-rotate-90'"
                style="color: rgba(252,129,129,0.6);"
              />
            </button>
            <div v-show="openMenus.inventory" class="mt-0.5 space-y-0.5 pl-4">
              <div class="pl-3 border-l" style="border-color: rgba(252,129,129,0.2);">
                <router-link
                  to="/dashboard/inventory/management"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Manajemen Stok
                </router-link>
              </div>
            </div>
          </div>

          <!-- Penjualan & Faktur -->
          <div class="mt-0.5">
            <button
              @click="toggleMenu('sales')"
              class="sidebar-link group w-full flex items-center justify-between px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            >
              <div class="flex items-center">
                <ShoppingCart class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
                Penjualan &amp; Faktur
              </div>
              <ChevronDown
                class="w-4 h-4 flex-shrink-0 transition-transform duration-200"
                :class="openMenus.sales ? 'rotate-0' : '-rotate-90'"
                style="color: rgba(252,129,129,0.6);"
              />
            </button>
            <div v-show="openMenus.sales" class="mt-0.5 space-y-0.5 pl-4">
              <div class="pl-3 border-l" style="border-color: rgba(252,129,129,0.2);">
                <router-link
                  to="/dashboard/sales/management"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Manajemen Penjualan
                </router-link>
                <router-link
                  to="/dashboard/sales/create"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Buat Faktur
                </router-link>
                <router-link
                  to="/dashboard/sales/list"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Daftar Faktur
                </router-link>
              </div>
            </div>
          </div>

          <!-- Laporan -->
          <div class="mt-0.5">
            <button
              @click="toggleMenu('reports')"
              class="sidebar-link group w-full flex items-center justify-between px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            >
              <div class="flex items-center">
                <BarChart3 class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
                Laporan
              </div>
              <ChevronDown
                class="w-4 h-4 flex-shrink-0 transition-transform duration-200"
                :class="openMenus.reports ? 'rotate-0' : '-rotate-90'"
                style="color: rgba(252,129,129,0.6);"
              />
            </button>
            <div v-show="openMenus.reports" class="mt-0.5 space-y-0.5 pl-4">
              <div class="pl-3 border-l" style="border-color: rgba(252,129,129,0.2);">
                <router-link
                  to="/dashboard/reports/profit"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Laba Kotor
                </router-link>
                <router-link
                  to="/dashboard/reports/tax"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Pajak Keluaran
                </router-link>
              </div>
            </div>
          </div>

          <!-- Pengaturan -->
          <div class="mt-0.5">
            <button
              @click="toggleMenu('settings')"
              class="sidebar-link group w-full flex items-center justify-between px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            >
              <div class="flex items-center">
                <Settings class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
                Pengaturan
              </div>
              <ChevronDown
                class="w-4 h-4 flex-shrink-0 transition-transform duration-200"
                :class="openMenus.settings ? 'rotate-0' : '-rotate-90'"
                style="color: rgba(252,129,129,0.6);"
              />
            </button>
            <div v-show="openMenus.settings" class="mt-0.5 space-y-0.5 pl-4">
              <div class="pl-3 border-l" style="border-color: rgba(252,129,129,0.2);">
                <router-link
                  to="/dashboard/settings/company"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Profil Perusahaan
                </router-link>
                <router-link
                  to="/dashboard/settings/users"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Pengguna
                </router-link>
              </div>
            </div>
          </div>

        </nav>
      </div>

      <!-- Footer User Info -->
      <div class="px-3 py-3 flex-shrink-0" style="border-top: 1px solid rgba(252,129,129,0.15);">
        <button
          @click="handleLogout"
          class="w-full group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
          style="color: rgba(252,165,165,0.8);"
          @mouseover="($event.currentTarget as HTMLElement).style.background='rgba(252,129,129,0.12)'; ($event.currentTarget as HTMLElement).style.color='#FECDD3'"
          @mouseleave="($event.currentTarget as HTMLElement).style.background='transparent'; ($event.currentTarget as HTMLElement).style.color='rgba(252,165,165,0.8)'"
        >
          <LogOut class="w-5 h-5 mr-3 flex-shrink-0" />
          Logout
        </button>
      </div>
    </aside>

    <!-- Main Content wrapper -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Header -->
      <header class="bg-white border-b border-slate-200 h-16 flex items-center justify-between px-4 sm:px-6 lg:px-8 z-10" style="box-shadow: 0 1px 3px rgba(0,0,0,0.06);">
        <div class="flex items-center">
          <button
            @click="isSidebarOpen = !isSidebarOpen"
            class="p-2 rounded-md text-slate-500 hover:text-slate-700 hover:bg-slate-100 focus:outline-none transition-colors duration-150"
          >
            <svg class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
        <div class="flex items-center gap-3">
          <div class="flex items-center gap-2 text-sm text-slate-600">
            <div class="w-7 h-7 rounded-full flex items-center justify-center" style="background: #5C1A1A;">
              <User class="w-4 h-4 text-red-200" />
            </div>
            <span class="font-medium">{{ authStore.user?.username }}</span>
            <span class="text-xs px-2 py-0.5 rounded-full bg-red-50 text-red-700 font-medium border border-red-200">{{ authStore.user?.role }}</span>
          </div>
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
import {
  LayoutDashboard,
  Database,
  Warehouse,
  ShoppingCart,
  BarChart3,
  Settings,
  ChevronDown,
  LogOut,
  User,
} from '@lucide/vue'

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

<style scoped>
/* Sidebar link base */
.sidebar-link {
  color: rgba(252, 165, 165, 0.85);
}

.sidebar-link:hover {
  background: rgba(252, 129, 129, 0.12);
  color: #FECDD3;
}

/* Active state with left border */
.sidebar-link-active {
  background: rgba(252, 129, 129, 0.18) !important;
  color: #FECDD3 !important;
  box-shadow: inset 3px 0 0 #FC8181;
}

.sidebar-link-active .sidebar-icon {
  color: #FCA5A5 !important;
}

/* Icon color */
.sidebar-icon {
  color: rgba(252, 129, 129, 0.7);
  transition: color 0.15s ease;
}

.sidebar-link:hover .sidebar-icon {
  color: #FCA5A5;
}

/* Sub-link styles */
.sidebar-sublink {
  color: rgba(252, 165, 165, 0.7);
}

.sidebar-sublink:hover {
  background: rgba(252, 129, 129, 0.10);
  color: #FECDD3;
}

.sidebar-sublink-active {
  color: #FECDD3 !important;
  background: rgba(252, 129, 129, 0.15) !important;
}

.sidebar-sublink-active span {
  background: #FC8181 !important;
}

/* Thin scrollbar for nav */
.scrollbar-thin::-webkit-scrollbar {
  width: 3px;
}
.scrollbar-thin::-webkit-scrollbar-track {
  background: transparent;
}
.scrollbar-thin::-webkit-scrollbar-thumb {
  background: rgba(252, 129, 129, 0.2);
  border-radius: 999px;
}
</style>
