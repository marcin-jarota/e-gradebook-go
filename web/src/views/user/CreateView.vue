<template>
  <MainLayout>
    <div class="container">
      <h3>Dodaj użytkownika</h3>
      <form @submit.prevent="handleSubmit">
        <InputText v-model="data.name" :required="true" name="user.name" label="Imię" type="text" />
        <InputText v-model="data.surname" :required="true" name="user.surname" label="Nazwisko" type="text" />
        <InputText @change="validate('email')" v-model="data.email" :required="true" :error="errors.email"
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
import { useSnackbarStore } from '@/stores/snackbar'

const snackbarStore = useSnackbarStore()
type FormData<TRole> = {
  name: string
  surname: string
  email: string
  role: TRole
}

const formSchema = yup.object().shape({
  name: yup.string().min(3, 'Imie minimum 3 znaki').max(255).required('Imie wymagane'),
  surname: yup.string().required('Nazwisko wymagane').min(3, 'Nazwisko minimum 3 znaki').max(255),
  email: yup.string().email('Niepoprawny adres email').required('Adres e-mail wymagany'),
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

const urls: Record<Role, string> = {
  [Role.Student]: import.meta.env.VITE_API_BASE_URL + '/student/create',
  [Role.Teacher]: '',
  [Role.Admin]: import.meta.env.VITE_API_BASE_URL + '/user/create'
}

const handleSubmit = async () => {
  try {
    formSchema.validateSync(data)

    if (!urls[data.role]) {
      return
    }

    const res = await fetch(urls[data.role], {
      method: 'POST',
      body: JSON.stringify({ ...data, password: 'devpass' }),
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('token')
      }
    })

    const json = await res.json()

    if (json.error) {
      snackbarStore.createsnackbar(json.error, { autoHideDelay: 5000 })
      return
    }
    snackbarStore.createsnackbar('Użytkownik dodany pomyślnie!', { autoHideDelay: 5000 })
    resetForm()
  } catch (err) {
    snackbarStore.createsnackbar('Nie udało się zapisać danych', { autoHideDelay: 5000 })
    console.log(JSON.stringify(err))
  }
}

const resetForm = () => {
  data.name = ''
  data.surname = ''
  data.email = ''
  data.role = Role.Student
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
