<template>
  <div class="space-y-6">
    <div>
      <h2 class="text-2xl font-bold text-slate-800">Dashboard Overview</h2>
      <p class="text-slate-500">Ringkasan bisnis bulan ini.</p>
    </div>

    <!-- Alert / Low Stock (data dari API) -->
    <div v-if="lowStockAlerts.length > 0" class="bg-amber-50 border-l-4 border-amber-500 p-4 rounded-md">
      <div class="flex">
        <div class="flex-shrink-0">
          <span class="text-amber-500">⚠️</span>
        </div>
        <div class="ml-3">
          <h3 class="text-sm font-medium text-amber-800">Perhatian: Stok Menipis</h3>
          <div class="mt-2 text-sm text-amber-700">
            <ul class="list-disc pl-5 space-y-1">
              <li v-for="item in lowStockAlerts" :key="item.id">{{ item.name }} - sisa {{ item.stock }} unit</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Stat Cards -->
    <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-3">
      <!-- Total Penjualan Bulan Ini -->
      <div class="bg-white overflow-hidden shadow rounded-lg border border-slate-200">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-indigo-100 rounded-md p-3">
              <span class="text-indigo-600 text-xl">💰</span>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-slate-500 truncate">Total Penjualan (Bulan Ini)</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-slate-900">
                    <span v-if="loading" class="animate-pulse text-slate-400">Memuat...</span>
                    <span v-else>Rp {{ formatNumber(stats?.total_revenue_this_month ?? 0) }}</span>
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <!-- Estimasi PPN Terkumpul -->
      <div class="bg-white overflow-hidden shadow rounded-lg border border-slate-200">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-green-100 rounded-md p-3">
              <span class="text-green-600 text-xl">🧾</span>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-slate-500 truncate">Estimasi PPN Terkumpul (Bulan Ini)</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-slate-900">
                    <span v-if="loading" class="animate-pulse text-slate-400">Memuat...</span>
                    <span v-else>Rp {{ formatNumber(stats?.total_ppn_this_month ?? 0) }}</span>
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <!-- Total Faktur -->
      <div class="bg-white overflow-hidden shadow rounded-lg border border-slate-200">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-blue-100 rounded-md p-3">
              <span class="text-blue-600 text-xl">📦</span>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-slate-500 truncate">Total Faktur (Bulan Ini)</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-slate-900">
                    <span v-if="loading" class="animate-pulse text-slate-400">Memuat...</span>
                    <span v-else>{{ stats?.total_invoices_this_month ?? 0 }}</span>
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Chart -->
    <div class="bg-white shadow rounded-lg border border-slate-200 p-6">
      <h3 class="text-lg font-medium text-slate-900 mb-4">Grafik Penjualan 30 Hari Terakhir</h3>
      <div v-if="loading" class="h-96 flex items-center justify-center text-slate-400">
        <span class="animate-pulse">Memuat grafik...</span>
      </div>
      <div v-else-if="chartOption.series[0].data.length === 0" class="h-96 flex items-center justify-center text-slate-400">
        <span>Belum ada data penjualan dalam 30 hari terakhir.</span>
      </div>
      <div v-else class="h-96 w-full">
        <v-chart class="chart" :option="chartOption" autoresize />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
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

const formatNumber = (num: number) => Math.round(num).toLocaleString('id-ID')

// --- Chart ---
const chartOption = ref({
  tooltip: { trigger: 'axis' },
  legend: { data: ['Penjualan (Rp)'] },
  grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
  xAxis: { type: 'category', boundaryGap: false, data: [] as string[] },
  yAxis: {
    type: 'value',
    axisLabel: { formatter: (v: number) => 'Rp ' + (v / 1_000_000).toFixed(0) + ' Jt' }
  },
  series: [
    {
      name: 'Penjualan (Rp)',
      type: 'line',
      smooth: true,
      areaStyle: {},
      itemStyle: { color: '#4f46e5' },
      data: [] as number[]
    }
  ]
})

const updateChart = (chartData: MonthlySalePoint[]) => {
  chartOption.value.xAxis.data = chartData.map(d => d.date)
  chartOption.value.series[0].data = chartData.map(d => d.total)
}

// --- Fetch ---
const fetchStats = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/dashboard')
    stats.value = res.data
    lowStockAlerts.value = res.data.low_stock_products || []
    updateChart(res.data.monthly_sales_chart || [])
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
