<template>
  <MainLayout>
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Dodaj lekcje</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="p-3">
            <DynamicSelector :items="subjects" @on-select="handleSubjectSelect" title="Wybierz przedmiot"
              placeholder="Informatyka" />
            <div class="py-2" v-if="selectedSubject">
              Prowadzący: {{ selectedSubject.teacher.name }} {{ selectedSubject.teacher.surname }}
            </div>
            <div class="form-group pt-2">
              <label for="startTime" class="pb-2">Dzień tygodnia</label>
              <select class="form-select" v-model="dayOfWeek" aria-label="Wybierz dzień tygodnia">
                <option value="1">Poniedziałek</option>
                <option value="2">Wtorek</option>
                <option value="3">Środa</option>
                <option value="4">Cwartek</option>
                <option value="5">Piątek</option>
              </select>
            </div>

            <div class="form-group pt-2">
              <label for="startTime" class="pb-2">Godzina rozpoczęcia</label>
              <input type="time" id="startTime" placeholder="" class="form-control" v-model="startTime" />
            </div>
            <div class="form-group pt-2">
              <label for="endTime" class="pb-2">Godzina zakończenia</label>
              <input type="time" id="endTime" class="form-control" placeholder="" v-model="endTime" />
            </div>
            <VButton type="button" class="mt-4" variant="primary" @click="save">Zapisz</VButton>
          </div>
        </div>
      </div>
    </div>
    <div class="container d-flex justify-content-between mb-4">
      <h3>Plan zajęć</h3>
      <VButton variant="light" type="button" @click="openModal(modal)">Dodaj lekcję</VButton>
    </div>
    <LessonsPlan :lessons="classEvents" />
  </MainLayout>
</template>

<script setup lang="ts">
import VButton from '@/components/atoms/VButton.vue'
import DynamicSelector from '@/components/molecules/DynamicSelector.vue'
import MainLayout from '@/layouts/MainLayout.vue'
import { classGroupResource } from '@/resources/classGroup'
import { lessonResource } from '@/resources/lesson'
import type { Lesson, TeacherSubject } from '@/types/ClassGroup'
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import LessonsPlan from '@/components/organisms/LessonsPlan.vue'
import { Modal } from 'bootstrap'
import { useSnackbar } from '@/composables/useSnackbar'

const route = useRoute()

const { successSnackbar, errorSnackbar } = useSnackbar()
const subjects = ref<{ id: number; name: string }[]>([])
const teacherSubject = ref<TeacherSubject[]>([])
const selectedSubject = ref<TeacherSubject | null>()
const startTime = ref<string>()
const endTime = ref<string>()
const dayOfWeek = ref<string>()

const modal = ref<HTMLDivElement | null>(null)

const openModal = (e: HTMLDivElement | null) => {
  if (e) {
    const bModal = new Modal(e)
    bModal.show()
  }
}

const classEvents = ref<Lesson[]>([])

const handleSubjectSelect = (subject: { id: number; name: string }) => {
  if (!subject) {
    selectedSubject.value = null
  } else {
    selectedSubject.value = teacherSubject.value.find((sub) => sub.subject.id === subject.id)
  }
}
const fetchSubjects = async () => {
  const { data } = await classGroupResource.teachersSubjects(Number(route.params.id))
  teacherSubject.value = data
  subjects.value = data.map((el) => ({
    id: el.subject.id,
    name: el.subject.name
  }))
}

const fetchLessons = async () => {
  const { data } = await classGroupResource.lessons(Number(route.params.id))

  classEvents.value = data
}

fetchLessons()

const save = async () => {
  try {
    if (
      !selectedSubject.value ||
      !selectedSubject.value.teacher ||
      !selectedSubject.value.subject ||
      !startTime.value ||
      !endTime.value ||
      !dayOfWeek.value ||
      isNaN(parseInt(dayOfWeek.value))
    )
      return
    await lessonResource.create({
      teacherID: selectedSubject.value.teacher.id,
      subjectID: selectedSubject.value.subject.id,
      classGroupID: Number(route.params.id),
      startTime: startTime.value,
      endTime: endTime.value,
      dayOfWeek: parseInt(dayOfWeek.value)
    })

    successSnackbar('Lekcja dodana pomyślnie', 3000)
  } catch (err) {
    errorSnackbar('Nie udało dodać się lekcji', 3000)
    alert(err)
  }

  try {
    await fetchLessons()
  } catch (err) {
    errorSnackbar('Nie udało się odświeżyć widoku lekcji')
  }
}
fetchSubjects()
</script>

<style scoped lang="scss">
.calendar-grid {
  position: relative;
}

.hour-slot {
  height: 60px;
  /* Adjust as needed */
  border-bottom: 1px solid #ddd;
}

.time-slot {
  height: 60px;
  /* Should match .hour-slot height */
  text-align: right;
  padding-right: 10px;
  display: flex;
  align-items: center;
  border-bottom: 1px solid #ddd;
}
</style>
