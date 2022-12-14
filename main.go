package main

import (
	"comm-api/connector"
	"comm-api/models"
	"comm-api/router"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Starting Communicating Api....")

	err1 := godotenv.Load()

	if err1 != nil {
		println(err1)
		return

	}

	models.COLLECTION_NAME = os.Getenv("COLLECTION_NAME")
	models.DB_NAME = os.Getenv("DB_NAME")
	models.DB_URL = os.Getenv("DB_URL")

	// fmt.Println(models.DB_URL)
	// fmt.println(models.DB_URL)

	connector.Init()

	r := router.Router()
	err := http.ListenAndServe(":4000", r)

	if err != nil {
		panic("Problem starting the server")
	}

}
