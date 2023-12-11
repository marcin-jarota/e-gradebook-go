<template>
  <VButton type="button" variant="primary" @click="openModal(modal)">Przypisz przedmiot</VButton>
  <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Przypisz przedmiot i nauczyciela</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body modal-body-inner">
          <DynamicSelector :items="subjects" @on-select="(s) => {
              selectedSubject = s
              selectedTeacher = null
            }
            " title="Wybierz przedmiot" placeholder="Informatyka" />
          <DynamicSelector :items="selectedSubject?.teachers || []" @on-select="(s) => (selectedTeacher = s)"
            title="Wybierz nauczyciela prowadzącego" placeholder="Imię lub nazwisko" />
        </div>
        <div class="modal-footer">
          <VButton @click="save" variant="primary" type="button">Przypisz</VButton>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import DynamicSelector from '../molecules/DynamicSelector.vue'
import type { Subject } from '@/types/Subject'
import { subjectResource } from '@/resources/subject'
import { classGroupResource } from '@/resources/classGroup'
import { Modal } from 'bootstrap'
import VButton from '@/components/atoms/VButton.vue'
import { useSnackbar } from '@/composables/useSnackbar'

const props = defineProps<{ classGroupId: number }>()
const $emit = defineEmits(['save-success'])
const subjects = ref<Subject[]>([])
const selectedSubject = ref<Subject>()
const selectedTeacher = ref<{ id: number; name: string; surname: string } | null>()
const { errorSnackbar, successSnackbar } = useSnackbar()

const modal = ref<HTMLDivElement | null>(null)

const openModal = (e: HTMLDivElement | null) => {
  if (e) {
    const bModal = new Modal(e)
    bModal.show()
  }
}

const closeModal = (e: HTMLDivElement | null) => {
  if (e) {
    const bModal = Modal.getInstance(e)
    if (bModal) {
      bModal.hide()
    }
  }
}

const fetchSubjects = async () => {
  const { data } = await subjectResource.list()
  subjects.value = data
}

const save = async () => {
  try {
    if (!selectedTeacher.value || !selectedSubject.value) return
    await classGroupResource.assignTeacherSubject({
      classGroupID: props.classGroupId,
      teacherID: selectedTeacher.value?.id,
      subjectID: selectedSubject.value.id
    })
    closeModal(modal.value)
    successSnackbar('Przedmiot przypisany pomyślnie', 3000)
    $emit('save-success')
  } catch (err) {
    console.log(err)
    errorSnackbar('Nie udało się przypisać przedmiotu', 3000)
  }
}

fetchSubjects()
</script>

<style scoped></style>
