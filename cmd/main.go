package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gustasousaghs/carros-pcd-service/config"
	"github.com/gustasousaghs/carros-pcd-service/router"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Arquivo .env não encontrado, continuando com variáveis de ambiente do sistema")
		}
	}
	config.MysqlConnect()

	port := os.Getenv("PORT") // Railway define essa variável automaticamente
	if port == "" {
		port = "8080"
	}

	r := router.InitRouter()
	fmt.Println("Servidor rodando na porta " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
