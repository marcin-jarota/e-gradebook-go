<script setup lang="ts">
import MainSidebar from '@/components/organisms/MainSidebar.vue'
import Snackbar from '@/components/common/SnackbarMessage.vue'
</script>

<template>
  <div class="d-flex">
    <aside>
      <MainSidebar />
    </aside>
    <main class="p-4 container-fluid">
      <slot />
    </main>
    <Snackbar />
  </div>
</template>
