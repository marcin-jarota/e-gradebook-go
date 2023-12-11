<template>
  <MainLayout>
    <div class="container">
      <h3>Przedmioty</h3>
      <table class="table table-hover mt-2">
        <thead>
          <tr>
            <th scope="col">Przedmiot</th>
            <th scope="col">Nauczyciel prowadzący</th>
            <th scole="col">Kontakt</th>
          </tr>
        </thead>
        <tbody v-if="teacherSubject?.length">
          <tr v-for="row in teacherSubject" :key="row.teacher.surname">
            <th scope="row">
              {{ row.subject.name }}
            </th>
            <td>
              {{ row.teacher.name }}
              {{ row.teacher.surname }}
            </td>
            <td>
              <a :href="'mailto:' + row.teacher.email">{{ row.teacher.email }}</a>
            </td>
          </tr>
        </tbody>
        <tbody v-else>
          <tr>
            <td colspan="4">Brak Przedmiotów</td>
          </tr>
        </tbody>
      </table>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import { useCurrentUser } from '@/composables/useCurrentUser'
import { userResource } from '@/resources/user'
import MainLayout from '@/layouts/MainLayout.vue'
import { classGroupResource } from '@/resources/classGroup'
import { ref } from 'vue'
import type { TeacherSubject } from '@/types/ClassGroup'

const { user } = useCurrentUser()

const teacherSubject = ref<TeacherSubject[]>([])
const getStudentSubjects = async () => {
  const { data: studentData } = await userResource.studentDataByUserID(user.id)
  const { data: subjects } = await classGroupResource.teachersSubjects(studentData.classGroupID)
  teacherSubject.value = subjects
}

getStudentSubjects()
</script>

<style scoped></style>
