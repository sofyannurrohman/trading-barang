<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Data Mitra</h2>
        <p class="text-slate-500">Kelola daftar Klien dan Supplier beserta informasi kontak dan identitas.</p>
      </div>
      <button
        @click="openDialog()"
        class="inline-flex items-center gap-2 px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 font-medium shadow-sm transition-colors"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Tambah Mitra
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white shadow-sm rounded-xl border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-slate-200">
          <thead class="bg-slate-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Tipe</th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Nama</th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">NPWP / NIK</th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Alamat</th>
              <th scope="col" class="px-6 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Aksi</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="5" class="px-6 py-10 text-center">
                <div class="flex flex-col items-center gap-2 text-slate-400">
                  <svg class="animate-spin w-6 h-6" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  <span class="text-sm">Memuat data...</span>
                </div>
              </td>
            </tr>
            <tr v-else-if="partners.length === 0">
              <td colspan="5" class="px-6 py-10 text-center text-slate-400 text-sm">
                Tidak ada mitra ditemukan.
              </td>
            </tr>
            <tr v-else v-for="partner in partners" :key="partner.ID" class="hover:bg-slate-50 transition-colors">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <span
                  :class="partner.type === 'client' ? 'bg-blue-100 text-blue-800' : 'bg-purple-100 text-purple-800'"
                  class="px-2.5 py-0.5 inline-flex text-xs leading-5 font-semibold rounded-full uppercase"
                >
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
                <div class="flex items-center justify-end gap-2">
                  <button
                    @click="openDialog(partner)"
                    class="p-1.5 text-indigo-600 hover:text-indigo-800 hover:bg-indigo-50 rounded-md transition-colors"
                    title="Edit Mitra"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button
                    @click="openDeleteConfirm(partner)"
                    class="p-1.5 text-red-600 hover:text-red-800 hover:bg-red-50 rounded-md transition-colors"
                    title="Hapus Mitra"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL: TAMBAH / EDIT MITRA     -->
    <!-- ============================== -->
    <div v-if="isDialogOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isDialogOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-lg z-10">
          <!-- Modal Header -->
          <div class="px-6 pt-6 pb-4 border-b border-slate-200">
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-bold text-slate-900">
                {{ isEditing ? 'Edit Mitra' : 'Tambah Mitra Baru' }}
              </h3>
              <button
                @click="isDialogOpen = false"
                class="p-2 text-slate-400 hover:text-slate-600 hover:bg-slate-100 rounded-lg transition-colors"
              >
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>

          <form @submit.prevent="savePartner">
            <div class="px-6 py-5 space-y-4">
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Tipe Mitra <span class="text-red-500">*</span></label>
                <select
                  v-model="form.type"
                  required
                  class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                >
                  <option value="client">Client (Pembeli)</option>
                  <option value="supplier">Supplier (Pemasok)</option>
                </select>
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Nama <span class="text-red-500">*</span></label>
                <input
                  type="text"
                  v-model="form.name"
                  required
                  placeholder="Nama perusahaan atau perorangan"
                  class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">NPWP</label>
                <input
                  type="text"
                  v-model="form.npwp"
                  placeholder="Opsional"
                  class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">
                  NIK <span class="text-slate-400 font-normal">(Wajib jika tidak ada NPWP)</span>
                </label>
                <input
                  type="text"
                  v-model="form.nik"
                  placeholder="Opsional jika ada NPWP"
                  class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Alamat Lengkap</label>
                <textarea
                  v-model="form.address"
                  rows="3"
                  placeholder="Alamat lengkap mitra"
                  class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent resize-none"
                ></textarea>
              </div>
            </div>

            <div class="px-6 py-4 border-t border-slate-200 flex justify-end gap-3">
              <button
                type="button"
                @click="isDialogOpen = false"
                class="px-4 py-2 text-sm font-medium text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors"
              >
                Batal
              </button>
              <button
                type="submit"
                :disabled="saveLoading"
                class="px-5 py-2 text-sm font-semibold bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition-colors flex items-center gap-2"
              >
                <svg v-if="saveLoading" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                </svg>
                {{ saveLoading ? 'Menyimpan...' : 'Simpan Mitra' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL: KONFIRMASI HAPUS MITRA  -->
    <!-- ============================== -->
    <div v-if="isDeleteConfirmOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isDeleteConfirmOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm z-10">
          <!-- Modal Header (Danger) -->
          <div class="px-6 pt-6 pb-4 border-b border-red-100 bg-red-50 rounded-t-2xl">
            <div class="flex items-start gap-3">
              <div class="flex-shrink-0 w-10 h-10 bg-red-100 border border-red-200 rounded-xl flex items-center justify-center">
                <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-bold text-red-800">Hapus Mitra</h3>
                <p class="text-sm text-red-600 mt-0.5">Tindakan ini tidak dapat dibatalkan.</p>
              </div>
            </div>
          </div>

          <!-- Body -->
          <div class="px-6 py-5">
            <p class="text-sm text-slate-700">
              Anda akan menghapus mitra
              <span class="font-semibold text-slate-900">{{ deletingPartnerName }}</span>.
              Data ini akan dihapus secara permanen dari sistem.
            </p>
          </div>

          <!-- Footer -->
          <div class="px-6 py-4 border-t border-slate-200 flex justify-end gap-3">
            <button
              @click="isDeleteConfirmOpen = false"
              class="px-4 py-2 text-sm font-medium text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors"
            >
              Batal
            </button>
            <button
              @click="submitDelete"
              :disabled="deleteLoading"
              class="px-5 py-2 text-sm font-semibold bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 transition-colors flex items-center gap-2"
            >
              <svg v-if="deleteLoading" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              {{ deleteLoading ? 'Menghapus...' : 'Ya, Hapus Mitra' }}
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

interface Partner {
  ID?: number
  type: string
  name: string
  address: string
  npwp: string
  nik: string
}

// ─── State ────────────────────────────────────────────────────────────────────
const partners = ref<Partner[]>([])
const loading = ref(false)

// Create / Edit Modal
const isDialogOpen = ref(false)
const isEditing = ref(false)
const saveLoading = ref(false)
const form = ref<Partner>({
  type: 'client',
  name: '',
  address: '',
  npwp: '',
  nik: ''
})

// Delete Confirm Modal
const isDeleteConfirmOpen = ref(false)
const deleteLoading = ref(false)
const deletingPartnerId = ref<number | undefined>(undefined)
const deletingPartnerName = ref('')

// ─── Fetch ────────────────────────────────────────────────────────────────────
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

// ─── Create & Edit ────────────────────────────────────────────────────────────
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
  // Validasi: minimal salah satu dari NPWP atau NIK harus diisi
  if (!form.value.npwp && !form.value.nik) {
    toast.warning('Harap isi minimal salah satu dari NPWP atau NIK untuk kebutuhan pajak.')
    return
  }

  saveLoading.value = true
  const wasEditing = isEditing.value
  try {
    if (isEditing.value && form.value.ID) {
      await api.put(`/api/partners/${form.value.ID}`, form.value, { skipToast: true } as any)
    } else {
      await api.post('/api/partners', form.value, { skipToast: true } as any)
    }
    // Tutup modal dulu, baru tampilkan toast agar tidak terhalangi backdrop
    isDialogOpen.value = false
    toast.success(wasEditing ? 'Mitra berhasil diperbarui' : 'Mitra berhasil ditambahkan')
    await fetchPartners()
  } catch (error: any) {
    console.error('Failed to save partner', error)
    const errMsg = error.response?.data?.error ?? 'Gagal menyimpan data mitra.'
    toast.error(errMsg)
  } finally {
    saveLoading.value = false
  }
}

// ─── Delete ───────────────────────────────────────────────────────────────────
const openDeleteConfirm = (partner: Partner) => {
  deletingPartnerId.value = partner.ID
  deletingPartnerName.value = partner.name
  isDeleteConfirmOpen.value = true
}

const submitDelete = async () => {
  if (!deletingPartnerId.value) return
  deleteLoading.value = true
  try {
    await api.delete(`/api/partners/${deletingPartnerId.value}`, { skipToast: true } as any)
    // Tutup modal dulu, baru tampilkan toast agar tidak terhalangi backdrop
    isDeleteConfirmOpen.value = false
    toast.success('Mitra berhasil dihapus')
    await fetchPartners()
  } catch (error: any) {
    console.error('Failed to delete partner', error)
    toast.error(error.response?.data?.error ?? 'Gagal menghapus mitra.')
  } finally {
    deleteLoading.value = false
  }
}

// ─── Init ─────────────────────────────────────────────────────────────────────
onMounted(() => {
  fetchPartners()
})
</script>
