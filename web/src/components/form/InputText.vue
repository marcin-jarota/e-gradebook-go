<script lang="ts" setup>
defineProps<{
  type: 'text' | 'textarea' | 'email' | 'password'
  label: string
  name: string
  modelValue: string
  placeholder?: string
}>()
defineEmits(['update:modelValue'])
</script>
<template>
  <div class="mb-">
    <label :for="name" class="pb-2">{{ label }}</label>
    <input class="form-control" :placeholder="placeholder" :name="name" :type="type" :autocomplete="'textinput_' + name"
      :value="modelValue" @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)" />
  </div>
</template>
