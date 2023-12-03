<template>
  <button :type="type" class="btn" @click="$emit('click', $event)" :class="classObject">
    <slot />
  </button>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

defineEmits(['click'])
const props = defineProps<{
  variant: 'primary' | 'danger' | 'secondary' | 'light'
  type: 'submit' | 'button' | 'reset'
}>()

const classObject = reactive({
  'btn-primary': props.variant === 'primary',
  'btn-danger': props.variant === 'danger',
  'btn-secondary': props.variant === 'secondary',
  'btn-light': props.variant === 'light'
})
</script>

<style scoped></style>
