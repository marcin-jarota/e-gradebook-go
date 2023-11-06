<template>
  <MainLayout>
    <div class="container">
      <h2>Dodawanie studenta</h2>
      <form>
        <InputText v-model="email" type="email" name="student-email" label="Email" />
      </form>
    </div>
  </MainLayout>
</template>
<script setup lang="ts">
import InputText from '@/components/form/InputText.vue'
import MainLayout from '@/layouts/MainLayout.vue'
import { ref } from 'vue'

const email = ref('')
</script>
