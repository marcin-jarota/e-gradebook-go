<template>
  <slot :openModal="() => openModal(modal)"> </slot>
  <Teleport to="body">
    <div class="modal fade" ref="modal" tabindex="-1" id="exampleModal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{ title }}</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body modal-body-inner">
            <DynamicSelector :items="selectableItems" @on-select="(s) => (selected = s)" :title="dropdownTitle"
              placeholder="Imię lub nazwisko" />
          </div>
          <div class="modal-footer">
            <VButton @click="save" variant="primary" type="button" :disabled="!selected">Przypisz</VButton>
          </div>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts" generic="T extends { id: number; name: string }">
import { ref } from 'vue'
import VButton from '@/components/atoms/VButton.vue'
import { Modal } from 'bootstrap'
import DynamicSelector from '@/components/molecules/DynamicSelector.vue'

defineProps<{
  title: string
  selectableItems: T[]
  dropdownTitle: string
  dropdownPlaceholder: string
}>()
const $emit = defineEmits<{ (e: 'save-request', item: T): any }>()
const selected = ref<T | null>(null)
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

const save = () => {
  closeModal(modal.value)
  if (!selected.value) return
  $emit('save-request', selected.value as T)
}
</script>

<style scoped lang="scss">
.modal-body-inner {
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
