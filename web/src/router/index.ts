import { createRouter, createWebHistory } from 'vue-router'
import jwtDecode from 'jwt-decode'
import { Role } from '@/types'
import type { SessionUser } from '@/types'
import { useSessionStore } from '@/stores/session'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect(to) {
        return { name: 'login' }
      }
    },
    {
      path: '/login',
      name: 'login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('@/views/LoginView.vue')
    },
    {
      path: '/student',
      name: 'student',
      component: () => import('@/views/StudentView.vue'),
      meta: { requiresAuth: true, roles: [Role.Student] }
    },
    {
      path: '/denied',
      name: 'denied',
      component: () => import('@/views/DeniedView.vue')
    },
    {
      path: '/:pathMatch(.*)*',
      component: () => import('@/views/PageNotFound.vue')
    }
  ]
})

router.beforeEach((to, _, next) => {
  const meta = to.meta as { requiresAuth?: boolean; requiresRole?: Role[] }
  const token = localStorage.getItem('token')

  if (!meta.requiresAuth) return next()

  if (meta.requiresAuth && token === 'undefined') {
    next('/login')
    return
  }

  const session = token ? jwtDecode<{ sessionUser: SessionUser }>(token) : null
  const user = session?.sessionUser

  const { updateUser } = useSessionStore()

  if (meta?.requiresRole?.length && (!user || !meta?.requiresRole?.includes(user.role))) {
    next('/denied')
  } else {
    updateUser(user as SessionUser)
    next()
  }
})

export default router
