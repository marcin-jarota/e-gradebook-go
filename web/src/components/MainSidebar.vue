<script setup lang="ts">
import { RouterLink } from 'vue-router'
import router, { routes } from '@/router'
import { useCurrentUser } from '@/composables/useCurrentUser'

const { userInitials, isAdmin, isStudent } = useCurrentUser()

const handleLogOut = async () => {
  try {
    await fetch(import.meta.env.VITE_API_BASE_URL + '/logout')
    localStorage.setItem('token', '')
    router.push('/login')
  } catch (err) {
    console.error(err)
  }
}
</script>

<template>
  <div class="sidebar d-flex flex-column flex-shrink-0 bg-light">
    <nav class="d-flex flex-column h-100">
      <RouterLink :to="routes.start.path" class="d-block p-3 link-dark border-bottom text-decoration-none">
        <img width="50" height="50" src="/apple-touch-icon.png" class="logo m-auto d-block" alt="App logo" />
      </RouterLink>
      <ul class="nav nav-pills nav-flush flex-column flex-grow-1 mb-auto text-center">
        <li v-if="isStudent" class="sidebar__link p-3 nav-link">
          <RouterLink :class="{ active: $route.name === routes.studentMarks.name }" class="nav-link py-2"
            :to="routes.studentMarks.path">
            <font-awesome-icon size="lg" icon="fa-solid fa-graduation-cap" />
            <span class="popover-content sidebar__hint">Oceny</span>
          </RouterLink>
        </li>

        <li v-if="isAdmin" class="sidebar__link p-3 nav-link">
          <RouterLink :to="routes.createUser.path">
            <font-awesome-icon icon="user-plus" />
            <span class="popover-content sidebar__hint">Dodaj użytkownika</span>
          </RouterLink>
        </li>
        <li v-if="isAdmin" class="sidebar__link p-3 nav-link">
          <RouterLink :to="routes.userList.path">
            <font-awesome-icon icon="fa-solid fa-users" />
            <span class="popover-content sidebar__hint">Lista użytkowników</span>
          </RouterLink>
        </li>

        <li v-if="isAdmin" class="sidebar__link p-3 nav-link">
          <RouterLink :to="routes.subjectList.path">
            <font-awesome-icon icon="fa-solid fa-chalkboard" />
            <span class="popover-content sidebar__hint">Lista przedmiotów</span>
          </RouterLink>
        </li>
        <li v-if="isAdmin" class="sidebar__link p-3 nav-link">
          <RouterLink :to="routes.classGroupList.path">
            <font-awesome-icon icon="fa-solid fa-school" />
            <span class="popover-content sidebar__hint">Lista klas</span>
          </RouterLink>
        </li>

        <li class="sidebar__link p-3 nav-link">
          <a href="javascript:void(0)" class="text-decoration-none cursor-pointer" @click.prevent="handleLogOut">
            <font-awesome-icon size="lg" class="logout-icon" icon="fa-solid fa-right-from-bracket" />
            <span class="popover-content sidebar__hint">Wyloguj</span>
          </a>
        </li>
        <li class="d-flex mt-auto align-items-center justify-content-center py-4 cursor-pointer">
          <RouterLink :to="routes.profile.path" class="text-decoration-none">
            <span class="sidebar__user-initials">{{ userInitials }}</span>
          </RouterLink>
        </li>
      </ul>
    </nav>
  </div>
</template>

<style lang="scss">
.sidebar {
  height: 100vh;
  position: sticky;
  top: 0;
  left: 0;

  &__user-initials {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    color: #fff;
    background-color: #999;
    margin-top: auto;
  }

  &__link {
    position: relative;
    /* @extend .p-3;
        @extend .nav-link; */
  }
}

.nav-link:hover {
  .popover-content {
    visibility: visible;
    opacity: 1;
  }
}

.popover-content {
  display: flex;
  visibility: hidden;
  white-space: pre;
  opacity: 0;
  position: absolute;
  background-color: #3498db;
  color: #fff;
  padding: 10px;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
  left: 85px;
  top: 50%;
  transform: translateY(-50%);
  padding: 4px 8px;
  font-size: 12px;
  transition: 0.3s;
}

.popover-content::before {
  content: '';
  position: absolute;
  top: 50%;
  right: 100%;
  margin-top: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: transparent transparent transparent #3498db;
  transform: rotate(180deg);
}

.logo {
  max-width: 50px;
}

.logout-icon {
  transform: rotate(180deg);
}
</style>
