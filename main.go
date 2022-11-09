package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SOOA-swarch-2022ii/sooa-subjects/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("Inicializando microservicio de usuarios")
	enrutador := routes.Routes()

	http.ListenAndServe(os.Getenv("PORT"), enrutador)
}
