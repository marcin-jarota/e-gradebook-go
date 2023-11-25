<template>
  <MainLayout>
    <div class="container">
      <div class="row">
        <div class="col-8">
          <h2 class="pb-2">{{ details?.name }}</h2>
          <table class="table table-hover">
            <thead>
              <tr>
                <th scope="col">ID</th>
                <th scope="col">Imię</th>
                <th>Nazwisko</th>
                <th scope="col">Średnia ocen</th>
                <th></th>
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
                <td>{{ student.avgMark || 'brak ocen' }}</td>
                <td>
                  <AddMark @on-add="refresh" :studentID="student.id" />
                </td>
              </tr>
            </tbody>
            <tbody v-else>
              <tr>
                <td colspan="4">Brak uczniów w klasie</td>
              </tr>
            </tbody>
          </table>
          <AddStudent :class-group-id="classGroupID" @on-add="getClassGroupStudents" />
        </div>
        <div class="col-4">
          <h3 class="pb-4">Oceny w klasie</h3>
          <apex-chart width="500" type="pie" :options="options" :series="series"></apex-chart>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import { classGroupResource } from '@/resources/classGroup'
import type { ClassGroupOutput, ClassGroupStudent } from '@/types/ClassGroup'
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'
import AddMark from '@/components/organisms/AddMark.vue'
import AddStudent from '@/components/organisms/AddStudent.vue'

const route = useRoute()
const students = ref<ClassGroupStudent[]>()
const details = ref<ClassGroupOutput>()
const marks = ref<{ value: number; id: number }[]>()
const series = ref<number[]>([])
const labels = ref<string[]>([])

const options = computed(() => ({
  chart: {
    width: 380,
    type: 'pie'
  },
  labels: labels.value,
  responsive: [
    {
      breakpoint: 480,
      options: {
        chart: {
          width: 200
        },
        legend: {
          position: 'bottom'
        }
      }
    }
  ]
}))

const refresh = () => {
  getClassGroupStudents()
  getMarks()
}

const classGroupID = computed(() => Number(route.params.id))

const getClassGroupStudents = async () => {
  const { data } = await classGroupResource.students(classGroupID.value)

  students.value = data
}

const getClassGroupDetails = async () => {
  const { data } = await classGroupResource.getOne(classGroupID.value)

  details.value = data
}

const getMarks = async () => {
  const { data } = await classGroupResource.getMarks(classGroupID.value)
  marks.value = data

  const x = data.reduce(
    (acc, mark) => {
      const value = acc[mark.value]
      if (value === undefined) {
        acc[mark.value] = 1
      } else {
        acc[mark.value] += 1
      }

      return acc
    },
    {} as Record<number, number>
  )

  series.value = []
  labels.value = []
  Object.entries(x).forEach(([key, value]) => {
    series.value = [...series.value, value]
    labels.value = [...labels.value, `Ocena ${key}`]
  })
  console.log(series.value)
  console.log(labels.value)
}

getClassGroupStudents()
getClassGroupDetails()
getMarks()
</script>

<style scoped></style>
