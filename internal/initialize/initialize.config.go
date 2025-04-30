package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/Youknow2509/discord_webhook/internal/global"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

// get config in ./envaironment/.env
func GetConfig() {
	// Load file
	err := godotenv.Load("./environment/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the value of the environment variable
	urlWebhook := os.Getenv("URL_WEBHOOK")
	if urlWebhook == "" {
		log.Fatal("URL_WEBHOOK is not set in the .env file")
	}
	fmt.Println("URL_WEBHOOK:", urlWebhook)

	// Get port
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		log.Fatal("SERVER_PORT is not set in the .env file")
	}
	fmt.Println("SERVER_PORT:", serverPort)

	// set the variable to global variable
	global.URL_WEBHOOK = urlWebhook
	global.SERVER_PORT = serverPort
}
