<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Data Produk</h2>
        <p class="text-slate-500">Kelola master data produk, SKU, dan harga standar.</p>
      </div>
      <button @click="openDialog()" class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 font-medium">
        + Tambah Produk
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-slate-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">SKU</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Nama Produk</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Kategori</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Harga Standar</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Stok Saat Ini</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="products.length === 0">
            <td colspan="6" class="px-6 py-4 text-center text-slate-500">Tidak ada produk ditemukan.</td>
          </tr>
          <tr v-else v-for="product in products" :key="product.ID">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-900">{{ product.sku }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ product.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ product.category }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">Rp {{ formatNumber(product.standard_price) }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
              <span :class="{'text-red-600 font-bold': product.current_stock < 10}">{{ product.current_stock }}</span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button @click="openDialog(product)" class="text-indigo-600 hover:text-indigo-900 mr-3">Edit</button>
              <button @click="deleteProduct(product.ID)" class="text-red-600 hover:text-red-900">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal Form (Simplified) -->
    <div v-if="isDialogOpen" class="fixed inset-0 z-50 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <div class="fixed inset-0 bg-slate-500 bg-opacity-75 transition-opacity" aria-hidden="true" @click="isDialogOpen = false"></div>
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg w-full">
          <form @submit.prevent="saveProduct">
            <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <h3 class="text-lg leading-6 font-medium text-slate-900 mb-4" id="modal-title">
                {{ isEditing ? 'Edit Produk' : 'Tambah Produk Baru' }}
              </h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-slate-700">SKU</label>
                  <input type="text" v-model="form.sku" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">Nama Produk</label>
                  <input type="text" v-model="form.name" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">Kategori</label>
                  <input type="text" v-model="form.category" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">Harga Jual Standar (Rp)</label>
                  <input type="number" v-model="form.standard_price" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
              </div>
            </div>
            <div class="bg-slate-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse">
              <button type="submit" class="w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:ml-3 sm:w-auto sm:text-sm">
                Simpan
              </button>
              <button type="button" @click="isDialogOpen = false" class="mt-3 w-full inline-flex justify-center rounded-md border border-slate-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-slate-700 hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">
                Batal
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

interface Product {
  ID?: number
  sku: string
  name: string
  category: string
  standard_price: number
  average_hpp: number
  current_stock: number
}

const products = ref<Product[]>([])
const loading = ref(false)
const isDialogOpen = ref(false)
const isEditing = ref(false)

const form = ref<Product>({
  sku: '',
  name: '',
  category: '',
  standard_price: 0,
  average_hpp: 0,
  current_stock: 0
})

const formatNumber = (num: number) => {
  return num.toLocaleString('id-ID')
}

const fetchProducts = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/products')
    products.value = res.data
  } catch (error) {
    console.error('Failed to fetch products', error)
  } finally {
    loading.value = false
  }
}

const openDialog = (product?: Product) => {
  if (product) {
    isEditing.value = true
    form.value = { ...product }
  } else {
    isEditing.value = false
    form.value = {
      sku: '',
      name: '',
      category: '',
      standard_price: 0,
      average_hpp: 0,
      current_stock: 0
    }
  }
  isDialogOpen.value = true
}

const saveProduct = async () => {
  try {
    if (isEditing.value && form.value.ID) {
      await api.put(`/api/products/${form.value.ID}`, form.value)
    } else {
      await api.post('/api/products', form.value)
    }
    isDialogOpen.value = false
    fetchProducts()
  } catch (error) {
    console.error('Failed to save product', error)
    alert('Failed to save product. Make sure SKU is unique.')
  }
}

const deleteProduct = async (id?: number) => {
  if (!id) return
  if (!confirm('Apakah anda yakin ingin menghapus produk ini?')) return
  
  try {
    await api.delete(`/api/products/${id}`)
    fetchProducts()
  } catch (error) {
    console.error('Failed to delete product', error)
  }
}

onMounted(() => {
  fetchProducts()
})
</script>
