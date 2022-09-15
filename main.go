package main

import (
	"github.com/rafaelbraga/api-go-gin/database"
	"github.com/rafaelbraga/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()

}
