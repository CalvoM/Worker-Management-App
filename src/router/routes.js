const routes = [{
        path: '/',
        name: 'home',
        component: () => import( /* webpackChunkName: "home" */ '../views/Home.vue')
    },
    {
        path: '/worker_list/',
        name: 'about',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import( /* webpackChunkName: "about" */ '../views/WorkersList.vue')
    }
]
export{
    routes
}