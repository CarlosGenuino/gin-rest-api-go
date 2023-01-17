package database

import (
	"log"

	"github.com/CarlosGenuino/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBancoDeDados() {
	stringDeConexao := "host=localhost user=develop password=1234 dbname=alunos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar no banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
