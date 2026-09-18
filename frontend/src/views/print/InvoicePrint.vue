<template>
  <div class="bg-slate-100 min-h-screen p-2 sm:p-6 md:p-8 flex flex-col items-center justify-start print:p-0 print:bg-white text-slate-900 selection:bg-rose-500 selection:text-white">
    
    <!-- Loading State -->
    <div v-if="loading" class="text-center py-16 print:hidden flex flex-col items-center gap-3 text-slate-500">
      <Loader2 class="w-8 h-8 animate-spin text-rose-700" />
      <span class="text-sm font-medium">Memuat data faktur UD Duo Srikandi...</span>
    </div>
    
    <div v-else-if="invoice" class="w-full max-w-4xl space-y-6">

      <!-- Top Action Bar (Print Hidden) -->
      <div class="bg-white p-4 rounded-2xl shadow-sm border border-slate-200/80 flex flex-wrap items-center justify-between gap-3 print:hidden">
        <div class="flex items-center gap-2">
          <router-link
            to="/dashboard/sales/list"
            class="inline-flex items-center gap-1.5 px-3.5 py-2 text-xs font-semibold text-slate-700 hover:text-slate-900 bg-slate-100 hover:bg-slate-200 rounded-xl transition-colors cursor-pointer"
          >
            <ArrowLeft class="w-4 h-4" />
            <span>Kembali ke Daftar</span>
          </router-link>
          <span class="text-xs font-bold text-slate-400">|</span>
          <span class="text-xs font-bold text-rose-800 bg-rose-50 px-2.5 py-1 rounded-lg border border-rose-200">
            Faktur: {{ invoice.invoice_number }}
          </span>
        </div>

        <div class="flex items-center gap-2">
          <button
            @click="isCustomizeOpen = !isCustomizeOpen"
            class="inline-flex items-center gap-1.5 px-3.5 py-2 text-xs font-semibold text-slate-700 hover:text-slate-900 bg-slate-100 hover:bg-slate-200 rounded-xl transition-colors cursor-pointer border border-slate-300"
          >
            <SlidersHorizontal class="w-4 h-4 text-rose-700" />
            <span>{{ isCustomizeOpen ? 'Tutup Pengaturan' : 'Sesuaikan Tampilan' }}</span>
          </button>
          
          <button
            @click="doPrint"
            class="inline-flex items-center gap-2 px-5 py-2 text-xs sm:text-sm font-bold text-white bg-gradient-to-r from-rose-700 to-red-800 hover:from-rose-600 hover:to-red-700 rounded-xl shadow-md shadow-rose-950/20 transition-all cursor-pointer"
          >
            <Printer class="w-4 h-4" />
            <span>Cetak / Simpan PDF</span>
          </button>
        </div>
      </div>

      <!-- Live Customization Drawer / Panel (Print Hidden) -->
      <div v-show="isCustomizeOpen" class="bg-white p-5 rounded-2xl shadow-md border border-rose-200 print:hidden space-y-4 transition-all">
        <div class="flex items-center justify-between pb-3 border-b border-slate-100">
          <h4 class="text-sm font-bold text-slate-900 flex items-center gap-2">
            <SlidersHorizontal class="w-4 h-4 text-rose-700" />
            <span>Kustomisasi Pratinjau Faktur (Live Edit)</span>
          </h4>
          <span class="text-xs text-slate-400">Perubahan ini hanya berlaku untuk sesi cetak saat ini</span>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 text-xs">
          <div>
            <label class="block font-semibold text-slate-700 mb-1">Judul Dokumen</label>
            <select v-model="customDocTitle" class="w-full border border-slate-300 rounded-lg p-2 focus:ring-2 focus:ring-rose-500 outline-none bg-white">
              <option value="FAKTUR PENJUALAN">FAKTUR PENJUALAN</option>
              <option value="INVOICE &amp; KWITANSI">INVOICE &amp; KWITANSI</option>
              <option value="SURAT JALAN / PENGIRIMAN">SURAT JALAN / PENGIRIMAN</option>
              <option value="NOTA PENJUALAN">NOTA PENJUALAN</option>
            </select>
          </div>

          <div>
            <label class="block font-semibold text-slate-700 mb-1">Pihak Penandatangan (Kanan)</label>
            <input 
              type="text" 
              v-model="customSignerTitle" 
              placeholder="Contoh: Pimpinan / Bagian Keuangan" 
              class="w-full border border-slate-300 rounded-lg p-2 focus:ring-2 focus:ring-rose-500 outline-none" 
            />
          </div>

          <div>
            <label class="block font-semibold text-slate-700 mb-1">Opsi Tampilan</label>
            <div class="space-y-1.5 pt-1">
              <label class="flex items-center gap-2 cursor-pointer">
                <input type="checkbox" v-model="showBankInfo" class="rounded text-rose-700 focus:ring-rose-500" />
                <span>Tampilkan Info Rekening Bank</span>
              </label>
              <label class="flex items-center gap-2 cursor-pointer">
                <input type="checkbox" v-model="showNotes" class="rounded text-rose-700 focus:ring-rose-500" />
                <span>Tampilkan Syarat / Catatan</span>
              </label>
            </div>
          </div>

          <div class="sm:col-span-3">
            <label class="block font-semibold text-slate-700 mb-1">Catatan Khusus Faktur Ini</label>
            <textarea 
              v-model="customNotes" 
              rows="2" 
              placeholder="Tambahkan catatan khusus untuk pelanggan pada faktur ini..." 
              class="w-full border border-slate-300 rounded-lg p-2 focus:ring-2 focus:ring-rose-500 outline-none"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- Printable Invoice Sheet Area -->
      <div 
        class="print-container bg-white text-black font-sans p-6 sm:p-10 rounded-2xl shadow-xl border border-slate-200 print:border-0 print:shadow-none print:rounded-none w-full mx-auto"
        style="min-height: 297mm;"
      >
        <!-- Kop Surat Resmi UD Duo Srikandi -->
        <div class="border-b-4 border-double border-slate-900 pb-5 mb-6">
          <div class="flex flex-col sm:flex-row justify-between items-start gap-4">
            
            <!-- Left: Company Identity -->
            <div class="space-y-1 max-w-md">
              <div class="flex items-center gap-2">
                <div class="w-8 h-8 rounded-lg bg-slate-900 text-white flex items-center justify-center font-black text-sm print:border print:border-black">
                  DS
                </div>
                <h2 class="text-2xl font-black uppercase tracking-wider text-slate-950">
                  {{ company.name || 'UD DUO SRIKANDI' }}
                </h2>
              </div>
              <p class="text-xs text-slate-800 leading-relaxed font-medium whitespace-pre-line mt-1">
                {{ company.address || 'Jl. Perdagangan Raya No. 88, Jawa Timur' }}
              </p>
              <div class="text-[11px] text-slate-700 flex flex-wrap gap-x-3 gap-y-0.5 pt-0.5">
                <span v-if="company.phone"><strong>Telp/WA:</strong> {{ company.phone }}</span>
                <span v-if="company.email"><strong>Email:</strong> {{ company.email }}</span>
                <span v-if="company.npwp"><strong>NPWP:</strong> {{ company.npwp }}</span>
              </div>
            </div>

            <!-- Right: Document & Invoice Number Box -->
            <div class="text-left sm:text-right w-full sm:w-auto">
              <h1 class="text-xl sm:text-2xl font-black uppercase tracking-widest text-slate-950 mb-1">
                {{ customDocTitle }}
              </h1>
              <div class="inline-block bg-slate-50 print:bg-transparent px-3 py-1.5 rounded-lg border border-slate-300 print:border-black text-left text-xs space-y-0.5">
                <div class="flex justify-between sm:justify-start gap-4">
                  <span class="text-slate-600 font-medium">No. Faktur:</span>
                  <strong class="text-slate-950 font-mono">{{ invoice.invoice_number }}</strong>
                </div>
                <div class="flex justify-between sm:justify-start gap-4">
                  <span class="text-slate-600 font-medium">Tanggal:</span>
                  <strong class="text-slate-950">{{ formatDate(invoice.CreatedAt) }}</strong>
                </div>
                <div class="flex justify-between sm:justify-start gap-4">
                  <span class="text-slate-600 font-medium">Status:</span>
                  <span 
                    class="font-bold uppercase" 
                    :class="invoice.status === 'PAID' ? 'text-emerald-700' : 'text-amber-700'"
                  >
                    {{ invoice.status === 'PAID' ? 'LUNAS (PAID)' : 'BELUM LUNAS' }}
                  </span>
                </div>
              </div>
            </div>

          </div>
        </div>

        <!-- Client Destination Info -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6 text-xs sm:text-sm">
          <div class="p-3.5 rounded-xl border border-slate-300 print:border-black bg-slate-50/50 print:bg-transparent space-y-1">
            <span class="text-[11px] font-bold text-slate-500 print:text-black uppercase tracking-wider block border-b border-slate-200 print:border-black pb-1 mb-1.5">
              KEPADA YTH. (PEMBELI):
            </span>
            <p class="font-black text-sm sm:text-base text-slate-950">{{ invoice.partner?.name }}</p>
            <p class="text-slate-700 whitespace-pre-line leading-relaxed">{{ invoice.partner?.address || 'Alamat tidak dicantumkan' }}</p>
            <p class="pt-1 text-slate-800" v-if="invoice.partner?.npwp"><strong>NPWP:</strong> {{ invoice.partner?.npwp }}</p>
            <p class="pt-1 text-slate-800" v-else-if="invoice.partner?.nik"><strong>NIK:</strong> {{ invoice.partner?.nik }}</p>
          </div>

          <div class="p-3.5 rounded-xl border border-slate-300 print:border-black bg-slate-50/50 print:bg-transparent space-y-1 flex flex-col justify-between">
            <div>
              <span class="text-[11px] font-bold text-slate-500 print:text-black uppercase tracking-wider block border-b border-slate-200 print:border-black pb-1 mb-1.5">
                KETERANGAN TRANSAKSI:
              </span>
              <p class="text-slate-800 leading-relaxed">
                Penjualan komoditas barang dagang resmi dari <strong>{{ company.name || 'UD DUO SRIKANDI' }}</strong>.
              </p>
            </div>
            <div class="text-[11px] text-slate-600 pt-2 border-t border-slate-200 print:border-black">
              <span>Mata Uang: <strong>IDR (Rupiah)</strong></span>
            </div>
          </div>
        </div>

        <!-- Items Table -->
        <div class="overflow-x-auto mb-6">
          <table class="w-full border-collapse border border-slate-900 text-xs sm:text-sm">
            <thead>
              <tr class="bg-slate-100 print:bg-slate-200 text-slate-950 font-bold">
                <th class="border border-slate-900 px-3 py-2.5 text-center w-10">No</th>
                <th class="border border-slate-900 px-3 py-2.5 text-left">Nama / Deskripsi Barang</th>
                <th class="border border-slate-900 px-3 py-2.5 text-center w-20">Qty</th>
                <th class="border border-slate-900 px-3 py-2.5 text-right w-32">Harga Satuan (Rp)</th>
                <th class="border border-slate-900 px-3 py-2.5 text-right w-36">Subtotal (Rp)</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, index) in invoice.items" :key="item.ID" class="border-b border-slate-900">
                <td class="border-r border-slate-900 px-3 py-2 text-center">{{ Number(index) + 1 }}</td>
                <td class="border-r border-slate-900 px-3 py-2 font-medium text-slate-950">{{ item.product?.name }}</td>
                <td class="border-r border-slate-900 px-3 py-2 text-center font-semibold">{{ item.quantity }}</td>
                <td class="border-r border-slate-900 px-3 py-2 text-right">{{ formatNumber(item.unit_price) }}</td>
                <td class="px-3 py-2 text-right font-bold text-slate-950">{{ formatNumber(item.subtotal) }}</td>
              </tr>
            </tbody>
            <tfoot>
              <!-- DPP -->
              <tr class="border-t-2 border-slate-900">
                <td colspan="3" class="border-r border-slate-900 p-2 text-right"></td>
                <td class="border-r border-slate-900 px-3 py-1.5 font-bold text-right text-slate-800">DPP (Dasar Pajak)</td>
                <td class="px-3 py-1.5 font-bold text-right text-slate-950">Rp {{ formatNumber(invoice.total_dpp) }}</td>
              </tr>
              <!-- PPN -->
              <tr>
                <td colspan="3" class="border-r border-slate-900 p-2 text-right"></td>
                <td class="border-r border-slate-900 px-3 py-1.5 font-bold text-right text-slate-800">
                  PPN ({{ invoice.is_taxable ? (company.ppn_rate || 11) : 0 }}%)
                </td>
                <td class="px-3 py-1.5 font-bold text-right text-slate-950">Rp {{ formatNumber(invoice.total_ppn) }}</td>
              </tr>
              <!-- Ongkir (if any) -->
              <tr v-if="invoice.shipping_cost > 0">
                <td colspan="3" class="border-r border-slate-900 p-2 text-right"></td>
                <td class="border-r border-slate-900 px-3 py-1.5 font-bold text-right text-slate-800">Ongkos Kirim</td>
                <td class="px-3 py-1.5 font-bold text-right text-slate-950">Rp {{ formatNumber(invoice.shipping_cost) }}</td>
              </tr>
              <!-- Diskon (if any) -->
              <tr v-if="invoice.discount > 0">
                <td colspan="3" class="border-r border-slate-900 p-2 text-right"></td>
                <td class="border-r border-slate-900 px-3 py-1.5 font-bold text-right text-slate-800">Potongan Diskon</td>
                <td class="px-3 py-1.5 font-bold text-right text-rose-700">- Rp {{ formatNumber(invoice.discount) }}</td>
              </tr>
              <!-- GRAND TOTAL -->
              <tr class="border-t-2 border-slate-900 bg-slate-100 print:bg-slate-200">
                <td colspan="3" class="border-r border-slate-900 p-2 text-right font-bold text-xs">
                  <span>Terbilang: </span><em class="capitalize font-normal text-slate-800"># {{ terbilang(invoice.grand_total) }} Rupiah #</em>
                </td>
                <td class="border-r border-slate-900 px-3 py-2 font-black text-right text-sm uppercase">GRAND TOTAL</td>
                <td class="px-3 py-2 font-black text-right text-sm sm:text-base text-slate-950">Rp {{ formatNumber(invoice.grand_total) }}</td>
              </tr>
            </tfoot>
          </table>
        </div>

        <!-- Footer Info: Bank & Notes -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 text-xs mb-8">
          <!-- Bank Info -->
          <div v-if="showBankInfo && (company.bank_info || defaultBankInfo)" class="p-3 rounded-xl border border-slate-300 print:border-black">
            <span class="font-bold block uppercase text-[10px] text-slate-500 print:text-black mb-1">
              PEMBAYARAN VIA TRANSFER BANK:
            </span>
            <p class="font-mono text-xs whitespace-pre-line text-slate-900 font-semibold leading-relaxed">
              {{ company.bank_info || defaultBankInfo }}
            </p>
          </div>

          <!-- Notes / Syarat Ketentuan -->
          <div v-if="showNotes" class="p-3 rounded-xl border border-slate-300 print:border-black">
            <span class="font-bold block uppercase text-[10px] text-slate-500 print:text-black mb-1">
              SYARAT &amp; KETENTUAN:
            </span>
            <p class="text-[11px] text-slate-700 leading-relaxed">
              {{ customNotes || company.invoice_footer || defaultNotes }}
            </p>
          </div>
        </div>

        <!-- Dual Formal Signatures -->
        <div class="flex justify-between items-end pt-4 text-xs sm:text-sm">
          <div class="text-center w-40 sm:w-52 space-y-16">
            <p class="font-semibold text-slate-800">Tanda Terima Pelanggan,</p>
            <div>
              <div class="border-b border-black w-36 sm:w-44 mx-auto"></div>
              <p class="text-[11px] text-slate-600 mt-1">(Nama Terang &amp; Cap)</p>
            </div>
          </div>

          <div class="text-center w-40 sm:w-52 space-y-16">
            <div>
              <p class="font-semibold text-slate-800">Hormat Kami,</p>
              <strong class="block text-xs uppercase text-slate-950">{{ company.name || 'UD DUO SRIKANDI' }}</strong>
            </div>
            <div>
              <div class="border-b border-black w-36 sm:w-44 mx-auto"></div>
              <p class="text-[11px] font-bold text-slate-900 mt-1">{{ customSignerTitle }}</p>
            </div>
          </div>
        </div>

      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import api from '@/plugins/axios'
import {
  Printer,
  ArrowLeft,
  SlidersHorizontal,
  Loader2
} from '@lucide/vue'

const route = useRoute()
const invoice = ref<any>(null)
const loading = ref(true)
const company = ref<any>({})

// Customizer State
const isCustomizeOpen = ref(false)
const customDocTitle = ref('FAKTUR PENJUALAN')
const customSignerTitle = ref('Pimpinan / Bagian Keuangan')
const showBankInfo = ref(true)
const showNotes = ref(true)
const customNotes = ref('')

const defaultBankInfo = `BCA: 8870-123-456 a/n UD DUO SRIKANDI\nBRI: 0123-01-000456-50-1 a/n UD DUO SRIKANDI`
const defaultNotes = `Barang yang sudah dibeli tidak dapat ditukar/dikembalikan kecuali ada perjanjian tertulis sebelumnya.`

const formatNumber = (num: number) => {
  if (num === undefined || num === null) return '0'
  return Math.round(num).toLocaleString('id-ID')
}

const formatDate = (dateString: string) => {
  if (!dateString) return '-'
  return new Date(dateString).toLocaleDateString('id-ID', {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// Function to convert number to words in Indonesian (Terbilang)
const terbilang = (angka: number): string => {
  if (!angka || angka <= 0) return 'Nol'
  const bilangan = ['', 'Satu', 'Dua', 'Tiga', 'Empat', 'Lima', 'Enam', 'Tujuh', 'Delapan', 'Sembilan', 'Sepuluh', 'Sebelas']
  
  const konversi = (n: number): string => {
    if (n < 12) return bilangan[n]
    if (n < 20) return konversi(n - 10) + ' Belas'
    if (n < 100) return konversi(Math.floor(n / 10)) + ' Puluh ' + konversi(n % 10)
    if (n < 200) return 'Seratus ' + konversi(n - 100)
    if (n < 1000) return konversi(Math.floor(n / 100)) + ' Ratus ' + konversi(n % 100)
    if (n < 2000) return 'Seribu ' + konversi(n - 1000)
    if (n < 1000000) return konversi(Math.floor(n / 1000)) + ' Ribu ' + konversi(n % 1000)
    if (n < 1000000000) return konversi(Math.floor(n / 1000000)) + ' Juta ' + konversi(n % 1000000)
    if (n < 1000000000000) return konversi(Math.floor(n / 1000000000)) + ' Miliar ' + konversi(n % 1000000000)
    return ''
  }

  return konversi(Math.floor(angka)).replace(/\s+/g, ' ').trim()
}

const doPrint = () => {
  window.print()
}

onMounted(async () => {
  try {
    const [invRes, compRes] = await Promise.all([
      api.get(`/api/sales/invoices/${route.params.id}`),
      api.get('/api/settings/company')
    ])
    invoice.value = invRes.data
    company.value = compRes.data || {}
    if (company.value.invoice_footer) {
      customNotes.value = company.value.invoice_footer
    }
  } catch (error) {
    console.error('Failed to load invoice', error)
  } finally {
    loading.value = false
  }
})
</script>

<style>
@media print {
  @page {
    size: A4 portrait;
    margin: 10mm;
  }
  body {
    background: white !important;
    color: black !important;
    font-size: 11pt;
  }
  .print\:hidden {
    display: none !important;
  }
  .print-container {
    padding: 0 !important;
    border: none !important;
    box-shadow: none !important;
    width: 100% !important;
    max-width: 100% !important;
  }
}
</style>
