<script setup lang="ts">
import { RouterView, RouterLink } from 'vue-router';
</script>

<template>
  <div class="flex flex-col min-h-screen bg-gray-100 text-gray-900">
    <!-- Navbar -->
    <nav class="bg-blue-600 text-white p-4 shadow-md">
      <div class="container mx-auto flex justify-between items-center">
        <h1 class="text-lg font-bold">Stock App</h1>
        <div class="space-x-4">
          <RouterLink to="/" class="hover:underline">Inicio</RouterLink>
          <RouterLink to="/stocks" class="hover:underline">Stocks</RouterLink>
          <RouterLink to="/recommendations" class="hover:underline">Recomendaciones</RouterLink>
        </div>
      </div>
    </nav>

    <!-- Contenido dinámico -->
    <main class="flex-grow container mx-auto p-6">
      <RouterView />
    </main>

    <!-- Footer -->
    <footer class="bg-gray-800 text-white text-center p-4">
      © 2025 Stock App - Todos los derechos reservados
    </footer>
  </div>
</template>
