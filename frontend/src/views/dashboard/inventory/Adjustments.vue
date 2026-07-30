<template>
  <div class="space-y-6 max-w-3xl">
    <div>
      <h2 class="text-2xl font-bold text-slate-800">Penyesuaian Stok</h2>
      <p class="text-slate-500">Catat barang rusak, hilang, atau stock opname manual.</p>
    </div>

    <div class="bg-white shadow rounded-lg border border-slate-200 p-6">
      <form @submit.prevent="submitAdjustment" class="space-y-6">
        
        <div v-if="successMsg" class="p-4 bg-green-50 text-green-700 rounded-md">
          {{ successMsg }}
        </div>
        <div v-if="errorMsg" class="p-4 bg-red-50 text-red-700 rounded-md">
          {{ errorMsg }}
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700">Pilih Produk</label>
          <select v-model="form.product_id" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
            <option value="" disabled>-- Pilih Produk --</option>
            <option v-for="product in products" :key="product.ID" :value="product.ID">
              {{ product.sku }} - {{ product.name }} (Stok Saat Ini: {{ product.current_stock }})
            </option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700">Kuantitas Penyesuaian</label>
          <p class="text-xs text-slate-500 mb-1">Gunakan angka minus (-) untuk barang hilang/rusak. Angka plus (+) untuk penambahan manual.</p>
          <input type="number" v-model="form.quantity" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700">Alasan / Referensi</label>
          <input type="text" v-model="form.reference" required placeholder="Contoh: Barang rusak karena bocor" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
        </div>

        <div class="pt-4 border-t border-slate-200">
          <button type="submit" :disabled="loading" class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 disabled:opacity-50">
            {{ loading ? 'Memproses...' : 'Simpan Penyesuaian' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

interface Product {
  ID: number
  sku: string
  name: string
  current_stock: number
}

const products = ref<Product[]>([])
const loading = ref(false)
const successMsg = ref('')
const errorMsg = ref('')

const form = ref({
  product_id: '' as string | number,
  quantity: 0,
  reference: ''
})

const fetchProducts = async () => {
  try {
    const res = await api.get('/api/products')
    products.value = res.data
  } catch (error) {
    console.error('Failed to fetch products', error)
  }
}

const submitAdjustment = async () => {
  if (form.value.quantity === 0) {
    errorMsg.value = 'Kuantitas tidak boleh 0.'
    return
  }

  loading.value = true
  successMsg.value = ''
  errorMsg.value = ''
  
  try {
    await api.post('/api/inventory/adjustment', {
      product_id: Number(form.value.product_id),
      quantity: form.value.quantity,
      reference: form.value.reference
    })
    
    successMsg.value = 'Penyesuaian stok berhasil disimpan.'
    
    form.value = {
      product_id: '',
      quantity: 0,
      reference: ''
    }
    
    fetchProducts()
    setTimeout(() => { successMsg.value = '' }, 5000)
  } catch (error: any) {
    console.error('Failed adjustment', error)
    errorMsg.value = error.response?.data?.error || 'Gagal menyimpan penyesuaian.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProducts()
})
</script>
