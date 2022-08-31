import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/Login/Login.vue'),
      meta: {
        title: 'login'
      }
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('@/views/Home/Home.vue'),

      meta: {
        title: 'home'
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  // 动态title
  document.title = to.meta.title
  next()
})

export default router
