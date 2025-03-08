package routes

import (
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	routes := router.Group("/books")
	saveBookAvailable := dependencies.GetBooksAvailable().Execute

	routes.POST("/",saveBookAvailable)
}