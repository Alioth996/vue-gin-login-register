import { ElMessage } from 'element-plus'
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
      path: '/home',
      name: 'home',
      component: () => import('@/views/Home/Home.vue'),
      redirect: '/dashboard',
      children: [
        {
          name: 'dashboard',
          path: '/dashboard',
          component: () => import('@/views/Dashboard/index.vue')
        }
      ],
      meta: {
        title: 'home'
      }
    },
    {
      path: '/404',
      name: 'PageNoExist',
      component: () => import('../views/NotFound.vue'),
      meta: {
        title: 'Page not found'
      }
    },
    {
      path: '/:catchAll(.*)',
      redirect: {
        name: 'PageNoExist'
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  // 动态title
  document.title = to.meta.title ?? 'welcome'
  // 非跳转登录页并且没有登录,强制跳转login
  const re = localStorage.getItem('token')
  if (to.fullPath != '/login' && !re) {
    ElMessage.error('未登录!!')
    next('/login')
  } else next()
})

export default router
