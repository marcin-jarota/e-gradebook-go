import { createRouter, createWebHistory } from 'vue-router'
import jwtDecode from 'jwt-decode'
import { Role } from '@/types'
import type { SessionUser } from '@/types'
import { useSessionStore } from '@/stores/session'

type Route =
  | 'login'
  | 'student'
  | 'start'
  | 'profile'
  | 'studentMarks'
  | 'subjectList'
  | 'studentList'

export const routes: Record<Route, { path: string; name: string }> = Object.freeze({
  login: {
    path: '/login',
    name: 'login'
  },
  student: {
    path: '/student',
    name: 'student'
  },
  studentMarks: {
    path: '/student/marks',
    name: 'student-marks'
  },
  start: {
    path: '/start',
    name: 'start'
  },
  profile: {
    path: '/profile',
    name: 'profile'
  },
  subjectList: {
    path: '/subject/list',
    name: 'subject-list'
  },
  studentList: {
    path: '/student/list',
    name: 'student-list'
  }
})

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
      path: routes.login.path,
      name: routes.login.name,
      component: () => import('@/views/LoginView.vue')
    },
    {
      path: '/student',
      name: 'student',
      component: () => import('@/views/StudentView.vue'),
      meta: { requiresAuth: true, roles: [Role.Student], title: 'Student' }
    },
    {
      path: '/denied',
      name: 'denied',
      component: () => import('@/views/DeniedView.vue')
    },
    {
      path: routes.start.path,
      name: routes.start.name,
      component: () => import('@/views/StartView.vue'),
      meta: { requiresAuth: true, roles: [Role.Student, Role.Admin], title: 'Start' }
    },
    {
      path: routes.studentMarks.path,
      name: routes.studentMarks.name,
      component: () => import('@/views/student/MarksView.vue'),
      meta: { requiresAuth: true, roles: [Role.Student], title: 'Oceny' }
    },
    {
      path: routes.subjectList.path,
      name: routes.subjectList.name,
      component: () => import('@/views/subject/ListView.vue'),
      meta: { requiresAuth: true, roles: [Role.Admin], title: 'Przedmioty' }
    },
    {
      path: routes.studentList.path,
      name: routes.studentList.name,
      component: () => import('@/views/student/ListView.vue'),
      meta: { requiresAuth: true, roles: [Role.Admin] }
    },
    {
      path: '/:pathMatch(.*)*',
      component: () => import('@/views/PageNotFound.vue')
    }
  ]
})

router.beforeEach((to, _, next) => {
  const meta = to.meta as { requiresAuth?: boolean; roles?: Role[]; title?: string }
  const token = localStorage.getItem('token')

  if (!meta.requiresAuth) return next()

  if (meta.requiresAuth && (token === 'undefined' || !token)) {
    next('/login')
    return
  }

  const session = token ? jwtDecode<{ sessionUser: SessionUser }>(token) : null
  const user = session?.sessionUser

  const { updateUser } = useSessionStore()

  if (meta.title) {
    document.title = meta.title
  }

  if (meta?.roles?.length && (!user || !meta?.roles?.includes(user.role))) {
    next('/denied')
  } else {
    updateUser(user as SessionUser)
    next()
  }
})

export default router
