<template>
  <div class="space-y-6">
    <div>
      <h2 class="text-2xl font-bold text-slate-800">Buat Faktur Baru</h2>
      <p class="text-slate-500">Pilih klien dan masukkan barang yang dibeli.</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      
      <!-- Cart Form -->
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-white shadow rounded-lg border border-slate-200 p-6">
          <h3 class="text-lg font-medium text-slate-800 mb-4">Informasi Penjualan</h3>
          
          <div class="mb-4">
            <label class="block text-sm font-medium text-slate-700">Klien (Pembeli)</label>
            <select v-model="form.partner_id" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
              <option value="" disabled>-- Pilih Klien --</option>
              <option v-for="client in clients" :key="client.ID" :value="client.ID">
                {{ client.name }} ({{ client.npwp || client.nik }})
              </option>
            </select>
          </div>

          <div class="border-t border-slate-200 pt-4">
            <h4 class="text-md font-medium text-slate-700 mb-3">Tambah Barang</h4>
            <div class="grid grid-cols-1 sm:grid-cols-12 gap-4 items-end">
              <div class="sm:col-span-5">
                <label class="block text-xs font-medium text-slate-500">Produk</label>
                <select v-model="itemForm.product_id" @change="onProductSelect" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 sm:text-sm">
                  <option value="" disabled>-- Pilih Produk --</option>
                  <option v-for="p in products" :key="p.ID" :value="p.ID">
                    {{ p.name }} (Stok: {{ p.current_stock }})
                  </option>
                </select>
              </div>
              <div class="sm:col-span-2">
                <label class="block text-xs font-medium text-slate-500">Qty</label>
                <input type="number" v-model="itemForm.quantity" min="1" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 sm:text-sm" />
              </div>
              <div class="sm:col-span-3">
                <label class="block text-xs font-medium text-slate-500">Harga (Rp)</label>
                <input type="number" v-model="itemForm.unit_price" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 sm:text-sm" />
              </div>
              <div class="sm:col-span-2">
                <button @click="addItem" type="button" class="w-full bg-indigo-100 text-indigo-700 hover:bg-indigo-200 py-2 rounded-md text-sm font-medium">Tambah</button>
              </div>
            </div>
          </div>
        </div>

        <!-- Items Table -->
        <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
          <table class="min-w-full divide-y divide-slate-200">
            <thead class="bg-slate-50">
              <tr>
                <th class="px-4 py-3 text-left text-xs font-medium text-slate-500 uppercase">Produk</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-slate-500 uppercase">Harga</th>
                <th class="px-4 py-3 text-center text-xs font-medium text-slate-500 uppercase">Qty</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-slate-500 uppercase">Subtotal</th>
                <th class="px-4 py-3 text-right text-xs font-medium text-slate-500 uppercase">Aksi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-200">
              <tr v-if="cart.length === 0">
                <td colspan="5" class="px-4 py-6 text-center text-slate-500 text-sm">Belum ada barang yang ditambahkan.</td>
              </tr>
              <tr v-for="(item, index) in cart" :key="index">
                <td class="px-4 py-3 text-sm text-slate-900">{{ item.product_name }}</td>
                <td class="px-4 py-3 text-sm text-slate-500 text-right">Rp {{ formatNumber(item.unit_price) }}</td>
                <td class="px-4 py-3 text-sm text-slate-900 text-center">{{ item.quantity }}</td>
                <td class="px-4 py-3 text-sm font-medium text-slate-900 text-right">Rp {{ formatNumber(item.quantity * item.unit_price) }}</td>
                <td class="px-4 py-3 text-right text-sm font-medium">
                  <button @click="removeItem(index)" class="text-red-600 hover:text-red-900">Hapus</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Summary Panel -->
      <div class="lg:col-span-1">
        <div class="bg-slate-900 rounded-lg shadow border border-slate-800 text-white p-6 sticky top-24">
          <h3 class="text-lg font-medium mb-4">Ringkasan Faktur</h3>
          
          <div class="space-y-3 text-sm text-slate-300">
            <div class="flex justify-between items-center mb-2 pb-2 border-b border-slate-700">
              <span class="text-sm">Kenakan PPN</span>
              <input type="checkbox" v-model="form.is_taxable" class="h-4 w-4 rounded border-slate-300 text-indigo-600 focus:ring-indigo-500 bg-slate-800" />
            </div>
            <div class="flex justify-between">
              <span>Dasar Pengenaan Pajak (DPP)</span>
              <span class="text-white font-medium">Rp {{ formatNumber(totalDPP) }}</span>
            </div>
            <div class="flex justify-between">
              <span>PPN (11%)</span>
              <span class="text-white font-medium">Rp {{ formatNumber(totalPPN) }}</span>
            </div>
            <div class="border-t border-slate-700 pt-3 mt-3 flex justify-between text-lg font-bold text-white">
              <span>Total Akhir</span>
              <span>Rp {{ formatNumber(grandTotal) }}</span>
            </div>
          </div>

          <div v-if="errorMsg" class="mt-4 p-3 bg-red-900/50 text-red-200 text-sm rounded border border-red-800">
            {{ errorMsg }}
          </div>

          <button @click="submitInvoice" :disabled="loading || cart.length === 0 || !form.partner_id" class="mt-6 w-full py-3 px-4 rounded-md shadow text-sm font-bold bg-indigo-500 hover:bg-indigo-400 text-white disabled:opacity-50 transition-colors">
            {{ loading ? 'Menyimpan...' : 'Buat Faktur Penjualan' }}
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '@/plugins/axios'
import { toast } from 'vue-sonner'

const router = useRouter()

interface Product { ID: number; name: string; standard_price: number; current_stock: number }
interface Partner { ID: number; name: string; npwp: string; nik: string }
interface CartItem { product_id: number; product_name: string; quantity: number; unit_price: number }

const products = ref<Product[]>([])
const clients = ref<Partner[]>([])
const cart = ref<CartItem[]>([])

const loading = ref(false)
const errorMsg = ref('')

const form = ref({
  partner_id: '',
  is_taxable: true
})

const itemForm = ref({
  product_id: '',
  quantity: 1,
  unit_price: 0
})

const fetchDependencies = async () => {
  try {
    const [prodRes, partRes] = await Promise.all([
      api.get('/api/products'),
      api.get('/api/partners?type=client')
    ])
    products.value = prodRes.data
    clients.value = partRes.data
  } catch (error) {
    console.error('Failed to fetch data', error)
  }
}

const onProductSelect = () => {
  const p = products.value.find(x => x.ID === Number(itemForm.value.product_id))
  if (p) {
    itemForm.value.unit_price = p.standard_price
    itemForm.value.quantity = 1
  }
}

const addItem = () => {
  if (!itemForm.value.product_id || itemForm.value.quantity < 1 || itemForm.value.unit_price < 0) return
  
  const p = products.value.find(x => x.ID === Number(itemForm.value.product_id))
  if (p) {
    cart.value.push({
      product_id: p.ID,
      product_name: p.name,
      quantity: itemForm.value.quantity,
      unit_price: itemForm.value.unit_price
    })
    // Reset item form
    itemForm.value = { product_id: '', quantity: 1, unit_price: 0 }
  }
}

const removeItem = (index: number) => {
  cart.value.splice(index, 1)
}

const formatNumber = (num: number) => num.toLocaleString('id-ID')

const totalDPP = computed(() => {
  return cart.value.reduce((sum, item) => sum + (item.quantity * item.unit_price), 0)
})
const totalPPN = computed(() => form.value.is_taxable ? totalDPP.value * 0.11 : 0) // Using 11% fixed for UI simplicity right now, backend recalculates
const grandTotal = computed(() => totalDPP.value + totalPPN.value)

const submitInvoice = async () => {
  loading.value = true
  errorMsg.value = ''
  try {
    await api.post('/api/sales/invoice', {
      partner_id: Number(form.value.partner_id),
      is_taxable: form.value.is_taxable,
      items: cart.value.map(i => ({
        product_id: i.product_id,
        quantity: i.quantity,
        unit_price: i.unit_price
      }))
    }, { skipToast: true } as any)
    
    toast.success('Faktur penjualan berhasil dibuat')
    // Redirect to list
    router.push('/dashboard/sales/list')
  } catch (error: any) {
    console.error(error)
    const errMsg = error.response?.data?.error || 'Gagal membuat faktur.'
    errorMsg.value = errMsg
    toast.error(errMsg)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchDependencies()
})
</script>
