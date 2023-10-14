<script setup lang="ts">
import { computed } from 'vue';
import { useSessionStore } from '@/stores/session';
import { RouterLink } from 'vue-router';
import { Role } from '@/types';

const sessionStore = useSessionStore()
const user = sessionStore.user

const userInitials = computed(() => {
    return `${user?.name[0]}${user?.surname[0]}`
})

const isStudent = computed(() => {
    return user?.role === Role.Student 
})

</script>

<template>
    <nav class="sidebar">
        <ul class="navbar-nav">
            <li v-if="isStudent">
                <RouterLink to="/student/marks">Oceny</RouterLink>
            </li>
            <li v-if="sessionStore.user">
                <span class="sidebar__user-initials">{{ userInitials }}</span>
            </li>
        </ul>
    </nav>
</template>

<style lang="scss" scoped>
.sidebar {

}

.sidebar__user-initials {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    color: #fff;
    background-color: #999;
}
</style>