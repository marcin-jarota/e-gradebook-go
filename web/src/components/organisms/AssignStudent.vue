<template>
  <SelectDialog
    @save-request="save"
    :selectable-items="studentsToAssigne"
    title="Przypisz ucznia"
    dropdown-title="Wybierz ucznia"
    dropdown-placeholder="Jan Kowalski"
    v-slot="props"
  >
    <VButton type="button" variant="primary" @click="props.openModal">Przypisz ucznia</VButton>
  </SelectDialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useSnackbar } from '@/composables/useSnackbar'
import VButton from '@/components/atoms/VButton.vue'
import SelectDialog from '@/components/organisms/SelectDialog.vue'
import { studentResource } from '@/resources/student'
import { classGroupResource } from '@/resources/classGroup'

const { successSnackbar } = useSnackbar()
const props = defineProps<{ classGroupId: number }>()
const $emit = defineEmits(['save-success'])

const errorCode = ref('')

const studentsToAssigne = ref<{ id: number; name: string }[]>([])

const getStudents = async () => {
  const { data } = await studentResource.getAll()
  studentsToAssigne.value = data.map(({ id, fullName }) => ({ id, name: fullName }))
}

getStudents()

const save = async (student: (typeof studentsToAssigne.value)[0]) => {
  try {
    if (!student) return

    await classGroupResource.assignStudent({
      studentID: student.id,
      classGroupID: props.classGroupId
    })
    successSnackbar('Student przypisany', 3000)
    $emit('save-success')
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
