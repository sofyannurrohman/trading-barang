<template>
  <div class="space-y-6 sm:space-y-8">
    
    <!-- Hero Welcome Banner -->
    <div 
      class="rounded-2xl p-5 sm:p-7 text-white shadow-xl relative overflow-hidden flex flex-col md:flex-row items-start md:items-center justify-between gap-6"
      style="background: radial-gradient(circle at 10% 20%, #681A1A 0%, #3B0D0D 60%, #1A0505 100%); border: 1px solid rgba(254, 205, 211, 0.15);"
    >
      <!-- Ambient Glow & Decor -->
      <div class="absolute -right-16 -top-16 w-64 h-64 rounded-full bg-rose-500/10 blur-3xl pointer-events-none"></div>
      <div class="absolute right-32 bottom-0 w-48 h-48 rounded-full bg-amber-500/10 blur-2xl pointer-events-none"></div>

      <!-- Welcome Left Content -->
      <div class="relative z-10 space-y-2 max-w-xl">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-white/10 text-rose-200 text-xs font-semibold backdrop-blur-sm border border-white/10">
          <Sparkles class="w-3.5 h-3.5 text-amber-300" />
          <span>{{ greetingTime }} &bull; Pusat Kendali ERP</span>
        </div>
        <h1 class="text-2xl sm:text-3xl font-extrabold tracking-tight text-white leading-tight">
          Selamat Datang, <span class="bg-gradient-to-r from-rose-200 via-pink-100 to-amber-200 bg-clip-text text-transparent">{{ authStore.user?.username || 'Admin' }}</span>!
        </h1>
        <p class="text-sm sm:text-base text-rose-100/80 leading-relaxed">
          Semua transaksi, persediaan barang di gudang, dan laporan performa penjualan siap dipantau secara real-time.
        </p>
      </div>

      <!-- Quick Action Shortcuts -->
      <div class="relative z-10 flex flex-wrap sm:flex-nowrap items-center gap-3 w-full md:w-auto">
        <router-link
          to="/dashboard/sales/create"
          class="flex-1 sm:flex-initial inline-flex items-center justify-center gap-2 px-4 py-2.5 rounded-xl bg-gradient-to-r from-rose-500 to-red-600 hover:from-rose-400 hover:to-red-500 text-white text-xs sm:text-sm font-bold shadow-lg shadow-rose-950/40 transition-all cursor-pointer group"
        >
          <PlusCircle class="w-4 h-4 text-rose-200 group-hover:rotate-90 transition-transform duration-200" />
          <span>Faktur Baru</span>
        </router-link>
        <router-link
          to="/dashboard/inventory/management"
          class="flex-1 sm:flex-initial inline-flex items-center justify-center gap-2 px-4 py-2.5 rounded-xl bg-white/10 hover:bg-white/15 text-rose-100 text-xs sm:text-sm font-semibold border border-white/15 backdrop-blur-sm transition-colors cursor-pointer"
        >
          <Boxes class="w-4 h-4 text-rose-300" />
          <span>Kelola Stok</span>
        </router-link>
      </div>
    </div>

    <!-- Smart Stock Alert Banner (If Any) -->
    <div v-if="lowStockAlerts.length > 0" class="bg-amber-500/10 border border-amber-500/30 rounded-2xl p-4 sm:p-5 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div class="flex items-start gap-3.5">
        <div class="p-2.5 bg-amber-500/20 text-amber-600 rounded-xl flex-shrink-0 mt-0.5">
          <AlertTriangle class="w-5 h-5" />
        </div>
        <div>
          <h3 class="text-sm font-bold text-amber-900">Perhatian: {{ lowStockAlerts.length }} Produk Membutuhkan Penambahan Stok</h3>
          <div class="text-xs text-amber-800/80 mt-1 flex flex-wrap gap-x-3 gap-y-1">
            <span v-for="item in lowStockAlerts.slice(0, 4)" :key="item.id" class="inline-flex items-center gap-1 font-medium">
              &bull; {{ item.name }} (sisa <strong class="text-amber-950">{{ item.stock }}</strong>)
            </span>
            <span v-if="lowStockAlerts.length > 4" class="italic text-amber-700 font-semibold">
              + {{ lowStockAlerts.length - 4 }} produk lainnya
            </span>
          </div>
        </div>
      </div>
      <router-link
        to="/dashboard/inventory/management"
        class="inline-flex items-center gap-1.5 px-3.5 py-1.5 rounded-lg bg-amber-600 hover:bg-amber-700 text-white text-xs font-bold transition-colors cursor-pointer flex-shrink-0 w-full sm:w-auto justify-center"
      >
        <span>Restok Sekarang</span>
        <ArrowUpRight class="w-3.5 h-3.5" />
      </router-link>
    </div>

    <!-- Modern Key Metrics Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5 sm:gap-6">
      
      <!-- Total Penjualan Bulan Ini -->
      <div class="bg-white rounded-2xl p-5 sm:p-6 shadow-sm border border-slate-200/80 hover:shadow-md transition-shadow relative overflow-hidden group">
        <div class="flex items-center justify-between">
          <div class="w-12 h-12 rounded-xl bg-indigo-50 flex items-center justify-center text-indigo-600 border border-indigo-100">
            <TrendingUp class="w-6 h-6" />
          </div>
          <span class="text-[11px] font-semibold text-indigo-700 bg-indigo-50 px-2.5 py-1 rounded-full border border-indigo-100">
            Bulan Ini
          </span>
        </div>
        <div class="mt-4">
          <p class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Total Penjualan</p>
          <div class="text-2xl sm:text-3xl font-extrabold text-slate-900 mt-1 tracking-tight">
            <span v-if="loading" class="animate-pulse text-slate-400 text-xl font-normal">Memuat data...</span>
            <span v-else>Rp {{ formatNumber(stats?.total_revenue_this_month ?? 0) }}</span>
          </div>
          <div class="mt-2 flex items-center gap-1.5 text-xs text-slate-500">
            <CheckCircle2 class="w-3.5 h-3.5 text-emerald-500" />
            <span>Terhitung dari transaksi penjualan</span>
          </div>
        </div>
      </div>

      <!-- Total Pendapatan Penjualan -->
      <div class="bg-white rounded-2xl p-5 sm:p-6 shadow-sm border border-slate-200/80 hover:shadow-md transition-shadow relative overflow-hidden group">
        <div class="flex items-center justify-between">
          <div class="w-12 h-12 rounded-xl bg-emerald-50 flex items-center justify-center text-emerald-600 border border-emerald-100">
            <Receipt class="w-6 h-6" />
          </div>
          <span class="text-[11px] font-semibold text-emerald-700 bg-emerald-50 px-2.5 py-1 rounded-full border border-emerald-100">
            Akumulasi
          </span>
        </div>
        <div class="mt-4">
          <p class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Total Transaksi</p>
          <div class="text-2xl sm:text-3xl font-extrabold text-emerald-600 mt-1 tracking-tight">
            <span v-if="loading" class="animate-pulse text-slate-400 text-xl font-normal">Memuat data...</span>
            <span v-else>Rp {{ formatNumber(stats?.total_revenue_this_month ?? 0) }}</span>
          </div>
          <div class="mt-2 flex items-center gap-1.5 text-xs text-slate-500">
            <ShieldCheck class="w-3.5 h-3.5 text-emerald-500" />
            <span>Rekapitulasi penjualan bulan ini</span>
          </div>
        </div>
      </div>

      <!-- Total Faktur Diterbitkan -->
      <div class="bg-white rounded-2xl p-5 sm:p-6 shadow-sm border border-slate-200/80 hover:shadow-md transition-shadow relative overflow-hidden group sm:col-span-2 lg:col-span-1">
        <div class="flex items-center justify-between">
          <div class="w-12 h-12 rounded-xl bg-rose-50 flex items-center justify-center text-rose-700 border border-rose-100">
            <FileSpreadsheet class="w-6 h-6" />
          </div>
          <span class="text-[11px] font-semibold text-rose-700 bg-rose-50 px-2.5 py-1 rounded-full border border-rose-100">
            Faktur Dibuat
          </span>
        </div>
        <div class="mt-4">
          <p class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Total Faktur Penjualan</p>
          <div class="text-2xl sm:text-3xl font-extrabold text-slate-900 mt-1 tracking-tight">
            <span v-if="loading" class="animate-pulse text-slate-400 text-xl font-normal">Memuat data...</span>
            <span v-else>{{ stats?.total_invoices_this_month ?? 0 }} <span class="text-base font-semibold text-slate-500">faktur</span></span>
          </div>
          <div class="mt-2 flex items-center gap-1.5 text-xs text-slate-500">
            <router-link to="/dashboard/sales/list" class="text-rose-700 hover:text-rose-800 font-semibold inline-flex items-center gap-1">
              <span>Buka daftar faktur</span>
              <ArrowUpRight class="w-3 h-3" />
            </router-link>
          </div>
        </div>
      </div>

    </div>

    <!-- Chart & Analytics Section -->
    <div class="bg-white rounded-2xl shadow-sm border border-slate-200/80 p-5 sm:p-7">
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6 pb-4 border-b border-slate-100">
        <div>
          <h3 class="text-base sm:text-lg font-bold text-slate-900 flex items-center gap-2">
            <BarChart3 class="w-5 h-5 text-rose-700" />
            <span>Grafik Penjualan 30 Hari Terakhir</span>
          </h3>
          <p class="text-xs sm:text-sm text-slate-500 mt-0.5">Tren omzet transaksi harian (data riil dari sistem).</p>
        </div>
        <div class="flex items-center gap-2">
          <button
            @click="fetchStats"
            :disabled="loading"
            class="inline-flex items-center gap-1.5 text-xs font-semibold text-slate-600 hover:text-slate-900 bg-slate-100 hover:bg-slate-200 px-3 py-1.5 rounded-lg border border-slate-200/60 transition-colors cursor-pointer disabled:opacity-50"
            title="Muat Ulang Data"
          >
            <RefreshCw class="w-3.5 h-3.5" :class="{ 'animate-spin': loading }" />
            <span>Refresh</span>
          </button>
          <span class="inline-flex items-center gap-1.5 text-xs font-semibold text-rose-900 bg-rose-50 px-3 py-1.5 rounded-lg border border-rose-200">
            <span class="w-2.5 h-2.5 rounded-full" style="background: #9B2C2C;"></span>
            <span>Omzet (Rp)</span>
          </span>
        </div>
      </div>

      <div v-if="loading" class="h-80 flex flex-col items-center justify-center text-slate-400 gap-2">
        <Loader2 class="w-6 h-6 animate-spin text-rose-600" />
        <span class="text-xs font-medium">Mengambil data grafik dari server...</span>
      </div>
      <div v-else class="h-80 sm:h-96 w-full">
        <v-chart class="chart" :option="chartOption" autoresize />
      </div>
    </div>

    <!-- Quick Module Navigation Hub -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-5">
      
      <!-- Master Data Hub -->
      <router-link to="/dashboard/master-data/products" class="p-5 rounded-2xl bg-white border border-slate-200/80 shadow-sm hover:border-rose-300 hover:shadow-md transition-all group">
        <div class="flex items-center justify-between">
          <div class="w-10 h-10 rounded-xl bg-rose-50 text-rose-700 flex items-center justify-center group-hover:bg-rose-700 group-hover:text-white transition-colors">
            <Database class="w-5 h-5" />
          </div>
          <ArrowUpRight class="w-4 h-4 text-slate-400 group-hover:text-rose-700 transition-colors" />
        </div>
        <h4 class="text-sm font-bold text-slate-900 mt-4">Master Data &amp; SKU</h4>
        <p class="text-xs text-slate-500 mt-1">Katalog produk, harga standar, serta database mitra klien &amp; supplier.</p>
      </router-link>

      <!-- Gudang & Stok Hub -->
      <router-link to="/dashboard/inventory/management" class="p-5 rounded-2xl bg-white border border-slate-200/80 shadow-sm hover:border-amber-300 hover:shadow-md transition-all group">
        <div class="flex items-center justify-between">
          <div class="w-10 h-10 rounded-xl bg-amber-50 text-amber-700 flex items-center justify-center group-hover:bg-amber-600 group-hover:text-white transition-colors">
            <Warehouse class="w-5 h-5" />
          </div>
          <ArrowUpRight class="w-4 h-4 text-slate-400 group-hover:text-amber-600 transition-colors" />
        </div>
        <h4 class="text-sm font-bold text-slate-900 mt-4">Logistik &amp; Gudang</h4>
        <p class="text-xs text-slate-500 mt-1">Penerimaan barang masuk (inbound), penyesuaian stock opname, dan kartu stok.</p>
      </router-link>

      <!-- Laporan Finansial Hub -->
      <router-link to="/dashboard/reports/profit" class="p-5 rounded-2xl bg-white border border-slate-200/80 shadow-sm hover:border-emerald-300 hover:shadow-md transition-all group">
        <div class="flex items-center justify-between">
          <div class="w-10 h-10 rounded-xl bg-emerald-50 text-emerald-700 flex items-center justify-center group-hover:bg-emerald-600 group-hover:text-white transition-colors">
            <BarChart3 class="w-5 h-5" />
          </div>
          <ArrowUpRight class="w-4 h-4 text-slate-400 group-hover:text-emerald-600 transition-colors" />
        </div>
        <h4 class="text-sm font-bold text-slate-900 mt-4">Laba Rugi &amp; Analitik</h4>
        <p class="text-xs text-slate-500 mt-1">Estimasi laba kotor berbasis HPP aktual dan analisa margin penjualan.</p>
      </router-link>

    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import VChart, { THEME_KEY } from 'vue-echarts'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart, BarChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
} from 'echarts/components'
import { provide } from 'vue'
import api from '@/plugins/axios'
import {
  Sparkles,
  TrendingUp,
  Receipt,
  FileSpreadsheet,
  AlertTriangle,
  Boxes,
  PlusCircle,
  ArrowUpRight,
  ShieldCheck,
  CheckCircle2,
  Database,
  Warehouse,
  BarChart3,
  Loader2,
  RefreshCw
} from '@lucide/vue'

use([
  CanvasRenderer,
  LineChart,
  BarChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
])

provide(THEME_KEY, 'light')

const authStore = useAuthStore()

// Greeting based on client time
const greetingTime = computed(() => {
  const hour = new Date().getHours()
  if (hour >= 4 && hour < 11) return 'Selamat Pagi'
  if (hour >= 11 && hour < 15) return 'Selamat Siang'
  if (hour >= 15 && hour < 18) return 'Selamat Sore'
  return 'Selamat Malam'
})

// --- Types ---
interface LowStockItem { id: number; name: string; stock: number }
interface MonthlySalePoint { date: string; total: number }
interface DashboardStats {
  total_revenue_this_month: number
  total_ppn_this_month: number
  total_invoices_this_month: number
  low_stock_products: LowStockItem[]
  monthly_sales_chart: MonthlySalePoint[]
}

// --- State ---
const stats = ref<DashboardStats | null>(null)
const loading = ref(true)
const lowStockAlerts = ref<LowStockItem[]>([])
const monthlySales = ref<MonthlySalePoint[]>([])

const formatNumber = (num: number) => Math.round(num).toLocaleString('id-ID')

const formatDateLabel = (dateStr: string) => {
  if (!dateStr) return ''
  const parts = dateStr.split('-')
  if (parts.length === 3) {
    const day = parts[2]
    const months = ['Jan', 'Feb', 'Mar', 'Apr', 'Mei', 'Jun', 'Jul', 'Agu', 'Sep', 'Okt', 'Nov', 'Des']
    const monthIndex = parseInt(parts[1], 10) - 1
    return `${day} ${months[monthIndex] || ''}`
  }
  return dateStr
}

// --- Dynamic Computed Chart Options ---
const chartOption = computed(() => {
  const dates = monthlySales.value.map(d => formatDateLabel(d.date))
  const rawDates = monthlySales.value.map(d => d.date)
  const totals = monthlySales.value.map(d => d.total)

  return {
    tooltip: {
      trigger: 'axis',
      backgroundColor: '#1e293b',
      borderColor: '#334155',
      textStyle: { color: '#f8fafc', fontSize: 12 },
      formatter: (params: any) => {
        const p = Array.isArray(params) ? params[0] : params
        const originalDate = rawDates[p.dataIndex] || p.name
        return `<div class="font-sans">
          <div class="text-[11px] text-slate-400 font-medium mb-1">${originalDate}</div>
          <div class="font-bold text-rose-300">Rp ${Number(p.value).toLocaleString('id-ID')}</div>
        </div>`
      }
    },
    grid: { left: '2%', right: '3%', bottom: '2%', top: '10%', containLabel: true },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: dates,
      axisLine: { lineStyle: { color: '#cbd5e1' } },
      axisLabel: { color: '#64748b', fontSize: 10, interval: 3 }
    },
    yAxis: {
      type: 'value',
      splitLine: { lineStyle: { color: '#f1f5f9', type: 'dashed' } },
      axisLabel: {
        color: '#64748b',
        fontSize: 11,
        formatter: (v: number) => {
          if (v >= 1_000_000) return 'Rp ' + (v / 1_000_000).toFixed(0) + ' Jt'
          if (v >= 1_000) return 'Rp ' + (v / 1_000).toFixed(0) + ' Rb'
          return 'Rp ' + v
        }
      }
    },
    series: [
      {
        name: 'Penjualan (Rp)',
        type: 'line',
        smooth: true,
        showSymbol: false,
        lineStyle: {
          width: 3,
          color: '#9B2C2C'
        },
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              { offset: 0, color: 'rgba(155, 44, 44, 0.35)' },
              { offset: 1, color: 'rgba(155, 44, 44, 0.00)' }
            ]
          }
        },
        itemStyle: { color: '#9B2C2C' },
        data: totals
      }
    ]
  }
})

// --- Fetch Stats & Chart ---
const fetchStats = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/dashboard')
    stats.value = res.data
    lowStockAlerts.value = res.data.low_stock_products || []
    monthlySales.value = res.data.monthly_sales_chart || []
  } catch (e) {
    console.error('Failed to fetch dashboard stats', e)
  } finally {
    loading.value = false
  }
}

onMounted(fetchStats)
</script>

<style scoped>
.chart {
  width: 100%;
  height: 100%;
}
</style>
