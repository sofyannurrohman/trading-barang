<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <div>
        <h2 class="text-xl sm:text-2xl font-bold text-slate-800">Daftar Faktur Penjualan</h2>
        <p class="text-slate-500 text-sm">Riwayat seluruh transaksi penjualan dan cetak faktur pajak.</p>
      </div>
      <router-link
        to="/dashboard/sales/create"
        class="inline-flex items-center justify-center px-4 py-2.5 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 font-medium text-sm shadow-sm transition-colors cursor-pointer w-full sm:w-auto"
      >
        + Buat Faktur Baru
      </router-link>
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
              <th scope="col" class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Total Akhir</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Status</th>
              <th scope="col" class="px-4 sm:px-6 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Aksi</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="6" class="px-6 py-8 text-center text-slate-400 text-sm">Memuat data faktur...</td>
            </tr>
            <tr v-else-if="invoices.length === 0">
              <td colspan="6" class="px-6 py-8 text-center text-slate-400 text-sm">Belum ada faktur penjualan.</td>
            </tr>
            <tr v-else v-for="invoice in invoices" :key="invoice.ID" class="hover:bg-slate-50/60 transition-colors">
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-semibold text-indigo-600">{{ invoice.invoice_number }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ formatDate(invoice.CreatedAt) }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm text-slate-900 font-medium">{{ invoice.partner?.name }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-bold text-slate-900">Rp {{ formatNumber(invoice.grand_total) }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm">
                <span class="px-2.5 py-0.5 inline-flex text-xs leading-5 font-semibold rounded-full" :class="invoice.status === 'PAID' ? 'bg-emerald-100 text-emerald-800' : 'bg-amber-100 text-amber-800'">
                  {{ invoice.status }}
                </span>
              </td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-3">
                  <a :href="`/print/invoice/${invoice.ID}`" target="_blank" class="text-blue-600 hover:text-blue-900 font-semibold">Cetak A4</a>
                  <button @click="openEditModal(invoice)" class="text-indigo-600 hover:text-indigo-900 font-semibold cursor-pointer">Edit</button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- MODAL: EDIT INVOICE (Mobile Optimized Scrollable) -->
    <div v-if="isEditModalOpen && editingInvoice" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen p-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isEditModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md z-10 p-5 sm:p-6 max-h-[90vh] overflow-y-auto">
          <h3 class="text-lg font-bold text-slate-800 mb-4">Edit Invoice {{ editingInvoice.invoice_number }}</h3>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-2">Status Pembayaran</label>
              <div class="flex gap-4">
                <label class="flex items-center gap-2 cursor-pointer text-sm">
                  <input type="radio" v-model="editForm.status" value="UNPAID" class="text-indigo-600 focus:ring-indigo-500" /> Belum Lunas (UNPAID)
                </label>
                <label class="flex items-center gap-2 cursor-pointer text-sm">
                  <input type="radio" v-model="editForm.status" value="PAID" class="text-indigo-600 focus:ring-indigo-500" /> Lunas (PAID)
                </label>
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Ongkos Kirim (Rp)</label>
              <input type="number" v-model="editForm.shipping_cost" class="w-full border border-slate-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Diskon (Rp)</label>
              <input type="number" v-model="editForm.discount" class="w-full border border-slate-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
          </div>
          
          <div class="mt-6 flex justify-end gap-3 pt-3 border-t border-slate-100">
            <button @click="isEditModalOpen = false" class="px-4 py-2 text-sm border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors cursor-pointer">Batal</button>
            <button @click="submitEdit" :disabled="editLoading" class="px-4 py-2 text-sm bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 font-medium transition-colors disabled:opacity-50 cursor-pointer">
              {{ editLoading ? 'Menyimpan...' : 'Simpan' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'
import { toast } from 'vue-sonner'

interface Invoice {
  ID: number
  CreatedAt: string
  invoice_number: string
  grand_total: number
  status: string
  partner: { name: string }
  shipping_cost?: number
  discount?: number
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
    invoices.value = res.data || []
  } catch (error) {
    console.error('Failed to fetch invoices', error)
  } finally {
    loading.value = false
  }
}

const isEditModalOpen = ref(false)
const editLoading = ref(false)
const editingInvoice = ref<Invoice | null>(null)
const editForm = ref({ status: 'UNPAID', shipping_cost: 0, discount: 0 })

const openEditModal = (invoice: Invoice) => {
  editingInvoice.value = invoice
  editForm.value = {
    status: invoice.status,
    shipping_cost: invoice.shipping_cost ?? 0,
    discount: invoice.discount ?? 0,
  }
  isEditModalOpen.value = true
}

const submitEdit = async () => {
  if (!editingInvoice.value) return
  editLoading.value = true
  const invoiceNumber = editingInvoice.value.invoice_number
  try {
    await api.put(`/api/sales/invoices/${editingInvoice.value.ID}`, {
      status: editForm.value.status,
      shipping_cost: Number(editForm.value.shipping_cost) || 0,
      discount: Number(editForm.value.discount) || 0,
    }, { skipToast: true } as any)
    
    isEditModalOpen.value = false
    toast.success(`Invoice ${invoiceNumber} berhasil diperbarui`)
    await fetchInvoices()
  } catch (error: any) {
    toast.error(error.response?.data?.error ?? 'Gagal memperbarui invoice.')
  } finally {
    editLoading.value = false
  }
}

onMounted(() => {
  fetchInvoices()
})
</script>
