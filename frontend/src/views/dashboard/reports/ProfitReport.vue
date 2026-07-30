<template>
  <div class="space-y-6">
    <div>
      <h2 class="text-2xl font-bold text-slate-800">Laporan Laba / Rugi</h2>
      <p class="text-slate-500">Ringkasan estimasi laba kotor berdasarkan HPP pada saat barang terjual.</p>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div class="bg-white rounded-lg shadow border border-slate-200 p-6">
        <h3 class="text-sm font-medium text-slate-500">Total Pendapatan (DPP)</h3>
        <p class="mt-2 text-3xl font-bold text-slate-900">Rp {{ formatNumber(totalRevenue) }}</p>
      </div>
      <div class="bg-white rounded-lg shadow border border-slate-200 p-6">
        <h3 class="text-sm font-medium text-slate-500">Total Modal (HPP Terjual)</h3>
        <p class="mt-2 text-3xl font-bold text-slate-900">Rp {{ formatNumber(totalCost) }}</p>
      </div>
      <div class="bg-white rounded-lg shadow border border-slate-200 p-6">
        <h3 class="text-sm font-medium text-slate-500">Total Laba Kotor</h3>
        <p class="mt-2 text-3xl font-bold" :class="totalProfit >= 0 ? 'text-green-600' : 'text-red-600'">
          Rp {{ formatNumber(totalProfit) }}
        </p>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-slate-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">No. Faktur</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Tanggal</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Pendapatan</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Modal (HPP)</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Laba Kotor</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Margin</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="reports.length === 0">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Belum ada data penjualan.</td>
          </tr>
          <tr v-else v-for="row in reports" :key="row.invoice_number">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-indigo-600">{{ row.invoice_number }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ row.date }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-900 text-right">Rp {{ formatNumber(row.revenue) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500 text-right">Rp {{ formatNumber(row.total_cost) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-right" :class="row.gross_profit >= 0 ? 'text-green-600' : 'text-red-600'">
              Rp {{ formatNumber(row.gross_profit) }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-right font-medium text-slate-500">
              {{ row.margin_percent.toFixed(2) }}%
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import api from '@/plugins/axios'

interface ProfitRow {
  invoice_number: string
  date: string
  revenue: number
  total_cost: number
  gross_profit: number
  margin_percent: number
}

const reports = ref<ProfitRow[]>([])
const loading = ref(false)

const formatNumber = (num: number) => num.toLocaleString('id-ID')

const totalRevenue = computed(() => reports.value.reduce((sum, r) => sum + r.revenue, 0))
const totalCost = computed(() => reports.value.reduce((sum, r) => sum + r.total_cost, 0))
const totalProfit = computed(() => reports.value.reduce((sum, r) => sum + r.gross_profit, 0))

const fetchReport = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/reports/profit')
    reports.value = res.data || []
  } catch (error) {
    console.error('Failed to fetch profit report', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchReport()
})
</script>
