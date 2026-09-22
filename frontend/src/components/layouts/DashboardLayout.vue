<template>
  <div class="min-h-screen bg-slate-100 flex font-sans relative overflow-x-hidden">
    
    <!-- Mobile Backdrop Overlay -->
    <div
      v-if="isSidebarOpen"
      class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm z-40 lg:hidden transition-opacity duration-300"
      @click="closeSidebarOnMobile"
    ></div>

    <!-- Sidebar -->
    <aside
      class="fixed lg:static inset-y-0 left-0 z-50 w-72 lg:w-64 flex flex-col transition-transform duration-300 ease-in-out flex-shrink-0 shadow-2xl lg:shadow-none"
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:-ml-64 lg:translate-x-0'"
      style="background: linear-gradient(180deg, #4A1010 0%, #5C1A1A 40%, #3B0D0D 100%);"
    >
      <!-- Logo Header -->
      <div class="h-16 flex items-center justify-between px-4 flex-shrink-0" style="border-bottom: 1px solid rgba(252,129,129,0.15);">
        <h1 class="text-xl font-bold flex items-center gap-3" style="color: #FECDD3;">
          <div
            class="w-9 h-9 rounded-lg flex items-center justify-center flex-shrink-0 shadow-lg"
            style="background: linear-gradient(135deg, #C53030, #9B2C2C); border: 1px solid rgba(252,129,129,0.3);"
          >
            <LayoutDashboard class="w-5 h-5" style="color: #FECDD3;" />
          </div>
          <div>
            <span class="text-base font-bold tracking-tight block leading-none" style="color: #FECDD3;">Trading ERP</span>
            <span class="text-[10px] text-rose-300/60 font-medium tracking-wider uppercase">Enterprise Suite</span>
          </div>
        </h1>

        <!-- Close button for mobile screen -->
        <button
          @click="isSidebarOpen = false"
          class="lg:hidden p-1.5 rounded-lg text-rose-200/70 hover:text-rose-100 hover:bg-white/10 transition-colors"
          title="Tutup Menu"
        >
          <X class="w-5 h-5" />
        </button>
      </div>

      <!-- Navigation -->
      <div class="flex-1 overflow-y-auto py-4 scrollbar-thin">
        <nav class="space-y-0.5 px-3">

          <!-- Dashboard -->
          <router-link
            to="/dashboard"
            class="sidebar-link group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            active-class="sidebar-link-active"
            @click="handleNavClick"
          >
            <LayoutDashboard class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
            Dashboard Overview
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
                  @click="handleNavClick"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Data Produk
                </router-link>
                <router-link
                  to="/dashboard/master-data/partners"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                  @click="handleNavClick"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Data Mitra
                </router-link>
              </div>
            </div>
          </div>

          <!-- Gudang & Stok (1-Click Direct Link) -->
          <router-link
            to="/dashboard/inventory/management"
            class="sidebar-link group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150 mt-0.5"
            active-class="sidebar-link-active"
            @click="handleNavClick"
          >
            <Warehouse class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
            Gudang &amp; Stok
          </router-link>

          <!-- Penjualan & Faktur (1-Click Direct Link) -->
          <router-link
            to="/dashboard/sales/management"
            class="sidebar-link group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150 mt-0.5"
            active-class="sidebar-link-active"
            @click="handleNavClick"
          >
            <ShoppingCart class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
            Penjualan &amp; Faktur
          </router-link>

          <!-- Laporan Laba Rugi (1-Click Direct Link) -->
          <router-link
            to="/dashboard/reports/profit"
            class="sidebar-link group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150 mt-0.5"
            active-class="sidebar-link-active"
            @click="handleNavClick"
          >
            <BarChart3 class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
            Laporan Laba Rugi
          </router-link>

          <!-- Pengaturan -->
          <div class="mt-0.5">
            <button
              @click="toggleMenu('settings')"
              class="sidebar-link group w-full flex items-center justify-between px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150"
            >
              <div class="flex items-center">
                <Settings class="sidebar-icon w-5 h-5 mr-3 flex-shrink-0" />
                Pengaturan Sistem
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
                  @click="handleNavClick"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Profil Perusahaan
                </router-link>
                <router-link
                  to="/dashboard/settings/users"
                  class="sidebar-sublink group flex items-center px-3 py-2 text-sm rounded-lg transition-all duration-150"
                  active-class="sidebar-sublink-active"
                  @click="handleNavClick"
                >
                  <span class="w-1.5 h-1.5 rounded-full mr-2.5 flex-shrink-0" style="background: rgba(252,129,129,0.5);"></span>
                  Kelola Pengguna
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
          class="w-full group flex items-center px-3 py-2.5 text-sm font-medium rounded-lg transition-all duration-150 cursor-pointer"
          style="color: rgba(252,165,165,0.8);"
          @mouseover="($event.currentTarget as HTMLElement).style.background='rgba(252,129,129,0.12)'; ($event.currentTarget as HTMLElement).style.color='#FECDD3'"
          @mouseleave="($event.currentTarget as HTMLElement).style.background='transparent'; ($event.currentTarget as HTMLElement).style.color='rgba(252,165,165,0.8)'"
        >
          <LogOut class="w-5 h-5 mr-3 flex-shrink-0" />
          Keluar Sistem
        </button>
      </div>
    </aside>

    <!-- Main Content wrapper -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden w-full">
      <!-- Modern Redesigned Header -->
      <header class="bg-white/95 backdrop-blur border-b border-slate-200/80 h-16 flex items-center justify-between px-3 sm:px-6 lg:px-8 z-30 sticky top-0 shadow-sm">
        
        <!-- Left: Toggle & Dynamic Breadcrumb -->
        <div class="flex items-center gap-3">
          <button
            @click="isSidebarOpen = !isSidebarOpen"
            class="p-2 rounded-xl text-slate-600 hover:text-slate-900 hover:bg-slate-100 focus:outline-none transition-colors duration-150 cursor-pointer"
            aria-label="Toggle navigation menu"
          >
            <Menu class="h-5 w-5" />
          </button>
          
          <!-- Breadcrumb Path -->
          <div class="flex items-center gap-2 text-xs sm:text-sm font-medium text-slate-500">
            <span class="hidden md:inline-flex items-center gap-1.5 text-slate-400">
              <Building2 class="w-4 h-4 text-rose-800" />
              <span>ERP</span>
            </span>
            <ChevronRight class="hidden md:inline-block w-3.5 h-3.5 text-slate-300" />
            <span class="text-slate-800 font-semibold truncate max-w-[160px] sm:max-w-xs">
              {{ currentRouteTitle }}
            </span>
          </div>
        </div>
        
        <!-- Right: Status, Date, User Info & Quick Action -->
        <div class="flex items-center gap-2.5 sm:gap-4">
          
          <!-- Live Status Badge (Hidden on very small screens) -->
          <div class="hidden sm:inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full bg-emerald-50 border border-emerald-200 text-emerald-700 text-xs font-medium">
            <span class="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
            <span>Online</span>
          </div>

          <!-- Current Date Indicator (Hidden on mobile) -->
          <div class="hidden md:flex items-center gap-1.5 text-xs text-slate-500 font-medium px-2 py-1 bg-slate-50 rounded-lg border border-slate-200/60">
            <Calendar class="w-3.5 h-3.5 text-slate-400" />
            <span>{{ currentDateStr }}</span>
          </div>

          <!-- User Profile Badge & Quick Menu -->
          <div class="flex items-center gap-2.5 pl-2 sm:border-l sm:border-slate-200">
            <div class="flex items-center gap-2">
              <div class="w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0 shadow-sm border border-rose-200" style="background: linear-gradient(135deg, #7B1D1D, #4A1010);">
                <User class="w-4 h-4 text-rose-100" />
              </div>
              <div class="hidden sm:flex flex-col text-left leading-tight">
                <span class="text-xs font-bold text-slate-800 truncate max-w-[110px]">{{ authStore.user?.username }}</span>
                <span class="text-[10px] font-semibold text-rose-700 uppercase tracking-wide">{{ authStore.user?.role }}</span>
              </div>
            </div>

            <!-- Logout button right in header for quick access -->
            <button
              @click="handleLogout"
              class="p-1.5 rounded-lg text-slate-400 hover:text-rose-600 hover:bg-rose-50 transition-colors cursor-pointer"
              title="Keluar"
            >
              <LogOut class="w-4 h-4" />
            </button>
          </div>

        </div>
      </header>

      <!-- Main viewport -->
      <main class="flex-1 relative overflow-y-auto focus:outline-none w-full bg-slate-50/70">
        <div class="py-5 sm:py-7">
          <div class="max-w-7xl mx-auto px-3 sm:px-6 lg:px-8">
            <router-view />
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import {
  LayoutDashboard,
  Database,
  Warehouse,
  ShoppingCart,
  BarChart3,
  Settings,
  ChevronDown,
  ChevronRight,
  LogOut,
  User,
  Menu,
  X,
  Building2,
  Calendar
} from '@lucide/vue'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

// Dynamic Route Title
const currentRouteTitle = computed(() => {
  const name = route.name as string
  switch (name) {
    case 'DashboardOverview': return 'Dashboard Overview'
    case 'Products': return 'Data Master Produk'
    case 'Partners': return 'Data Mitra (Klien & Supplier)'
    case 'StockManagement': return 'Gudang & Stok'
    case 'SalesManagement': return 'Penjualan & Faktur'
    case 'ProfitReport': return 'Laporan Laba Rugi'
    case 'CompanySettings': return 'Profil Perusahaan'
    case 'UsersManagement': return 'Manajemen Pengguna'
    default: return 'ERP Dashboard'
  }
})

// Current Date in Indonesian
const currentDateStr = computed(() => {
  return new Intl.DateTimeFormat('id-ID', {
    weekday: 'short',
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  }).format(new Date())
})

// On mobile (<1024px), sidebar defaults to closed. On desktop, defaults to open.
const isSidebarOpen = ref(window.innerWidth >= 1024)

const checkWindowSize = () => {
  if (window.innerWidth >= 1024) {
    isSidebarOpen.value = true
  } else {
    isSidebarOpen.value = false
  }
}

onMounted(() => {
  window.addEventListener('resize', checkWindowSize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', checkWindowSize)
})

const closeSidebarOnMobile = () => {
  if (window.innerWidth < 1024) {
    isSidebarOpen.value = false
  }
}

const handleNavClick = () => {
  if (window.innerWidth < 1024) {
    isSidebarOpen.value = false
  }
}

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
