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
  faUsers
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
    faUsers
  )
  app.component('font-awesome-icon', FontAwesomeIcon)
}
