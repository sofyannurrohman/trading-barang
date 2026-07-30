<template>
  <div class="space-y-6">
    <div>
      <h2 class="text-2xl font-bold text-slate-800">Dashboard Overview</h2>
      <p class="text-slate-500">Ringkasan bisnis bulan ini.</p>
    </div>

    <!-- Alert / Low Stock -->
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
      <div class="bg-white overflow-hidden shadow rounded-lg border border-slate-200">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-indigo-100 rounded-md p-3">
              <span class="text-indigo-600 text-xl">💰</span>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-slate-500 truncate">Total Penjualan</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-slate-900">Rp 125.000.000</div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white overflow-hidden shadow rounded-lg border border-slate-200">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-green-100 rounded-md p-3">
              <span class="text-green-600 text-xl">🧾</span>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-slate-500 truncate">Estimasi PPN Terkumpul</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-slate-900">Rp 13.750.000</div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white overflow-hidden shadow rounded-lg border border-slate-200">
        <div class="p-5">
          <div class="flex items-center">
            <div class="flex-shrink-0 bg-blue-100 rounded-md p-3">
              <span class="text-blue-600 text-xl">📦</span>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-slate-500 truncate">Total Faktur</dt>
                <dd class="flex items-baseline">
                  <div class="text-2xl font-semibold text-slate-900">45</div>
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
      <div class="h-96 w-full">
        <v-chart class="chart" :option="chartOption" autoresize />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
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

use([
  CanvasRenderer,
  LineChart,
  BarChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
])

// Provide the theme for echarts (optional, can be 'dark')
provide(THEME_KEY, 'light')
// const authStore = useAuthStore()
const lowStockAlerts = ref([
  { id: 1, name: 'Semen Tiga Roda 50kg', stock: 5 },
  { id: 2, name: 'Besi Beton 10mm', stock: 12 },
])

const chartOption = ref({
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    data: ['Penjualan (Rp)']
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: ['1 Sep', '5 Sep', '10 Sep', '15 Sep', '20 Sep', '25 Sep', '30 Sep']
  },
  yAxis: {
    type: 'value',
    axisLabel: {
      formatter: '{value} Jt'
    }
  },
  series: [
    {
      name: 'Penjualan (Rp)',
      type: 'line',
      stack: 'Total',
      smooth: true,
      areaStyle: {},
      itemStyle: { color: '#4f46e5' }, // Indigo-600
      data: [12, 25, 43, 33, 56, 75, 125]
    }
  ]
})
</script>

<style scoped>
.chart {
  width: 100%;
  height: 100%;
}
</style>
