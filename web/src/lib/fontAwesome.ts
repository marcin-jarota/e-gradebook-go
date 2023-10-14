import { library } from '@fortawesome/fontawesome-svg-core'
import {
  faRightFromBracket,
  faGraduationCap,
  faUserGraduate
} from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import type { App } from 'vue'

export const setupIcons = (app: App<Element>) => {
  library.add(faRightFromBracket, faGraduationCap, faUserGraduate)
  app.component('font-awesome-icon', FontAwesomeIcon)
}
