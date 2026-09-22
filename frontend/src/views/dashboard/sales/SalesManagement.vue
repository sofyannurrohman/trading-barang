<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h2 class="text-2xl font-bold text-slate-800">Manajemen Penjualan</h2>
        <p class="text-slate-500">Kelola seluruh data transaksi penjualan dengan mudah.</p>
      </div>
      <button
        @click="openCreateModal"
        class="inline-flex items-center gap-2 px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 font-medium shadow-sm transition-colors"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Buat Transaksi
      </button>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
      <!-- Total Penjualan -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-indigo-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Total Penjualan</p>
          <p class="text-lg font-bold text-slate-900 truncate">Rp {{ formatNumber(summary.totalSales) }}</p>
          <p class="text-xs text-slate-400">{{ filteredInvoices.length }} faktur</p>
        </div>
      </div>

      <!-- Total Transaksi per Mitra -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-emerald-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Total Transaksi Mitra</p>
          <p class="text-lg font-bold text-slate-900 truncate">{{ summary.totalPartnerTx }} transaksi</p>
          <p class="text-xs text-slate-400">{{ summary.uniquePartners }} mitra aktif</p>
        </div>
      </div>

      <!-- Total Item Terjual -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-blue-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Total Item Terjual</p>
          <p class="text-lg font-bold text-slate-900 truncate">{{ formatNumber(summary.totalItems) }} unit</p>
          <p class="text-xs text-slate-400">dari semua faktur</p>
        </div>
      </div>

      <!-- Total Laba -->
      <div class="bg-white rounded-xl shadow-sm border border-slate-200 p-5 flex items-center gap-4">
        <div class="flex-shrink-0 w-12 h-12 bg-purple-100 rounded-xl flex items-center justify-center">
          <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
          </svg>
        </div>
        <div class="min-w-0 flex-1">
          <p class="text-xs font-medium text-slate-500 truncate">Total Laba</p>
          <p class="text-lg font-bold text-slate-900 truncate">Rp {{ formatNumber(summary.totalProfit) }}</p>
          <p class="text-xs text-slate-400">laba kotor (Subtotal - HPP)</p>
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
            placeholder="Cari no. faktur atau mitra..."
            class="pl-9 pr-3 py-2 w-full border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent"
          />
        </div>
      </div>
      <div class="flex gap-2 flex-wrap">
        <input
          v-model="filterDateFrom"
          type="date"
          class="px-3 py-2 border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <span class="self-center text-slate-400 text-sm">s/d</span>
        <input
          v-model="filterDateTo"
          type="date"
          class="px-3 py-2 border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
        />
        <select
          v-model="filterStatus"
          class="px-3 py-2 border border-slate-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500"
        >
          <option value="">Semua Status</option>
          <option value="UNPAID">UNPAID</option>
          <option value="PAID">PAID</option>
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
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Tanggal</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">No. Faktur</th>
              <th class="px-4 py-3 text-left text-xs font-semibold text-slate-500 uppercase tracking-wider">Mitra</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Penjualan</th>
              <th class="px-4 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Transaksi Mitra</th>
              <th class="px-4 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Item</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Ongkos Kirim</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Diskon</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Laba</th>
              <th class="px-4 py-3 text-center text-xs font-semibold text-slate-500 uppercase tracking-wider">Status</th>
              <th class="px-4 py-3 text-right text-xs font-semibold text-slate-500 uppercase tracking-wider">Aksi</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-slate-100">
            <tr v-if="loading">
              <td colspan="11" class="px-4 py-10 text-center">
                <div class="flex flex-col items-center gap-2 text-slate-400">
                  <svg class="animate-spin w-6 h-6" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
                  </svg>
                  <span class="text-sm">Memuat data...</span>
                </div>
              </td>
            </tr>
            <tr v-else-if="paginatedInvoices.length === 0">
              <td colspan="11" class="px-4 py-10 text-center text-slate-400 text-sm">
                Tidak ada data transaksi yang ditemukan.
              </td>
            </tr>
            <tr
              v-else
              v-for="invoice in paginatedInvoices"
              :key="invoice.ID"
              class="hover:bg-slate-50 transition-colors"
            >
              <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-600">{{ formatDate(invoice.CreatedAt) }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm font-medium text-indigo-600">{{ invoice.invoice_number }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-800">{{ invoice.partner?.name ?? '-' }}</td>
              <td class="px-4 py-3 whitespace-nowrap text-sm font-semibold text-slate-900 text-right">
                Rp {{ formatNumber(invoice.grand_total) }}
              </td>
              <!-- Transaksi = total invoice count untuk mitra ini -->
              <td class="px-4 py-3 whitespace-nowrap text-sm text-center">
                <span class="inline-flex items-center justify-center w-8 h-8 rounded-full bg-emerald-100 text-emerald-700 font-semibold text-xs">
                  {{ partnerTxCount[invoice.partner?.name ?? ''] ?? 1 }}
                </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-600 text-center">
                {{ getTotalItems(invoice) }} unit
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-slate-600 text-right">
                <span v-if="invoice.shipping_cost > 0">Rp {{ formatNumber(invoice.shipping_cost) }}</span>
                <span v-else class="text-slate-400">—</span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-right">
                <span v-if="invoice.discount > 0" class="text-red-600 font-medium">- Rp {{ formatNumber(invoice.discount) }}</span>
                <span v-else class="text-slate-400">—</span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm font-semibold text-right">
                <span :class="calcProfit(invoice) >= 0 ? 'text-emerald-600' : 'text-red-600'">
                  Rp {{ formatNumber(calcProfit(invoice)) }}
                </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-sm text-center">
                <span
                  class="px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full"
                  :class="invoice.status === 'PAID' ? 'bg-green-100 text-green-800' : 'bg-amber-100 text-amber-700'"
                >
                  {{ invoice.status }}
                </span>
              </td>
              <td class="px-4 py-3 whitespace-nowrap text-right text-sm font-medium">
                <div class="flex items-center justify-end gap-2">
                  <a
                    :href="`/print/invoice/${invoice.ID}`"
                    target="_blank"
                    class="p-1.5 text-blue-600 hover:text-blue-800 hover:bg-blue-50 rounded-md transition-colors"
                    title="Cetak Faktur"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17h2a2 2 0 002-2v-4a2 2 0 00-2-2H5a2 2 0 00-2 2v4a2 2 0 002 2h2m2 4h6a2 2 0 002-2v-4a2 2 0 00-2-2H9a2 2 0 00-2 2v4a2 2 0 002 2zm8-12V5a2 2 0 00-2-2H9a2 2 0 00-2 2v4h10z" />
                    </svg>
                  </a>
                  <button
                    @click="openEditModal(invoice)"
                    class="p-1.5 text-indigo-600 hover:text-indigo-800 hover:bg-indigo-50 rounded-md transition-colors"
                    title="Edit Invoice"
                  >
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
                    </svg>
                  </button>
                  <button
                    @click="openDeleteModal(invoice)"
                    class="p-1.5 text-red-600 hover:text-red-800 hover:bg-red-50 rounded-md transition-colors"
                    title="Hapus Invoice"
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
          Menampilkan {{ paginationStart }}-{{ paginationEnd }} dari {{ filteredInvoices.length }} data
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
    <!-- MODAL: BUAT TRANSAKSI BARU     -->
    <!-- ============================== -->
    <div v-if="isCreateModalOpen" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-start justify-center min-h-screen pt-8 px-4 pb-20">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isCreateModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-2xl z-10">
          <!-- Modal Header -->
          <div class="px-6 pt-6 pb-4 border-b border-slate-200">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-bold text-slate-900">Buat Transaksi Penjualan Baru</h3>
                <p class="text-sm text-slate-500 mt-0.5">Pilih mitra, tambahkan produk, lalu simpan faktur.</p>
              </div>
              <button @click="isCreateModalOpen = false" class="p-2 text-slate-400 hover:text-slate-600 hover:bg-slate-100 rounded-lg transition-colors">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>

          <div class="px-6 py-4 space-y-4 max-h-[70vh] overflow-y-auto">
            <!-- Pilih Mitra -->
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Mitra / Klien <span class="text-red-500">*</span></label>
              <select v-model="createForm.partner_id" class="block w-full border border-slate-300 rounded-lg shadow-sm py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
                <option value="">— Pilih Mitra —</option>
                <option v-for="p in partners" :key="p.ID" :value="p.ID">{{ p.name }}</option>
              </select>
            </div>

            <!-- Tambah Item -->
            <div class="bg-slate-50 rounded-xl p-4 border border-slate-200">
              <h4 class="text-sm font-semibold text-slate-700 mb-3">Tambah Produk ke Keranjang</h4>
              <div class="grid grid-cols-12 gap-2 items-end">
                <div class="col-span-5">
                  <label class="block text-xs font-medium text-slate-500 mb-1">Produk</label>
                  <select v-model="itemForm.product_id" @change="onProductSelect" class="block w-full border border-slate-300 rounded-lg py-2 px-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500">
                    <option value="">— Pilih —</option>
                    <option v-for="p in products" :key="p.ID" :value="p.ID">{{ p.name }} ({{ p.current_stock }})</option>
                  </select>
                </div>
                <div class="col-span-2">
                  <label class="block text-xs font-medium text-slate-500 mb-1">Qty</label>
                  <input type="number" v-model="itemForm.quantity" min="1" class="block w-full border border-slate-300 rounded-lg py-2 px-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
                </div>
                <div class="col-span-3">
                  <label class="block text-xs font-medium text-slate-500 mb-1">Harga (Rp)</label>
                  <input type="number" v-model="itemForm.unit_price" class="block w-full border border-slate-300 rounded-lg py-2 px-2 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
                </div>
                <div class="col-span-2">
                  <button @click="addItemToCart" type="button" class="w-full bg-indigo-600 text-white py-2 rounded-lg text-sm font-medium hover:bg-indigo-700 transition-colors">
                    + Tambah
                  </button>
                </div>
              </div>

              <!-- Cart Table -->
              <div v-if="cart.length > 0" class="mt-3 rounded-lg border border-slate-200 overflow-hidden">
                <table class="min-w-full divide-y divide-slate-200">
                  <thead class="bg-white">
                    <tr>
                      <th class="px-3 py-2 text-left text-xs font-medium text-slate-500">Produk</th>
                      <th class="px-3 py-2 text-center text-xs font-medium text-slate-500">Qty</th>
                      <th class="px-3 py-2 text-right text-xs font-medium text-slate-500">Subtotal</th>
                      <th class="px-3 py-2"></th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-slate-100 bg-white">
                    <tr v-for="(item, idx) in cart" :key="idx">
                      <td class="px-3 py-2 text-sm text-slate-700">{{ item.product_name }}</td>
                      <td class="px-3 py-2 text-sm text-center text-slate-600">{{ item.quantity }}</td>
                      <td class="px-3 py-2 text-sm text-right font-medium text-slate-800">Rp {{ formatNumber(item.quantity * item.unit_price) }}</td>
                      <td class="px-3 py-2 text-right">
                        <button @click="cart.splice(idx, 1)" class="text-red-500 hover:text-red-700 text-xs">✕</button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <p v-else class="mt-3 text-xs text-slate-400 text-center py-2">Belum ada produk ditambahkan.</p>
            </div>

            <!-- Ongkos Kirim & Diskon -->
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Ongkos Kirim (Rp)</label>
                <input type="number" v-model="createForm.shipping_cost" min="0" class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Diskon (Rp)</label>
                <input type="number" v-model="createForm.discount" min="0" class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
              </div>
            </div>

            <!-- Summary Mini -->
            <div class="bg-slate-900 rounded-xl p-4 text-white text-sm space-y-2">
              <div class="flex justify-between text-slate-300">
                <span>DPP (Subtotal Produk)</span>
                <span>Rp {{ formatNumber(cartTotal) }}</span>
              </div>
              <div class="flex justify-between text-slate-300">
                <span>Ongkos Kirim</span>
                <span>Rp {{ formatNumber(Number(createForm.shipping_cost) || 0) }}</span>
              </div>
              <div class="flex justify-between text-slate-300">
                <span>Diskon</span>
                <span class="text-red-400">- Rp {{ formatNumber(Number(createForm.discount) || 0) }}</span>
              </div>
              <div class="border-t border-slate-700 pt-2 flex justify-between font-bold text-base">
                <span>Total Akhir</span>
                <span class="text-indigo-400">Rp {{ formatNumber(cartGrandTotal) }}</span>
              </div>
            </div>

            <p v-if="createError" class="text-red-600 text-sm bg-red-50 border border-red-200 rounded-lg p-3">{{ createError }}</p>
          </div>

          <div class="px-6 py-4 border-t border-slate-200 flex justify-end gap-3">
            <button @click="isCreateModalOpen = false" class="px-4 py-2 text-sm font-medium text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">
              Batal
            </button>
            <button
              @click="submitCreate"
              :disabled="createLoading || cart.length === 0 || !createForm.partner_id"
              class="px-5 py-2 text-sm font-semibold bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              {{ createLoading ? 'Menyimpan...' : 'Simpan Transaksi' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL: EDIT INVOICE            -->
    <!-- ============================== -->
    <div v-if="isEditModalOpen && editingInvoice" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm" @click="isEditModalOpen = false"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-md z-10">
          <!-- Modal Header -->
          <div class="px-6 pt-6 pb-4 border-b border-slate-200">
            <div class="flex items-center justify-between">
              <div>
                <h3 class="text-lg font-bold text-slate-900">Edit Invoice</h3>
                <p class="text-sm text-indigo-600 font-medium mt-0.5">{{ editingInvoice.invoice_number }}</p>
              </div>
              <button @click="isEditModalOpen = false" class="p-2 text-slate-400 hover:text-slate-600 hover:bg-slate-100 rounded-lg transition-colors">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>
          </div>

          <div class="px-6 py-5 space-y-4">
            <!-- Info Konteks -->
            <div class="bg-blue-50 border border-blue-200 rounded-xl p-4 text-sm">
              <div class="flex gap-2 items-start">
                <svg class="w-5 h-5 text-blue-500 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <div class="text-blue-800 space-y-1">
                  <p class="font-medium">Yang dapat diubah:</p>
                  <ul class="list-disc list-inside space-y-0.5 text-blue-700">
                    <li>Status pembayaran</li>
                    <li>Ongkos kirim</li>
                    <li>Diskon</li>
                  </ul>
                  <p class="text-blue-600 text-xs mt-1">* Item produk tidak dapat diubah untuk menjaga konsistensi stok.</p>
                </div>
              </div>
            </div>

            <!-- Detail Invoice (readonly) -->
            <div class="bg-slate-50 rounded-xl p-4 grid grid-cols-2 gap-3 text-sm">
              <div>
                <span class="text-slate-500">Mitra</span>
                <p class="font-medium text-slate-800">{{ editingInvoice.partner?.name }}</p>
              </div>
              <div>
                <span class="text-slate-500">Tanggal</span>
                <p class="font-medium text-slate-800">{{ formatDate(editingInvoice.CreatedAt) }}</p>
              </div>
              <div class="col-span-2">
                <span class="text-slate-500">Subtotal Produk</span>
                <p class="font-medium text-slate-800">Rp {{ formatNumber(editingInvoice.total_dpp) }}</p>
              </div>
            </div>

            <!-- Form Edit -->
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-1">Status Pembayaran</label>
              <div class="flex gap-3">
                <label class="flex items-center gap-2 cursor-pointer">
                  <input type="radio" v-model="editForm.status" value="UNPAID" class="text-indigo-600" />
                  <span class="text-sm font-medium text-amber-700 bg-amber-100 px-3 py-1 rounded-full">UNPAID</span>
                </label>
                <label class="flex items-center gap-2 cursor-pointer">
                  <input type="radio" v-model="editForm.status" value="PAID" class="text-indigo-600" />
                  <span class="text-sm font-medium text-green-700 bg-green-100 px-3 py-1 rounded-full">PAID</span>
                </label>
              </div>
            </div>
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Ongkos Kirim (Rp)</label>
                <input type="number" v-model="editForm.shipping_cost" min="0" class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
              </div>
              <div>
                <label class="block text-sm font-medium text-slate-700 mb-1">Diskon (Rp)</label>
                <input type="number" v-model="editForm.discount" min="0" class="block w-full border border-slate-300 rounded-lg py-2 px-3 text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500" />
              </div>
            </div>

            <!-- Preview Total Baru -->
            <div class="bg-slate-900 rounded-xl p-3 text-white text-sm flex justify-between">
              <span class="text-slate-400">Total Akhir (Baru)</span>
              <span class="font-bold text-indigo-400">
                Rp {{ formatNumber(
                  editingInvoice.total_dpp
                  + Number(editForm.shipping_cost || 0)
                  - Number(editForm.discount || 0)
                ) }}
              </span>
            </div>
          </div>

          <div class="px-6 py-4 border-t border-slate-200 flex justify-end gap-3">
            <button @click="isEditModalOpen = false" class="px-4 py-2 text-sm font-medium text-slate-700 border border-slate-300 rounded-lg hover:bg-slate-50 transition-colors">
              Batal
            </button>
            <button
              @click="submitEdit"
              :disabled="editLoading"
              class="px-5 py-2 text-sm font-semibold bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition-colors"
            >
              {{ editLoading ? 'Menyimpan...' : 'Simpan Perubahan' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- ============================== -->
    <!-- MODAL: HAPUS INVOICE (Informatif) -->
    <!-- ============================== -->
    <div v-if="isDeleteModalOpen && deletingInvoice" class="fixed inset-0 z-50 overflow-y-auto">
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
                <h3 class="text-lg font-bold text-red-800">Hapus Invoice</h3>
                <p class="text-sm text-red-600 mt-0.5">Tindakan ini tidak dapat dibatalkan.</p>
              </div>
            </div>
          </div>

          <div class="px-6 py-5 space-y-4">
            <!-- Info Invoice -->
            <div class="bg-slate-50 rounded-xl p-4 border border-slate-200">
              <p class="text-xs font-semibold text-slate-500 uppercase tracking-wide mb-2">Detail Invoice yang Akan Dihapus</p>
              <div class="grid grid-cols-2 gap-2 text-sm">
                <div>
                  <span class="text-slate-500">No. Faktur</span>
                  <p class="font-semibold text-indigo-600">{{ deletingInvoice.invoice_number }}</p>
                </div>
                <div>
                  <span class="text-slate-500">Mitra</span>
                  <p class="font-medium text-slate-800">{{ deletingInvoice.partner?.name }}</p>
                </div>
                <div>
                  <span class="text-slate-500">Tanggal</span>
                  <p class="font-medium text-slate-800">{{ formatDate(deletingInvoice.CreatedAt) }}</p>
                </div>
                <div>
                  <span class="text-slate-500">Total</span>
                  <p class="font-bold text-slate-800">Rp {{ formatNumber(deletingInvoice.grand_total) }}</p>
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
                  <span>Invoice <strong>{{ deletingInvoice.invoice_number }}</strong> akan dihapus secara permanen.</span>
                </li>
                <li class="flex items-start gap-2">
                  <svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                  </svg>
                  <span>Stok <strong>{{ getTotalItems(deletingInvoice) }} unit produk</strong> akan dikembalikan ke gudang.</span>
                </li>
                <li class="flex items-start gap-2">
                  <svg class="w-4 h-4 text-amber-500 flex-shrink-0 mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                  </svg>
                  <span>Riwayat transaksi senilai <strong>Rp {{ formatNumber(deletingInvoice.grand_total) }}</strong> akan hilang dari laporan.</span>
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
              {{ deleteLoading ? 'Menghapus...' : 'Ya, Hapus Invoice' }}
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

// ─── Types ────────────────────────────────────────────────────────────────────
interface Partner { ID: number; name: string; npwp: string; nik: string }
interface Product { ID: number; name: string; standard_price: number; current_stock: number }

interface InvoiceItem {
  ID: number
  product_id: number
  quantity: number
  unit_price: number
  subtotal: number
  hpp_at_sale: number
  product?: { name: string }
}

interface Invoice {
  ID: number
  CreatedAt: string
  invoice_number: string
  partner_id: number
  partner?: { name: string }
  total_dpp: number
  total_ppn: number
  grand_total: number
  shipping_cost: number
  discount: number
  status: string
  items?: InvoiceItem[]
}

interface CartItem {
  product_id: number
  product_name: string
  quantity: number
  unit_price: number
}

// ─── State ────────────────────────────────────────────────────────────────────
const invoices  = ref<Invoice[]>([])
const partners  = ref<Partner[]>([])
const products  = ref<Product[]>([])
const loading   = ref(false)

// Filter
const searchQuery  = ref('')
const filterStatus = ref('')
const filterDateFrom = ref('')
const filterDateTo   = ref('')

// Pagination
const currentPage = ref(1)
const pageSize    = 10

// Create Modal
const isCreateModalOpen = ref(false)
const createLoading     = ref(false)
const createError       = ref('')
const cart              = ref<CartItem[]>([])
const createForm = ref({ partner_id: '', shipping_cost: 0, discount: 0 })
const itemForm   = ref({ product_id: '', quantity: 1, unit_price: 0 })

// Edit Modal
const isEditModalOpen = ref(false)
const editLoading     = ref(false)
const editingInvoice  = ref<Invoice | null>(null)
const editForm = ref({ status: 'UNPAID', shipping_cost: 0, discount: 0 })

// Delete Modal
const isDeleteModalOpen = ref(false)
const deleteLoading     = ref(false)
const deletingInvoice   = ref<Invoice | null>(null)
const deleteError       = ref('')

// ─── Computed ────────────────────────────────────────────────────────────────
const filteredInvoices = computed(() => {
  let list = [...invoices.value]
  const q = searchQuery.value.toLowerCase()

  if (q) {
    list = list.filter(inv =>
      inv.invoice_number.toLowerCase().includes(q) ||
      (inv.partner?.name ?? '').toLowerCase().includes(q)
    )
  }
  if (filterStatus.value) {
    list = list.filter(inv => inv.status === filterStatus.value)
  }
  if (filterDateFrom.value) {
    const from = new Date(filterDateFrom.value)
    list = list.filter(inv => new Date(inv.CreatedAt) >= from)
  }
  if (filterDateTo.value) {
    const to = new Date(filterDateTo.value)
    to.setHours(23, 59, 59)
    list = list.filter(inv => new Date(inv.CreatedAt) <= to)
  }
  return list
})

const totalPages = computed(() => Math.max(1, Math.ceil(filteredInvoices.value.length / pageSize)))

const paginatedInvoices = computed(() => {
  const start = (currentPage.value - 1) * pageSize
  return filteredInvoices.value.slice(start, start + pageSize)
})

const paginationStart = computed(() => {
  if (filteredInvoices.value.length === 0) return 0
  return (currentPage.value - 1) * pageSize + 1
})

const paginationEnd = computed(() =>
  Math.min(currentPage.value * pageSize, filteredInvoices.value.length)
)

// partnerTxCount: { [partnerName]: count invoice di semua data (bukan hanya filtered) }
const partnerTxCount = computed(() => {
  const map: Record<string, number> = {}
  for (const inv of invoices.value) {
    const name = inv.partner?.name ?? ''
    if (name) map[name] = (map[name] ?? 0) + 1
  }
  return map
})

// Summary cards
const summary = computed(() => {
  const list = filteredInvoices.value
  const totalSales = list.reduce((s, inv) => s + inv.grand_total, 0)
  const totalItems = list.reduce((s, inv) => s + getTotalItems(inv), 0)
  const totalProfit = list.reduce((s, inv) => s + calcProfit(inv), 0)

  // Total Transaksi Mitra: sum of each invoice's partner tx count (total count per row for mitra)
  const uniquePartners = new Set(list.map(inv => inv.partner?.name).filter(Boolean)).size
  const totalPartnerTx = list.length // total transaksi keseluruhan yang ditampilkan

  return { totalSales, totalItems, totalProfit, uniquePartners, totalPartnerTx }
})

// Cart computed
const cartTotal = computed(() =>
  cart.value.reduce((s, i) => s + i.quantity * i.unit_price, 0)
)
const cartGrandTotal = computed(() => {
  const ship = Number(createForm.value.shipping_cost) || 0
  const disc = Number(createForm.value.discount) || 0
  return cartTotal.value + ship - disc
})

// ─── Helpers ─────────────────────────────────────────────────────────────────
const formatNumber = (num: number) =>
  Math.round(num).toLocaleString('id-ID')

const formatDate = (dateStr: string) =>
  new Date(dateStr).toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' })

const getTotalItems = (invoice: Invoice): number =>
  (invoice.items ?? []).reduce((s, item) => s + item.quantity, 0)

const calcProfit = (invoice: Invoice): number => {
  const totalHPP = (invoice.items ?? []).reduce(
    (s, item) => s + item.quantity * (item.hpp_at_sale ?? 0), 0
  )
  // Laba = Pendapatan (DPP) - HPP
  return invoice.total_dpp - totalHPP
}

// Reset page when filter changes
watch([searchQuery, filterStatus, filterDateFrom, filterDateTo], () => {
  currentPage.value = 1
})

// ─── Fetch Data ───────────────────────────────────────────────────────────────
const fetchAll = async () => {
  loading.value = true
  try {
    const [invRes, partRes, prodRes] = await Promise.all([
      api.get('/api/sales/invoices?include_items=true'),
      api.get('/api/partners'),
      api.get('/api/products'),
    ])
    // Fetch detail (with items) for each invoice to get items for profit calc
    const basicInvoices: Invoice[] = invRes.data ?? []
    // Fetch details in batches for items
    const detailed = await Promise.all(
      basicInvoices.map(inv => api.get(`/api/sales/invoices/${inv.ID}`).then(r => r.data).catch(() => inv))
    )
    invoices.value = detailed
    partners.value = partRes.data ?? []
    products.value = prodRes.data ?? []
  } catch (err) {
    console.error('Fetch error:', err)
  } finally {
    loading.value = false
  }
}

const resetFilters = () => {
  searchQuery.value = ''
  filterStatus.value = ''
  filterDateFrom.value = ''
  filterDateTo.value = ''
  currentPage.value = 1
}

// ─── Create ───────────────────────────────────────────────────────────────────
const openCreateModal = () => {
  createForm.value = { partner_id: '', shipping_cost: 0, discount: 0 }
  itemForm.value = { product_id: '', quantity: 1, unit_price: 0 }
  cart.value = []
  createError.value = ''
  isCreateModalOpen.value = true
}

const onProductSelect = () => {
  const p = products.value.find(x => x.ID === Number(itemForm.value.product_id))
  if (p) {
    itemForm.value.unit_price = p.standard_price
    itemForm.value.quantity = 1
  }
}

const addItemToCart = () => {
  if (!itemForm.value.product_id || itemForm.value.quantity < 1) return
  const p = products.value.find(x => x.ID === Number(itemForm.value.product_id))
  if (!p) return
  cart.value.push({
    product_id: p.ID,
    product_name: p.name,
    quantity: itemForm.value.quantity,
    unit_price: itemForm.value.unit_price,
  })
  itemForm.value = { product_id: '', quantity: 1, unit_price: 0 }
}

const submitCreate = async () => {
  if (!createForm.value.partner_id || cart.value.length === 0) return
  createLoading.value = true
  createError.value = ''
  try {
    await api.post('/api/sales/invoice', {
      partner_id: Number(createForm.value.partner_id),
      is_taxable: false,
      shipping_cost: Number(createForm.value.shipping_cost) || 0,
      discount: Number(createForm.value.discount) || 0,
      items: cart.value.map(i => ({
        product_id: i.product_id,
        quantity: i.quantity,
        unit_price: i.unit_price,
      })),
    }, { skipToast: true } as any)
    // Tutup modal dulu, baru toast agar tidak terhalangi backdrop
    isCreateModalOpen.value = false
    toast.success('Transaksi penjualan berhasil dibuat')
    await fetchAll()
  } catch (err: any) {
    const errMsg = err.response?.data?.error ?? 'Gagal membuat transaksi.'
    createError.value = errMsg
    toast.error(errMsg)
  } finally {
    createLoading.value = false
  }
}

// ─── Edit ─────────────────────────────────────────────────────────────────────
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
    // Tutup modal dulu, baru toast agar tidak terhalangi backdrop
    isEditModalOpen.value = false
    toast.success(`Invoice ${invoiceNumber} berhasil diperbarui`)
    await fetchAll()
  } catch (err: any) {
    toast.error(err.response?.data?.error ?? 'Gagal memperbarui invoice.')
  } finally {
    editLoading.value = false
  }
}

// ─── Delete ───────────────────────────────────────────────────────────────────
const openDeleteModal = (invoice: Invoice) => {
  deletingInvoice.value = invoice
  deleteError.value = ''
  isDeleteModalOpen.value = true
}

const submitDelete = async () => {
  if (!deletingInvoice.value) return
  deleteLoading.value = true
  deleteError.value = ''
  const invoiceNumber = deletingInvoice.value.invoice_number
  try {
    await api.delete(`/api/sales/invoices/${deletingInvoice.value.ID}`, { skipToast: true } as any)
    // Tutup modal dulu, baru toast agar tidak terhalangi backdrop
    isDeleteModalOpen.value = false
    toast.success(`Invoice ${invoiceNumber} berhasil dihapus`)
    await fetchAll()
  } catch (err: any) {
    const errMsg = err.response?.data?.error ?? 'Gagal menghapus invoice.'
    deleteError.value = errMsg
    toast.error(errMsg)
  } finally {
    deleteLoading.value = false
  }
}

// ─── Init ─────────────────────────────────────────────────────────────────────
onMounted(fetchAll)
</script>
