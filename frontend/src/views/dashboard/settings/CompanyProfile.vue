<template>
  <div class="space-y-6 max-w-2xl">
    <div>
      <h2 class="text-xl sm:text-2xl font-bold text-slate-800">Profil Perusahaan</h2>
      <p class="text-slate-500 text-sm">Kelola identitas perusahaan yang akan tampil pada kop faktur cetak.</p>
    </div>

    <div class="bg-white shadow-sm rounded-xl border border-slate-200 p-4 sm:p-6">
      <form @submit.prevent="saveProfile" class="space-y-5">
        
        <div v-if="successMsg" class="p-3.5 bg-emerald-50 text-emerald-700 text-sm rounded-lg border border-emerald-200 flex items-center gap-2">
          <span>✓</span> {{ successMsg }}
        </div>
        <div v-if="errorMsg" class="p-3.5 bg-rose-50 text-rose-700 text-sm rounded-lg border border-rose-200 flex items-center gap-2">
          <span>✕</span> {{ errorMsg }}
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700 mb-1">Nama Perusahaan <span class="text-red-500">*</span></label>
          <input type="text" v-model="form.name" required class="block w-full border border-slate-300 rounded-lg py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm" />
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700 mb-1">NPWP Perusahaan <span class="text-red-500">*</span></label>
          <input type="text" v-model="form.npwp" required class="block w-full border border-slate-300 rounded-lg py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm" />
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700 mb-1">Alamat Lengkap <span class="text-red-500">*</span></label>
          <textarea v-model="form.address" rows="3" required class="block w-full border border-slate-300 rounded-lg py-2.5 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm"></textarea>
        </div>

        <div class="border-t border-slate-200 pt-5 mt-5">
          <h3 class="text-base font-semibold text-slate-900 mb-2">Pengaturan Pajak Global</h3>
          <p class="text-xs text-slate-500 mb-3">Persentase PPN default yang akan diterapkan pada pembuatan faktur baru.</p>
          <div class="flex items-center gap-2">
            <input type="number" step="0.1" v-model="form.ppn_rate" required class="block w-32 border border-slate-300 rounded-lg py-2 px-3 focus:outline-none focus:ring-2 focus:ring-indigo-500 text-sm" />
            <span class="text-slate-600 font-semibold text-sm">%</span>
          </div>
        </div>

        <div class="pt-4 flex justify-end border-t border-slate-100">
          <button 
            type="submit" 
            :disabled="loading" 
            class="w-full sm:w-auto inline-flex justify-center items-center py-2.5 px-6 rounded-lg text-sm font-semibold text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 transition-colors shadow-sm cursor-pointer"
          >
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
import { toast } from 'vue-sonner'

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
    toast.success('Profil perusahaan berhasil diperbarui')
    setTimeout(() => { successMsg.value = '' }, 3000)
  } catch (error: any) {
    console.error('Failed to save profile', error)
    errorMsg.value = error.response?.data?.error || 'Gagal menyimpan profil.'
    toast.error(errorMsg.value)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProfile()
})
</script>
