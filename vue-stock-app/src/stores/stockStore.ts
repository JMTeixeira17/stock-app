import { defineStore } from 'pinia';
import axios from 'axios';


interface Stock {
  id: number;
  Ticker: string;
  Company: string;
  TargetFrom: string;
  TargetTo: string;
}



export const useStockStore = defineStore('stockStore', {
  state: () => ({
    stocks: { data:[] as Stock[] } // Estado inicial, asegurando que stocks.data es un array vacío
  }),

  actions: {
    async fetchStocks() {
      try {
        const response = await axios.get('http://localhost:8080/stocks?page=1&limit=10');
        console.log("API Response:", response.data); // Verifica que la respuesta contiene 'data'
        this.stocks = response.data; // Asigna los datos recibidos a 'stocks'
      } catch (error) {
        console.error("Error fetching stocks:", error);
      }
    }
  }
});