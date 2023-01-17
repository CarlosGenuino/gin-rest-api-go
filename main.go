package main

import (
	"github.com/CarlosGenuino/gin-api-rest/database"
	"github.com/CarlosGenuino/gin-api-rest/routes"
)

func main() {
	database.ConectaBancoDeDados()
	routes.HandleRequests()
}
