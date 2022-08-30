import { createRouter, createWebHistory } from 'vue-router'
const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/login' },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login/index.vue')
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/Home/index.vue')
    }
  ]
})

// router.beforeEach((to, from, next) => {
//   if (to.fullPath != '/login') {
//     next('/login')
//   } else next()
// })

export default router
