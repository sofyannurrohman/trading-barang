<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <div>
        <h2 class="text-xl sm:text-2xl font-bold text-slate-800">Laporan Pajak (PPN)</h2>
        <p class="text-slate-500 text-sm">Rekapitulasi PPN keluaran berdasarkan transaksi penjualan.</p>
      </div>
      <button 
        @click="exportCSV" 
        class="inline-flex items-center justify-center gap-2 px-4 py-2.5 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 font-medium text-sm shadow-sm transition-colors cursor-pointer w-full sm:w-auto"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/></svg>
        Export CSV
      </button>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6">
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5">
        <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Total Dasar Pengenaan Pajak (DPP)</h3>
        <p class="mt-2 text-xl sm:text-2xl font-bold text-slate-900">Rp {{ formatNumber(totalDPP) }}</p>
      </div>
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5">
        <h3 class="text-xs font-semibold text-slate-500 uppercase tracking-wider">Total PPN Terpungut</h3>
        <p class="mt-2 text-xl sm:text-2xl font-bold text-indigo-600">Rp {{ formatNumber(totalPPN) }}</p>
      </div>
    </div>

    <!-- Table Container with Horizontal Scroll for Mobile -->
    <div class="bg-white shadow rounded-xl border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-slate-200">
          <thead class="bg-slate-50">
            <tr>
              <th scope="col" class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">No. Faktur</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Tanggal</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Klien</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">NPWP / NIK</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">DPP</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">PPN</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="6" class="px-6 py-8 text-center text-slate-400 text-sm">Memuat data pajak...</td>
            </tr>
            <tr v-else-if="reports.length === 0">
              <td colspan="6" class="px-6 py-8 text-center text-slate-400 text-sm">Belum ada data pajak penjualan.</td>
            </tr>
            <tr v-else v-for="row in reports" :key="row.invoice_number" class="hover:bg-slate-50/60 transition-colors">
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-900">{{ row.invoice_number }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ row.date }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm text-slate-800 font-medium">{{ row.client_name }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-600">{{ row.npwp_nik || '-' }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-900 text-right">Rp {{ formatNumber(row.dpp) }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-bold text-indigo-600 text-right">Rp {{ formatNumber(row.ppn) }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import api from '@/plugins/axios'
import { toast } from 'vue-sonner'

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

const exportCSV = () => {
  if (reports.value.length === 0) {
    toast.error('Tidak ada data untuk diekspor')
    return
  }
  let csvContent = 'data:text/csv;charset=utf-8,No. Faktur,Tanggal,Klien,NPWP/NIK,DPP,PPN\n'
  reports.value.forEach(r => {
    csvContent += `"${r.invoice_number}","${r.date}","${r.client_name}","${r.npwp_nik}",${r.dpp},${r.ppn}\n`
  })
  const encodedUri = encodeURI(csvContent)
  const link = document.createElement('a')
  link.setAttribute('href', encodedUri)
  link.setAttribute('download', `Laporan_PPN_${new Date().toISOString().slice(0, 10)}.csv`)
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  toast.success('Laporan berhasil diekspor')
}

onMounted(() => {
  fetchReport()
})
</script>
