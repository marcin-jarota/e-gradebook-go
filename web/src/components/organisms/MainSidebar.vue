<script setup lang="ts">
import { RouterLink } from 'vue-router'
import router, { routes } from '@/router'
import { useCurrentUser } from '@/composables/useCurrentUser'
import { userResource } from '@/resources/user'

const { userInitials, hasPermission } = useCurrentUser()

const handleLogOut = async () => {
  try {
    await userResource.logout()
    localStorage.removeItem('token')
    router.push('/login')
  } catch (err) {
    console.error(err)
  }
}

const sidebarRoutes = Object.values(routes).filter((r) => r?.meta?.sidebar)

console.log(sidebarRoutes)
</script>

<template>
  <div class="sidebar d-flex flex-column flex-shrink-0 bg-light">
    <nav class="d-flex flex-column h-100">
      <RouterLink :to="routes.start.path" class="d-block p-3 link-dark border-bottom text-decoration-none">
        <img width="50" height="50" src="/apple-touch-icon.png" class="logo m-auto d-block" alt="App logo" />
      </RouterLink>
      <ul class="nav nav-pills nav-flush flex-column flex-grow-1 mb-auto text-center">
        <template v-for="route in sidebarRoutes">
          <li v-if="route.meta?.roles?.some(hasPermission)" :key="route.name" class="sidebar__link p-3 nav-link">
            <RouterLink :class="{ active: $route.name === route.name }" class="nav-link py-2" :to="route.path">
              <font-awesome-icon v-if="route?.meta?.icon" size="lg" :icon="['fa-solid', route.meta?.icon]" />
              <span class="popover-content sidebar__hint">{{ route?.meta?.title }}</span>
            </RouterLink>
          </li>
        </template>

        <li class="sidebar__link p-3 nav-link">
          <a href="javascript:void(0)" class="text-decoration-none cursor-pointer" @click.prevent="handleLogOut">
            <font-awesome-icon size="lg" class="logout-icon" icon="fa-solid fa-right-from-bracket" />
            <span class="popover-content sidebar__hint">Wyloguj</span>
          </a>
        </li>
        <li class="d-flex mt-auto align-items-center justify-content-center py-4 cursor-pointer">
          <RouterLink :to="'/'" class="text-decoration-none">
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

.nav-link>a:hover {
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
  font-size: 14px;
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
