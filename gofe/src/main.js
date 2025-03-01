import { createApp } from 'vue'
import App from './App.vue'
import VueMathjaxNext from 'vue-mathjax-next'
import VueCookies from 'vue-cookies'
import './assets/styles.css'
import router from '.'

const app = createApp(App)
app.use(router)
   .use(VueMathjaxNext)
   .use(VueCookies)
   .mount('#app')
