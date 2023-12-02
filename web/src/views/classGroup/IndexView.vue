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
          <AddStudent :class-group-id="classGroupID" @on-add="refresh" />
        </div>
        <div class="col-4">
          <h3 class="pb-4">Oceny w klasie</h3>
          <apex-chart width="500" type="pie" :options="options.settings" :series="options.series"></apex-chart>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import { classGroupResource } from '@/resources/classGroup'
import type { ClassGroupOutput, ClassGroupStudent } from '@/types/ClassGroup'
import { computed, reactive, ref } from 'vue'
import { useRoute } from 'vue-router'
import AddMark from '@/components/organisms/AddMark.vue'
import AddStudent from '@/components/organisms/AddStudent.vue'

const route = useRoute()
const students = ref<ClassGroupStudent[]>()
const details = ref<ClassGroupOutput>()
const marks = reactive<{ list: { value: number; id: number }[] }>({ list: [] })

const options = computed(() => {
  const val = marks.list.reduce(
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

  const series: number[] = []
  const labels: string[] = []

  if (val) {
    Object.entries(val)
      .sort(([keyA], [keyB]) => keyA - keyB)
      .forEach(([key, value]) => {
        series.push(value)
        labels.push(`${marksMap[key] || key}`)
      })
  }

  return {
    series,
    settings: {
      chart: {
        width: 380,
        type: 'pie'
      },
      labels,
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
    }
  }
})

const refresh = async () => {
  await getClassGroupStudents()
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

const marksMap = {
  1: 'niedotateczny',
  1.5: 'niedostateczny +',
  2: 'dopuszczający',
  3: 'dostateczny',
  3.5: 'dostateczny +',
  4: 'dobry',
  4.5: 'dobry +',
  5: 'bardzo dobry',
  5.5: 'bardzo dobry +',
  6: 'celujący'
}
const getMarks = async () => {
  const { data } = await classGroupResource.getMarks(classGroupID.value)
  marks.list = data
}

getClassGroupStudents()
getClassGroupDetails()
getMarks()
</script>

<style scoped></style>
