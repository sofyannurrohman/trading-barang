<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Daftar Faktur Penjualan</h2>
        <p class="text-slate-500">Riwayat seluruh transaksi penjualan dan cetak faktur pajak.</p>
      </div>
      <router-link to="/dashboard/sales/create" class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 font-medium">
        + Buat Faktur Baru
      </router-link>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-slate-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">No. Faktur</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Tanggal</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Klien</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Total Akhir</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Status</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="invoices.length === 0">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Belum ada faktur.</td>
          </tr>
          <tr v-else v-for="invoice in invoices" :key="invoice.ID">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-indigo-600">{{ invoice.invoice_number }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ formatDate(invoice.CreatedAt) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-900">{{ invoice.partner?.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-slate-900">Rp {{ formatNumber(invoice.grand_total) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full" :class="invoice.status === 'PAID' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'">
                {{ invoice.status }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <a :href="`/print/invoice/${invoice.ID}`" target="_blank" class="text-blue-600 hover:text-blue-900">Cetak A4</a>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

interface Invoice {
  ID: number
  CreatedAt: string
  invoice_number: string
  grand_total: number
  status: string
  partner: { name: string }
}

const invoices = ref<Invoice[]>([])
const loading = ref(false)

const formatNumber = (num: number) => num.toLocaleString('id-ID')
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('id-ID', {
    year: 'numeric', month: 'long', day: 'numeric'
  })
}

const fetchInvoices = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/sales/invoices')
    invoices.value = res.data
  } catch (error) {
    console.error('Failed to fetch invoices', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchInvoices()
})
</script>
