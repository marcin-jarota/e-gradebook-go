<template>
  <MainLayout>
    <div class="container">
      <h3>Dodaj użytkownika</h3>
      <form @submit.prevent="handleSubmit">
        <InputText v-model="data.name" :required="true" name="user.name" label="Imię" type="text" />
        <InputText v-model="data.surname" :required="true" name="user.surname" label="Nazwisko" type="text" />
        <InputText @keydown="validate('email')" v-model="data.email" :required="true" :error="errors.email"
          name="user.email" label="E-mail" type="email" />
        <div class="mb-3">
          <label class="pb-2">Rola</label>
          <select class="form-select" v-model="data.role" aria-label="Default select example">
            <option selected disabled>Wybierz role</option>
            <option :value="Role.Admin">Admin</option>
            <option :value="Role.Student">Uczeń</option>
            <option :value="Role.Teacher">Nauczyciel</option>
          </select>
        </div>
        <Button variant="primary" type="submit">Zapisz</Button>
      </form>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import * as yup from 'yup'
import { Role } from '@/types'
import MainLayout from '@/layouts/MainLayout.vue'
import InputText from '@/components/atoms/InputText.vue'
import Button from '@/components/atoms/VButton.vue'

type FormData<TRole> = {
  name: string
  surname: string
  email: string
  role: TRole
}

const formSchema = yup.object().shape({
  name: yup.string().required('Imie wymagane').min(3).max(255),
  surname: yup.string().required('Nazwisko wymagane').min(3).max(255),
  email: yup.string().email('Niepoprawny adres email').required(),
  role: yup.mixed().required('Rola wymagana')
})

const data = reactive<FormData<Role>>({
  name: '',
  surname: '',
  email: '',
  role: Role.Student
})

const errors = reactive<FormData<string>>({
  name: '',
  surname: '',
  email: '',
  role: ''
})

const handleSubmit = async () => {
  console.log(data)
  try {
    formSchema.validateSync(data)
  } catch (err) {
    console.log(JSON.stringify(err))
  }
}

const validate = (field: keyof FormData<string>) => {
  try {
    const result = formSchema.validateSyncAt(field, data)
    errors[field] = ''
    console.log(result)
  } catch (err) {
    //@ts-ignore
    console.log(err.message)
    errors[field] = (err as yup.ValidationError).message
  }
}
</script>

<style scoped></style>
