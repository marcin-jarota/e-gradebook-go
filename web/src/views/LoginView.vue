<template>
  <div class="container login-container">
    <div v-if="errorCode" class="alert alert-danger">
      {{ errorCode }}
    </div>
    <form @submit.prevent="handleLogin">
      <InputText v-model="email" type="email" name="email" label="Email" />
      <InputText v-model="password" type="password" name="password" label="Hasło" />
      <button class="btn btn-primary my-2">Zaloguj</button>
    </form>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import InputText from '@/components/form/InputText.vue'
import router, { routes } from '@/router';

const email = ref('')
const password = ref('')
const errorCode = ref('')

const handleLogin = async () => {
  console.log(email.value)

  fetch('http://localhost:8080/login', { method: 'POST', headers: {'Content-Type': 'application/json'}, body: JSON.stringify({ email: email.value, password: password.value }) })
      .then(res => res.json())
    .then(({ token }) => {
      localStorage.setItem('token', token)
        router.push(routes.start.path)
    }).catch(err => {
      if (!err.message) {
        errorCode.value = 'asdasdas'
      } else {
        errorCode.value = err.message
       }

       console.log('err', err)
      })
}
</script>
<style lang="scss" scoped>

</style>
