<template>
  <MainLayout>
    <div class="container">
      <h3>Spis lat edukacyjnych</h3>
      <table class="table table-hover">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Rok szkolny</th>
            <th scope="col">Liczba klas</th>
            <th scope="col">Rozpoczęcie</th>
            <th scope="col">Zakończenie</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="schoolYear in schoolYears" :key="schoolYear.id">
            <th scope="row">
              {{ schoolYear.id }}
            </th>
            <td>
              {{ schoolYear.name }}
            </td>
            <td>
              {{ schoolYear.classGroupsCount }}
            </td>
            <td>{{ formatDateToMMDDYYYY(new Date(schoolYear.start)) }}</td>
            <td>{{ formatDateToMMDDYYYY(new Date(schoolYear.end)) }}</td>
          </tr>
        </tbody>
      </table>

      <OpenSchoolYear @on-add="fetchYears" />
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import VButton from '@/components/atoms/VButton.vue'
import OpenSchoolYear from '@/components/organisms/OpenSchoolYear.vue'
import { schoolYearResource } from '@/resources/schoolYear'
import type { SchoolYearDetailed } from '@/types/SchoolYear'
import { ref } from 'vue'

const schoolYears = ref<SchoolYearDetailed[]>([])

const fetchYears = async () => {
  const { data } = await schoolYearResource.list()
  schoolYears.value = data
}

function formatDateToMMDDYYYY(date: Date) {
  const month = String(date.getMonth() + 1).padStart(2, '0') // Months are zero-based
  const day = String(date.getDate()).padStart(2, '0')
  const year = date.getFullYear()

  return `${month}-${day}-${year}`
}

fetchYears()
</script>

<style scoped></style>
