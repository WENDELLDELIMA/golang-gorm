package main

import (
	"fmt"

	"github.com/WENDELLDELIMA/go-rest-api/database"
	"github.com/WENDELLDELIMA/go-rest-api/routes"
)



func main() {
	
	database.ConectaComBancoDeDados()
	fmt.Println("Starting")
	routes.HandleRequest()
}