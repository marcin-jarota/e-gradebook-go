<script setup lang="ts">
import MainSidebar from '../components/MainSidebar.vue'
</script>

<template>
  <div class="d-flex">
    <aside>
      <MainSidebar />
    </aside>
    <main class="p-4">
      <slot />
    </main>
  </div>
</template>
