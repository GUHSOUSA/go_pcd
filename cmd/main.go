package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gustasousaghs/carros-pcd-service/config"
	"github.com/gustasousaghs/carros-pcd-service/router"
)

func main() {
	config.MysqlConnect()

	port := os.Getenv("PORT") // Railway define essa variável automaticamente
	if port == "" {
		port = "8080"
	}

	r := router.InitRouter()
	fmt.Println("Servidor rodando na porta " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
