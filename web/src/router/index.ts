import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import jwtDecode from 'jwt-decode'
import { Role } from '@/types'
import type { SessionUser } from '@/types'
import { useSessionStore } from '@/stores/session'
import { userResource } from '@/resources/user'

export const routes: Record<
  string,
  RouteRecordRaw & {
    meta?: { roles: Role[]; title: string; requiresAuth: boolean; sidebar?: boolean; icon?: string }
  }
> = {
  home: {
    path: '/',
    name: 'home',
    redirect() {
      return { name: 'login' }
    }
  },
  login: {
    path: '/login',
    name: 'login',
    component: () => import('@/views/LoginView.vue')
  },
  setupPassword: {
    path: '/setup-password',
    name: 'setup-password',
    component: () => import('@/views/user/SetupPasswordView.vue')
  },
  start: {
    path: '/start',
    name: '/start',
    meta: { requiresAuth: true, roles: [], title: 'Start' },
    component: () => import('@/views/StartView.vue')
  },
  studentMarks: {
    path: '/student/marks',
    name: 'student.marks',
    meta: {
      requiresAuth: true,
      mainNav: true,
      roles: [Role.Student],
      sidebar: true,
      title: 'Oceny',
      icon: 'fa-graduation-cap'
    },
    component: () => import('@/views/student/MarksView.vue')
  },
  studentSubjects: {
    path: '/student/subjects',
    name: 'student.subjects',
    meta: {
      requiresAuth: true,
      mainNav: true,
      roles: [Role.Student],
      sidebar: true,
      title: 'Przedmioty',
      icon: 'fa-chalkboard'
    },
    component: () => import('@/views/student/SubjectsView.vue')
  },
  subjectList: {
    path: '/subject/list',
    name: 'subject.list',
    meta: {
      requiresAuth: true,
      roles: [Role.Admin],
      sidebar: true,
      title: 'Przedmioty',
      icon: 'fa-chalkboard'
    },
    component: () => import('@/views/subject/ListView.vue')
  },
  usetList: {
    path: '/user/list',
    name: 'user.list',
    meta: {
      requiresAuth: true,
      roles: [Role.Admin],
      sidebar: true,
      title: 'Użytkownicy',
      icon: 'fa-users'
    },
    component: () => import('@/views/user/ListView.vue')
  },
  userCreate: {
    path: '/user/create',
    name: 'user.create',
    meta: {
      requiresAuth: true,
      roles: [Role.Admin],
      sidebar: true,
      title: 'Dodawanie użytkownika',
      icon: 'user-plus'
    },
    component: () => import('@/views/user/CreateView.vue')
  },
  classGroupList: {
    path: '/class-group/list',
    name: 'class-group-list',
    meta: {
      requiresAuth: true,
      roles: [Role.Admin],
      sidebar: true,
      title: 'Klasy',
      icon: 'fa-school'
    },
    component: () => import('@/views/classGroup/ListView.vue')
  },
  classGroup: {
    path: '/class-group/:id',
    name: 'class-group',
    meta: {
      title: 'Dane klasy',
      roles: [Role.Admin, Role.Teacher],
      requiresAuth: true
    },
    component: () => import('@/views/classGroup/IndexView.vue')
  },
  accessDenied: {
    path: '/denied',
    name: 'denied',
    component: () => import('@/views/DeniedView.vue')
  },
  notFound: {
    path: '/:pathMatch(.*)*',
    component: () => import('@/views/PageNotFound.vue')
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: Object.values(routes)
})

router.beforeEach(async (to, _, next) => {
  const meta = to.meta as { requiresAuth?: boolean; roles?: Role[]; title?: string }
  const token = localStorage.getItem('token')

  if (!meta.requiresAuth) return next()

  if (meta.requiresAuth && (token === 'undefined' || !token)) {
    next('/login')
    return
  }

  const session = token ? jwtDecode<{ sessionUser: SessionUser }>(token) : null
  const user = session?.sessionUser

  try {
    await userResource.tokenValid(token as string)
  } catch (err) {
    next('/login')
    return
  }

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
