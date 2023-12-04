<template>
  <table class="table">
    <thead>
      <tr>
        <th scope="col">ID</th>
        <th scope="col">Nazwa Przedmiotu</th>
        <th scope="col">Nauczyciele uczący</th>
        <th scope="col">Akcje</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="subject in subjects" :key="subject.id">
        <th scope="row">{{ subject.id }}</th>
        <td>{{ subject.name }}</td>
        <td>
          <span v-for="teacher in subject.teachers" :key="teacher.id + 't'">{{ teacher.name }} {{ teacher.surname }},
          </span>
        </td>
        <td>
          <AssignTeacherSubject :subject-id="subject.id" @save-success="$emit('refresh-request')" />

          <VButton variant="danger" class="ms-2" type="button" @click="$emit('deleteClick', subject.id)">
            <font-awesome-icon icon="fa-solid fa-trash" /></VButton>
        </td>
      </tr>
    </tbody>
  </table>
</template>
<script setup lang="ts">
import type { Subject } from '@/types/Subject'
import VButton from './atoms/VButton.vue'
import AssignTeacherSubject from './organisms/AssignTeacherSubject.vue'
defineProps<{ subjects: Subject[] }>()
defineEmits<{ (e: 'deleteClick', id: number): void; (e: 'refresh-request'): void }>()
</script>
