<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Manajemen Produk</h2>
        <p class="text-slate-500">Kelola master data produk, SKU, kategori, dan harga standar.</p>
      </div>
      <button
        @click="openCreateModal"
        class="inline-flex items-center gap-2 px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 font-medium shadow-sm transition-colors"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Tambah Produk
      </button>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
      <!-- Total Produk -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-indigo-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Total Produk</p>
          <p class="text-lg font-bold text-slate-900 truncate">{{ formatNumber(summary.totalProducts) }}</p>
          <p class="text-xs text-slate-400">item di master data</p>
        </div>
      </div>

      <!-- Total Nilai Stok (HPP) -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-emerald-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Total Nilai Stok (HPP)</p>
          <p class="text-lg font-bold text-slate-900 truncate">Rp {{ formatNumber(summary.totalStockValue) }}</p>
          <p class="text-xs text-slate-400">modal di gudang</p>
        </div>
      </div>

      <!-- Stok Rendah -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-amber-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Peringatan Stok</p>
          <p class="text-lg font-bold text-slate-900 truncate">{{ formatNumber(summary.lowStockCount) }}</p>
          <p class="text-xs text-slate-400">produk stok &lt; 10</p>
        </div>
      </div>

      <!-- Nilai Jual Total -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-purple-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Nilai Jual Stok</p>
          <p class="text-lg font-bold text-slate-900 truncate">Rp {{ formatNumber(summary.totalSalesValue) }}</p>
          <p class="text-xs text-slate-400">potensi pendapatan</p>
        </div>
      </div>
    </div>

    <!-- Filter Bar -->
    <div class="bg-white rounded-xl border border-slate-200 shadow-sm p-4 flex flex-wrap gap-3 items-center">
      <div class="flex-1 min-w-48">
        <div class="relative">
          <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Cari SKU atau nama produk..."
            class="pl-9 pr-3 py-2 w-full border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
          />
        </div>
      </div>
      <div class="flex gap-2 flex-wrap">
        <select
          v-model="filterCategory"
          class="px-3 py-2 border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
        >
          <option value="">Semua Kategori</option>
          <option v-for="cat in uniqueCategories" :key="cat" :value="cat">{{ cat }}</option>
        </select>
        <select
          v-model="filterStock"
          class="px-3 py-2 border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
        >
          <option value="">Semua Stok</option>
          <option value="low">Stok Rendah (&lt; 10)</option>
          <option value="empty">Stok Kosong (= 0)</option>
          <option value="available">Tersedia (&gt; 0)</option>
        </select>
        <button
          @click="resetFilters"
          class="px-3 py-2 text-sm text-slate-600 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors"
        >
          Reset
        </button>
      </div>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-slate-200">
          <thead class="bg-slate-50">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Foto</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">SKU</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Nama Produk</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Kategori</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Harga Standar</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">HPP Rata-rata</th>
              <th class="px-4 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Stok</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Aksi</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="8" class="px-4 py-10 text-center">
                <div class="flex flex-col items-center gap-2 text-slate-400">
                  <svg class="animate-spin w-6 h-6" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  <span class="text-sm">Memuat data...</span>
                </div>
              </td>
            </tr>
            <tr v-else-if="paginatedProducts.length === 0">
              <td colspan="8" class="px-4 py-10 text-center text-slate-400 text-sm">
                Tidak ada data produk yang ditemukan.
              </td>
            </tr>
            <tr
              v-else
              v-for="product in paginatedProducts"
              :key="product.ID"
              class="hover:bg-slate-50 transition-colors"
            >
              <!-- Foto Thumbnail -->
              <td class="px-4 py-3 whitespace-nowrap">
                <div class="w-10 h-10 rounded-lg overflow-hidden bg-slate-100 border border-slate-200 flex items-center justify-center flex-shrink-0">
                  <img
                    v-if="product.image_url"
                    :src="`${apiBase}${product.image_url}`"
                    class="w-full h-full object-cover"
                    :alt="product.name"
                  />
                  <ImageOff v-else class="w-5 h-5 text-slate-400" />
                </div>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm font-mono text-slate-600">{{ product.sku }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm font-medium text-slate-900">{{ product.name }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm">
                <span v-if="product.category" class="px-2 py-1 bg-slate-100 text-slate-600 text-xs rounded-md">{{ product.category }}</span>
                <span v-else class="text-slate-400 text-xs">-</span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-700 text-right">
                Rp {{ formatNumber(product.standard_price) }}
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-500 text-right">
                Rp {{ formatNumber(product.average_hpp) }}
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-center">
                <span
                  class="font-semibold"
                  :class="{
                    'text-red-600': product.current_stock === 0,
                    'text-amber-600': product.current_stock > 0 && product.current_stock < 10,
                    'text-emerald-600': product.current_stock >= 10
                  }"
                >
                  {{ product.current_stock }}
                </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-2">
                  <button
                    @click="openEditModal(product)"
                    class="p-1.5 text-indigo-600 hover:text-indigo-800 hover:bg-indigo-50 rounded-md transition-colors"
                    title="Edit Produk"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button
                    @click="openDeleteModal(product)"
                    class="p-1.5 text-red-600 hover:text-red-800 hover:bg-red-50 rounded-md transition-colors"
                    title="Hapus Produk"
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

      <!-- Pagination -->
      <div class="px-4 py-3 border-t border-slate-200 flex items-center justify-between bg-slate-50">
        <p class="text-xs text-slate-500">
          Menampilkan {{ paginationStart }}-{{ paginationEnd }} dari {{ filteredProducts.length }} data
        </p>
        <div class="flex gap-1">
          <button
            @click="currentPage--"
            :disabled="currentPage <= 1"
            class="px-3 py-1 text-sm border border-slate-300 rounded-md disabled:opacity-40 disabled:cursor-not-allowed hover:bg-white transition-colors"
          >
            ‹
          </button>
          <button
            v-for="page in totalPages"
            :key="page"
            @click="currentPage = page"
            class="px-3 py-1 text-sm border rounded-md transition-colors"
            :class="currentPage === page ? 'bg-indigo-600 text-white border-indigo-600' : 'border-slate-300 hover:bg-white'"
          >
            {{ page }}
          </button>
          <button
            @click="currentPage++"
            :disabled="currentPage >= totalPages"
            class="px-3 py-1 text-sm border border-slate-300 rounded-md disabled:opacity-40 disabled:cursor-not-allowed hover:bg-white transition-colors"
          >
            ›
          </button>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL: BUAT/EDIT PRODUK        -->
    <!-- ============================== -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-md z-10">
          <!-- Modal Header -->
          <div class="px-6 pt-6 pb-4 border-b border-slate-200">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-bold text-slate-900">
                  {{ isEditing ? 'Edit Produk' : 'Tambah Produk Baru' }}
                </h3>
              </div>
              <button @click="closeModal" class="p-2 text-slate-400 hover:text-slate-600 hover:bg-slate-100 rounded-lg transition-colors">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>

          <div class="px-6 py-5 space-y-4 max-h-[75vh] overflow-y-auto">
            <form @submit.prevent="submitForm">
              <div class="space-y-4">
                <!-- Upload Foto Produk -->
                <div>
                  <label class="block text-sm font-medium text-slate-700 mb-1.5">Foto Produk</label>

                  <!-- Preview gambar (jika ada file baru atau foto lama) -->
                  <div
                    v-if="previewUrl || form.image_url"
                    class="relative group w-full h-40 rounded-xl overflow-hidden border border-slate-200 mb-2 cursor-pointer"
                    @click="triggerFileInput"
                  >
                    <img
                      :src="previewUrl || `${apiBase}${form.image_url}`"
                      class="w-full h-full object-cover"
                      alt="Preview foto produk"
                    />
                    <!-- Overlay hover -->
                    <div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity flex flex-col items-center justify-center gap-2">
                      <Camera class="w-7 h-7 text-white" />
                      <p class="text-white text-sm font-medium">Klik untuk ganti foto</p>
                    </div>
                    <!-- Tombol hapus foto -->
                    <button
                      type="button"
                      @click.stop="clearImage"
                      class="absolute top-2 right-2 p-1.5 bg-red-600 hover:bg-red-700 text-white rounded-full shadow-lg opacity-0 group-hover:opacity-100 transition-all"
                      title="Hapus foto"
                    >
                      <X class="w-3.5 h-3.5" />
                    </button>
                  </div>

                  <!-- Drag & Drop Zone (jika belum ada foto) -->
                  <div
                    v-else
                    @dragover.prevent="isDragging = true"
                    @dragleave.prevent="isDragging = false"
                    @drop.prevent="onDrop"
                    @click="triggerFileInput"
                    class="border-2 border-dashed rounded-xl p-6 text-center cursor-pointer transition-all duration-200 select-none"
                    :class="isDragging
                      ? 'border-indigo-400 bg-indigo-50 scale-[1.01]'
                      : 'border-slate-300 hover:border-indigo-400 hover:bg-slate-50'"
                  >
                    <div class="flex flex-col items-center gap-2">
                      <div
                        class="w-12 h-12 rounded-xl flex items-center justify-center transition-colors"
                        :class="isDragging ? 'bg-indigo-100' : 'bg-slate-100'"
                      >
                        <ImagePlus class="w-6 h-6" :class="isDragging ? 'text-indigo-500' : 'text-slate-400'" />
                      </div>
                      <div>
                        <p class="text-sm text-slate-600">
                          Drag & drop foto, atau
                          <span class="text-indigo-600 font-medium">klik untuk pilih</span>
                        </p>
                        <p class="text-xs text-slate-400 mt-0.5">JPG, PNG, WebP — Maksimal 2MB</p>
                      </div>
                    </div>
                  </div>

                  <!-- Hidden file input -->
                  <input
                    ref="fileInputRef"
                    type="file"
                    accept="image/jpeg,image/png,image/webp"
                    class="hidden"
                    @change="onFileInputChange"
                  />

                  <!-- Indikator file terpilih -->
                  <p v-if="selectedFile" class="mt-1.5 text-xs text-indigo-600 flex items-center gap-1">
                    <Check class="w-3.5 h-3.5" />
                    {{ selectedFile.name }} ({{ (selectedFile.size / 1024).toFixed(0) }} KB) — akan diupload saat simpan
                  </p>
                </div>

                <!-- Divider -->
                <div class="border-t border-slate-100 pt-1"></div>

                <div>
                  <label class="block text-sm font-medium text-slate-700 mb-1">SKU <span class="text-red-500">*</span></label>
                  <input
                    type="text"
                    v-model="form.sku"
                    required
                    placeholder="Contoh: PRD-001"
                    class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 font-mono"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700 mb-1">Nama Produk <span class="text-red-500">*</span></label>
                  <input
                    type="text"
                    v-model="form.name"
                    required
                    placeholder="Contoh: Semen Tiga Roda 50kg"
                    class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700 mb-1">Kategori</label>
                  <input
                    type="text"
                    v-model="form.category"
                    placeholder="Contoh: Material Bangunan"
                    class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
                  />
                </div>
                <div>
                  <label class="block text-sm font-medium text-slate-700 mb-1">Harga Jual Standar (Rp) <span class="text-red-500">*</span></label>
                  <input
                    type="number"
                    v-model="form.standard_price"
                    required
                    min="0"
                    class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
                  />
                </div>

                <!-- Info Box about HPP and Stock -->
                <div class="bg-blue-50 border border-blue-200 rounded-xl p-4 text-sm mt-4">
                  <div class="flex gap-2 items-start">
                    <svg class="w-5 h-5 text-blue-500 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <div class="text-blue-800 space-y-1">
                      <p class="font-medium">HPP dan Stok</p>
                      <p class="text-blue-700 text-xs">
                        HPP (Harga Pokok Pembelian) rata-rata dan Stok Saat Ini dikelola secara otomatis oleh sistem saat transaksi barang masuk (Inbound).
                      </p>
                      <div v-if="isEditing" class="mt-2 text-xs">
                        <span class="font-semibold text-blue-900">Saat ini: </span>
                        Stok {{ form.current_stock }} unit — HPP Rata-rata Rp {{ formatNumber(form.average_hpp) }}
                      </div>
                    </div>
                  </div>
                </div>

                <p v-if="formError" class="text-red-600 text-sm bg-red-50 border border-red-200 rounded-lg p-3">{{ formError }}</p>
              </div>

              <div class="mt-6 flex justify-end gap-3 pt-4 border-t border-slate-200">
                <button
                  type="button"
                  @click="closeModal"
                  class="px-4 py-2 text-sm font-medium text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors"
                >
                  Batal
                </button>
                <button
                  type="submit"
                  :disabled="submitLoading"
                  class="px-5 py-2 text-sm font-semibold bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition-colors flex items-center gap-2"
                >
                  <svg v-if="submitLoading" class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  {{ submitLoading ? 'Menyimpan...' : 'Simpan Produk' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL: HAPUS PRODUK (Informatif) -->
    <!-- ============================== -->
    <div v-if="isDeleteModalOpen && deletingProduct" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isDeleteModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-md z-10">
          <!-- Modal Header (Danger) -->
          <div class="px-6 pt-6 pb-4 border-b border-red-100 bg-red-50 rounded-t-2xl">
            <div class="flex items-start gap-3">
              <div class="flex-shrink-0 w-10 h-10 bg-red-100 border border-red-200 rounded-xl flex items-center justify-center">
                <svg class="w-5 h-5 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                </svg>
              </div>
              <div>
                <h3 class="text-lg font-bold text-red-800">Hapus Produk</h3>
                <p class="text-sm text-red-600 mt-0.5">Tindakan ini tidak dapat dibatalkan.</p>
              </div>
            </div>
          </div>

          <div class="px-6 py-5 space-y-4">
            <!-- Info Produk -->
            <div class="bg-slate-50 rounded-xl p-4 border border-slate-200">
              <p class="text-xs font-semibold text-slate-500 uppercase tracking-wide mb-2">Detail Produk yang Akan Dihapus</p>
              <div class="grid grid-cols-2 gap-2 text-sm">
                <div class="col-span-2">
                  <span class="text-slate-500">Nama Produk</span>
                  <p class="font-medium text-slate-800">{{ deletingProduct.name }}</p>
                </div>
                <div>
                  <span class="text-slate-500">SKU</span>
                  <p class="font-semibold text-indigo-600">{{ deletingProduct.sku }}</p>
                </div>
                <div>
                  <span class="text-slate-500">Stok Saat Ini</span>
                  <p class="font-bold" :class="deletingProduct.current_stock > 0 ? 'text-amber-600' : 'text-slate-800'">
                    {{ deletingProduct.current_stock }} unit
                  </p>
                </div>
              </div>
            </div>

            <!-- Dampak Penghapusan -->
            <div class="bg-amber-50 border border-amber-200 rounded-xl p-4">
              <p class="text-xs font-semibold text-amber-700 uppercase tracking-wide mb-2">⚠ Dampak Penghapusan</p>
              <ul class="text-sm text-amber-800 space-y-1.5">
                <li class="flex items-start gap-2">
                  <svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                  </svg>
                  <span>Produk akan dihapus dari master data secara permanen.</span>
                </li>
                <li class="flex items-start gap-2">
                  <svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                  </svg>
                  <span>Produk ini tidak akan bisa lagi dipilih saat membuat faktur penjualan baru.</span>
                </li>
                <li v-if="deletingProduct.current_stock > 0" class="flex items-start gap-2 font-medium text-red-700">
                  <svg class="w-4 h-4 text-red-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
                  </svg>
                  <span>Masih ada stok {{ deletingProduct.current_stock }} unit! Pastikan Anda sudah menyesuaikan stok atau menghabiskannya.</span>
                </li>
              </ul>
            </div>

            <p v-if="deleteError" class="text-red-600 text-sm bg-red-50 border border-red-200 rounded-lg p-3">{{ deleteError }}</p>
          </div>

          <div class="px-6 py-4 border-t border-slate-200 flex justify-end gap-3">
            <button @click="isDeleteModalOpen = false" class="px-4 py-2 text-sm font-medium text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">
              Batal
            </button>
            <button
              @click="submitDelete"
              :disabled="deleteLoading"
              class="px-5 py-2 text-sm font-semibold bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50 transition-colors"
            >
              {{ deleteLoading ? 'Menghapus...' : 'Ya, Hapus Produk' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import api from '@/plugins/axios'
import { toast } from 'vue-sonner'
import { ImagePlus, ImageOff, Camera, X, Check } from '@lucide/vue'

// ─── Config ───────────────────────────────────────────────────────────────────
const apiBase = 'http://localhost:8080'

// ─── Types ────────────────────────────────────────────────────────────────────
interface Product {
  ID?: number
  sku: string
  name: string
  category: string
  standard_price: number
  average_hpp: number
  current_stock: number
  image_url?: string
}

// ─── State ────────────────────────────────────────────────────────────────────
const products = ref<Product[]>([])
const loading = ref(false)

// Filter
const searchQuery = ref('')
const filterCategory = ref('')
const filterStock = ref('')

// Pagination
const currentPage = ref(1)
const pageSize = 10

// Create/Edit Modal
const isModalOpen = ref(false)
const isEditing = ref(false)
const submitLoading = ref(false)
const formError = ref('')
const form = ref<Product>({
  sku: '',
  name: '',
  category: '',
  standard_price: 0,
  average_hpp: 0,
  current_stock: 0,
  image_url: ''
})

// Upload foto
const selectedFile = ref<File | null>(null)
const previewUrl = ref<string>('')
const isDragging = ref(false)
const fileInputRef = ref<HTMLInputElement | null>(null)

// Delete Modal
const isDeleteModalOpen = ref(false)
const deleteLoading = ref(false)
const deletingProduct = ref<Product | null>(null)
const deleteError = ref('')

// ─── Computed ────────────────────────────────────────────────────────────────
const uniqueCategories = computed(() => {
  const cats = new Set(products.value.map(p => p.category).filter(Boolean))
  return Array.from(cats).sort()
})

const filteredProducts = computed(() => {
  let list = [...products.value]
  const q = searchQuery.value.toLowerCase()

  if (q) {
    list = list.filter(p =>
      p.name.toLowerCase().includes(q) ||
      p.sku.toLowerCase().includes(q)
    )
  }
  if (filterCategory.value) {
    list = list.filter(p => p.category === filterCategory.value)
  }
  if (filterStock.value) {
    if (filterStock.value === 'low') {
      list = list.filter(p => p.current_stock < 10 && p.current_stock > 0)
    } else if (filterStock.value === 'empty') {
      list = list.filter(p => p.current_stock === 0)
    } else if (filterStock.value === 'available') {
      list = list.filter(p => p.current_stock > 0)
    }
  }
  return list
})

const totalPages = computed(() => Math.max(1, Math.ceil(filteredProducts.value.length / pageSize)))

const paginatedProducts = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return filteredProducts.value.slice(start, start + pageSize)
})

const paginationStart = computed(() => {
  if (filteredProducts.value.length === 0) return 0
  return (currentPage.value - 1) * pageSize + 1
})

const paginationEnd = computed(() =>
  Math.min(currentPage.value * pageSize, filteredProducts.value.length)
)

// Summary cards
const summary = computed(() => {
  const list = filteredProducts.value
  const totalProducts = list.length
  const totalStockValue = list.reduce((s, p) => s + (p.current_stock * p.average_hpp), 0)
  const totalSalesValue = list.reduce((s, p) => s + (p.current_stock * p.standard_price), 0)
  const lowStockCount = list.filter(p => p.current_stock < 10).length

  return { totalProducts, totalStockValue, totalSalesValue, lowStockCount }
})

// ─── Helpers ─────────────────────────────────────────────────────────────────
const formatNumber = (num: number) => Math.round(num).toLocaleString('id-ID')

watch([searchQuery, filterCategory, filterStock], () => {
  currentPage.value = 1
})

const resetFilters = () => {
  searchQuery.value = ''
  filterCategory.value = ''
  filterStock.value = ''
  currentPage.value = 1
}

// ─── Upload Foto Helpers ──────────────────────────────────────────────────────
const triggerFileInput = () => {
  fileInputRef.value?.click()
}

const onFileSelected = (file: File) => {
  if (file.size > 2 * 1024 * 1024) {
    formError.value = 'Ukuran file melebihi 2MB. Pilih gambar yang lebih kecil.'
    return
  }
  const allowed = ['image/jpeg', 'image/png', 'image/webp']
  if (!allowed.includes(file.type)) {
    formError.value = 'Format file tidak didukung. Gunakan JPG, PNG, atau WebP.'
    return
  }
  // Revoke URL sebelumnya untuk cegah memory leak
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
  }
  selectedFile.value = file
  previewUrl.value = URL.createObjectURL(file)
  formError.value = ''
}

const onFileInputChange = (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (file) onFileSelected(file)
}

const onDrop = (e: DragEvent) => {
  isDragging.value = false
  const file = e.dataTransfer?.files[0]
  if (file) onFileSelected(file)
}

const clearImage = () => {
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
  }
  selectedFile.value = null
  previewUrl.value = ''
  form.value.image_url = ''
  if (fileInputRef.value) fileInputRef.value.value = ''
}

// ─── Fetch Data ───────────────────────────────────────────────────────────────
const fetchProducts = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/products')
    products.value = res.data ?? []
  } catch (error) {
    console.error('Failed to fetch products', error)
  } finally {
    loading.value = false
  }
}

// ─── Create & Edit ────────────────────────────────────────────────────────────
const openCreateModal = () => {
  isEditing.value = false
  formError.value = ''
  selectedFile.value = null
  previewUrl.value = ''
  form.value = {
    sku: '',
    name: '',
    category: '',
    standard_price: 0,
    average_hpp: 0,
    current_stock: 0,
    image_url: ''
  }
  isModalOpen.value = true
}

const openEditModal = (product: Product) => {
  isEditing.value = true
  formError.value = ''
  selectedFile.value = null
  previewUrl.value = ''
  form.value = { ...product }
  isModalOpen.value = true
}

const closeModal = () => {
  // Revoke blob URL untuk cegah memory leak
  if (previewUrl.value) {
    URL.revokeObjectURL(previewUrl.value)
    previewUrl.value = ''
  }
  selectedFile.value = null
  if (fileInputRef.value) fileInputRef.value.value = ''
  isModalOpen.value = false
}

const submitForm = async () => {
  submitLoading.value = true
  formError.value = ''
  try {
    let productId: number

    const wasEditing = isEditing.value
    if (isEditing.value && form.value.ID) {
      await api.put(`/api/products/${form.value.ID}`, {
        sku: form.value.sku,
        name: form.value.name,
        category: form.value.category,
        standard_price: Number(form.value.standard_price) || 0
      }, { skipToast: true } as any)
      productId = form.value.ID
    } else {
      const res = await api.post('/api/products', {
        sku: form.value.sku,
        name: form.value.name,
        category: form.value.category,
        standard_price: Number(form.value.standard_price) || 0
      }, { skipToast: true } as any)
      productId = res.data.ID
    }

    // Upload foto jika ada file yang dipilih
    if (selectedFile.value) {
      const formData = new FormData()
      formData.append('image', selectedFile.value)
      await api.post(`/api/products/${productId}/image`, formData, {
        headers: { 'Content-Type': 'multipart/form-data' },
        skipToast: true
      } as any)
    }

    // Tutup modal terlebih dahulu
    isModalOpen.value = false
    // Bersihkan state upload
    if (previewUrl.value) URL.revokeObjectURL(previewUrl.value)
    previewUrl.value = ''
    selectedFile.value = null

    // Tampilkan toast setelah modal tertutup agar tidak terhalangi backdrop
    toast.success(wasEditing ? 'Produk berhasil diperbarui' : 'Produk berhasil ditambahkan')

    await fetchProducts()
  } catch (error: any) {
    console.error('Failed to save product', error)
    const errMsg = error.response?.data?.error ?? 'Gagal menyimpan produk. Pastikan SKU unik.'
    formError.value = errMsg
    toast.error(errMsg)
  } finally {
    submitLoading.value = false
  }
}

// ─── Delete ───────────────────────────────────────────────────────────────────
const openDeleteModal = (product: Product) => {
  deletingProduct.value = product
  deleteError.value = ''
  isDeleteModalOpen.value = true
}

const submitDelete = async () => {
  if (!deletingProduct.value?.ID) return
  deleteLoading.value = true
  deleteError.value = ''
  try {
    await api.delete(`/api/products/${deletingProduct.value.ID}`, { skipToast: true } as any)
    // Tutup modal terlebih dahulu
    isDeleteModalOpen.value = false
    // Tampilkan toast setelah modal tertutup agar tidak terhalangi backdrop
    toast.success('Produk berhasil dihapus')
    await fetchProducts()
  } catch (error: any) {
    console.error('Failed to delete product', error)
    const errMsg = error.response?.data?.error ?? 'Gagal menghapus produk.'
    deleteError.value = errMsg
    toast.error(errMsg)
  } finally {
    deleteLoading.value = false
  }
}

// ─── Init ─────────────────────────────────────────────────────────────────────
onMounted(() => {
  fetchProducts()
})
</script>
