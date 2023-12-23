<script setup lang="ts">
import { RouterLink } from 'vue-router'
import router, { routes } from '@/router'
import { useCurrentUser } from '@/composables/useCurrentUser'
import { userResource } from '@/resources/user'
import NotificationsList from '@/components/organisms/NotificationsList.vue'
import { computed, onMounted, ref } from 'vue'
const { userInitials, hasPermission, user } = useCurrentUser()

const handleLogOut = async () => {
  try {
    await userResource.logout()
    localStorage.removeItem('token')
    router.push('/login')
  } catch (err) {
    console.error(err)
  }
}

const notificationsVisible = ref(false)

const notifications = ref<{ id: number; message: string; createdAt: string; read: boolean }[]>([])

const fetchNotifications = async () => {
  const { data: notificationsList } = await userResource.getNotifications(user.id)

  notifications.value = notificationsList || []
}

fetchNotifications()

const sidebarRoutes = Object.values(routes).filter((r) => r?.meta?.sidebar)

const unreadNotifications = computed(() => notifications.value.filter((n) => !n.read).length)
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
        <li class="sidebar__link px-3 d-flex justify-content-center align-items-center nav-notification-link mt-auto">
          <button @click="notificationsVisible = !notificationsVisible" class="nav-button nav-link">
            <font-awesome-icon size="lg" :icon="['fa-solid', 'fa-bell']" />
            <span class="notification-count bg-warning text-white" v-if="unreadNotifications">{{
              unreadNotifications
            }}</span>
          </button>
        </li>
        <li class="d-flex align-items-center justify-content-center py-4 cursor-pointer">
          <RouterLink :to="'/'" class="text-decoration-none">
            <span class="sidebar__user-initials">{{ userInitials }}</span>
          </RouterLink>
        </li>
      </ul>
    </nav>
    <transition name="slide-fade">
      <div v-if="notificationsVisible" class="notifications-list-container bg-light shadow p-2 rounded">
        <NotificationsList :notifications="notifications" />
      </div>
    </transition>
  </div>
</template>

<style lang="scss">
.notifications-list-container {
  position: fixed; // używamy fixed, aby lista była "pływająca"
  top: 50px; // dostosuj zgodnie z położeniem ikony dzwonka
  right: 50px; // dostosuj, aby lista pojawiała się poza sidebar'em
  z-index: 100; // upewnij się, że lista jest na wierzchu
  background-color: white;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  border-radius: 4px;
  overflow: hidden;
  border-radius: 8px;
}

.slide-fade-enter-active,
.slide-fade-leave-active {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}

.slide-fade-enter,
.slide-fade-leave-to

/* .slide-fade-leave-active in <2.1.8 */
  {
  transform: translateX(100%);
  /* Wjeżdża z prawej */
  opacity: 0;
}

.notification-popover-enter-active,
.notification-popover-leave-active {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}

.notification-popover-enter,
.notification-popover-leave-to

/* .notification-popover-leave-active in <2.1.8 */
  {
  transform: translateX(0);
  /* Wraca w prawo */
  opacity: 1;
}

.sidebar {
  height: 100vh;
  position: sticky;
  top: 0;
  left: 0;
  z-index: 1000;

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

.notification-count {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 9px;
  border-radius: 50%;
  padding: 5px;
  top: -8px;
  right: -15px;
  width: 18px;
  height: 18px;
}

button.nav-button {
  position: relative;
  background-color: transparent;
  padding: 0 !important;
  margin: 0;
  border: none;
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
