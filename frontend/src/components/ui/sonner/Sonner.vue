<script lang="ts" setup>
import { Toaster as Sonner, type ToasterProps } from 'vue-sonner'

const props = withDefaults(defineProps<ToasterProps>(), {
  position: 'top-right',
  richColors: true,
  closeButton: true,
  duration: 4000,
  visibleToasts: 5,
})
</script>

<template>
  <Sonner
    v-bind="props"
    class="toaster group"
    :toast-options="{
      classes: {
        toast: 'group toast group-[.toaster]:bg-white group-[.toaster]:text-slate-900 group-[.toaster]:border-slate-200 group-[.toaster]:shadow-lg group-[.toaster]:rounded-xl group-[.toaster]:font-sans text-sm font-medium',
        description: 'group-[.toast]:text-slate-500 text-xs font-normal mt-0.5',
        actionButton: 'group-[.toast]:bg-primary group-[.toast]:text-primary-foreground text-xs font-medium rounded-lg px-2.5 py-1',
        cancelButton: 'group-[.toast]:bg-slate-100 group-[.toast]:text-slate-600 text-xs font-medium rounded-lg px-2.5 py-1',
        closeButton: 'group-[.toast]:bg-white group-[.toast]:text-slate-400 hover:group-[.toast]:text-slate-700 group-[.toast]:border-slate-200',
        error: 'group-[.toaster]:!bg-rose-50 group-[.toaster]:!text-rose-900 group-[.toaster]:!border-rose-200',
        success: 'group-[.toaster]:!bg-emerald-50 group-[.toaster]:!text-emerald-900 group-[.toaster]:!border-emerald-200',
        warning: 'group-[.toaster]:!bg-amber-50 group-[.toaster]:!text-amber-900 group-[.toaster]:!border-amber-200',
        info: 'group-[.toaster]:!bg-sky-50 group-[.toaster]:!text-sky-900 group-[.toaster]:!border-sky-200',
      },
      style: {
        zIndex: 99999,
      },
      ...props.toastOptions,
    }"
  />
</template>
