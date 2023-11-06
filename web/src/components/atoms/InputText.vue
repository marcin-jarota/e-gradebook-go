<script lang="ts" setup>
defineProps<{
  type: 'text' | 'textarea' | 'email' | 'password'
  label: string
  name: string
  modelValue: string
  placeholder?: string
  error?: string
  required?: boolean
}>()
defineEmits(['update:modelValue', 'keypress', 'keydown'])
</script>
<template>
  <div class="mb-3">
    <label :for="name" class="pb-2">{{ label }} <span v-if="required" class="text-danger">*</span></label>
    <input class="form-control" :required="required" :placeholder="placeholder" :type="type" :value="modelValue"
      @keydown="$emit('keydown', $event)" @keypress="$emit('keypress', $event)" :autocomplete="name + '_autocomplete'"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)" />
    <span v-if="error">{{ error }}</span>
  </div>
</template>
