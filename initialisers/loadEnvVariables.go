package initialisers

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file... oh dear..")
		os.Setenv("PORT", "80")
	}
}
