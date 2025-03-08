package controller

import (
	"ArquitecturaHexagonal_02/src/Loan/application/UseCase"
	"ArquitecturaHexagonal_02/src/Loan/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveBookAvailableController struct {
	saveBookAvailable *UseCase.BooksAvailable
}
func NewSaveBookAvailableController(saveBookAvailable *UseCase.BooksAvailable) *SaveBookAvailableController {
	return &SaveBookAvailableController{saveBookAvailable: saveBookAvailable}
}

func (controller *SaveBookAvailableController) Execute(ctx *gin.Context){
	var book domain.Book

	if err := ctx.ShouldBindJSON(&book); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	 err := controller.saveBookAvailable.Execute(book)

	 if err != nil{
		ctx.JSON(500, gin.H{"error": err.Error()})
	}else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Book saved"})
	}

	



}