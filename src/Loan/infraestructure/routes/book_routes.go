package routes

import (
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	routes := router.Group("/books")
	saveBookAvailable := dependencies.SaveBooksAvailable().Execute
	getBooksAvailable := dependencies.GetBooksAvailables().Execute
	getNewBooksIsAdded := dependencies.GetNewBookIdAddedController().Execute

	routes.POST("/",saveBookAvailable)
	routes.GET("/", getBooksAvailable )
	routes.GET("/newBookIsAdded",getNewBooksIsAdded)
}