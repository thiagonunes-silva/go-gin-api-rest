package main

import (
	"github.com/thiagonunes.silva/go-gin-api-rest/database"
	"github.com/thiagonunes.silva/go-gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
