<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useStockStore } from '@/stores/stockStore';

const stockStore = useStockStore();
const loading = ref(true);

onMounted(async () => {
  await stockStore.fetchStocks();
  console.log(stockStore.stocks.data[0]);
  loading.value = false;
});
</script>

<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Lista de Stocks</h1>
    <!-- Mensaje de Cargando -->
    <div v-if="loading">Cargando...</div>
    
    <!-- Mostrar la tabla solo si hay datos -->
    <div v-else>
      <table v-if="stockStore.stocks.data.length" class="w-full border-collapse border border-gray-300">
        <thead>
          <tr class="bg-gray-100">
            <th class="border p-2">Ticker</th>
            <th class="border p-2">Company</th>
            <th class="border p-2">Target From</th>
            <th class="border p-2">Target To</th>
          </tr>
        </thead>
        <tbody>
          <!-- Iterar sobre los datos de stock -->
          <tr v-for="stock in stockStore.stocks.data" :key="stock.id" class="hover:bg-gray-50">
            <td class="border p-2">{{ stock.Ticker }}</td>
            <td class="border p-2">{{ stock.Company }}</td>
            <td class="border p-2">{{ stock.TargetFrom }}</td>
            <td class="border p-2">{{ stock.TargetTo }}</td>
          </tr>
        </tbody>
      </table>
      <!-- Mostrar mensaje si no hay datos -->
      <div v-else>No se encontraron stocks</div>
    </div>
  </div>
</template>

<style scoped>
/* Estilos personalizados para la tabla */
table {
  border-collapse: collapse;
  width: 100%;
}
th, td {
  padding: 0.75rem;
  text-align: left;
}
th {
  background-color: #f0f0f0;
}
tr:hover {
  background-color: #f9f9f9;
}
</style>

