<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Manajemen Stok & Opname</h2>
        <p class="text-slate-500 text-sm mt-1">Pusat kendali gudang: Inbound, Stock Opname, dan pemantauan pergerakan barang.</p>
      </div>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-5">
      <!-- Total Nilai Aset -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4 relative overflow-hidden">
        <div class="absolute -right-4 -bottom-4 opacity-10">
          <svg class="w-24 h-24 text-indigo-600" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/></svg>
        </div>
        <div class="flex-shrink-0 w-12 h-12 bg-indigo-50 rounded-xl flex items-center justify-center border border-indigo-100 relative z-10">
          <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        </div>
        <div class="min-w-0 flex-1 relative z-10">
          <p class="text-xs font-medium text-slate-500 uppercase tracking-wide">Total Nilai Stok (HPP)</p>
          <p class="text-xl font-bold text-slate-900 mt-1 truncate">Rp {{ formatNumber(summary.totalStockValue) }}</p>
        </div>
      </div>

      <!-- Peringatan Stok -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4 relative overflow-hidden">
        <div class="absolute -right-4 -bottom-4 opacity-10">
          <svg class="w-24 h-24 text-amber-600" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L1 21h22M12 6l7.5 13h-15M11 10h2v4h-2M11 16h2v2h-2"/></svg>
        </div>
        <div class="flex-shrink-0 w-12 h-12 bg-amber-50 rounded-xl flex items-center justify-center border border-amber-100 relative z-10">
          <svg class="w-6 h-6 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
        </div>
        <div class="min-w-0 flex-1 relative z-10">
          <p class="text-xs font-medium text-slate-500 uppercase tracking-wide">Stok Menipis &lt; 10</p>
          <p class="text-xl font-bold text-slate-900 mt-1 truncate">{{ summary.lowStockCount }} <span class="text-sm font-normal text-slate-500">produk</span></p>
        </div>
      </div>

      <!-- Stok Habis -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4 relative overflow-hidden">
        <div class="absolute -right-4 -bottom-4 opacity-10">
          <svg class="w-24 h-24 text-red-600" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm5 11H7v-2h10v2z"/></svg>
        </div>
        <div class="flex-shrink-0 w-12 h-12 bg-red-50 rounded-xl flex items-center justify-center border border-red-100 relative z-10">
          <svg class="w-6 h-6 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636" /></svg>
        </div>
        <div class="min-w-0 flex-1 relative z-10">
          <p class="text-xs font-medium text-slate-500 uppercase tracking-wide">Stok Kosong / Habis</p>
          <p class="text-xl font-bold text-red-600 mt-1 truncate">{{ summary.emptyStockCount }} <span class="text-sm font-normal text-red-400">produk</span></p>
        </div>
      </div>
    </div>

    <!-- Filter Bar -->
    <div class="bg-white rounded-xl border border-slate-200 shadow-sm p-4 flex flex-col sm:flex-row gap-3 items-center">
      <div class="flex-1 w-full min-w-48 relative">
        <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
        <input v-model="searchQuery" type="text" placeholder="Cari SKU atau nama produk..." class="pl-9 pr-3 py-2 w-full border border-slate-300 rounded-lg text-sm focus:ring-2 focus:ring-indigo-500 focus:border-transparent outline-none transition-all" />
      </div>
      <div class="flex gap-2 w-full sm:w-auto overflow-x-auto">
        <select v-model="filterStock" class="px-3 py-2 border border-slate-300 rounded-lg text-sm outline-none focus:ring-2 focus:ring-indigo-500 bg-white min-w-32">
          <option value="">Semua Kondisi Stok</option>
          <option value="low">Menipis (&lt; 10)</option>
          <option value="empty">Kosong (= 0)</option>
          <option value="available">Tersedia (&gt; 0)</option>
        </select>
      </div>
    </div>

    <!-- Main Table -->
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-slate-200">
          <thead class="bg-slate-50">
            <tr>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Info Produk</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">HPP Rata-rata</th>
              <th class="px-4 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Stok Sistem</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider min-w-[200px]">Aksi &amp; Mutasi</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="4" class="px-4 py-12 text-center text-slate-400">
                <svg class="animate-spin w-8 h-8 mx-auto mb-3" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                </svg>
                Memuat data inventory...
              </td>
            </tr>
            <tr v-else-if="paginatedProducts.length === 0">
              <td colspan="4" class="px-4 py-12 text-center text-slate-500 text-sm">Tidak ada data produk yang sesuai filter.</td>
            </tr>
            <tr v-else v-for="product in paginatedProducts" :key="product.ID" class="hover:bg-slate-50/70 transition-colors group">
              <td class="px-4 py-3">
                <div class="flex flex-col">
                  <span class="text-sm font-bold text-slate-900">{{ product.name }}</span>
                  <span class="text-xs font-mono text-slate-500">{{ product.sku }}</span>
                </div>
              </td>
              <td class="px-4 py-3 text-right text-sm text-slate-600">
                Rp {{ formatNumber(product.average_hpp) }}
              </td>
              <td class="px-4 py-3 text-center">
                <span class="inline-flex items-center justify-center min-w-10 px-2.5 py-0.5 rounded-full text-sm font-bold"
                  :class="{
                    'bg-red-100 text-red-700': product.current_stock === 0,
                    'bg-amber-100 text-amber-700': product.current_stock > 0 && product.current_stock < 10,
                    'bg-emerald-100 text-emerald-700': product.current_stock >= 10
                  }">
                  {{ product.current_stock }}
                </span>
              </td>
              <td class="px-4 py-3 text-right whitespace-nowrap">
                <div class="flex items-center justify-end gap-1.5 opacity-90 group-hover:opacity-100 transition-opacity">
                  <!-- Inbound Button -->
                  <button @click="openInboundModal(product)" class="inline-flex items-center gap-1 px-2.5 py-1.5 bg-indigo-50 text-indigo-700 hover:bg-indigo-100 rounded-md text-xs font-medium transition-colors border border-indigo-200" title="Terima Barang Masuk (Inbound)">
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"/></svg>
                    Inbound
                  </button>
                  <!-- Opname Button -->
                  <button @click="openOpnameModal(product)" class="inline-flex items-center gap-1 px-2.5 py-1.5 bg-amber-50 text-amber-700 hover:bg-amber-100 rounded-md text-xs font-medium transition-colors border border-amber-200" title="Sesuaikan Fisik (Stock Opname)">
                    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3"/></svg>
                    Opname
                  </button>
                  <!-- History Button -->
                  <button @click="openHistoryModal(product)" class="p-1.5 text-slate-500 hover:bg-slate-100 hover:text-slate-800 rounded-md transition-colors border border-transparent hover:border-slate-300" title="Lihat Kartu Stok (Riwayat)">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      <!-- Pagination -->
      <div class="px-4 py-3 border-t border-slate-200 flex flex-wrap gap-3 items-center justify-between bg-slate-50/50">
        <p class="text-xs text-slate-500">Menampilkan {{ paginationStart }}-{{ paginationEnd }} dari {{ filteredProducts.length }} data</p>
        <div class="flex gap-1">
          <button @click="currentPage--" :disabled="currentPage <= 1" class="px-3 py-1 text-sm border border-slate-300 rounded-md disabled:opacity-40 hover:bg-white transition-colors">‹</button>
          <button v-for="page in totalPages" :key="page" @click="currentPage = page" class="px-3 py-1 text-sm border rounded-md transition-colors" :class="currentPage === page ? 'bg-indigo-600 text-white border-indigo-600' : 'border-slate-300 hover:bg-white'">{{ page }}</button>
          <button @click="currentPage++" :disabled="currentPage >= totalPages" class="px-3 py-1 text-sm border border-slate-300 rounded-md disabled:opacity-40 hover:bg-white transition-colors">›</button>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL 1: INBOUND               -->
    <!-- ============================== -->
    <div v-if="activeModal === 'inbound' && selectedProduct" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md z-10">
          <div class="px-6 pt-6 pb-4 border-b border-slate-100 flex justify-between items-center bg-indigo-50/50 rounded-t-2xl">
            <h3 class="text-lg font-bold text-indigo-900 flex items-center gap-2">
              <svg class="w-5 h-5 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"/></svg>
              Barang Masuk (Inbound)
            </h3>
            <button @click="closeModal" class="text-slate-400 hover:text-slate-600"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg></button>
          </div>
          <div class="px-6 py-5">
            <form @submit.prevent="submitInbound" class="space-y-4">
              <!-- Info Produk Pendek -->
              <div class="p-3 bg-slate-50 border border-slate-200 rounded-lg flex justify-between items-center">
                <div>
                  <p class="text-xs font-mono text-slate-500">{{ selectedProduct.sku }}</p>
                  <p class="text-sm font-bold text-slate-800">{{ selectedProduct.name }}</p>
                </div>
                <div class="text-right">
                  <p class="text-xs text-slate-500">Stok Saat Ini</p>
                  <p class="text-sm font-bold text-slate-800">{{ selectedProduct.current_stock }}</p>
                </div>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-xs font-medium text-slate-700 mb-1">Kuantitas Masuk <span class="text-red-500">*</span></label>
                  <input type="number" v-model.number="inboundForm.quantity" required min="1" class="w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:ring-2 focus:ring-indigo-500 outline-none" />
                </div>
                <div>
                  <label class="block text-xs font-medium text-slate-700 mb-1">Harga Beli/Unit <span class="text-red-500">*</span></label>
                  <input type="number" v-model.number="inboundForm.unit_price" required min="1" class="w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:ring-2 focus:ring-indigo-500 outline-none" />
                </div>
              </div>
              
              <div>
                <label class="block text-xs font-medium text-slate-700 mb-1">Referensi / Nomor PO</label>
                <input type="text" v-model="inboundForm.reference" placeholder="Contoh: INV-SUP-202301" class="w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:ring-2 focus:ring-indigo-500 outline-none" />
              </div>

              <div class="mt-4 pt-4 border-t border-slate-100 flex justify-between items-center">
                <span class="text-sm font-medium text-slate-500">Total Pembelian:</span>
                <span class="text-lg font-bold text-indigo-700">Rp {{ formatNumber(inboundTotal) }}</span>
              </div>
              <p v-if="actionError" class="text-xs text-red-600 bg-red-50 p-2 rounded">{{ actionError }}</p>

              <div class="pt-2 flex justify-end gap-2">
                <button type="button" @click="closeModal" class="px-4 py-2 text-sm font-medium text-slate-600 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">Batal</button>
                <button type="submit" :disabled="actionLoading" class="px-4 py-2 text-sm font-bold text-white bg-indigo-600 rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition-colors">
                  {{ actionLoading ? 'Menyimpan...' : 'Simpan Inbound' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL 2: OPNAME (ADJUSTMENT)   -->
    <!-- ============================== -->
    <div v-if="activeModal === 'opname' && selectedProduct" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-md z-10">
          <div class="px-6 pt-6 pb-4 border-b border-slate-100 flex justify-between items-center bg-amber-50/50 rounded-t-2xl">
            <h3 class="text-lg font-bold text-amber-900 flex items-center gap-2">
              <svg class="w-5 h-5 text-amber-600" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3"/></svg>
              Stock Opname
            </h3>
            <button @click="closeModal" class="text-slate-400 hover:text-slate-600"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg></button>
          </div>
          <div class="px-6 py-5">
            <form @submit.prevent="submitOpname" class="space-y-4">
              <div class="p-3 bg-slate-50 border border-slate-200 rounded-lg grid grid-cols-2 gap-4">
                <div class="col-span-2">
                  <p class="text-xs font-mono text-slate-500">{{ selectedProduct.sku }}</p>
                  <p class="text-sm font-bold text-slate-800">{{ selectedProduct.name }}</p>
                </div>
                <div>
                  <p class="text-xs text-slate-500">Stok Sistem</p>
                  <p class="text-xl font-bold text-slate-800">{{ selectedProduct.current_stock }}</p>
                </div>
                <div>
                  <p class="text-xs text-slate-500">Selisih Opname</p>
                  <p class="text-xl font-bold" :class="opnameDiff === 0 ? 'text-slate-400' : (opnameDiff > 0 ? 'text-emerald-600' : 'text-red-600')">
                    {{ opnameDiff > 0 ? '+' : '' }}{{ opnameDiff }}
                  </p>
                </div>
              </div>

              <div>
                <label class="block text-xs font-medium text-slate-700 mb-1">Stok Fisik Aktual <span class="text-red-500">*</span></label>
                <div class="relative">
                  <input type="number" v-model.number="opnameForm.physical_stock" required min="0" class="w-full border border-slate-300 rounded-lg py-2.5 px-3 text-lg font-bold focus:ring-2 focus:ring-amber-500 outline-none bg-amber-50/30" />
                  <span class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 text-sm font-medium">unit</span>
                </div>
                <p class="text-xs text-slate-500 mt-1">Masukkan jumlah barang fisik yang ada di gudang saat ini.</p>
              </div>
              
              <div>
                <label class="block text-xs font-medium text-slate-700 mb-1">Referensi / Alasan Penyesuaian <span class="text-red-500">*</span></label>
                <input type="text" v-model="opnameForm.reference" required class="w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:ring-2 focus:ring-amber-500 outline-none" />
              </div>

              <p v-if="actionError" class="text-xs text-red-600 bg-red-50 p-2 rounded">{{ actionError }}</p>

              <div class="pt-4 flex justify-end gap-2 border-t border-slate-100">
                <button type="button" @click="closeModal" class="px-4 py-2 text-sm font-medium text-slate-600 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">Batal</button>
                <button type="submit" :disabled="actionLoading || opnameDiff === 0" class="px-4 py-2 text-sm font-bold text-white bg-amber-600 rounded-lg hover:bg-amber-700 disabled:opacity-50 transition-colors">
                  {{ actionLoading ? 'Memproses...' : 'Simpan Opname' }}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL 3: KARTU STOK (HISTORY)  -->
    <!-- ============================== -->
    <div v-if="activeModal === 'history' && selectedProduct" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-2 sm:px-4 py-10">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="closeModal"></div>
        <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-2xl z-10 flex flex-col max-h-[85vh]">
          <!-- Header -->
          <div class="px-6 py-4 border-b border-slate-100 flex justify-between items-center shrink-0">
            <div>
              <h3 class="text-lg font-bold text-slate-800">Kartu Stok (Riwayat)</h3>
              <p class="text-sm font-mono text-slate-500 mt-0.5">{{ selectedProduct.sku }} — {{ selectedProduct.name }}</p>
            </div>
            <button @click="closeModal" class="p-2 text-slate-400 hover:text-slate-600 hover:bg-slate-100 rounded-lg transition-colors shrink-0">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" /></svg>
            </button>
          </div>
          <!-- Body -->
          <div class="p-0 overflow-y-auto bg-slate-50 relative flex-1">
            <div v-if="historyLoading" class="p-12 text-center text-slate-400">
              <svg class="animate-spin w-8 h-8 mx-auto mb-3" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              Memuat riwayat transaksi...
            </div>
            <div v-else-if="stockHistory.length === 0" class="p-12 text-center text-slate-500 text-sm">
              Belum ada riwayat transaksi untuk produk ini.
            </div>
            <div v-else class="p-4 sm:p-6 relative">
              <!-- Timeline vertical line -->
              <div class="absolute left-6 sm:left-8 top-6 bottom-6 w-0.5 bg-slate-200 z-0"></div>
              
              <div class="space-y-6 relative z-10">
                <div v-for="item in reversedHistory" :key="item.ID" class="flex gap-4">
                  <!-- Icon indicator -->
                  <div class="flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center shadow-sm border-2 border-white ring-1 ring-slate-100 mt-1"
                    :class="{
                      'bg-indigo-100 text-indigo-600': item.type === 'INBOUND',
                      'bg-amber-100 text-amber-600': item.type === 'ADJUSTMENT',
                      'bg-emerald-100 text-emerald-600': item.type === 'SALE'
                    }">
                    <svg v-if="item.type === 'INBOUND'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"/></svg>
                    <svg v-else-if="item.type === 'ADJUSTMENT'" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6l3 1m0 0l-3 9a5.002 5.002 0 006.001 0M6 7l3 9M6 7l6-2m6 2l3-1m-3 1l-3 9a5.002 5.002 0 006.001 0M18 7l3 9m-3-9l-6-2m0-2v2m0 16V5m0 16H9m3 0h3"/></svg>
                    <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"/></svg>
                  </div>
                  <!-- Content card -->
                  <div class="flex-1 bg-white p-4 rounded-xl shadow-sm border border-slate-200">
                    <div class="flex justify-between items-start mb-2">
                      <div>
                        <span class="inline-flex px-2 py-0.5 rounded text-[10px] font-bold uppercase tracking-wider mb-1"
                          :class="{
                            'bg-indigo-50 text-indigo-700': item.type === 'INBOUND',
                            'bg-amber-50 text-amber-700': item.type === 'ADJUSTMENT',
                            'bg-emerald-50 text-emerald-700': item.type === 'SALE'
                          }">{{ item.type }}</span>
                        <p class="text-sm font-bold text-slate-800">{{ item.reference || 'Tanpa Referensi' }}</p>
                      </div>
                      <span class="text-xs text-slate-500 whitespace-nowrap">{{ formatDate(item.CreatedAt) }}</span>
                    </div>
                    <div class="flex items-center gap-4 mt-3 pt-3 border-t border-slate-100">
                      <div>
                        <p class="text-[10px] text-slate-500 uppercase">Mutasi (Qty)</p>
                        <p class="text-sm font-bold" :class="item.quantity > 0 ? 'text-emerald-600' : 'text-red-600'">
                          {{ item.quantity > 0 ? '+' : '' }}{{ item.quantity }}
                        </p>
                      </div>
                      <div v-if="item.type === 'INBOUND'">
                        <p class="text-[10px] text-slate-500 uppercase">Harga Beli</p>
                        <p class="text-sm font-medium text-slate-700">Rp {{ formatNumber(item.unit_price) }}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import api from '@/plugins/axios'

// ─── Interfaces ──────────────────────────────────────────────────────────────
interface Product {
  ID: number
  sku: string
  name: string
  category: string
  standard_price: number
  average_hpp: number
  current_stock: number
}

interface InventoryTransaction {
  ID: number
  CreatedAt: string
  product_id: number
  type: string
  quantity: number
  unit_price: number
  reference: string
}

// ─── State ───────────────────────────────────────────────────────────────────
const products = ref<Product[]>([])
const loading = ref(false)

const searchQuery = ref('')
const filterStock = ref('')
const currentPage = ref(1)
const pageSize = 10

// Modals
const activeModal = ref<'inbound' | 'opname' | 'history' | null>(null)
const selectedProduct = ref<Product | null>(null)
const actionLoading = ref(false)
const actionError = ref('')

// Inbound Form
const inboundForm = ref({ quantity: 1, unit_price: 0, reference: '' })
const inboundTotal = computed(() => inboundForm.value.quantity * inboundForm.value.unit_price)

// Opname Form
const opnameForm = ref({ physical_stock: 0, reference: 'Stock Opname' })
const opnameDiff = computed(() => {
  if (!selectedProduct.value) return 0
  return (Number(opnameForm.value.physical_stock) || 0) - selectedProduct.value.current_stock
})

// History State
const stockHistory = ref<InventoryTransaction[]>([])
const historyLoading = ref(false)
const reversedHistory = computed(() => [...stockHistory.value].reverse()) // Terbaru di atas

// ─── Computed (Filtering & Pagination) ───────────────────────────────────────
const filteredProducts = computed(() => {
  let list = [...products.value]
  const q = searchQuery.value.toLowerCase()

  if (q) {
    list = list.filter(p => p.name.toLowerCase().includes(q) || p.sku.toLowerCase().includes(q))
  }
  if (filterStock.value) {
    if (filterStock.value === 'low') list = list.filter(p => p.current_stock < 10 && p.current_stock > 0)
    else if (filterStock.value === 'empty') list = list.filter(p => p.current_stock === 0)
    else if (filterStock.value === 'available') list = list.filter(p => p.current_stock > 0)
  }
  return list
})

const totalPages = computed(() => Math.max(1, Math.ceil(filteredProducts.value.length / pageSize)))
const paginatedProducts = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return filteredProducts.value.slice(start, start + pageSize)
})
const paginationStart = computed(() => filteredProducts.value.length === 0 ? 0 : (currentPage.value - 1) * pageSize + 1)
const paginationEnd = computed(() => Math.min(currentPage.value * pageSize, filteredProducts.value.length))

const summary = computed(() => {
  const list = products.value
  return {
    totalStockValue: list.reduce((sum, p) => sum + (p.current_stock * p.average_hpp), 0),
    lowStockCount: list.filter(p => p.current_stock > 0 && p.current_stock < 10).length,
    emptyStockCount: list.filter(p => p.current_stock === 0).length
  }
})

// ─── Formatting ──────────────────────────────────────────────────────────────
const formatNumber = (num: number) => Math.round(num).toLocaleString('id-ID')
const formatDate = (dateString: string) => {
  const d = new Date(dateString)
  return new Intl.DateTimeFormat('id-ID', { day: 'numeric', month: 'short', year: 'numeric', hour: '2-digit', minute: '2-digit' }).format(d)
}

// ─── Actions ─────────────────────────────────────────────────────────────────
const fetchProducts = async () => {
  loading.value = true
  try {
    const res = await api.get('/api/products')
    products.value = res.data ?? []
  } catch (error) {
    console.error('Failed fetching products', error)
  } finally {
    loading.value = false
  }
}

const openInboundModal = (product: Product) => {
  selectedProduct.value = product
  inboundForm.value = { quantity: 1, unit_price: product.average_hpp > 0 ? product.average_hpp : 0, reference: '' }
  actionError.value = ''
  activeModal.value = 'inbound'
}

const submitInbound = async () => {
  if (!selectedProduct.value) return
  actionLoading.value = true
  actionError.value = ''
  try {
    await api.post('/api/inventory/inbound', {
      product_id: selectedProduct.value.ID,
      quantity: inboundForm.value.quantity,
      unit_price: inboundForm.value.unit_price,
      reference: inboundForm.value.reference
    })
    closeModal()
    await fetchProducts()
  } catch (error: any) {
    actionError.value = error.response?.data?.error || 'Gagal menyimpan Inbound.'
  } finally {
    actionLoading.value = false
  }
}

const openOpnameModal = (product: Product) => {
  selectedProduct.value = product
  opnameForm.value = { physical_stock: product.current_stock, reference: 'Stock Opname' }
  actionError.value = ''
  activeModal.value = 'opname'
}

const submitOpname = async () => {
  if (!selectedProduct.value || opnameDiff.value === 0) return
  actionLoading.value = true
  actionError.value = ''
  try {
    await api.post('/api/inventory/adjustment', {
      product_id: selectedProduct.value.ID,
      quantity: opnameDiff.value, // Kirim selisih, bukan physical_stock
      reference: opnameForm.value.reference
    })
    closeModal()
    await fetchProducts()
  } catch (error: any) {
    actionError.value = error.response?.data?.error || 'Gagal menyimpan Opname.'
  } finally {
    actionLoading.value = false
  }
}

const openHistoryModal = async (product: Product) => {
  selectedProduct.value = product
  activeModal.value = 'history'
  historyLoading.value = true
  stockHistory.value = []
  try {
    const res = await api.get(`/api/inventory/stock-card/${product.ID}`)
    stockHistory.value = res.data ?? []
  } catch (error) {
    console.error('Failed fetching history', error)
  } finally {
    historyLoading.value = false
  }
}

const closeModal = () => {
  activeModal.value = null
  selectedProduct.value = null
}

onMounted(() => {
  fetchProducts()
})
</script>
