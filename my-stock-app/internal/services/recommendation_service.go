package services

import (
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/JMTeixeira17/my-stock-app/internal/models"
	"github.com/JMTeixeira17/my-stock-app/internal/storage"
)

// StockRecommendation contiene la recomendación
type StockRecommendation struct {
	Ticker          string  `json:"ticker"`
	Company         string  `json:"company"`
	PotentialUpside float64 `json:"potential_upside"`
	Rating          string  `json:"rating"`
}

func cleanAndParsePrice(priceStr string) (float64, error) {
	// Expresión regular para eliminar "$" y espacios
	re := regexp.MustCompile(`[$,\s]`)
	cleaned := re.ReplaceAllString(priceStr, "")

	// Convertir a float
	return strconv.ParseFloat(cleaned, 64)
}

// RecommendStocks analiza los datos de la base y sugiere las mejores inversiones
func RecommendStocks() ([]StockRecommendation, error) {
	db := storage.DB
	var stocks []models.Stock

	// Obtener los datos de la base de datos
	result := db.Find(&stocks)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(stocks) == 0 {
		return nil, nil
	}

	// Crear un slice de recomendaciones
	var recommendations []StockRecommendation

	for _, stock := range stocks {
		// Convertir TargetFrom y TargetTo a float64
		targetFrom, err1 := cleanAndParsePrice(stock.TargetFrom)
		targetTo, err2 := cleanAndParsePrice(stock.TargetTo)

		if err1 != nil || err2 != nil {
			log.Printf("Error parsing target prices for %s: %v, %v", stock.Ticker, err1, err2)
			continue // Saltamos este stock si hay error en la conversión
		}

		// Calcular el potencial de crecimiento
		potentialUpside := (targetTo - targetFrom) / targetFrom * 100

		// Agregar a la lista de recomendaciones
		recommendations = append(recommendations, StockRecommendation{
			Ticker:          stock.Ticker,
			Company:         stock.Company,
			PotentialUpside: math.Round(potentialUpside*100) / 100, // Redondear a 2 decimales
			Rating:          stock.RatingTo,
		})
	}

	// Ordenar las recomendaciones por mayor potencial de crecimiento
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].PotentialUpside > recommendations[j].PotentialUpside
	})

	if len(recommendations) > 10 {
		recommendations = recommendations[:10]
	}

	return recommendations, nil
}
