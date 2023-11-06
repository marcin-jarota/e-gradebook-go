<template>
  <div aria-live="polite" aria-atomic="true" class="d-flex"
    style="position: fixed; right: 10px; top: 10px; min-height: 200px">
    <TransitionGroup name="list" tag="ul" style="position: relative; top: 0; right: 0">
      <div class="toast show my-2 align-items-center text-bg-success border-0" v-for="s in snackbarStore.snackbars"
        :key="s.id" role="alert" aria-live="assertive" aria-atomic="true">
        <div class="d-flex">
          <div class="toast-body">{{ s.content }}</div>
          <button type="button" @click="snackbarStore.closeSnackbar(s.id)" class="btn-close btn-close-white me-2 m-auto"
            data-bs-dismiss="toast" aria-label="Close"></button>
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
import { useSnackbarStore } from '@/stores/snackbar'

const snackbarStore = useSnackbarStore()
</script>

<style lang="scss">
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
