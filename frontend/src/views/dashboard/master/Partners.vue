<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Data Mitra</h2>
        <p class="text-slate-500">Kelola daftar Klien dan Supplier, beserta NPWP/NIK untuk kebutuhan faktur pajak.</p>
      </div>
      <button @click="openDialog()" class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 font-medium">
        + Tambah Mitra
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-slate-50">
          <tr>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Tipe</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Nama</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">NPWP / NIK</th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Alamat</th>
            <th scope="col" class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="5" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="partners.length === 0">
            <td colspan="5" class="px-6 py-4 text-center text-slate-500">Tidak ada mitra ditemukan.</td>
          </tr>
          <tr v-else v-for="partner in partners" :key="partner.ID">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
              <span :class="partner.type === 'client' ? 'bg-blue-100 text-blue-800' : 'bg-purple-100 text-purple-800'" class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full uppercase">
                {{ partner.type }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-900 font-medium">{{ partner.name }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">
              <div v-if="partner.npwp">NPWP: {{ partner.npwp }}</div>
              <div v-else-if="partner.nik">NIK: {{ partner.nik }}</div>
              <div v-else class="text-red-400 italic">Tidak ada</div>
            </td>
            <td class="px-6 py-4 text-sm text-slate-500 max-w-xs truncate">{{ partner.address }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <button @click="openDialog(partner)" class="text-indigo-600 hover:text-indigo-900 mr-3">Edit</button>
              <button @click="deletePartner(partner.ID)" class="text-red-600 hover:text-red-900">Hapus</button>
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
          <form @submit.prevent="savePartner">
            <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <h3 class="text-lg leading-6 font-medium text-slate-900 mb-4" id="modal-title">
                {{ isEditing ? 'Edit Mitra' : 'Tambah Mitra Baru' }}
              </h3>
              <div class="space-y-4">
                <div>
                  <label class="block text-sm font-medium text-slate-700">Tipe Mitra</label>
                  <select v-model="form.type" required class="mt-1 block w-full pl-3 pr-10 py-2 text-base border border-slate-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md">
                    <option value="client">Client (Pembeli)</option>
                    <option value="supplier">Supplier (Pemasok)</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">Nama</label>
                  <input type="text" v-model="form.name" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">NPWP</label>
                  <input type="text" v-model="form.npwp" placeholder="Opsional" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">NIK <span class="text-slate-400 font-normal">(Wajib jika tidak ada NPWP)</span></label>
                  <input type="text" v-model="form.nik" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700">Alamat Lengkap</label>
                  <textarea v-model="form.address" rows="3" class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"></textarea>
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

interface Partner {
  ID?: number
  type: string
  name: string
  address: string
  npwp: string
  nik: string
}

const partners = ref<Partner[]>([])
const loading = ref(false)
const isDialogOpen = ref(false)
const isEditing = ref(false)

const form = ref<Partner>({
  type: 'client',
  name: '',
  address: '',
  npwp: '',
  nik: ''
})

const fetchPartners = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/partners')
    partners.value = res.data
  } catch (error) {
    console.error('Failed to fetch partners', error)
  } finally {
    loading.value = false
  }
}

const openDialog = (partner?: Partner) => {
  if (partner) {
    isEditing.value = true
    form.value = { ...partner }
  } else {
    isEditing.value = false
    form.value = {
      type: 'client',
      name: '',
      address: '',
      npwp: '',
      nik: ''
    }
  }
  isDialogOpen.value = true
}

const savePartner = async () => {
  // Validate either NPWP or NIK is present
  if (!form.value.npwp && !form.value.nik) {
    alert('Harap isi minimal salah satu dari NPWP atau NIK untuk kebutuhan pajak.')
    return
  }

  try {
    if (isEditing.value && form.value.ID) {
      await api.put(`/api/partners/${form.value.ID}`, form.value)
    } else {
      await api.post('/api/partners', form.value)
    }
    isDialogOpen.value = false
    fetchPartners()
  } catch (error) {
    console.error('Failed to save partner', error)
    alert('Failed to save partner.')
  }
}

const deletePartner = async (id?: number) => {
  if (!id) return
  if (!confirm('Apakah anda yakin ingin menghapus mitra ini?')) return
  
  try {
    await api.delete(`/api/partners/${id}`)
    fetchPartners()
  } catch (error) {
    console.error('Failed to delete partner', error)
  }
}

onMounted(() => {
  fetchPartners()
})
</script>
