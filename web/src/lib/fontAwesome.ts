import { library } from '@fortawesome/fontawesome-svg-core'
import {
  faRightFromBracket,
  faGraduationCap,
  faUserGraduate,
  faUser,
  faCheck,
  faXmark,
  faUserPlus,
  faChalkboard,
  faUsers,
  faSchool,
  faTrash,
  faCalendar,
  faBell,
  faTriangleExclamation,
  faAward,
  faChalkboardUser,
  faChalkboardTeacher
} from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import type { App } from 'vue'

export const setupIcons = (app: App<Element>) => {
  library.add(
    faRightFromBracket,
    faGraduationCap,
    faUserGraduate,
    faUser,
    faCheck,
    faXmark,
    faUserPlus,
    faChalkboard,
    faUsers,
    faSchool,
    faXmark,
    faTrash,
    faCalendar,
    faBell,
    faTriangleExclamation,
    faAward,
    faChalkboardTeacher,
    faChalkboardUser
  )
  app.component('font-awesome-icon', FontAwesomeIcon)
}
