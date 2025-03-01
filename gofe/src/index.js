import { createRouter, createWebHistory } from 'vue-router'
 
const routes = [
  { 
    path: '/', 
    name: 'Home',
    component: () => import('./App2.vue')
  },
  {
    path: '/terms',
    name: 'Terms',
    component: () => import('./views/TermsPage.vue')
  },
  {
    path: '/privacy',
    name: 'Privacy',
    component: () => import('./views/PrivacyPage.vue')
  },
  {
    path: '/disclaimer',
    name: 'Disclaimer',
    component: () => import('./views/DisclaimerPage.vue')
  },
  {
    path: '/verify-email',
    name: 'VerifyEmail',
    component: () => import('./views/VerifyEmail.vue')
  },
  {
    path: '/reset-password',
    name: 'ResetPassword',
    component: () => import('./views/ResetPassword.vue')
  }
]  

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
