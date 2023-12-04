<template>
  <SelectDialog @save-request="save" :selectable-items="teachersToAssign" title="Przypisz nauczyciela uczącego"
    dropdown-title="Wybierz nauczyciela" dropdown-placeholder="Jan Kowalski" v-slot="props">
    <VButton type="button" variant="light" @click="props.openModal"><font-awesome-icon icon="fa-solid fa-user-plus" />
    </VButton>
  </SelectDialog>
</template>

<script setup lang="ts">
import { inject, ref } from 'vue'
import { useSnackbar } from '@/composables/useSnackbar'
import VButton from '@/components/atoms/VButton.vue'
import SelectDialog from '@/components/organisms/SelectDialog.vue'
import { teacherResource } from '@/resources/teacher'
import { subjectResource } from '@/resources/subject'

const { successSnackbar, errorSnackbar } = useSnackbar()
const props = defineProps<{ subjectId: number }>()
const $emit = defineEmits(['save-success'])
const translate = inject('translate') as (code: string) => string

const teachersToAssign = ref<{ id: number; name: string }[]>([])

const getTeachers = async () => {
  const { data } = await teacherResource.list()
  teachersToAssign.value = data.map(({ id, name, surname }) => ({ id, name: `${name} ${surname}` }))
}

getTeachers()

const save = async (teacher: (typeof teachersToAssign.value)[0]) => {
  try {
    if (!teacher) return

    await subjectResource.assignTeacher({
      teacherID: teacher.id,
      subjectID: props.subjectId
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
