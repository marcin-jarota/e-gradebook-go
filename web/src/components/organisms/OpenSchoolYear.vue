<template>
  <VButton type="button" variant="light" @click="openModal(modal)">Otwórz rok szkolny</VButton>
  <Teleport to="body">
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Otwórz rok szkolny</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body modal-body-inner">
            <div v-if="errorCode" class="alert alert-danger">
              {{ $translate(errorCode) }}
            </div>

            <InputText v-model="name" required placeholder="2024/2025" type="text" name="name" label="Nazwa" />
            <div>
              <label for="startDate">Data rozpiczęcia</label>
              <input id="startDate" @change="startDate = ($event.target as HTMLInputElement).value" class="form-control"
                type="date" />
            </div>
            <div>
              <label for="startDate">Data zakończenia</label>
              <input id="startDate" @change="endDate = ($event.target as HTMLInputElement).value" class="form-control"
                type="date" />
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
import InputText from '@/components/form/InputText.vue'
import { Modal } from 'bootstrap'
import { useSnackbar } from '@/composables/useSnackbar'
import { schoolYearResource } from '@/resources/schoolYear'
const { successSnackbar } = useSnackbar()
const $emit = defineEmits(['on-add'])

const errorCode = ref('')
const startDate = ref('')
const endDate = ref('')
const name = ref('')

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
    if (!name.value === undefined || startDate.value === undefined || endDate.value === undefined)
      return
    await schoolYearResource.openSchoolYear({
      name: name.value,
      start: formatDateToMMDDYYYY(new Date(startDate.value)),
      end: formatDateToMMDDYYYY(new Date(endDate.value))
    })
    successSnackbar('Nowy rok szkolny otwarty', 3000)
    closeModal(modal.value)
    $emit('on-add')
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
