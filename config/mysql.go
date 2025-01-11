package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gustasousaghs/carros-pcd-service/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MysqlConnect() {
	// Evitar carregar .env em produção
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Erro ao carregar .env")
		}
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var dbErr error
	DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", dbErr)
	}
	if err := DB.AutoMigrate(&models.Car{}); err != nil {
		log.Fatal("Erro ao migrar o banco de dados:", err)
	}
	fmt.Println("Banco de dados conectado com sucesso!")
}
