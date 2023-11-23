<template>
  <VButton type="button" variant="primary" @click="openModal(modal)">Dodaj ocenę</VButton>
  <Teleport to="body">
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Dodaj ocenę</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body modal-body-inner">
            <div v-if="errorCode" class="alert alert-danger">
              {{ $translate(errorCode) }}
            </div>
            <SubjectSelector @on-subject-select="(s) => (subject = s)" />

            <InputText v-model="comment" required placeholder="Sprawdzian" type="text" name="mark-comment"
              label="Komentarz do oceny" />
            <div>
              <label for="startDate">Data wystawienia</label>
              <input id="startDate" @change="markDate = ($event.target as HTMLInputElement).value" class="form-control"
                :max="new Date().toISOString().substring(0, 10)" type="date" />
            </div>
            <div class="">
              <label class="pb-2">Ocena</label>

              <select class="form-select" aria-label="Default select example"
                @change="markValue = Number(($event.target as HTMLInputElement).value)">
                <option selected disabled>Wybierz ocenę</option>
                <option v-for="(m, v) in marksMap" :key="m" :value="m">{{ v }}</option>
              </select>
            </div>
          </div>
          <div class="modal-footer">
            <VButton @click="save" variant="primary" type="button">Zapisz</VButton>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import VButton from '@/components/atoms/VButton.vue'
import SubjectSelector from '@/components/molecules/SubjectSelector.vue'
import InputText from '@/components/form/InputText.vue'
import { Modal } from 'bootstrap'
import { type Subject } from '@/types/Subject'
import { markResource } from '@/resources/mark'
import { useCurrentUser } from '@/composables/useCurrentUser'
import { useSnackbar } from '@/composables/useSnackbar'

const { user } = useCurrentUser()
const { successSnackbar } = useSnackbar()
const props = defineProps<{ studentID: number }>()
const $emit = defineEmits(['on-add'])

const marksMap = {
  '1': 1,
  '1+': 1.5,
  '2': 2,
  '2+': 2.5,
  '3': 3,
  '3+': 3.5,
  '4': 4,
  '4+': 4.5,
  5: 5,
  '5+': 5.5,
  '6': 6
}
const errorCode = ref('')
const subject = ref<Subject | null>(null)
const markDate = ref('')
const comment = ref('')
const markValue = ref<number>()

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
function formatDateToMMDDYYYY(date: Date) {
  const month = String(date.getMonth() + 1).padStart(2, '0') // Months are zero-based
  const day = String(date.getDate()).padStart(2, '0')
  const year = date.getFullYear()

  return `${month}-${day}-${year}`
}

const save = async () => {
  try {
    if (subject.value?.id === undefined || markValue.value === undefined) return
    await markResource.addMark({
      subjectID: subject.value?.id,
      comment: comment.value,
      date: markDate.value && formatDateToMMDDYYYY(new Date(markDate.value)),
      value: markValue.value,
      studentID: props.studentID,
      teacherID: user.id
    })
    successSnackbar('Ocena dodana', 3000)
    closeModal(modal.value)
    $emit('on-add')
  } catch (err) {
    errorCode.value = (err as any).response?.data?.error
  }
  console.log(subject.value, markDate.value, comment.value, markValue.value)
}
</script>

<style scoped lang="scss">
.modal-body-inner {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
