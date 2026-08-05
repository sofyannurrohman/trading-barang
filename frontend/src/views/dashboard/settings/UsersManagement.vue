<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Manajemen Pengguna</h2>
        <p class="text-slate-500">Kelola akses pengguna ke dalam sistem aplikasi.</p>
      </div>
      <button @click="openCreateModal" class="px-4 py-2 bg-indigo-600 text-white rounded-md hover:bg-indigo-700 font-medium transition-colors">
        + Tambah Pengguna
      </button>
    </div>

    <!-- Table -->
    <div class="bg-white shadow rounded-lg border border-slate-200 overflow-hidden">
      <table class="min-w-full divide-y divide-slate-200">
        <thead class="bg-slate-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Username</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Email</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-slate-500 uppercase tracking-wider">Role</th>
            <th class="px-6 py-3 text-right text-xs font-medium text-slate-500 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-slate-200">
          <tr v-if="loading">
            <td colspan="4" class="px-6 py-4 text-center text-slate-500">Loading data...</td>
          </tr>
          <tr v-else-if="users.length === 0">
            <td colspan="4" class="px-6 py-4 text-center text-slate-500">Belum ada pengguna.</td>
          </tr>
          <tr v-else v-for="user in users" :key="user.ID">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-slate-900">{{ user.username }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-slate-500">{{ user.email || '-' }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm">
              <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full" 
                :class="user.role === 'admin' ? 'bg-purple-100 text-purple-800' : 'bg-blue-100 text-blue-800'">
                {{ user.role }}
              </span>
            </td>
            <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
              <div class="flex items-center justify-end gap-3">
                <button @click="openEditModal(user)" class="text-indigo-600 hover:text-indigo-900">Edit</button>
                <button @click="openDeleteModal(user)" class="text-red-600 hover:text-red-900">Hapus</button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- CREATE/EDIT MODAL -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md z-10 p-6">
          <h3 class="text-lg font-bold mb-4">{{ isEditing ? 'Edit Pengguna' : 'Tambah Pengguna' }}</h3>
          
          <form @submit.prevent="submitForm" class="space-y-4">
            <div>
              <label class="block text-sm font-medium mb-1">Username <span class="text-red-500">*</span></label>
              <input type="text" v-model="form.username" required class="w-full border border-slate-300 rounded-lg px-3 py-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1">Email</label>
              <input type="email" v-model="form.email" class="w-full border border-slate-300 rounded-lg px-3 py-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1">
                Password 
                <span v-if="!isEditing" class="text-red-500">*</span>
                <span v-else class="text-xs text-slate-400 font-normal ml-1">(Kosongkan jika tidak diubah)</span>
              </label>
              <input type="password" v-model="form.password" :required="!isEditing" class="w-full border border-slate-300 rounded-lg px-3 py-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none" />
            </div>
            <div>
              <label class="block text-sm font-medium mb-1">Role <span class="text-red-500">*</span></label>
              <select v-model="form.role" required class="w-full border border-slate-300 rounded-lg px-3 py-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none bg-white">
                <option value="user">User</option>
                <option value="admin">Admin</option>
              </select>
            </div>

            <div class="mt-6 flex justify-end gap-3">
              <button type="button" @click="closeModal" class="px-4 py-2 text-sm border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">Batal</button>
              <button type="submit" :disabled="submitLoading" class="px-4 py-2 text-sm bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors disabled:opacity-50">
                {{ submitLoading ? 'Menyimpan...' : 'Simpan' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- DELETE MODAL -->
    <div v-if="isDeleteModalOpen && deletingUser" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isDeleteModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-sm z-10 p-6 text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-100 mb-4">
            <svg class="h-6 w-6 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
            </svg>
          </div>
          <h3 class="text-lg font-bold mb-2">Hapus Pengguna</h3>
          <p class="text-sm text-slate-500 mb-6">Apakah Anda yakin ingin menghapus <b>{{ deletingUser.username }}</b>? Tindakan ini tidak dapat dibatalkan.</p>
          <div class="flex justify-center gap-3">
            <button @click="isDeleteModalOpen = false" class="px-4 py-2 text-sm font-medium border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">Batal</button>
            <button @click="submitDelete" :disabled="deleteLoading" class="px-4 py-2 text-sm font-medium bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors disabled:opacity-50">
              {{ deleteLoading ? 'Menghapus...' : 'Ya, Hapus' }}
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

interface User {
  ID: number
  username: string
  email: string
  role: string
}

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

const fetchUsers = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/settings/users')
    users.value = res.data
  } catch (error: any) {
    toast.error(error.response?.data?.error || 'Gagal memuat data pengguna.')
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
    // Only send password if it's provided (required for create, optional for edit)
    if (form.value.password) {
      payload.password = form.value.password
    }

    if (isEditing.value) {
      await api.put(`/api/settings/users/${form.value.id}`, payload, { skipToast: true } as any)
      closeModal()
      toast.success('Pengguna berhasil diperbarui')
    } else {
      await api.post('/api/settings/users', payload, { skipToast: true } as any)
      closeModal()
      toast.success('Pengguna berhasil ditambahkan')
    }
    await fetchUsers()
  } catch (error: any) {
    toast.error(error.response?.data?.error || 'Gagal menyimpan pengguna.')
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
    await api.delete(`/api/settings/users/${deletingUser.value.ID}`, { skipToast: true } as any)
    isDeleteModalOpen.value = false
    toast.success(`Pengguna ${deletedUsername} berhasil dihapus`)
    await fetchUsers()
  } catch (error: any) {
    toast.error(error.response?.data?.error || 'Gagal menghapus pengguna.')
  } finally {
    deleteLoading.value = false
  }
}

onMounted(() => {
  fetchUsers()
})
</script>
