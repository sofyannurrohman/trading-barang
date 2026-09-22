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
      { path: 'inventory/management', name: 'StockManagement', component: () => import('../views/dashboard/inventory/StockManagement.vue') },
      
      // Penjualan & Faktur (Pusat Transaksi & Faktur Terpadu)
      { path: 'sales/management', name: 'SalesManagement', component: () => import('../views/dashboard/sales/SalesManagement.vue') },
      { path: 'sales/create', redirect: '/dashboard/sales/management' },
      { path: 'sales/list', redirect: '/dashboard/sales/management' },
      
      // Laporan
      { path: 'reports/profit', name: 'ProfitReport', component: () => import('../views/dashboard/reports/ProfitReport.vue') },
      { path: 'reports/tax', redirect: '/dashboard/reports/profit' },
      
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
