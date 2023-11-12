<template>
  <div class="container-md vh-100 d-flex justify-content-center">
    <div class="row align-items-center">
      <div class="col py-4">
        <div v-if="errorCode" class="alert alert-danger">
          {{ errorCode }}
        </div>
        <div v-if="setup" class="alert alert-success">Konto aktywowane, możesz się zalogować</div>
        <h1 class="py-4"><img width="50" height="50" src="/favicon-32x32.png" /> gradebook</h1>
        <form @submit.prevent="handleLogin">
          <InputText v-model="email" type="email" name="email" label="Email" />
          <InputText v-model="password" type="password" name="password" label="Hasło" />
          <button class="btn btn-primary my-2">Zaloguj</button>
        </form>
      </div>
      <div class="col">
        <img src="/login_pic.jpeg" class="rounded-4 img-fluid" alt="" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import InputText from '@/components/form/InputText.vue'
import router, { routes } from '@/router'
import { useQuery } from '@/composables/useQuery'

const { getParam } = useQuery()

const email = ref(getParam('email') || '')
const password = ref('')
const errorCode = ref('')

const setup = computed(() => {
  const params = new URL(location.href).searchParams
  return params.get('setup')
})
const handleLogin = async () => {
  fetch('http://localhost:8080/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email: email.value, password: password.value })
  })
    .then((res) => res.json())
    .then(({ token, Error }) => {
      if (Error) {
        errorCode.value = Error
        return
      }

      localStorage.setItem('token', token)
      router.push(routes.start.path)
    })
    .catch((err) => {
      if (!err.message) {
        errorCode.value = 'We could not handle request'
      } else {
        errorCode.value = err.message || err.Error
      }

      console.log('err', err)
    })
}
</script>
<style lang="scss" scoped></style>
