<template>
  <table class="table">
    <thead>
      <tr>
        <th scope="col">ID</th>
        <th scope="col">Nazwa Przedmiotu</th>
        <th scope="col">Akcje</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in list" :key="item.id">
        <th scope="row">{{ subject.id }}</th>
        <td>{{ subject.name }}</td>
        <td>
          <VButton variant="danger" type="button" @click="$emit('deleteClick', subject.id)">Usuń</VButton>
        </td>
      </tr>
    </tbody>
  </table>
</template>
<script setup lang="ts" generic="T extends { id: number }">
import VButton from '@/components/atoms/VButton.vue'
defineProps<{ list: T[] }>()
defineEmits<{ (e: 'deleteClick', id: number): void }>()
</script>
