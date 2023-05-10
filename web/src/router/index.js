import { createRouter,  createWebHistory } from 'vue-router'



const base = import.meta.env.VITE_PUBLIC_PATH

const routes = [
  {
      path: '/',
      name: 'login',
      component: () => import('@/views/login.vue')

  },
  {
      path:"/detail",
      name:"detail",
      component:()=> import('@/views/detail.vue'),
      meta:{
        requireAuth: true
      }
  }
]



const router = createRouter({
 history: createWebHistory(`${base}/`),
 routes: routes
})


router.beforeEach((to, from, next) => {

  if(to.meta.requireAuth){
    if (sessionStorage.getItem("search-token")) {
          if(to.path == '/'){
            sessionStorage.removeItem('search-token')
            next({path: '/'})
          }else{
            next()
            return
          }
    }else{
      sessionStorage.removeItem('search-token')
      next({path: '/'})
    }
  }else{
    next()
  }

    
})


export const setupRouter = async(app) => {
    app.use(router);
    await router.isReady();
  }

export default router
