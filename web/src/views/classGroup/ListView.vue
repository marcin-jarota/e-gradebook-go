<template>
  <MainLayout>
    <div class="container">
      <h2 class="pb-2">Lista klas</h2>
      <ClassGroupTable :class-groups="classGroups" @delete-click="openDeleteModal" />
      <VButton @click="openModal(modal)" variant="primary" type="button">Dodaj klasę</VButton>
    </div>

    <!-- Modal -->
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Dodaj klasę</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-if="errorCode" class="alert alert-danger">
              {{ $translate(errorCode) }}
            </div>
            <InputText v-model="classGroupName" required placeholder="2c Informatyczna" type="text" name="class-name"
              label="Nazwa klasy" />
          </div>
          <div class="modal-footer">
            <VButton @click="saveClassGroup" variant="primary" type="button">Zapisz</VButton>
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
              <span> Osoby przypisane do klasy zostaną z niej wypisane</span>
            </div>
          </div>
          <div class="modal-footer">
            <!-- <VButton @click="deleteSubject" variant="danger" type="button">Usuń</VButton> -->
          </div>
        </div>
      </div>
    </div>
  </MainLayout>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
import { Modal } from 'bootstrap'
import { classGroupResource } from '@/resources/classGroup'
import MainLayout from '@/layouts/MainLayout.vue'
import VButton from '@/components/atoms/VButton.vue'
import InputText from '@/components/form/InputText.vue'
import { useSnackbar } from '@/composables/useSnackbar'
import type { ClassGroupOutput } from '@/types/ClassGroup'
import ClassGroupTable from '@/components/organisms/ClassGroupTable.vue'

const { successSnackbar } = useSnackbar()
const classGroups = ref<ClassGroupOutput[]>([])
const classGroupName = ref('')
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

const saveClassGroup = async (e: Event) => {
  try {
    e.preventDefault()
    e.stopPropagation()
    await classGroupResource.create({ name: classGroupName.value })
    successSnackbar('Klasa stworzona', 3000)
    closeModal(modal.value)
    await getClassGroups()
  } catch (err) {
    const code = (err as any)?.response?.data?.error
    if (code) {
      errorCode.value = code
    }
  }
}
//
// const deleteSubject = async (e: Event) => {
//   try {
//     e.preventDefault()
//     e.stopPropagation()
//     await subjectResource.delete(subjectID.value as number)
//
//     closeDeleteModal()
//     successSnackbar('Przedmiot usunięty', 4000)
//
//     await getSubjects()
//   } catch (err) {
//     const code = (err as any)?.response?.data?.error
//     if (code) {
//       errorCode.value = code
//     }
//   }
// }

const getClassGroups = async () => {
  const { data } = await classGroupResource.list()
  classGroups.value = data.data
}

getClassGroups()
</script>

<style></style>
