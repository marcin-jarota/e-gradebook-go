import './assets/main.scss'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { setupIcons } from './lib/fontAwesome'
import TranslatePlugin from './plugins/translate'

const app = createApp(App)

app.use(TranslatePlugin, {
  subject: {
    exists: 'Podany przedmiot już istnieje',
    missingName: 'Proszę podać nazwę przedmiotu'
  }
})

app.use(createPinia())
app.use(router)

setupIcons(app)

app.mount('#app')
