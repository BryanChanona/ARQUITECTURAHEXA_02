package main

import (
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/dependencies"
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/routes"
	"ArquitecturaHexagonal_02/src/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()

	r := gin.Default()
	helpers.InitCORS(r)

	routes.BookRoutes(r)
	r.Run()
	

}