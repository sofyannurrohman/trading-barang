<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Kartu Stok</h2>
        <p class="text-slate-500">Lihat riwayat pergerakan (masuk/keluar) untuk setiap produk.</p>
      </div>
    </div>

    <!-- Product Selector -->
    <div class="bg-white shadow rounded-lg border border-slate-200 p-4">
      <div class="max-w-xs">
        <label class="block text-sm font-medium text-slate-700">Pilih Produk</label>
        <select v-model="selectedProductId" @change="fetchStockCard" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
          <option value="" disabled>-- Pilih Produk --</option>
          <option v-for="product in products" :key="product.ID" :value="product.ID">
            {{ product.sku }} - {{ product.name }}
          </option>
        </select>
      </div>
    </div>

    <!-- Table -->
    <div v-if="selectedProductId" class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <div class="px-6 py-4 border-b border-slate-200 flex justify-between items-center bg-slate-50">
        <h3 class="text-lg font-medium text-slate-800">Riwayat Transaksi</h3>
        <span v-if="currentProduct" class="text-sm font-medium px-3 py-1 bg-indigo-100 text-indigo-800 rounded-full">
          Stok Saat Ini: {{ currentProduct.current_stock }}
        </span>
      </div>
      
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-white">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Tanggal</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Tipe Transaksi</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Kuantitas</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Harga/HPP</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Referensi</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="5" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="transactions.length === 0">
            <td colspan="5" class="px-6 py-4 text-center text-slate-500">Belum ada transaksi untuk produk ini.</td>
          </tr>
          <tr v-else v-for="tx in transactions" :key="tx.ID">
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ formatDate(tx.CreatedAt) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <span v-if="tx.type === 'INBOUND'" class="text-green-600">BARANG MASUK</span>
              <span v-else-if="tx.type === 'ADJUSTMENT' && tx.quantity < 0" class="text-red-600">PENYESUAIAN (KELUAR)</span>
              <span v-else-if="tx.type === 'ADJUSTMENT' && tx.quantity > 0" class="text-blue-600">PENYESUAIAN (MASUK)</span>
              <span v-else-if="tx.type === 'SALE'" class="text-orange-600">PENJUALAN</span>
              <span v-else>{{ tx.type }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm font-bold" :class="tx.quantity > 0 ? 'text-green-600' : 'text-red-600'">
              {{ tx.quantity > 0 ? '+' : '' }}{{ tx.quantity }}
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">Rp {{ formatNumber(tx.unit_price) }}</td>
            <td class="px-6 py-4 text-sm text-slate-500 max-w-xs truncate" :title="tx.reference">{{ tx.reference || '-' }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import api from '@/plugins/axios'

interface Product {
  ID: number
  sku: string
  name: string
  current_stock: number
}

interface Transaction {
  ID: number
  CreatedAt: string
  type: string
  quantity: number
  unit_price: number
  reference: string
}

const products = ref<Product[]>([])
const selectedProductId = ref<string | number>('')
const transactions = ref<Transaction[]>([])
const loading = ref(false)

const formatNumber = (num: number) => num.toLocaleString('id-ID')
const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString('id-ID', {
    year: 'numeric', month: 'short', day: 'numeric',
    hour: '2-digit', minute: '2-digit'
  })
}

const currentProduct = computed(() => {
  return products.value.find(p => p.ID === Number(selectedProductId.value))
})

const fetchProducts = async () => {
  try {
    const res = await api.get('/api/products')
    products.value = res.data
  } catch (error) {
    console.error('Failed to fetch products', error)
  }
}

const fetchStockCard = async () => {
  if (!selectedProductId.value) return
  
  loading.value = true
  try {
    const res = await api.get(`/api/inventory/stock-card/${selectedProductId.value}`)
    transactions.value = res.data
  } catch (error) {
    console.error('Failed to fetch stock card', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProducts()
})
</script>
