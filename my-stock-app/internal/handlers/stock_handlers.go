package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/JMTeixeira17/my-stock-app/internal/models"
	"github.com/JMTeixeira17/my-stock-app/internal/services"
	"github.com/JMTeixeira17/my-stock-app/internal/storage"
	"github.com/gin-gonic/gin"
)

// FetchAndStoreStocksHandler obtiene los datos de la API y los almacena en la BD
func FetchAndStoreStocksHandler(c *gin.Context) {
	stocks, err := services.GetStocksFromAPI()
	if err != nil {
		log.Printf("Error fetching stocks from API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stocks"})
		return
	}

	err = services.SaveStocksToDB(stocks)
	if err != nil {
		log.Printf("Error saving stocks to DB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save stocks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stocks fetched and saved successfully!"})
}

func GetStocksFromDBHandler(c *gin.Context) {
	// Obtener parámetros de paginación
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	var stocks []models.Stock
	if err := storage.DB.Limit(limit).Offset(offset).Find(&stocks).Error; err != nil {
		log.Printf("Error retrieving stocks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving stocks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"data":  stocks,
	})
}

func RecommendStocksHandler(c *gin.Context) {
	recommendations, err := services.RecommendStocks()
	log.Printf("hola", recommendations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing recommendations"})
		return
	}

	if recommendations == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No stocks available for recommendation"})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}
