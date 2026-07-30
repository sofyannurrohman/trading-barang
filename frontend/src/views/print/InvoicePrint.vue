<template>
  <div class="print-container bg-white min-h-screen text-black font-sans p-8 max-w-4xl mx-auto">
    
    <div v-if="loading" class="text-center py-10 print:hidden">
      Memuat Faktur...
    </div>
    
    <div v-else-if="invoice">
      <!-- Header -->
      <div class="flex justify-between items-start border-b-2 border-black pb-6 mb-6">
        <div>
          <h1 class="text-3xl font-bold uppercase tracking-wider mb-2">FAKTUR PENJUALAN</h1>
          <p class="text-sm"><strong>No. Faktur:</strong> {{ invoice.invoice_number }}</p>
          <p class="text-sm"><strong>Tanggal:</strong> {{ formatDate(invoice.CreatedAt) }}</p>
        </div>
        <div class="text-right">
          <h2 class="text-xl font-bold">{{ company.Name || 'PT TRADING BARANG' }}</h2>
          <p class="text-sm whitespace-pre-line">{{ company.Address || 'Jl. Contoh Alamat No. 123\nJakarta Pusat, 10110' }}</p>
          <p class="text-sm mt-1"><strong>NPWP:</strong> {{ company.NPWP || '01.234.567.8-901.000' }}</p>
        </div>
      </div>

      <!-- Client Info -->
      <div class="mb-8">
        <h3 class="font-bold border-b border-black inline-block mb-2">KEPADA YTH:</h3>
        <p class="font-bold text-lg">{{ invoice.partner?.name }}</p>
        <p class="text-sm whitespace-pre-line">{{ invoice.partner?.address }}</p>
        <p class="text-sm mt-1" v-if="invoice.partner?.npwp"><strong>NPWP:</strong> {{ invoice.partner?.npwp }}</p>
        <p class="text-sm mt-1" v-else-if="invoice.partner?.nik"><strong>NIK:</strong> {{ invoice.partner?.nik }}</p>
      </div>

      <!-- Items Table -->
      <table class="w-full mb-8 border-collapse border border-black text-sm">
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
            <td class="border border-black p-2">{{ item.product?.name }}</td>
            <td class="border border-black p-2 text-center">{{ item.quantity }}</td>
            <td class="border border-black p-2 text-right">Rp {{ formatNumber(item.unit_price) }}</td>
            <td class="border border-black p-2 text-right">Rp {{ formatNumber(item.subtotal) }}</td>
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
            <td class="border border-black p-2 font-bold text-right text-lg">TOTAL</td>
            <td class="border border-black p-2 font-bold text-right text-lg">Rp {{ formatNumber(invoice.grand_total) }}</td>
          </tr>
        </tfoot>
      </table>

      <!-- Signatures -->
      <div class="flex justify-between mt-16 pt-8">
        <div class="text-center w-48">
          <p class="mb-20">Penerima,</p>
          <hr class="border-black"/>
          <p class="text-xs mt-1">(Nama Terang & Cap)</p>
        </div>
        <div class="text-center w-48">
          <p class="mb-20">Hormat Kami,</p>
          <hr class="border-black"/>
          <p class="text-xs mt-1">{{ company.Name || 'PT TRADING BARANG' }}</p>
        </div>
      </div>
      
    </div>
    
    <!-- Print Button (Hidden in Print Mode) -->
    <button v-if="invoice" @click="doPrint" class="fixed bottom-8 right-8 bg-black text-white px-6 py-3 rounded-full shadow-lg font-bold print:hidden hover:bg-gray-800 transition-colors">
      🖨️ Cetak PDF
    </button>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/plugins/axios'

const route = useRoute()
const invoice = ref<any>(null)
const loading = ref(true)

// Mock company settings, in a real app fetch this from /api/settings/company
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
    const res = await api.get(`/api/sales/invoices/${route.params.id}`)
    invoice.value = res.data
    // Optional: automatically trigger print dialog after 500ms
    setTimeout(() => { window.print() }, 500)
  } catch (error) {
    console.error('Failed to load invoice', error)
    alert('Faktur tidak ditemukan.')
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
