<template>
  <div class="py-2 alert alert-success" v-if="link"><a :href="link">Aktywuj konto</a></div>
  <form @submit.prevent="handleSubmit">
    <InputText v-model="data.name" :required="true" name="user.name" label="Imię" type="text" />
    <InputText v-model="data.surname" :required="true" name="user.surname" label="Nazwisko" type="text" />
    <InputText @change="validate('email')" v-model="data.email" :required="true" :error="errors.email" name="user.email"
      label="E-mail" type="email" />
    <div class="mb-3">
      <label class="pb-2">Rola</label>
      <select class="form-select" v-model="data.role" aria-label="Selektor roli użytkownika">
        <option selected disabled>Wybierz role</option>
        <option :value="Role.Admin">Admin</option>
        <option :value="Role.Student">Uczeń</option>
        <option :value="Role.Teacher">Nauczyciel</option>
      </select>
    </div>
    <Button variant="primary" type="submit">Zapisz</Button>
  </form>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import * as yup from 'yup'
import { Role, type UserInput } from '@/types'
import InputText from '@/components/atoms/InputText.vue'
import Button from '@/components/atoms/VButton.vue'
import { userResource } from '@/resources/user'
import { useSnackbar } from '@/composables/useSnackbar'

const { errorSnackbar, successSnackbar } = useSnackbar()

const props = defineProps<{ user?: UserInput }>()

const data = reactive<UserInput>({
  name: props.user?.name || '',
  surname: props.user?.surname || '',
  email: props.user?.email || '',
  role: props.user?.role || Role.Student
})

const link = ref('')

const errors = reactive<Omit<UserInput, 'role'>>({
  name: '',
  surname: '',
  email: ''
})

const validationSchema = yup.object().shape({
  name: yup.string().min(3, 'Imie minimum 3 znaki').max(255).required('Imie wymagane'),
  surname: yup.string().required('Nazwisko wymagane').min(3, 'Nazwisko minimum 3 znaki').max(255),
  email: yup.string().email('Niepoprawny adres email').required('Adres e-mail wymagany'),
  role: yup.mixed().required('Rola wymagana')
})

const resetForm = () => {
  data.name = ''
  data.surname = ''
  data.email = ''
  data.role = Role.Student
}

const validate = (field: keyof Omit<UserInput, 'role'>) => {
  try {
    const result = validationSchema.validateSyncAt(field, data)
    errors[field] = ''
    console.log(result)
  } catch (err) {
    errors[field] = (err as yup.ValidationError).message
  }
}

const handleSubmit = async () => {
  try {
    const { data: response } = await userResource.create(data)
    link.value = response.activationLink
    successSnackbar('Użytkownik zapisane pomyślnie!', 3000)
    resetForm()
  } catch (err) {
    errorSnackbar('Nie udało się zapisać użytkownika')
  }
}
</script>

<style scoped></style>
