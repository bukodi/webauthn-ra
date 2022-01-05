import Vue from 'vue';
import Router from 'vue-router';
import store from '../store';

import TopStories from '../views/TopStories.vue';
import CodeExamples from '../views/CodeExamples.vue';
import MyFavorites from '../views/MyFavorites.vue';
import Home from '../views/Home.vue';
import Login from '@/views/Login.vue';
import Register from '@/views/Register.vue';

Vue.use(Router);

class RouteMeta {
  title: string;

  constructor ({ title }: { title: string }) {
    this.title = title;
  }
}

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: new RouteMeta({ title: 'Home' })
    },
    {
      path: '/register',
      name: 'register',
      component: Register,
      meta: new RouteMeta({ title: 'Register' })
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: new RouteMeta({ title: 'Login' })
    },
    {
      path: '/top-stories',
      name: 'top-stories',
      component: TopStories,
      meta: new RouteMeta({ title: 'Top Stories' })
    },
    {
      path: '/code-examples',
      name: 'code-examples',
      component: CodeExamples,
      meta: new RouteMeta({ title: 'Code Examples' })
    },
    {
      path: '/my-favorites',
      name: 'my-favorites',
      component: MyFavorites,
      meta: new RouteMeta({ title: 'Favorites' })
    }
  ]
});

// This callback runs before every route change, including on initial load
router.beforeEach((to, from, next) => {
  const routeMeta = to.meta as RouteMeta;
  store.dispatch('topToolbar/changeTitle', routeMeta.title);
  document.title = routeMeta.title;
  next();
});

export default router;
