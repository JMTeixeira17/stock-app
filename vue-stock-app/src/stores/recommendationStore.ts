import { defineStore } from 'pinia';
import axios from 'axios';

export const useRecommendationStore = defineStore('recommendationStore', {
  state: () => ({
    recommendations: [] as Array<any>,  // Estado inicial, asegurando que recommendations es un array vacío
    loading: false,  // Estado de carga
  }),

  actions: {
    async fetchRecommendations() {
      this.loading = true;  // Establecer estado de carga
      try {
        const response = await axios.get('http://localhost:8080/recommendations'); // Llamada al endpoint de recomendaciones
        console.log("API Response:", response.data); // Verifica que la respuesta es la esperada
        this.recommendations = response.data; // Asigna los datos recibidos a 'recommendations'
      } catch (error) {
        console.error("Error fetching recommendations:", error); // Muestra el error si ocurre
      } finally {
        this.loading = false;  // Finaliza el estado de carga
      }
    },
  },
});
