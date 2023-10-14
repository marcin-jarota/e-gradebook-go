import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { SessionUser } from '@/types'

export const useSessionStore = defineStore('session', () => {
  const user = ref<SessionUser>()

  const updateUser = (newUser: SessionUser) => {
    user.value = newUser
  }

  return { user, updateUser }
})
