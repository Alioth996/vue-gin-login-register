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
      component: () => import('@/views/Login/index.vue'),
      meta: {
        title: 'login'
      }
    },
    {
      path: '/home/:username',
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
  // 非跳转登录页并且没有登录,强制跳转login
  const re = localStorage.getItem('token')
  if (to.fullPath != '/login' && !re) {
    next('/login')
  } else next()
})

export default router
