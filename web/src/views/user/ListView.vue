<script lang="ts" setup>
import MainLayout from '@/layouts/MainLayout.vue'
import IconStatus from '@/components/common/IconStatus.vue'
import { useCurrentUser } from '@/composables/useCurrentUser'
import { useSnackbarStore } from '@/stores/snackbar'
import { type UserOutput, type UserListResponse, Role } from '@/types'
import { ref } from 'vue'

const { user: currentUser } = useCurrentUser()
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

const roleName: Record<Role, string> = {
  [Role.Admin]: 'Administrator',
  [Role.Student]: 'Uczeń',
  [Role.Teacher]: 'Nauczyciel'
}

const activate = async (userID: number) => {
  try {
    const response = await fetch(import.meta.env.VITE_API_BASE_URL + '/user/activate/' + userID, {
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('token')
      }
    })

    const data = await response.json()
    if (data.error) {
      alert(data.error)
      return
    }

    alert('OK!')
    console.log(data)
    users.value = users.value.map((u) => {
      if (u.id === userID) u.isActive = true
      return u
    })
  } catch (err) { }
}

const destroySession = async (userID: number) => {
  try {
    const response = await fetch(
      import.meta.env.VITE_API_BASE_URL + '/user/destroy-session/' + userID,
      {
        headers: {
          'Content-Type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('token')
        }
      }
    )

    const data = await response.json()
    if (data.error) {
      alert(data.error)
      return
    }

    alert('OK!')
    console.log(data)
    users.value = users.value.map((u) => {
      if (u.id === userID) u.isActive = true
      return u
    })
  } catch (err) { }
}
</script>

<template>
  <MainLayout>
    <div class="container-lg">
      <h2>Lista użytkowników w systemie</h2>
      <div class="table-responsive py-3">
        <table class="table">
          <thead class="table-light">
            <tr>
              <th scope="col">ID</th>
              <th scope="col">Imie</th>
              <th scope="col">Nazwisko</th>
              <th scope="col" class="col-4">Email</th>
              <th scope="col">Rola</th>
              <th scope="col">Aktywny</th>
              <th scope="col">Zalogowany</th>
              <th scope="col">Akcje</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.id">
              <td scope="row">{{ user.id }}</td>
              <td>{{ user.name }}</td>
              <td>{{ user.surname }}</td>
              <td>{{ user.email }}</td>
              <td>{{ roleName[user.role] }}</td>
              <td>
                <IconStatus :active="user.isActive || false" />
              </td>
              <td>
                <IconStatus :active="user.sessionActive || false" />
              </td>
              <td>
                <button class="btn btn-primary">Edytuj</button>
                <button v-if="currentUser.id !== user.id && user.sessionActive" @click="destroySession(user.id)"
                  class="btn btn-secondary ms-2">
                  Wyloguj
                </button>
                <button v-if="currentUser.id !== user.id && !user.isActive" @click="activate(user.id)"
                  class="btn btn-success ms-2">
                  Aktywuj
                </button>
                <button v-if="currentUser.id !== user.id && user.isActive" class="btn btn-warning ms-2">
                  Deaktywuj
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </MainLayout>
</template>

<style lang="scss" scoped></style>
