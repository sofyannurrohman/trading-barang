import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { toast } from 'vue-sonner'

const routes = [
  {
    path: '/',
    redirect: '/dashboard'
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/dashboard',
    component: () => import('../components/layouts/DashboardLayout.vue'),
    meta: { requiresAuth: true, role: 'admin' }, // We'll keep admin requirement or relax it
    children: [
      { path: '', name: 'DashboardOverview', component: () => import('../views/dashboard/Overview.vue') },
      
      // Master Data
      { path: 'master-data/products', name: 'Products', component: () => import('../views/dashboard/master/Products.vue') },
      { path: 'master-data/partners', name: 'Partners', component: () => import('../views/dashboard/master/Partners.vue') },
      
      // Gudang & Inventori
      { path: 'inventory/inbound', name: 'Inbound', component: () => import('../views/dashboard/inventory/Inbound.vue') },
      { path: 'inventory/adjustments', name: 'Adjustments', component: () => import('../views/dashboard/inventory/Adjustments.vue') },
      { path: 'inventory/stock-card', name: 'StockCard', component: () => import('../views/dashboard/inventory/StockCard.vue') },
      
      // Penjualan & Faktur
      { path: 'sales/create', name: 'CreateInvoice', component: () => import('../views/dashboard/sales/CreateInvoice.vue') },
      { path: 'sales/list', name: 'InvoiceList', component: () => import('../views/dashboard/sales/InvoiceList.vue') },
      
      // Laporan
      { path: 'reports/profit', name: 'ProfitReport', component: () => import('../views/dashboard/reports/ProfitReport.vue') },
      { path: 'reports/tax', name: 'TaxReport', component: () => import('../views/dashboard/reports/TaxReport.vue') },
      
      // Pengaturan
      { path: 'settings/company', name: 'CompanySettings', component: () => import('../views/dashboard/settings/CompanyProfile.vue') },
      { path: 'settings/users', name: 'UsersManagement', component: () => import('../views/dashboard/settings/UsersManagement.vue') }
    ]
  },
  {
    path: '/print/invoice/:id',
    name: 'PrintInvoice',
    component: () => import('../views/print/InvoicePrint.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore()
  const isAuthenticated = authStore.isAuthenticated()
  const user = authStore.user

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'Login' })
  } else if (to.meta.role && user?.role !== to.meta.role) {
    toast.error("Anda tidak memiliki akses ke halaman ini. Silakan login kembali.")
    authStore.logout()
    next({ name: 'Login' })
  } else {
    next()
  }
})

export default router
