package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBancoDados() {
	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao), &gorm.Config{})
	if err != nil {
		log.Fatalf("Não foi possível conectar ao banco de dados: %v", err)
	}
}
