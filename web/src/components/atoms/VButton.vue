<template>
  <button :type="type" class="btn" @clik="$emit('click')" :class="classObject">
    <slot />
  </button>
</template>

<script setup lang="ts">
import { reactive } from 'vue'

defineEmits(['click'])
const props = defineProps<{ variant: 'primary' | 'danger'; type: 'submit' | 'button' | 'reset' }>()

const classObject = reactive({
  'btn-primary': props.variant === 'primary',
  'btn-danger': props.variant === 'danger'
})
</script>

<style scoped></style>
