<template>
  <div class="space-y-6 max-w-3xl">
    <div>
      <h2 class="text-xl sm:text-2xl font-bold text-slate-800">Profil Perusahaan &amp; Pengaturan Faktur</h2>
      <p class="text-slate-500 text-sm">Kelola identitas resmi perusahaan yang akan tercetak otomatis pada kop dan catatan faktur penjualan.</p>
    </div>

    <div class="bg-white shadow-sm rounded-2xl border border-slate-200 p-5 sm:p-7">
      <form @submit.prevent="saveProfile" class="space-y-5">
        
        <div v-if="successMsg" class="p-3.5 bg-emerald-50 text-emerald-700 text-sm rounded-xl border border-emerald-200 flex items-center gap-2">
          <span>✓</span> {{ successMsg }}
        </div>
        <div v-if="errorMsg" class="p-3.5 bg-rose-50 text-rose-700 text-sm rounded-xl border border-rose-200 flex items-center gap-2">
          <span>✕</span> {{ errorMsg }}
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <!-- Nama Perusahaan -->
          <div class="sm:col-span-2">
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1.5">
              Nama Perusahaan / Usaha Dagang <span class="text-red-500">*</span>
            </label>
            <input 
              type="text" 
              v-model="form.name" 
              required 
              placeholder="Contoh: UD DUO SRIKANDI" 
              class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm font-semibold text-slate-900" 
            />
          </div>

          <!-- Nomor Telepon / WA -->
          <div>
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1.5">
              No. Telepon / WhatsApp
            </label>
            <input 
              type="text" 
              v-model="form.phone" 
              placeholder="+62 812-3456-7890" 
              class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm" 
            />
          </div>

          <!-- Email Bisnis -->
          <div>
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1.5">
              Email Bisnis
            </label>
            <input 
              type="email" 
              v-model="form.email" 
              placeholder="kontak@duosrikandi.com" 
              class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm" 
            />
          </div>

          <!-- NPWP / NIK -->
          <div class="sm:col-span-2">
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1.5">
              Nomor NPWP / NIK Usaha <span class="text-red-500">*</span>
            </label>
            <input 
              type="text" 
              v-model="form.npwp" 
              required 
              placeholder="31.456.789.0-604.000" 
              class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm" 
            />
          </div>

          <!-- Alamat Lengkap -->
          <div class="sm:col-span-2">
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1.5">
              Alamat Lengkap Operasional <span class="text-red-500">*</span>
            </label>
            <textarea 
              v-model="form.address" 
              rows="3" 
              required 
              placeholder="Jl. Perdagangan Raya No. 88, Jawa Timur" 
              class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm"
            ></textarea>
          </div>
        </div>

        <!-- Rekening Pembayaran -->
        <div class="border-t border-slate-200 pt-5 mt-5">
          <h3 class="text-sm font-bold text-slate-900 mb-1">Informasi Rekening Pembayaran (Bank Transfer)</h3>
          <p class="text-xs text-slate-500 mb-2">Informasi ini akan tercetak di bagian bawah faktur agar klien dapat melakukan transfer.</p>
          <textarea 
            v-model="form.bank_info" 
            rows="2" 
            placeholder="Contoh: BCA: 8870-123-456 a/n UD DUO SRIKANDI&#10;BRI: 0123-01-000456-50-1 a/n UD DUO SRIKANDI" 
            class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm font-mono text-xs"
          ></textarea>
        </div>

        <!-- Catatan Standar Faktur -->
        <div class="border-t border-slate-200 pt-5 mt-5">
          <h3 class="text-sm font-bold text-slate-900 mb-1">Catatan Kaki &amp; Syarat Standar Faktur</h3>
          <p class="text-xs text-slate-500 mb-2">Teks keterangan / syarat pengembalian yang tercetak di bawah faktur.</p>
          <textarea 
            v-model="form.invoice_footer" 
            rows="2" 
            placeholder="Contoh: Barang yang sudah dibeli tidak dapat ditukar/dikembalikan kecuali ada perjanjian tertulis." 
            class="block w-full border border-slate-300 rounded-xl py-2.5 px-3.5 focus:outline-none focus:ring-2 focus:ring-rose-500/50 focus:border-rose-500 text-sm text-slate-700"
          ></textarea>
        </div>

        <!-- Pengaturan PPN -->
        <div class="border-t border-slate-200 pt-5 mt-5">
          <h3 class="text-sm font-bold text-slate-900 mb-1">Pengaturan Pajak Global</h3>
          <p class="text-xs text-slate-500 mb-2">Persentase PPN default yang diterapkan saat membuat faktur baru.</p>
          <div class="flex items-center gap-2">
            <input 
              type="number" 
              step="0.1" 
              v-model="form.ppn_rate" 
              required 
              class="block w-28 border border-slate-300 rounded-xl py-2 px-3 focus:outline-none focus:ring-2 focus:ring-rose-500/50 text-sm font-semibold" 
            />
            <span class="text-slate-600 font-semibold text-sm">%</span>
          </div>
        </div>

        <!-- Submit Button -->
        <div class="pt-5 flex justify-end border-t border-slate-100">
          <button 
            type="submit" 
            :disabled="loading" 
            class="w-full sm:w-auto inline-flex justify-center items-center py-2.5 px-6 rounded-xl text-sm font-bold text-white shadow-md transition-colors disabled:opacity-50 cursor-pointer"
            style="background: linear-gradient(135deg, #A82020, #7B1D1D);"
          >
            {{ loading ? 'Menyimpan...' : 'Simpan Profil Perusahaan' }}
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
  name: 'UD DUO SRIKANDI',
  phone: '+62 812-3456-7890',
  email: 'kontak@duosrikandi.com',
  npwp: '31.456.789.0-604.000',
  address: 'Jl. Perdagangan Raya No. 88, Jawa Timur',
  bank_info: 'BCA: 8870-123-456 a/n UD DUO SRIKANDI\nBRI: 0123-01-000456-50-1 a/n UD DUO SRIKANDI',
  invoice_footer: 'Barang yang sudah dibeli tidak dapat ditukar/dikembalikan kecuali ada perjanjian tertulis sebelumnya.',
  ppn_rate: 11
})

const fetchProfile = async () => {
  try {
    const res = await api.get('/api/settings/company')
    if (res.data) {
      form.value = {
        name: res.data.name || 'UD DUO SRIKANDI',
        phone: res.data.phone || '',
        email: res.data.email || '',
        npwp: res.data.npwp || '',
        address: res.data.address || '',
        bank_info: res.data.bank_info || '',
        invoice_footer: res.data.invoice_footer || '',
        ppn_rate: res.data.ppn_rate !== undefined ? res.data.ppn_rate : 11
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
    successMsg.value = 'Profil perusahaan dan template faktur berhasil diperbarui!'
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
