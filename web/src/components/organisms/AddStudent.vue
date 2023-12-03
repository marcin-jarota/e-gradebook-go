<template>
  <VButton type="button" variant="primary" @click="openModal(modal)">Przypisz ucznia</VButton>
  <Teleport to="body">
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Przypisz ucznia</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body modal-body-inner">
            <div v-if="errorCode" class="alert alert-danger">
              {{ $translate(errorCode) }}
            </div>
            <DynamicSelector :items="studentsToAssigne" @on-select="(s) => (selectedStudent = s)" title="Wybierz ucznia"
              placeholder="Imię lub nazwisko" />
          </div>
          <div class="modal-footer">
            <VButton @click="save" variant="primary" type="button">Przypisz</VButton>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import VButton from '@/components/atoms/VButton.vue'
import { Modal } from 'bootstrap'
import { useSnackbar } from '@/composables/useSnackbar'
import DynamicSelector from '@/components/molecules/DynamicSelector.vue'
import { studentResource } from '@/resources/student'
import { classGroupResource } from '@/resources/classGroup'

const { successSnackbar } = useSnackbar()
const props = defineProps<{ classGroupId: number }>()
const $emit = defineEmits(['on-add'])

const errorCode = ref('')

const studentsToAssigne = ref<{ id: number; name: string }[]>([])
const selectedStudent = ref<{ id: number; name: string } | null>()

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

const getStudents = async () => {
  const { data } = await studentResource.getAll()
  studentsToAssigne.value = data.map(({ id, fullName }) => ({ id, name: fullName }))
}

getStudents()

const save = async () => {
  try {
    if (!selectedStudent.value) return

    await classGroupResource.assignStudent({
      studentID: selectedStudent.value?.id,
      classGroupID: props.classGroupId
    })
    successSnackbar('Student przypisany', 3000)
    $emit('on-add')
    closeModal(modal.value)
  } catch (err) {
    errorCode.value = (err as any).response?.data?.error
  }
}
</script>

<style scoped lang="scss">
.modal-body-inner {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
