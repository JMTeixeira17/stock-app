<script setup lang="ts">
import { onMounted } from 'vue';
import { useRecommendationStore } from '@/stores/recommendationStore';

const recommendationStore = useRecommendationStore(); // Usamos el store de recomendaciones

// Llamamos a la acción fetchRecommendations al montar el componente
onMounted(() => {
  console.log(1)
  recommendationStore.fetchRecommendations();
});
</script>

<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Recomendaciones de Inversión</h1>
    
    <!-- Mensaje de carga -->
    <div v-if="recommendationStore.loading">Cargando recomendaciones...</div>
    
    <!-- Tabla con las recomendaciones -->
    <table v-else class="w-full border-collapse border border-gray-300">
      <thead>
        <tr class="bg-gray-100">
          <th class="border p-2">Ticker</th>
          <th class="border p-2">Compañía</th>
          <th class="border p-2">Potencial de Aumento</th>
          <th class="border p-2">Calificación</th>
        </tr>
      </thead>
      <tbody>
        <!-- Iteramos sobre las recomendaciones y las mostramos en la tabla -->
        <tr v-for="recommendation in recommendationStore.recommendations" :key="recommendation.ticker" class="hover:bg-gray-50">
          <td class="border p-2">{{ recommendation.ticker }}</td>
          <td class="border p-2">{{ recommendation.company }}</td>
          <td class="border p-2">{{ recommendation.potential_upside }}%</td>
          <td class="border p-2">{{ recommendation.rating }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
/* Aquí puedes agregar estilos adicionales si lo necesitas */
</style>
