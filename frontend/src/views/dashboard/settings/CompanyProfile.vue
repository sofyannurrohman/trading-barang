<template>
  <div class="space-y-6 max-w-2xl">
    <div>
      <h2 class="text-2xl font-bold text-slate-800">Profil Perusahaan</h2>
      <p class="text-slate-500">Kelola identitas perusahaan yang akan tampil pada kop faktur cetak.</p>
    </div>

    <div class="bg-white shadow rounded-lg border border-slate-200 p-6">
      <form @submit.prevent="saveProfile" class="space-y-6">
        
        <div v-if="successMsg" class="p-4 bg-green-50 text-green-700 rounded-md">
          {{ successMsg }}
        </div>
        <div v-if="errorMsg" class="p-4 bg-red-50 text-red-700 rounded-md">
          {{ errorMsg }}
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700">Nama Perusahaan</label>
          <input type="text" v-model="form.name" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700">NPWP Perusahaan</label>
          <input type="text" v-model="form.npwp" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700">Alamat Lengkap</label>
          <textarea v-model="form.address" rows="3" required class="mt-1 block w-full border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"></textarea>
        </div>

        <div class="border-t border-slate-200 pt-6 mt-6">
          <h3 class="text-lg font-medium text-slate-900 mb-4">Pengaturan Global</h3>
          <div>
            <label class="block text-sm font-medium text-slate-700">Persentase PPN (%)</label>
            <p class="text-xs text-slate-500 mb-2">Nilai ini akan digunakan sebagai default saat membuat faktur penjualan baru.</p>
            <div class="flex items-center">
              <input type="number" step="0.1" v-model="form.ppn_rate" required class="block w-32 border border-slate-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm" />
              <span class="ml-2 text-slate-500 font-medium">%</span>
            </div>
          </div>
        </div>

        <div class="pt-4 flex justify-end">
          <button type="submit" :disabled="loading" class="inline-flex justify-center py-2 px-6 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50">
            {{ loading ? 'Menyimpan...' : 'Simpan Perubahan' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import api from '@/plugins/axios'

const loading = ref(false)
const successMsg = ref('')
const errorMsg = ref('')

const form = ref({
  name: '',
  npwp: '',
  address: '',
  ppn_rate: 11
})

const fetchProfile = async () => {
  try {
    const res = await api.get('/api/settings/company')
    if (res.data) {
      form.value = {
        name: res.data.name || '',
        npwp: res.data.npwp || '',
        address: res.data.address || '',
        ppn_rate: res.data.ppn_rate || 11
      }
    }
  } catch (error) {
    console.error('Failed to fetch profile', error)
  }
}

const saveProfile = async () => {
  loading.value = true
  successMsg.value = ''
  errorMsg.value = ''
  try {
    await api.put('/api/settings/company', form.value)
    successMsg.value = 'Profil perusahaan berhasil diperbarui!'
    setTimeout(() => { successMsg.value = '' }, 3000)
  } catch (error: any) {
    console.error('Failed to save profile', error)
    errorMsg.value = error.response?.data?.error || 'Gagal menyimpan profil.'
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProfile()
})
</script>
