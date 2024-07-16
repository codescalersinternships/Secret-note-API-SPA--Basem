import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import HomePage from '@/views/Home.vue';
import UserLogin from '@/components/UserLogin.vue'
import UserRegister from '@/components/UserRegister.vue'
import CreateNote from '@/components/CreateNote.vue';
import ViewNote from '@/components/ViewNote.vue';
import UserDashboard from '@/components/UserDashboard.vue';



const routes: Array<RouteRecordRaw> = [
  { path: '/', component: HomePage },
  { path: '/login', component: UserLogin },
  { path: '/register', component: UserRegister },
  { path: '/note', component: CreateNote },
  { path: '/note/:key', component: ViewNote },
  { path: '/dashboard', component: UserDashboard }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router
