package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

// Configuración de la base de datos
func InitDB() {
	err := godotenv.Load() // Cargar variables de entorno
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Leer la URL de conexión a la base de datos desde las variables de entorno
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is required in .env file")
	}

	// Conectarse a la base de datos
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Asignar la conexión a la variable global DB
	DB = db
	fmt.Println("Database connected successfully!")
}

// Cerrar la base de datos
func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatal("Error closing database:", err)
	}
}
