package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conexao() {
	endereco := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(endereco))
	if err != nil {
		log.Panic("Erro na conexão com banco de dados!!")
	}
}
