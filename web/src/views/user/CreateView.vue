<template>
  <MainLayout>
    <div class="container">
      <h3>Dodaj użytkownika</h3>
      <UserForm />
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import UserForm from '@/components/organisms/UserForm.vue'
</script>

<style scoped></style>
