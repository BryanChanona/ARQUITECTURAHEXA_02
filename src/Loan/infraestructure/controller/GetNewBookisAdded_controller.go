package controller

import (
	"ArquitecturaHexagonal_02/src/Loan/application/UseCase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookIsAddedController struct {
	useCase *UseCase.BookIsAddedUc
}

func NewBookIsAddedController(useCase *UseCase.BookIsAddedUc) *BookIsAddedController {
	return &BookIsAddedController{useCase: useCase}
}

func (controller *BookIsAddedController) Execute(ctx *gin.Context) {
	added, books, err := controller.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "added": false})
		return
	}

	if added {
		ctx.JSON(http.StatusOK, gin.H{"added": added, "books": books})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"added": added})
	}
}
