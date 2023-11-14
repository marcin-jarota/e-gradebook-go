<template>
  <MainLayout>
    <div class="container">
      <h2 class="pb-2">Lista przedmiotów</h2>
      <SubjectList :subjects="subjects" />
      <VButton @click="openModal" variant="primary" type="button">Dodaj przedmiot</VButton>
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

const modal = ref<HTMLDivElement | null>(null)

const openModal = () => {
  if (modal.value) {
    const bModal = new Modal(modal.value)
    bModal.show()
  }
}

const closeModal = () => {
  if (modal.value) {
    const bModal = Modal.getInstance(modal.value)
    if (bModal) {
      bModal.hide()
    }
  }
}

const saveSubject = async (e: Event) => {
  try {
    e.preventDefault()
    e.stopPropagation()
    await subjectResource.create({ name: subjectName.value })
    successSnackbar('Przedmiot dodany', 4000)
    closeModal()
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
  subjects.value = data.data
}

getSubjects()
</script>

<style></style>
