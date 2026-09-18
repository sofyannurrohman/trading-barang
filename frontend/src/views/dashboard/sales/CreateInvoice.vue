<template>
  <div class="space-y-6">
    <div>
      <h2 class="text-xl sm:text-2xl font-bold text-slate-800">Buat Faktur Baru</h2>
      <p class="text-slate-500 text-sm">Pilih klien dan masukkan barang yang dibeli.</p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      
      <!-- Cart Form -->
      <div class="lg:col-span-2 space-y-6">
        <div class="bg-white shadow rounded-xl border border-slate-200 p-4 sm:p-6">
          <h3 class="text-base sm:text-lg font-semibold text-slate-800 mb-4">Informasi Penjualan</h3>
          
          <div class="mb-5">
            <label class="block text-sm font-medium text-slate-700 mb-1">Klien (Pembeli) <span class="text-red-500">*</span></label>
            <select v-model="form.partner_id" class="block w-full border border-slate-300 rounded-lg py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm bg-white">
              <option value="" disabled>-- Pilih Klien --</option>
              <option v-for="client in clients" :key="client.ID" :value="client.ID">
                {{ client.name }} ({{ client.npwp || client.nik || 'Tanpa NPWP/NIK' }})
              </option>
            </select>
          </div>

          <div class="border-t border-slate-200 pt-4">
            <h4 class="text-sm font-semibold text-slate-700 mb-3">Tambah Item Barang</h4>
            <div class="space-y-3 sm:space-y-0 sm:grid sm:grid-cols-12 sm:gap-3 sm:items-end">
              <div class="sm:col-span-5">
                <label class="block text-xs font-medium text-slate-500 mb-1">Produk</label>
                <select v-model="itemForm.product_id" @change="onProductSelect" class="block w-full border border-slate-300 rounded-lg py-2 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm bg-white">
                  <option value="" disabled>-- Pilih Produk --</option>
                  <option v-for="p in products" :key="p.ID" :value="p.ID">
                    {{ p.name }} (Stok: {{ p.current_stock }})
                  </option>
                </select>
              </div>
              <div class="grid grid-cols-2 gap-3 sm:contents">
                <div class="sm:col-span-2">
                  <label class="block text-xs font-medium text-slate-500 mb-1">Qty</label>
                  <input type="number" v-model="itemForm.quantity" min="1" class="block w-full border border-slate-300 rounded-lg py-2 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm" />
                </div>
                <div class="sm:col-span-3">
                  <label class="block text-xs font-medium text-slate-500 mb-1">Harga (Rp)</label>
                  <input type="number" v-model="itemForm.unit_price" class="block w-full border border-slate-300 rounded-lg py-2 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm" />
                </div>
              </div>
              <div class="sm:col-span-2 pt-1 sm:pt-0">
                <button @click="addItem" type="button" class="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-2 px-3 rounded-lg text-sm font-semibold shadow-sm transition-colors cursor-pointer">
                  + Tambah
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Items Table -->
        <div class="bg-white shadow rounded-xl border border-slate-200 overflow-hidden">
          <div class="p-4 border-b border-slate-100 flex items-center justify-between bg-slate-50">
            <h4 class="text-sm font-bold text-slate-700">Daftar Barang yang Dipesan</h4>
            <span class="text-xs text-slate-500 bg-white px-2.5 py-1 rounded-full border border-slate-200 font-medium">
              {{ cart.length }} item
            </span>
          </div>

          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-slate-200">
              <thead class="bg-slate-50">
                <tr>
                  <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Produk</th>
                  <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Harga</th>
                  <th class="px-4 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Qty</th>
                  <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Subtotal</th>
                  <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Aksi</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-100 bg-white">
                <tr v-if="cart.length === 0">
                  <td colspan="5" class="px-4 py-8 text-center text-slate-400 text-sm">
                    Belum ada barang di keranjang faktur.
                  </td>
                </tr>
                <tr v-for="(item, index) in cart" :key="index" class="hover:bg-slate-50/50">
                  <td class="px-4 py-3 text-sm text-slate-900 font-medium whitespace-nowrap">{{ item.product_name }}</td>
                  <td class="px-4 py-3 text-sm text-slate-500 text-right whitespace-nowrap">Rp {{ formatNumber(item.unit_price) }}</td>
                  <td class="px-4 py-3 text-sm text-slate-900 text-center whitespace-nowrap font-semibold">{{ item.quantity }}</td>
                  <td class="px-4 py-3 text-sm font-bold text-slate-900 text-right whitespace-nowrap">Rp {{ formatNumber(item.quantity * item.unit_price) }}</td>
                  <td class="px-4 py-3 text-right text-sm font-medium whitespace-nowrap">
                    <button @click="removeItem(index)" class="text-red-600 hover:text-red-900 font-medium cursor-pointer">Hapus</button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Summary Panel -->
      <div class="lg:col-span-1">
        <div class="bg-slate-900 rounded-xl shadow-lg border border-slate-800 text-white p-5 sm:p-6 sticky top-20">
          <h3 class="text-base sm:text-lg font-bold mb-4 flex items-center justify-between">
            <span>Ringkasan Faktur</span>
            <span class="text-xs px-2 py-0.5 rounded bg-slate-800 text-slate-300 font-normal">Draft</span>
          </h3>
          
          <div class="space-y-3 text-sm text-slate-300">
            <div class="flex justify-between items-center pb-3 border-b border-slate-800">
              <span class="text-sm font-medium">Kenakan PPN ({{ ppnRate }}%)</span>
              <input type="checkbox" v-model="form.is_taxable" class="h-4 w-4 rounded border-slate-700 text-indigo-500 focus:ring-indigo-500 bg-slate-800 cursor-pointer" />
            </div>
            <div class="flex justify-between">
              <span>DPP (Dasar Pengenaan Pajak)</span>
              <span class="text-white font-medium">Rp {{ formatNumber(totalDPP) }}</span>
            </div>
            <div class="flex justify-between">
              <span>Estimasi PPN</span>
              <span class="text-white font-medium">Rp {{ formatNumber(totalPPN) }}</span>
            </div>
            <div class="border-t border-slate-800 pt-3 mt-3 flex justify-between text-base sm:text-lg font-bold text-white">
              <span>Total Akhir</span>
              <span class="text-rose-400">Rp {{ formatNumber(grandTotal) }}</span>
            </div>
          </div>

          <div v-if="errorMsg" class="mt-4 p-3 bg-red-900/50 text-red-200 text-xs sm:text-sm rounded-lg border border-red-800">
            {{ errorMsg }}
          </div>

          <button
            @click="submitInvoice"
            :disabled="loading || cart.length === 0 || !form.partner_id"
            class="mt-6 w-full py-3 px-4 rounded-xl shadow text-sm font-bold bg-gradient-to-r from-rose-600 to-red-700 hover:from-rose-500 hover:to-red-600 text-white disabled:opacity-50 transition-all cursor-pointer"
          >
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
const ppnRate = ref(11)

const form = ref({
  partner_id: '',
  is_taxable: true
})

const itemForm = ref({
  product_id: '',
  quantity: 1,
  unit_price: 0
})

const formatNumber = (val: number) => {
  if (!val) return '0'
  return new Intl.NumberFormat('id-ID').format(val)
}

const fetchData = async () => {
  try {
    const [pRes, cRes, compRes] = await Promise.all([
      api.get('/api/master/products'),
      api.get('/api/master/partners?type=client'),
      api.get('/api/settings/company')
    ])
    products.value = pRes.data || []
    clients.value = cRes.data || []
    if (compRes.data && compRes.data.ppn_rate !== undefined) {
      ppnRate.value = compRes.data.ppn_rate
    }
  } catch (error) {
    console.error('Failed to load form data', error)
  }
}

const onProductSelect = () => {
  const selected = products.value.find(p => p.ID === Number(itemForm.value.product_id))
  if (selected) {
    itemForm.value.unit_price = selected.standard_price
  }
}

const addItem = () => {
  if (!itemForm.value.product_id) {
    toast.error('Pilih produk terlebih dahulu')
    return
  }
  if (itemForm.value.quantity <= 0) {
    toast.error('Jumlah kuantiti minimal 1')
    return
  }

  const selected = products.value.find(p => p.ID === Number(itemForm.value.product_id))
  if (!selected) return

  if (itemForm.value.quantity > selected.current_stock) {
    toast.error(`Stok tidak mencukupi. Sisa stok: ${selected.current_stock}`)
    return
  }

  const existingIndex = cart.value.findIndex(item => item.product_id === selected.ID)
  if (existingIndex > -1) {
    const newQty = cart.value[existingIndex].quantity + Number(itemForm.value.quantity)
    if (newQty > selected.current_stock) {
      toast.error(`Total kuantiti di keranjang (${newQty}) melebihi stok (${selected.current_stock})`)
      return
    }
    cart.value[existingIndex].quantity = newQty
    cart.value[existingIndex].unit_price = Number(itemForm.value.unit_price)
  } else {
    cart.value.push({
      product_id: selected.ID,
      product_name: selected.name,
      quantity: Number(itemForm.value.quantity),
      unit_price: Number(itemForm.value.unit_price)
    })
  }

  itemForm.value = {
    product_id: '',
    quantity: 1,
    unit_price: 0
  }
}

const removeItem = (index: number) => {
  cart.value.splice(index, 1)
}

const totalDPP = computed(() => {
  return cart.value.reduce((sum, item) => sum + (item.quantity * item.unit_price), 0)
})

const totalPPN = computed(() => {
  if (!form.value.is_taxable) return 0
  return totalDPP.value * (ppnRate.value / 100)
})

const grandTotal = computed(() => {
  return totalDPP.value + totalPPN.value
})

const submitInvoice = async () => {
  if (!form.value.partner_id) {
    errorMsg.value = 'Silakan pilih klien'
    return
  }
  if (cart.value.length === 0) {
    errorMsg.value = 'Keranjang belanja masih kosong'
    return
  }

  loading.value = true
  errorMsg.value = ''

  try {
    const payload = {
      partner_id: Number(form.value.partner_id),
      is_taxable: form.value.is_taxable,
      items: cart.value.map(item => ({
        product_id: item.product_id,
        quantity: item.quantity,
        unit_price: item.unit_price
      }))
    }

    const res = await api.post('/api/sales/invoices', payload, { skipToast: true } as any)
    toast.success('Faktur berhasil dibuat')
    if (res.data?.id || res.data?.ID) {
      router.push('/dashboard/sales/list')
    } else {
      router.push('/dashboard/sales/list')
    }
  } catch (err: any) {
    errorMsg.value = err.response?.data?.error || 'Gagal membuat faktur'
    toast.error(errorMsg.value)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>
