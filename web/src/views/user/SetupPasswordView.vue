<template>
  <div class="container-md vh-100 d-flex justify-content-center">
    <div class="row align-items-center">
      <div class="col py-4">
        <div v-if="error" class="alert alert-danger">
          {{ error }}
        </div>
        <form @submit.prevent="submit">
          <InputText v-model="password" type="password" name="password" label="Hasło" />
          <InputText v-model="passwordConfirm" type="password" name="passwordConfirm" label="Potwierdź hasło" />
          <button class="btn btn-primary my-2">Zapisz</button>
        </form>
      </div>
      <div class="col">
        <img src="/login_pic.jpeg" class="rounded-4 img-fluid" alt="" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InputText from '@/components/form/InputText.vue'
import { userResource } from '@/resources/user'
import { useQuery } from '@/composables/useQuery'
import { useRouter } from 'vue-router'
import { routes } from '@/router'

const { getParam } = useQuery()
const password = ref('')
const passwordConfirm = ref('')
const error = ref('')

const router = useRouter()
const submit = async () => {
  if (password.value !== passwordConfirm.value) {
    error.value = 'Hasła muszą być identyczne'
  }

  try {
    const token = getParam('token')

    if (!token) {
      error.value = 'Brak tokenu'
      return
    }

    await userResource.setupPassword(
      { password: password.value, passwordConfirm: passwordConfirm.value },
      token
    )

    router.push({
      path: routes.login.path,
      query: {
        setup: 'true',
        email: getParam('email')
      }
    })
  } catch (err) {
    error.value = (err as any).message
  }
}
</script>

<style scoped></style>
