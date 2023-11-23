import './assets/main.scss'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { setupIcons } from './lib/fontAwesome'
import { setupCharts } from './lib/apesCharts'
import translatePlugin from './plugins/translate'

const app = createApp(App)

app.use(translatePlugin, {
  subject: {
    error: {
      exists: 'Podany przedmiot już istnieje',
      missingName: 'Proszę podać nazwę przedmiotu',
      cantDelete: 'Nie udało się usunąć przedmiotu'
    }
  },
  login: {
    error: {
      mismatch: 'Niepoprawny e-mail lub hasło',
      userInactive: 'Twoje konto jest nieaktywne, skontaktuj się z administratorem',
      fallback: 'Nie udało się zalogować'
    }
  }
})

app.use(createPinia())
app.use(router)

setupIcons(app)
setupCharts(app)

app.mount('#app')
