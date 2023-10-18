import './assets/main.scss'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import { setupIcons } from './lib/fontAwesome'

const app = createApp(App)

app.use(createPinia())
app.use(router)

setupIcons(app)

app.mount('#app')