<template>
  <MainLayout>
    <h3>Plan zajęć</h3>
    <LessonsPlan :lessons="lessons" />
  </MainLayout>
</template>

<script setup lang="ts">
import LessonsPlan from '@/components/organisms/LessonsPlan.vue'
import { useCurrentUser } from '@/composables/useCurrentUser'
import MainLayout from '@/layouts/MainLayout.vue'
import { classGroupResource } from '@/resources/classGroup'
import { userResource } from '@/resources/user'
import type { Lesson } from '@/types/ClassGroup'
import { ref } from 'vue'

const { user } = useCurrentUser()

const lessons = ref<Lesson[]>([])

const fetchLessons = async () => {
  const { data: student } = await userResource.studentDataByUserID(user.id)
  const { data: lessonsList } = await classGroupResource.lessons(student.classGroupID)

  lessons.value = lessonsList
}

fetchLessons()
</script>

<style scoped></style>
