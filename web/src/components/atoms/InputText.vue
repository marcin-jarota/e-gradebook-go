<script lang="ts" setup>
defineProps<{
  type: 'text' | 'textarea' | 'email' | 'password'
  label: string
  name: string
  disabled?: boolean
  modelValue?: string
  placeholder?: string
  error?: string
  required?: boolean
}>()
defineEmits(['update:modelValue', 'blur', 'keypress', 'change', 'focusin', 'focusout'])
</script>
<template>
  <div class="">
    <label :for="name" class="pb-2">{{ label }} <span v-if="required" class="text-danger">*</span></label>
    <input class="form-control" :readonly="disabled" :required="required" :placeholder="placeholder" :type="type"
      :value="modelValue" @change="$emit('change', $event)" @focusin="$emit('focusin', $event)"
      @focusout="$emit('focusout', $event)" @keypress="$emit('keypress', $event)" :autocomplete="name + '_autocomplete'"
      @blur="$emit('blur', $event)" @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)" />
    <span v-if="error" class="text-danger d-flex my-1">{{ error }}</span>
  </div>
</template>
