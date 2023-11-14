<template>
  <div class="table-responsive py-3">
    <table class="table table-hover">
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
        <tr v-for="user in users" :key="user.id" role="button">
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
            <button v-if="user.sessionActive" @click="destroySession(user.id)" class="btn btn-secondary me-2">
              Wyloguj
            </button>
            <button v-if="user.isActive" @click="deactivate(user.id)" class="btn btn-warning">
              Deaktywuj
            </button>
            <button v-else @click="activate(user.id)" class="btn btn-success">Aktywuj</button>
          </td>
          <td v-else>-</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts" setup>
import IconStatus from '@/components/common/IconStatus.vue'
import { ref } from 'vue'

import { useCurrentUser } from '@/composables/useCurrentUser'
import { useSnackbar } from '@/composables/useSnackbar'
import { userResource } from '@/resources/user'
import { type UserOutput, Role } from '@/types'

const { user: currentUser } = useCurrentUser()
const { errorSnackbar, successSnackbar } = useSnackbar()

const users = ref<UserOutput[]>([])

const roleName: Record<Role, string> = {
  [Role.Admin]: 'Administrator',
  [Role.Student]: 'Uczeń',
  [Role.Teacher]: 'Nauczyciel'
}

const SNAKCBAR_DURATION = 2000

const getUsers = async () => {
  try {
    const { data } = await userResource.list()
    users.value = data.sort((a, b) => a.id - b.id)
  } catch (err) {
    errorSnackbar('Nie udało się pobrać listy', SNAKCBAR_DURATION)
  }
}

const activate = async (userID: number) => {
  try {
    await userResource.activate(userID)
    await getUsers()
    successSnackbar('Użytkownik aktywowany pomyślnie', SNAKCBAR_DURATION)
  } catch (err) {
    errorSnackbar('Nie udało się aktytować użytkownika', SNAKCBAR_DURATION)
  }
}

const deactivate = async (userID: number) => {
  try {
    await userResource.deactivate(userID)
    await getUsers()
    successSnackbar('Użytkownik deaktywowany pomyślnie', SNAKCBAR_DURATION)
  } catch (err) {
    errorSnackbar('Nie udało się deaktywować użytkownika', SNAKCBAR_DURATION)
  }
}

const destroySession = async (userID: number) => {
  try {
    await userResource.destroySession(userID)
    await getUsers()
    successSnackbar('Użytkownik wylogowany pomyślnie', 3000)
  } catch (err) {
    console.log(err)
    errorSnackbar('Nie udało się wylogować użytkownika', 3000)
  }
}

getUsers()
</script>

<style lang="scss" scoped></style>
