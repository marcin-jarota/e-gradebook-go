<template>
  <SelectDialog @save-request="save" :selectable-items="subjectsToAssigne" title="Przypisz przedmiot"
    dropdown-title="Wybierz ucznia" dropdown-placeholder="Informatyka" v-slot="props">
    <VButton type="button" variant="primary" @click="props.openModal">Przypisz przedmiotj</VButton>
  </SelectDialog>
</template>

<script setup lang="ts">
import { inject, ref } from 'vue'
import { useSnackbar } from '@/composables/useSnackbar'
import VButton from '@/components/atoms/VButton.vue'
import SelectDialog from '@/components/organisms/SelectDialog.vue'
import { classGroupResource } from '@/resources/classGroup'
import { subjectResource } from '@/resources/subject'

const { successSnackbar, errorSnackbar } = useSnackbar()
const props = defineProps<{ classGroupId: number }>()
const $emit = defineEmits(['save-success'])
const translate = inject('translate') as (code: string) => string

const subjectsToAssigne = ref<{ id: number; name: string }[]>([])

const getsubjects = async () => {
  const { data } = await subjectResource.list()
  subjectsToAssigne.value = data.map(({ id, name }) => ({ id, name: name }))
}

getsubjects()

const save = async (subject: (typeof subjectsToAssigne.value)[0]) => {
  try {
    if (!subject) return

    await classGroupResource.assignSubject({
      subjectID: subject.id,
      classGroupID: props.classGroupId
    })
    successSnackbar('Przedmiot przypisany', 3000)
    $emit('save-success')
  } catch (err) {
    errorSnackbar(translate((err as any).response?.data?.error))
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
