<template>
  <MainLayout>
    <div class="container-lg">
      <h2>Lista użytkowników w systemie</h2>
      <UsersTable />
    </div>
  </MainLayout>
</template>

<script lang="ts" setup>
import MainLayout from '@/layouts/MainLayout.vue'
import UsersTable from '@/components/organisms/UsersTable.vue'
</script>

<style lang="scss" scoped></style>
