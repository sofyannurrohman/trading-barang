<template>
  <div class="relative" ref="dropdownRef">
    <!-- Bell Trigger Button -->
    <button
      @click="toggleDropdown"
      class="relative p-2 rounded-xl text-slate-600 hover:text-slate-900 hover:bg-slate-100 transition-colors duration-150 focus:outline-none cursor-pointer"
      :class="{ 'bg-slate-100 text-slate-900': isOpen }"
      title="Notifikasi & Peringatan Sistem"
      aria-label="Notifikasi"
    >
      <Bell class="w-5 h-5" />
      
      <!-- Unread Indicator Badge -->
      <span
        v-if="unreadCount > 0"
        class="absolute top-1 right-1 flex h-4 min-w-4 px-1 items-center justify-center rounded-full bg-rose-600 text-[10px] font-bold text-white shadow-sm ring-2 ring-white animate-pulse"
      >
        {{ unreadCount > 9 ? '9+' : unreadCount }}
      </span>
    </button>

    <!-- Dropdown Menu -->
    <transition
      enter-active-class="transition ease-out duration-150 transform"
      enter-from-class="opacity-0 scale-95 -translate-y-1"
      enter-to-class="opacity-100 scale-100 translate-y-0"
      leave-active-class="transition ease-in duration-100 transform"
      leave-from-class="opacity-100 scale-100 translate-y-0"
      leave-to-class="opacity-0 scale-95 -translate-y-1"
    >
      <div
        v-if="isOpen"
        class="absolute right-0 mt-2 w-80 sm:w-96 bg-white rounded-2xl shadow-xl border border-slate-200/80 z-50 overflow-hidden flex flex-col max-h-[480px]"
      >
        <!-- Header -->
        <div class="p-4 border-b border-slate-100 flex items-center justify-between bg-slate-50/70">
          <div class="flex items-center gap-2">
            <h3 class="font-semibold text-slate-800 text-sm">Notifikasi Sistem</h3>
            <span
              v-if="unreadCount > 0"
              class="px-2 py-0.5 text-xs font-semibold bg-rose-100 text-rose-700 rounded-full"
            >
              {{ unreadCount }} Baru
            </span>
          </div>
          
          <button
            v-if="notifications.length > 0"
            @click="markAllAsRead"
            class="text-xs text-rose-600 hover:text-rose-700 font-medium transition-colors cursor-pointer"
          >
            Tandai dibaca
          </button>
        </div>

        <!-- Notification List -->
        <div class="overflow-y-auto flex-1 divide-y divide-slate-100 scrollbar-thin">
          <div
            v-if="loading"
            class="p-6 text-center text-slate-400 text-xs flex flex-col items-center justify-center gap-2"
          >
            <div class="w-5 h-5 border-2 border-rose-500 border-t-transparent rounded-full animate-spin"></div>
            <span>Memuat notifikasi...</span>
          </div>

          <div
            v-else-if="notifications.length === 0"
            class="p-8 text-center flex flex-col items-center justify-center text-slate-400"
          >
            <div class="w-10 h-10 rounded-full bg-slate-100 flex items-center justify-center mb-2">
              <CheckCircle2 class="w-5 h-5 text-slate-400" />
            </div>
            <p class="text-xs font-medium text-slate-600">Tidak ada notifikasi baru</p>
            <p class="text-[11px] text-slate-400 mt-0.5">Semua status operasional aman</p>
          </div>

          <div
            v-else
            v-for="item in notifications"
            :key="item.id"
            @click="handleNotificationClick(item)"
            class="p-3.5 hover:bg-slate-50 transition-colors cursor-pointer flex items-start gap-3 relative group"
            :class="{ 'bg-rose-50/30': !item.read }"
          >
            <!-- Severity Icon -->
            <div
              class="flex-shrink-0 w-8 h-8 rounded-lg flex items-center justify-center mt-0.5"
              :class="{
                'bg-amber-100 text-amber-600': item.type === 'warning',
                'bg-rose-100 text-rose-600': item.type === 'danger',
                'bg-blue-100 text-blue-600': item.type === 'info',
                'bg-emerald-100 text-emerald-600': item.type === 'success'
              }"
            >
              <AlertTriangle v-if="item.type === 'warning' || item.type === 'danger'" class="w-4 h-4" />
              <Info v-else-if="item.type === 'info'" class="w-4 h-4" />
              <CheckCircle2 v-else class="w-4 h-4" />
            </div>

            <!-- Content -->
            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between gap-1">
                <p class="text-xs font-semibold text-slate-800 truncate">{{ item.title }}</p>
                <span class="text-[10px] text-slate-400 flex-shrink-0">{{ item.time }}</span>
              </div>
              <p class="text-xs text-slate-600 mt-0.5 line-clamp-2 leading-relaxed">{{ item.message }}</p>
            </div>

            <!-- Unread Dot -->
            <span
              v-if="!item.read"
              class="w-2 h-2 rounded-full bg-rose-500 flex-shrink-0 mt-1.5"
            ></span>
          </div>
        </div>

        <!-- Footer -->
        <div class="p-2.5 border-t border-slate-100 bg-slate-50/50 text-center">
          <router-link
            to="/dashboard/inventory/management"
            @click="isOpen = false"
            class="text-xs font-medium text-rose-600 hover:text-rose-700 transition-colors inline-flex items-center gap-1"
          >
            <span>Cek Manajemen Gudang & Stok</span>
            <ChevronRight class="w-3.5 h-3.5" />
          </router-link>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { Bell, AlertTriangle, Info, CheckCircle2, ChevronRight } from '@lucide/vue'
import api from '@/plugins/axios'

interface NotificationItem {
  id: string
  title: string
  message: string
  type: 'warning' | 'danger' | 'info' | 'success'
  time: string
  read: boolean
  route?: string
}

const router = useRouter()
const isOpen = ref(false)
const loading = ref(false)
const dropdownRef = ref<HTMLElement | null>(null)
const notifications = ref<NotificationItem[]>([])

const unreadCount = computed(() => {
  return notifications.value.filter(n => !n.read).length
})

const toggleDropdown = () => {
  isOpen.value = !isOpen.value
  if (isOpen.value) {
    fetchAlerts()
  }
}

const fetchAlerts = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/dashboard/overview', { silent: true })
    const lowStock = res.data?.data?.low_stock_products || res.data?.low_stock_products || []
    
    const items: NotificationItem[] = []
    
    if (Array.isArray(lowStock) && lowStock.length > 0) {
      lowStock.forEach((prod: any) => {
        items.push({
          id: `stock-${prod.id}`,
          title: `Stok Kritis: ${prod.name}`,
          message: `Sisa stok hanya ${prod.stock} unit. Segera lakukan restock di gudang.`,
          type: prod.stock <= 3 ? 'danger' : 'warning',
          time: 'Hari ini',
          read: false,
          route: '/dashboard/inventory/management',
        })
      })
    }

    // Persist read state if previously marked
    const readIds = JSON.parse(localStorage.getItem('read_notifications') || '[]')
    notifications.value = items.map(item => ({
      ...item,
      read: readIds.includes(item.id)
    }))
  } catch (error) {
    console.error('Failed to load notifications', error)
  } finally {
    loading.value = false
  }
}

const markAllAsRead = () => {
  const readIds = notifications.value.map(n => n.id)
  localStorage.setItem('read_notifications', JSON.stringify(readIds))
  notifications.value = notifications.value.map(n => ({ ...n, read: true }))
}

const handleNotificationClick = (item: NotificationItem) => {
  item.read = true
  const readIds = JSON.parse(localStorage.getItem('read_notifications') || '[]')
  if (!readIds.includes(item.id)) {
    readIds.push(item.id)
    localStorage.setItem('read_notifications', JSON.stringify(readIds))
  }
  if (item.route) {
    isOpen.value = false
    router.push(item.route)
  }
}

const handleClickOutside = (event: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
    isOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  fetchAlerts()
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>
