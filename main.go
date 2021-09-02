package main

import (
	"lemonilo/controllers"
	"log"
	"os"

	"lemonilo/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
	}

	appName := os.Getenv("APP_NAME")

	log.Println("[*] Running the " + appName)

	controllers.Init()

	router.Route()

}
