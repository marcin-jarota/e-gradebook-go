<template>
  <font-awesome-icon :class="{ 'icon--active': active, 'icon--inactive': !active }"
    :icon="['fa-solid', active ? 'fa-check' : 'fa-xmark']" />
</template>

<script setup lang="ts">
defineProps<{ active: boolean }>()
</script>

<style lang="scss">
.icon {
  font-size: 24px;

  &--active {
    color: green;
  }

  &--inactive {
    color: red;
  }
}
</style>
