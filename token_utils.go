package alphavantage

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetAPIToken() (string, error) {
	var err error
	var token string

	err = godotenv.Load()
	if err != nil {
		return token, errors.New("Error loading .env file")
	}

	// Get the user's API token from env. var.
	token = os.Getenv("ALPHAVANTAGE_TOKEN")

	return token, err
}
