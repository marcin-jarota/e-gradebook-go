<script lang="ts" setup>
import MainLayout from '@/layouts/MainLayout.vue'
import IconStatus from '@/components/common/IconStatus.vue'
import { type UserOutput, type UserListResponse } from '@/types'
import { ref } from 'vue'

const users = ref<UserOutput[]>([])

const getUsers = () => {
  fetch(`${import.meta.env.VITE_API_BASE_URL}/user/list`, {
    headers: {
      Authorization: 'Bearer ' + localStorage.getItem('token')
    }
  })
    .then((res) => {
      console.log(`${import.meta.env.VITE_API_BASE_URL}/user/list`)
      return res.json()
    })
    .then((data: UserListResponse) => {
      users.value = data.data
    })
}

getUsers()
</script>

<template>
  <MainLayout>
    <h2>Lista użytkowników w systemie</h2>
    <div class="table-responsive py-3">
      <table class="table">
        <thead class="table-light">
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Imie</th>
            <th scope="col">Nazwisko</th>
            <th scope="col">Email</th>
            <th scope="col">Aktywny</th>
            <th scope="col">Zalogowany</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td scope="row">{{ user.id }}</td>
            <td>{{ user.name }}</td>
            <td>{{ user.surname }}</td>
            <td>{{ user.email }}</td>
            <td>
              <IconStatus :active="user.isActive || false" />
            </td>
            <td>
              <IconStatus :active="user.sessionActive || false" />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </MainLayout>
</template>

<style lang="scss" scoped></style>
