<template>
  <div class="bg-slate-100 min-h-screen p-2 sm:p-6 md:p-8 flex flex-col items-center justify-start print:p-0 print:bg-white">
    
    <div v-if="loading" class="text-center py-10 print:hidden text-slate-500 text-sm">
      Memuat Faktur...
    </div>
    
    <div v-else-if="invoice" class="print-container bg-white text-black font-sans p-4 sm:p-8 rounded-xl shadow-lg border border-slate-200 print:border-0 print:shadow-none print:rounded-none w-full max-w-4xl overflow-x-auto">
      <!-- Header -->
      <div class="flex flex-col sm:flex-row justify-between items-start border-b-2 border-black pb-4 sm:pb-6 mb-6 gap-4">
        <div>
          <h1 class="text-2xl sm:text-3xl font-bold uppercase tracking-wider mb-2">FAKTUR PENJUALAN</h1>
          <p class="text-xs sm:text-sm"><strong>No. Faktur:</strong> {{ invoice.invoice_number }}</p>
          <p class="text-xs sm:text-sm"><strong>Tanggal:</strong> {{ formatDate(invoice.CreatedAt) }}</p>
        </div>
        <div class="text-left sm:text-right">
          <h2 class="text-lg sm:text-xl font-bold">{{ company.name || 'PT TRADING BARANG DEMO' }}</h2>
          <p class="text-xs sm:text-sm whitespace-pre-line text-slate-700">{{ company.address || 'Jl. Sudirman No. 1\nJakarta Pusat' }}</p>
          <p class="text-xs sm:text-sm mt-1"><strong>NPWP:</strong> {{ company.npwp || '12.345.678.9-012.000' }}</p>
        </div>
      </div>

      <!-- Client Info -->
      <div class="mb-6 sm:mb-8">
        <h3 class="font-bold border-b border-black inline-block mb-2 text-xs uppercase tracking-wider">KEPADA YTH:</h3>
        <p class="font-bold text-base sm:text-lg">{{ invoice.partner?.name }}</p>
        <p class="text-xs sm:text-sm whitespace-pre-line text-slate-700">{{ invoice.partner?.address || '-' }}</p>
        <p class="text-xs sm:text-sm mt-1" v-if="invoice.partner?.npwp"><strong>NPWP:</strong> {{ invoice.partner?.npwp }}</p>
        <p class="text-xs sm:text-sm mt-1" v-else-if="invoice.partner?.nik"><strong>NIK:</strong> {{ invoice.partner?.nik }}</p>
      </div>

      <!-- Items Table (Scrollable in Mobile Screen) -->
      <div class="overflow-x-auto mb-6 sm:mb-8">
        <table class="w-full border-collapse border border-black text-xs sm:text-sm min-w-[500px]">
          <thead>
            <tr class="bg-gray-100">
              <th class="border border-black p-2 text-center w-10">No</th>
              <th class="border border-black p-2 text-left">Deskripsi Barang</th>
              <th class="border border-black p-2 text-center w-20">Qty</th>
              <th class="border border-black p-2 text-right w-32">Harga Satuan</th>
              <th class="border border-black p-2 text-right w-32">Jumlah</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in invoice.items" :key="item.ID">
              <td class="border border-black p-2 text-center">{{ Number(index) + 1 }}</td>
              <td class="border border-black p-2 font-medium">{{ item.product?.name }}</td>
              <td class="border border-black p-2 text-center">{{ item.quantity }}</td>
              <td class="border border-black p-2 text-right">Rp {{ formatNumber(item.unit_price) }}</td>
              <td class="border border-black p-2 text-right font-semibold">Rp {{ formatNumber(item.subtotal) }}</td>
            </tr>
          </tbody>
          <tfoot>
            <tr>
              <td colspan="3" class="border-t border-black p-2 text-right border-l-0 border-b-0"></td>
              <td class="border border-black p-2 font-bold text-right">DPP</td>
              <td class="border border-black p-2 font-bold text-right">Rp {{ formatNumber(invoice.total_dpp) }}</td>
            </tr>
            <tr>
              <td colspan="3" class="p-2 text-right border-0"></td>
              <td class="border border-black p-2 font-bold text-right">PPN</td>
              <td class="border border-black p-2 font-bold text-right">Rp {{ formatNumber(invoice.total_ppn) }}</td>
            </tr>
            <tr>
              <td colspan="3" class="p-2 text-right border-0"></td>
              <td class="border border-black p-2 font-bold text-right text-base sm:text-lg">TOTAL</td>
              <td class="border border-black p-2 font-bold text-right text-base sm:text-lg">Rp {{ formatNumber(invoice.grand_total) }}</td>
            </tr>
          </tfoot>
        </table>
      </div>

      <!-- Signatures -->
      <div class="flex justify-between mt-8 sm:mt-16 pt-4 sm:pt-8 text-xs sm:text-sm">
        <div class="text-center w-36 sm:w-48">
          <p class="mb-14 sm:mb-20">Penerima,</p>
          <hr class="border-black"/>
          <p class="text-[10px] sm:text-xs mt-1 text-slate-600">(Nama Terang &amp; Cap)</p>
        </div>
        <div class="text-center w-36 sm:w-48">
          <p class="mb-14 sm:mb-20">Hormat Kami,</p>
          <hr class="border-black"/>
          <p class="text-[10px] sm:text-xs mt-1 font-semibold">{{ company.name || 'PT TRADING BARANG DEMO' }}</p>
        </div>
      </div>
      
    </div>
    
    <!-- Action Floating Buttons -->
    <div v-if="invoice" class="fixed bottom-6 right-6 flex items-center gap-3 print:hidden z-20">
      <button 
        @click="doPrint" 
        class="bg-slate-900 hover:bg-slate-800 text-white px-5 py-3 rounded-full shadow-xl font-bold text-sm flex items-center gap-2 cursor-pointer transition-transform hover:scale-105"
      >
        <span>🖨️ Cetak / PDF</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/plugins/axios'

const route = useRoute()
const invoice = ref<any>(null)
const loading = ref(true)
const company = ref<any>({})

const formatNumber = (num: number) => num.toLocaleString('id-ID')
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('id-ID', {
    year: 'numeric', month: 'long', day: 'numeric'
  })
}

const doPrint = () => {
  window.print()
}

onMounted(async () => {
  try {
    const [invRes, compRes] = await Promise.all([
      api.get(`/api/sales/invoices/${route.params.id}`),
      api.get('/api/settings/company')
    ])
    invoice.value = invRes.data
    company.value = compRes.data || {}
  } catch (error) {
    console.error('Failed to load invoice', error)
  } finally {
    loading.value = false
  }
})
</script>

<style>
@media print {
  body {
    background: white;
  }
  .print\:hidden {
    display: none !important;
  }
}
</style>
