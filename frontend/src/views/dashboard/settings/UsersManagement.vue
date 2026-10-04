<template>
  <div class="space-y-6">
    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
      <div>
        <h2 class="text-xl sm:text-2xl font-bold text-slate-800">Manajemen Pengguna</h2>
        <p class="text-slate-500 text-sm">Kelola akses pengguna ke dalam sistem aplikasi.</p>
      </div>
      <button 
        @click="openCreateModal" 
        class="inline-flex items-center justify-center px-4 py-2.5 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 font-medium text-sm shadow-sm transition-colors cursor-pointer w-full sm:w-auto"
      >
        + Tambah Pengguna
      </button>
    </div>

    <!-- Tab Navigation -->
    <div class="flex items-center gap-2 border-b border-slate-200">
      <button
        @click="activeTab = 'active'"
        class="pb-3 px-4 text-sm font-semibold transition-colors border-b-2 flex items-center gap-2 cursor-pointer"
        :class="activeTab === 'active' ? 'border-indigo-600 text-indigo-600' : 'border-transparent text-slate-500 hover:text-slate-700'"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
        </svg>
        Pengguna Aktif
      </button>
      <button
        @click="activeTab = 'trash'"
        class="pb-3 px-4 text-sm font-semibold transition-colors border-b-2 flex items-center gap-2 cursor-pointer"
        :class="activeTab === 'trash' ? 'border-rose-600 text-rose-600' : 'border-transparent text-slate-500 hover:text-slate-700'"
      >
        <RotateCcw class="w-4 h-4" />
        Kotak Sampah (Terhapus)
      </button>
      <button
        @click="activeTab = 'all'"
        class="pb-3 px-4 text-sm font-semibold transition-colors border-b-2 flex items-center gap-2 cursor-pointer"
        :class="activeTab === 'all' ? 'border-slate-800 text-slate-800' : 'border-transparent text-slate-500 hover:text-slate-700'"
      >
        Semua Pengguna
      </button>
    </div>

    <!-- Table Container with Horizontal Scroll for Mobile -->
    <div class="bg-white shadow rounded-xl border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-slate-200">
          <thead class="bg-slate-50">
            <tr>
              <th class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Username</th>
              <th class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Email</th>
              <th class="px-4 sm:px-6 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Role</th>
              <th class="px-4 sm:px-6 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Status</th>
              <th class="px-4 sm:px-6 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Aksi</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="5" class="px-6 py-8 text-center text-slate-400 text-sm">Memuat data pengguna...</td>
            </tr>
            <tr v-else-if="filteredUsers.length === 0">
              <td colspan="5" class="px-6 py-8 text-center text-slate-400 text-sm">
                {{ activeTab === 'trash' ? 'Tidak ada pengguna di kotak sampah.' : 'Belum ada pengguna.' }}
              </td>
            </tr>
            <tr v-else v-for="user in filteredUsers" :key="user.ID" class="hover:bg-slate-50/60 transition-colors" :class="user.deleted_at ? 'bg-slate-50/50 opacity-80' : ''">
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm font-semibold text-slate-900">{{ user.username }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ user.email || '-' }}</td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-sm">
                <span class="px-2.5 py-0.5 inline-flex text-xs leading-5 font-semibold rounded-full capitalize" 
                  :class="user.role === 'admin' ? 'bg-purple-100 text-purple-800' : 'bg-blue-100 text-blue-800'">
                  {{ user.role }}
                </span>
              </td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-center text-xs">
                <span v-if="user.deleted_at" class="px-2.5 py-0.5 inline-flex font-semibold rounded-full bg-rose-100 text-rose-700">
                  Terhapus
                </span>
                <span v-else class="px-2.5 py-0.5 inline-flex font-semibold rounded-full bg-emerald-100 text-emerald-800">
                  Aktif
                </span>
              </td>
              <td class="px-4 sm:px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-1.5">
                  <button
                    v-if="!user.deleted_at"
                    @click="openEditModal(user)"
                    class="p-1.5 text-indigo-600 hover:text-indigo-800 hover:bg-indigo-50 rounded-md transition-colors cursor-pointer inline-flex items-center justify-center"
                    title="Edit Pengguna"
                  >
                    <Pencil class="w-4 h-4" />
                  </button>
                  <button
                    v-if="user.deleted_at"
                    @click="openRestoreModal(user)"
                    class="inline-flex items-center gap-1 px-2.5 py-1.5 text-xs font-semibold text-emerald-700 bg-emerald-50 hover:bg-emerald-100 border border-emerald-200 rounded-lg transition-colors cursor-pointer"
                    title="Pulihkan Pengguna"
                  >
                    <RotateCcw class="w-3.5 h-3.5" />
                    Pulihkan
                  </button>
                  <button
                    v-else
                    @click="openDeleteModal(user)"
                    class="p-1.5 text-rose-600 hover:text-rose-800 hover:bg-rose-50 rounded-md transition-colors cursor-pointer inline-flex items-center justify-center"
                    title="Hapus Pengguna"
                  >
                    <Trash2 class="w-4 h-4" />
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- CREATE/EDIT MODAL (Mobile Friendly Scrollable) -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen p-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md z-10 p-5 sm:p-6 max-h-[90vh] overflow-y-auto">
          <h3 class="text-lg font-bold text-slate-800 mb-4">{{ isEditing ? 'Edit Pengguna' : 'Tambah Pengguna' }}</h3>
          
          <form @submit.prevent="submitForm" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Username <span class="text-red-500">*</span></label>
              <input type="text" v-model="form.username" required class="w-full border border-slate-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Email</label>
              <input type="email" v-model="form.email" placeholder="user@domain.com" class="w-full border border-slate-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">
                Kata Sandi 
                <span v-if="!isEditing" class="text-red-500">*</span>
                <span v-else class="text-xs text-slate-400 font-normal ml-1">(Kosongkan jika tidak diubah)</span>
              </label>
              <input type="password" v-model="form.password" :required="!isEditing" class="w-full border border-slate-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Role Akses <span class="text-red-500">*</span></label>
              <select v-model="form.role" required class="w-full border border-slate-300 rounded-lg px-3 py-2 text-sm focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none bg-white">
                <option value="user">User</option>
                <option value="admin">Admin</option>
              </select>
            </div>

            <div class="mt-6 flex justify-end gap-3 pt-3 border-t border-slate-100">
              <button type="button" @click="closeModal" class="px-4 py-2 text-sm border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors cursor-pointer">Batal</button>
              <button type="submit" :disabled="submitLoading" class="px-4 py-2 text-sm bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 transition-colors disabled:opacity-50 cursor-pointer">
                {{ submitLoading ? 'Menyimpan...' : 'Simpan' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- DELETE MODAL -->
    <div v-if="isDeleteModalOpen && deletingUser" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen p-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isDeleteModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-sm z-10 p-6 text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-rose-100 mb-4">
            <svg class="h-6 w-6 text-rose-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900 mb-2">Hapus Pengguna</h3>
          <p class="text-sm text-slate-500 mb-6">Pengguna <b>{{ deletingUser.username }}</b> akan dipindahkan ke Kotak Sampah.</p>
          <div class="flex justify-center gap-3">
            <button @click="isDeleteModalOpen = false" class="px-4 py-2 text-sm font-medium border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors cursor-pointer">Batal</button>
            <button @click="submitDelete" :disabled="deleteLoading" class="px-4 py-2 text-sm font-medium bg-rose-600 text-white rounded-lg hover:bg-rose-700 transition-colors disabled:opacity-50 cursor-pointer">
              {{ deleteLoading ? 'Menghapus...' : 'Ya, Hapus' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- RESTORE MODAL -->
    <div v-if="isRestoreModalOpen && restoringUser" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen p-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isRestoreModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-sm z-10 p-6 text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-emerald-100 text-emerald-600 mb-4">
            <RotateCcw class="w-6 h-6" />
          </div>
          <h3 class="text-lg font-bold text-slate-900 mb-2">Pulihkan Pengguna</h3>
          <p class="text-sm text-slate-500 mb-6">Apakah Anda yakin ingin mengembalikan akses pengguna <b>{{ restoringUser.username }}</b>?</p>
          <div class="flex justify-center gap-3">
            <button @click="isRestoreModalOpen = false" class="px-4 py-2 text-sm font-medium border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors cursor-pointer">Batal</button>
            <button @click="submitRestore" :disabled="restoreLoading" class="px-4 py-2 text-sm font-medium bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors disabled:opacity-50 cursor-pointer">
              {{ restoreLoading ? 'Memulihkan...' : 'Ya, Pulihkan' }}
            </button>
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import api from '@/plugins/axios'
import { notify } from '@/lib/notify'
import { Pencil, Trash2, RotateCcw } from '@lucide/vue'

interface User {
  ID: number
  username: string
  email: string
  role: string
  deleted_at?: string | null
}

const activeTab = ref<'active' | 'trash' | 'all'>('active')
const users = ref<User[]>([])
const loading = ref(false)

const isModalOpen = ref(false)
const isEditing = ref(false)
const submitLoading = ref(false)

const form = ref({
  id: 0,
  username: '',
  email: '',
  password: '',
  role: 'user'
})

const isDeleteModalOpen = ref(false)
const deleteLoading = ref(false)
const deletingUser = ref<User | null>(null)

const isRestoreModalOpen = ref(false)
const restoreLoading = ref(false)
const restoringUser = ref<User | null>(null)

const filteredUsers = computed(() => {
  if (activeTab.value === 'trash') {
    return users.value.filter(u => !!u.deleted_at)
  }
  if (activeTab.value === 'active') {
    return users.value.filter(u => !u.deleted_at)
  }
  return users.value
})

const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/settings/users')
    users.value = res.data || []
  } catch (error: any) {
    notify.error('Gagal Memuat Pengguna', error.response?.data?.error || 'Gagal memuat data pengguna.')
  } finally {
    loading.value = false
  }
}

const openCreateModal = () => {
  isEditing.value = false
  form.value = { id: 0, username: '', email: '', password: '', role: 'user' }
  isModalOpen.value = true
}

const openEditModal = (user: User) => {
  isEditing.value = true
  form.value = {
    id: user.ID,
    username: user.username,
    email: user.email,
    password: '',
    role: user.role
  }
  isModalOpen.value = true
}

const closeModal = () => {
  isModalOpen.value = false
}

const submitForm = async () => {
  submitLoading.value = true
  try {
    const payload: any = {
      username: form.value.username,
      email: form.value.email,
      role: form.value.role,
    }
    if (form.value.password) {
      payload.password = form.value.password
    }

    if (isEditing.value) {
      await api.put(`/api/settings/users/${form.value.id}`, payload)
      closeModal()
      notify.success('Pengguna Diperbarui', `Akun ${form.value.username} berhasil diperbarui.`)
    } else {
      await api.post('/api/settings/users', payload)
      closeModal()
      notify.success('Pengguna Ditambahkan', `Akun ${form.value.username} berhasil dibuat.`)
    }
    await fetchUsers()
  } catch (error: any) {
    notify.error('Gagal Menyimpan', error.response?.data?.error || 'Gagal menyimpan pengguna.')
  } finally {
    submitLoading.value = false
  }
}

const openDeleteModal = (user: User) => {
  deletingUser.value = user
  isDeleteModalOpen.value = true
}

const submitDelete = async () => {
  if (!deletingUser.value) return
  deleteLoading.value = true
  const deletedUsername = deletingUser.value.username
  try {
    await api.delete(`/api/settings/users/${deletingUser.value.ID}`)
    isDeleteModalOpen.value = false
    notify.success('Pengguna Dihapus', `Pengguna ${deletedUsername} berhasil dipindahkan ke kotak sampah`)
    await fetchUsers()
  } catch (error: any) {
    notify.error('Gagal Menghapus', error.response?.data?.error || 'Gagal menghapus pengguna.')
  } finally {
    deleteLoading.value = false
  }
}

const openRestoreModal = (user: User) => {
  restoringUser.value = user
  isRestoreModalOpen.value = true
}

const submitRestore = async () => {
  if (!restoringUser.value) return
  restoreLoading.value = true
  const uName = restoringUser.value.username
  try {
    await api.post(`/api/settings/users/${restoringUser.value.ID}/restore`, {})
    isRestoreModalOpen.value = false
    notify.success('Pengguna Dipulihkan', `Pengguna ${uName} berhasil dipulihkan`)
    await fetchUsers()
  } catch (error: any) {
    notify.error('Gagal Memulihkan', error.response?.data?.error || 'Gagal memulihkan pengguna.')
  } finally {
    restoreLoading.value = false
  }
}

onMounted(() => {
  fetchUsers()
})
</script>
