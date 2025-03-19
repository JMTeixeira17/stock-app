package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JMTeixeira17/my-stock-app/internal/storage"
)

const apiURL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"

// Struct para el stock que vamos a manejar
type Stock struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type APIResponse struct {
	Items    []Stock `json:"items"`
	NextPage string  `json:"next_page"`
}

func cleanPrice(price string) (float64, error) {
	// Eliminar el signo de dólar
	price = strings.Replace(price, "$", "", -1)

	// Convertir el precio a float64
	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, err
	}

	return floatPrice, nil
}

// Función para consumir la API y obtener los datos de los stocks
func GetStocksFromAPI() ([]Stock, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	var allStocks []Stock
	nextPage := ""

	for {
		// Construir la URL con el parámetro next_page si existe
		url := apiURL
		if nextPage != "" {
			url = fmt.Sprintf("%s?next_page=%s", apiURL, nextPage)
		}

		log.Printf("Requesting page: %s", nextPage)

		// Realizar la solicitud GET
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		// Añadir el header de autenticación
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6OCwiZW1haWwiOiJqdGVpeGVpcmExMTAyQGdtYWlsLmNvbSIsImV4cCI6MTc0MTYxNDE3MywiaWQiOiIwIiwicGFzc3dvcmQiOiInIE9SICcxJz0nMSJ9.O1ORkshSItIUMM-waPvlskcHcSB2I6nGAthPCtV1yiE")

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Si no es un 200 OK, devolver un error
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
		}

		// Decodificar la respuesta
		var apiResponse APIResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResponse)
		if err != nil {
			return nil, err
		}

		log.Printf("Received %d items from page: %s", len(apiResponse.Items), nextPage)

		// Agregar los stocks obtenidos a la lista general
		allStocks = append(allStocks, apiResponse.Items...)

		// Si no hay más página (next_page es vacío), hemos terminado
		if apiResponse.NextPage == "" {
			break
		}

		// Actualizamos next_page para continuar con la siguiente página
		nextPage = apiResponse.NextPage
	}

	log.Printf("Total items fetched: %d", len(allStocks))

	return allStocks, nil
}

func SaveStocksToDB(stocks []Stock) error {
	for _, stock := range stocks {
		// Intentamos obtener el stock con el ticker proporcionado o lo creamos si no existe
		if err := storage.DB.FirstOrCreate(&stock, Stock{Ticker: stock.Ticker}).Error; err != nil {
			log.Printf("Error saving stock %s: %v", stock.Ticker, err)
			return err
		}

		// Log para saber si el stock fue insertado o ya existía
		log.Printf("Stock %s processed successfully.", stock.Ticker)
	}

	log.Println("Stocks saved to the database successfully!")
	return nil
}
