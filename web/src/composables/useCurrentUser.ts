import { useSessionStore } from '@/stores/session'
import { Role, type SessionUser } from '@/types'
import { computed } from 'vue'

export const useCurrentUser = () => {
  const store = useSessionStore()
  const user = store.user as SessionUser

  const userInitials = computed(() => {
    return `${user?.name[0]}${user?.surname[0]}`
  })

  const isStudent = computed(() => {
    console.log(user)
    return user?.role === Role.Student
  })

  const isAdmin = computed(() => {
    return user?.role === Role.Admin
  })

  const hasPermission = (requiredRole: Role) => user.role === requiredRole

  return {
    userInitials,
    isStudent,
    isAdmin,
    hasPermission,
    user
  }
}
