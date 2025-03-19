import { createRouter, createWebHistory } from 'vue-router';
import StockList from '@/components/StockList.vue';
import StockRecommendations from '@/components/StockRecommendations.vue';

const routes = [
  { path: '/', redirect: '/stocks' },
  { path: '/stocks', component: StockList },
  { path: '/recommendations', component: StockRecommendations }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
