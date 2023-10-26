<script setup lang="ts">
import MainSidebar from '../components/MainSidebar.vue'
</script>

<template>
  <div class="d-flex">
    <aside>
      <MainSidebar />
    </aside>
    <main class="p-4 container-fluid">
      <slot />
    </main>
  </div>
</template>
