<script lang="ts" setup>
import { useCurrentUser } from '@/composables/useCurrentUser'
import MainLayout from '@/layouts/MainLayout.vue'
import { studentResource } from '@/resources/student'
import { userResource } from '@/resources/user'
import { markSimpleValue, markColor, type StudentMark } from '@/types/Mark'
import { Tooltip } from 'bootstrap'
import { computed, ref } from 'vue'

const { user } = useCurrentUser()

const marks = ref<StudentMark[]>([])

const tableMarks = computed(() => {
  const bySubject = marks.value.reduce(
    (acc, item) => {
      const { value, comment, teacher, subject } = item
      const { id, name } = subject

      const newItem = {
        value,
        comment,
        teacher,
        tooltip: buildTooltip(teacher, comment)
      }

      if (!acc[name]) {
        acc[name] = { id, name, items: [newItem] }
      } else {
        acc[name].items.push(newItem)
      }

      return acc
    },
    {} as Record<string, { id: number; name: string; items: any[] }>
  )

  return Object.values(bySubject).map((el) => ({ ...el, avg: calculateAvg(el.items) }))
})

const buildTooltip = (teacher: any, comment: string) => {
  let base = `Nauczyciel: ${teacher.name} ${teacher.surname}`

  if (comment) base += `<br> Komentarz: ${comment}`

  return base
}
const bindTooltips = () => {
  const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle]'))
  tooltipTriggerList.map(function (tooltipTriggerEl) {
    return new Tooltip(tooltipTriggerEl)
  })
}

const calculateAvg = (marks: { value: number }[]) => {
  const sum = marks.reduce((acc, curr) => (acc += curr.value), 0)
  const avg = sum / marks.length
  return Math.round(avg * 100) / 100
}

const getStudentMarks = async () => {
  const { data: studentData } = await userResource.studentDataByUserID(user.id)
  const { data } = await studentResource.getMarks(studentData.studentID)

  marks.value = data

  setTimeout(() => bindTooltips(), 200)
}

getStudentMarks()
</script>

<template>
  <MainLayout>
    <div class="container">
      <h2>Lista ocen</h2>
    </div>
    <div class="container mt-4">
      <table class="table">
        <thead class="table-light">
          <tr>
            <th scope="col">Nazwa przedmiotu</th>
            <th scope="col">Oceny</th>
            <th scope="col">Średnia</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="subject in tableMarks" :key="subject.id">
            <td>{{ subject.name }}</td>
            <td>
              <span v-for="mark in subject.items" :key="mark.id" data-bs-toggle="tooltip" data-bs-placement="top"
                data-bs-html="true" :title="mark.tooltip" class="badge" :class="markColor[mark.value]">{{
                  markSimpleValue[mark.value] }}</span>
            </td>
            <td>{{ subject.avg }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </MainLayout>
</template>

<style lang="scss" scoped>
.tooltip-inner {
  white-space: pre-line;
  /* Wstawiałby nową linię, jeśli napotka '\n' */
}

/* Dodatkowe style dla responsywności tabeli */
@media (max-width: 575.98px) {

  th,
  td {
    padding: 0.75rem 0.5rem;
  }
}

/* Dodatkowy styl dla odstępów między ocenami */
.badge {
  margin-right: 5px;
}
</style>
