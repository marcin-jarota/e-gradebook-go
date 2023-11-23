<template>
  <MainLayout>
    <div class="container">
      <h2 class="pb-2">Klasa {{ details?.name }}</h2>
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Imię</th>
            <th>Nazwisko</th>
            <th scope="col">Średnia ocen</th>
          </tr>
        </thead>
        <tbody v-if="students?.length">
          <tr v-for="student in students" :key="student.id">
            <th scope="row">
              {{ student.id }}
            </th>
            <td>
              {{ student.name }}
            </td>
            <td>{{ student.surname }}</td>
            <td>{{ student.avgMark }}</td>
          </tr>
        </tbody>
        <tbody v-else>
          <tr>
            <td colspan="4">Brak uczniów w klasie</td>
          </tr>
        </tbody>
      </table>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import { classGroupResource } from '@/resources/classGroup'
import type { ClassGroupOutput, ClassGroupStudent } from '@/types/ClassGroup'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const students = ref<ClassGroupStudent[]>()
const details = ref<ClassGroupOutput>()
const getClassGroup = async () => {
  const { data } = await classGroupResource.students(Number(route.params.id))

  students.value = data
}

const getClassGroupDetails = async () => {
  const { data } = await classGroupResource.getOne(Number(route.params.id))

  details.value = data
}

getClassGroup()
getClassGroupDetails()
</script>

<style scoped></style>
