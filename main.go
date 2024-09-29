package main

import (
	"github.com/victornevesHS/go-rest-gin/database"
	"github.com/victornevesHS/go-rest-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
