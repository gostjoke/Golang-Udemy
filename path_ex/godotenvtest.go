package path_ex

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Godotenvtest() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("a:", os.Getenv("a"))
	log.Println("SCERT:", os.Getenv("SCERT"))
}
