<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Laporan Pajak (PPN)</h2>
        <p class="text-slate-500">Rekapitulasi PPN keluaran berdasarkan transaksi penjualan.</p>
      </div>
      <button class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 font-medium">
        ⬇ Export CSV
      </button>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <div class="bg-white rounded-lg shadow border border-slate-200 p-6">
        <h3 class="text-sm font-medium text-slate-500">Total Dasar Pengenaan Pajak (DPP)</h3>
        <p class="mt-2 text-3xl font-bold text-slate-900">Rp {{ formatNumber(totalDPP) }}</p>
      </div>
      <div class="bg-white rounded-lg shadow border border-slate-200 p-6">
        <h3 class="text-sm font-medium text-slate-500">Total PPN Terpungut</h3>
        <p class="mt-2 text-3xl font-bold text-indigo-600">Rp {{ formatNumber(totalPPN) }}</p>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-slate-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">No. Faktur</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Tanggal</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Klien</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">NPWP / NIK</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">DPP</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">PPN</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="reports.length === 0">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Belum ada data pajak.</td>
          </tr>
          <tr v-else v-for="row in reports" :key="row.invoice_number">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-900">{{ row.invoice_number }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ row.date }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ row.client_name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-700">{{ row.npwp_nik }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-900 text-right">Rp {{ formatNumber(row.dpp) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-indigo-600 text-right">Rp {{ formatNumber(row.ppn) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import api from '@/plugins/axios'

interface TaxRow {
  invoice_number: string
  date: string
  client_name: string
  npwp_nik: string
  dpp: number
  ppn: number
}

const reports = ref<TaxRow[]>([])
const loading = ref(false)

const formatNumber = (num: number) => num.toLocaleString('id-ID')

const totalDPP = computed(() => reports.value.reduce((sum, r) => sum + r.dpp, 0))
const totalPPN = computed(() => reports.value.reduce((sum, r) => sum + r.ppn, 0))

const fetchReport = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/reports/tax')
    reports.value = res.data || []
  } catch (error) {
    console.error('Failed to fetch tax report', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchReport()
})
</script>
