<template>
  <SelectDialog @save-request="save" :selectable-items="teachersToAssign" title="Przypisz nauczyciela"
    dropdown-title="Wybierz nauczyciela" dropdown-placeholder="Jan Kowalski" v-slot="props">
    <VButton type="button" variant="primary" @click="props.openModal">Przypisz nauczyciela</VButton>
  </SelectDialog>
</template>

<script setup lang="ts">
import { inject, ref } from 'vue'
import { useSnackbar } from '@/composables/useSnackbar'
import VButton from '@/components/atoms/VButton.vue'
import SelectDialog from '@/components/organisms/SelectDialog.vue'
import { classGroupResource } from '@/resources/classGroup'
import { teacherResource } from '@/resources/teacher'

const { successSnackbar, errorSnackbar } = useSnackbar()
const props = defineProps<{ classGroupId: number }>()
const $emit = defineEmits(['save-success'])
const translate = inject('translate') as (code: string) => string

const teachersToAssign = ref<{ id: number; name: string }[]>([])

const getTeachers = async () => {
  const { data } = await teacherResource.list()
  teachersToAssign.value = data.map(({ id, name, surname }) => ({
    id,
    name: `${name} ${surname}`
  }))
}

getTeachers()

const save = async (teacher: (typeof teachersToAssign.value)[0]) => {
  try {
    if (!teacher) return

    await classGroupResource.assignTeacher({
      teacherID: teacher.id,
      classGroupID: props.classGroupId
    })
    successSnackbar('Nauczyciel przypisany', 3000)
    $emit('save-success')
  } catch (err) {
    const code = (err as any).response?.data?.error
    errorSnackbar(translate(code))
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
