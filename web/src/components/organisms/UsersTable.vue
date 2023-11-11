<template>
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
            <IconStatus :active="Boolean(user.isActive)" />
          </td>
          <td>
            <IconStatus :active="Boolean(user.sessionActive)" />
          </td>
          <td v-if="currentUser.id !== user.id">
            <button class="btn btn-primary">Edytuj</button>
            <button v-if="user.sessionActive" @click="destroySession(user.id)" class="btn btn-secondary ms-2">
              Wyloguj
            </button>
            <button v-if="!user.isActive" @click="activate(user.id)" class="btn btn-success ms-2">
              Aktywuj
            </button>
            <button v-else @click="deactivate(user.id)" class="btn btn-warning ms-2">
              Deaktywuj
            </button>
          </td>
          <td v-else>-</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts" setup>
import IconStatus from '@/components/common/IconStatus.vue'
import { useCurrentUser } from '@/composables/useCurrentUser'
import { useSnackbarStore } from '@/stores/snackbar'
import { userResource } from '@/resources/user'
import { type UserOutput, Role } from '@/types'
import { ref } from 'vue'

const { user: currentUser } = useCurrentUser()
const snackbarStore = useSnackbarStore()
const users = ref<UserOutput[]>([])

const roleName: Record<Role, string> = {
  [Role.Admin]: 'Administrator',
  [Role.Student]: 'Uczeń',
  [Role.Teacher]: 'Nauczyciel'
}

const getUsers = async () => {
  try {
    users.value = (await userResource.list()).data
  } catch (err) {
    console.log(err)
    snackbarStore.createsnackbar('Nie udało się pobrać listy', { variant: 'danger' })
  }
}

const activate = async (userID: number) => {
  try {
    await userResource.activate(userID)

    users.value = users.value.map((u) => {
      if (u.id === userID) u.isActive = true
      return u
    })
    snackbarStore.createsnackbar('Użytkownik aktywowany pomyślnie', { variant: 'success' })
  } catch (err) {
    snackbarStore.createsnackbar('Nie udało się aktytować użytkownika', { variant: 'danger' })
  }
}

const deactivate = async (userID: number) => {
  try {
    await userResource.deactivate(userID)
    updateUser(userID, 'isActive', false)
    updateUser(userID, 'sessionActive', false)
    snackbarStore.createsnackbar('Użytkownik aktywowany pomyślnie', { variant: 'success' })
  } catch (err) {
    snackbarStore.createsnackbar('Nie udało się aktytować użytkownika', { variant: 'danger' })
  }
}

const updateUser = <K extends keyof UserOutput>(userID: number, key: K, value: UserOutput[K]) => {
  users.value = users.value.map((u) => {
    if (u.id === userID && u?.[key]) {
      u[key] = value
    }
    return u
  })
}

const destroySession = async (userID: number) => {
  try {
    await userResource.destroySession(userID)
    users.value = users.value.map((u) => {
      if (u.id === userID) u.sessionActive = false
      return u
    })
    snackbarStore.createsnackbar('Użytkownik wylogowany pomyślnie', { variant: 'success' })
  } catch (err) {
    console.log(err)
    snackbarStore.createsnackbar('Nie udało się wylogować użytkownika', { variant: 'danger' })
  }
}

getUsers()
</script>

<style lang="scss" scoped></style>
