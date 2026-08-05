<script setup lang="ts">
import { Toaster } from 'vue-sonner'
</script>

<template>
  <router-view />
  <Toaster
    position="top-right"
    richColors
    :toastOptions="{ style: { zIndex: '99999' } }"
  />
</template>
