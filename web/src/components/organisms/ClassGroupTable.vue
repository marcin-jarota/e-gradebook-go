<template>
  <table class="table table-hover">
    <thead>
      <tr>
        <th scope="col">ID</th>
        <th scope="col">Nazwa klasy</th>
        <th>Rok edukacji</th>
        <th>Liczba uczniów</th>
        <th scope="col">Akcje</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="classgroup in classGroups" :key="classgroup.id">
        <th scope="row">
          {{ classgroup.id }}
        </th>
        <td>
          <RouterLink :to="{ name: routes.classGroup.name, params: { id: classgroup.id } }">{{
            classgroup.name
          }}</RouterLink>
        </td>
        <td>{{ classgroup.educationYear }}</td>
        <td>{{ classgroup.studentsCount }}</td>
        <td>
          <VButton variant="danger" type="button" @click="$emit('deleteClick', classgroup.id)">Usuń</VButton>
          <RouterLink class="ms-3" :to="{ name: routes.classGroupLessons.name, params: { id: classgroup.id } }">Plan zajęć
          </RouterLink>
        </td>
      </tr>
    </tbody>
  </table>
</template>
<script setup lang="ts">
import type { ClassGroupOutput } from '@/types/ClassGroup'
import VButton from '@/components/atoms/VButton.vue'
import { routes } from '@/router'
defineProps<{ classGroups: ClassGroupOutput[] }>()
defineEmits<{ (e: 'deleteClick', id: number): void }>()
</script>
