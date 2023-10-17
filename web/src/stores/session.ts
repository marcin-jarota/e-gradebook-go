import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { SessionUser } from '@/types'

export const useSessionStore = defineStore('session', () => {
  const user = ref<SessionUser | null>()

  const updateUser = (newUser: SessionUser | null) => {
    user.value = newUser
  }

  return { user, updateUser }
})
