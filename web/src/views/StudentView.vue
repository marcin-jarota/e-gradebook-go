<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import { useCounterStore } from '@/stores/counter'

const { count } = useCounterStore()
</script>
<template>
  <MainLayout>
    <div class="login">
      <h1>Student {{ count }}</h1>
    </div>
  </MainLayout>
</template>

<style lang="scss">
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
