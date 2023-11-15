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
      <tr v-for="subject in subjects" :key="subject.id">
        <th scope="row">{{ subject.id }}</th>
        <td>{{ subject.name }}</td>
        <td>
          <VButton variant="danger" type="button" @click="$emit('deleteClick', subject.id)">Usuń</VButton>
        </td>
      </tr>
    </tbody>
  </table>
</template>
<script setup lang="ts">
import type { Subject } from '@/types/Subject'
import VButton from './atoms/VButton.vue'
defineProps<{ subjects: Subject[] }>()
defineEmits<{ (e: 'deleteClick', id: number): void }>()
</script>
