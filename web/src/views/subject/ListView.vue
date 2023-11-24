<template>
  <MainLayout>
    <div class="container">
      <h2 class="pb-2">Lista przedmiotów</h2>
      <SubjectList :subjects="subjects" @delete-click="openDeleteModal" />
      <VButton @click="openModal(modal)" variant="primary" type="button">Dodaj przedmiot</VButton>
    </div>

    <!-- Modal -->
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Dodaj przedmiot</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-if="errorCode" class="alert alert-danger">
              {{ $translate(errorCode) }}
            </div>
            <InputText v-model="subjectName" required placeholder="Informatyka" type="text" name="subject-name"
              label="Nazwa przedmiotu" />
          </div>
          <div class="modal-footer">
            <VButton @click="saveSubject" variant="primary" type="button">Zapisz</VButton>
          </div>
        </div>
      </div>
    </div>
    <!-- delete modal -->
    <div class="modal fade" ref="deleteModal" tabindex="-1" id="deleteModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Czy jesteś pewien?</h5>
            <button type="button" class="btn-close" data-bs-dismiss="deleteModal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-if="errorCode" class="alert alert-danger">
              {{ $translate(errorCode) }}
            </div>
            <div v-else>
              <span> Oceny powiązane z przedmiotem zostaną zarchiwizowane</span>
            </div>
          </div>
          <div class="modal-footer">
            <VButton @click="deleteSubject" variant="danger" type="button">Usuń</VButton>
          </div>
        </div>
      </div>
    </div>
  </MainLayout>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
import type { Subject } from '@/types/Subject'
import { Modal } from 'bootstrap'
import { subjectResource } from '@/resources/subject'
import MainLayout from '@/layouts/MainLayout.vue'
import SubjectList from '@/components/SubjectList.vue'
import VButton from '@/components/atoms/VButton.vue'
import InputText from '@/components/form/InputText.vue'
import { useSnackbar } from '@/composables/useSnackbar'

const { successSnackbar } = useSnackbar()
const subjects = ref<Subject[]>([])
const subjectName = ref('')
const errorCode = ref('')
const subjectID = ref<number | null>(null)

const modal = ref<HTMLDivElement | null>(null)
const deleteModal = ref<HTMLDivElement | null>(null)

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

const openDeleteModal = (id: number) => {
  openModal(deleteModal.value)
  subjectID.value = id
}

const closeDeleteModal = () => {
  closeModal(deleteModal.value)
  subjectID.value = null
}

const saveSubject = async (e: Event) => {
  try {
    e.preventDefault()
    e.stopPropagation()
    await subjectResource.create({ name: subjectName.value })
    successSnackbar('Przedmiot dodany', 4000)
    closeModal(modal.value)
    await getSubjects()
  } catch (err) {
    const code = (err as any)?.response?.data?.error
    if (code) {
      errorCode.value = code
    }
  }
}

const deleteSubject = async (e: Event) => {
  try {
    if (subjectID.value === null) return
    e.preventDefault()
    e.stopPropagation()
    await subjectResource.delete(subjectID.value)

    closeDeleteModal()
    successSnackbar('Przedmiot usunięty', 4000)

    await getSubjects()
  } catch (err) {
    const code = (err as any)?.response?.data?.error
    if (code) {
      errorCode.value = code
    }
  }
}

const getSubjects = async () => {
  const { data } = await subjectResource.list()
  subjects.value = data
}

getSubjects()
</script>

<style></style>
