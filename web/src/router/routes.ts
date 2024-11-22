import { RouteRecordRaw } from 'vue-router';

const inicio = () => { 
  if (process.env.SHOW_LANDING_PAGE) return '/login';
  return '/';
};

const inicio2 = () => {
  if (process.env.SHOW_LANDING_PAGE) return '/';
  return '/index';
};

const routes: RouteRecordRaw[] = [
  {
    path: inicio(),
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      {
        path: '/usuarios',
        component: async () =>
          await import('pages/xauth/usuarios/index_usuarios.vue'),
      },
      {
        path: '/roles',
        component: async () =>
          await import('pages/xauth/roles/roles-index.vue'),
      },
    ],
  },
  {
    path: inicio2(),
    component: () => import('layouts/LandingLayout.vue'),
    children: [],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
