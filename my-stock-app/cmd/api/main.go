package main

import (
	"log"
	"time"

	"github.com/JMTeixeira17/my-stock-app/internal/handlers"
	"github.com/JMTeixeira17/my-stock-app/internal/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	storage.InitDB()
	defer storage.CloseDB()

	// Inicializar el router de Gin
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Permite el frontend en Vite
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Definir rutas en el handler
	api := router.Group("/stocks")
	{
		api.GET("/fetch", handlers.FetchAndStoreStocksHandler) // Obtiene y guarda datos desde la API
		api.GET("", handlers.GetStocksFromDBHandler)
		router.GET("/recommendations", handlers.RecommendStocksHandler) // Obtiene datos de la BD (paginado)
	}

	// Iniciar el servidor
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
