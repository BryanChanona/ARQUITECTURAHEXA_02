package controller

import (
	service "ArquitecturaHexagonal_02/src/Loan/application/Service"
	"ArquitecturaHexagonal_02/src/Loan/application/UseCase"
	"ArquitecturaHexagonal_02/src/Loan/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveBookAvailableController struct {
	saveBookAvailable *UseCase.BooksAvailable
	event      *service.Event
}
func NewSaveBookAvailableController(saveBookAvailable *UseCase.BooksAvailable,event *service.Event) *SaveBookAvailableController {
	return &SaveBookAvailableController{
		saveBookAvailable: saveBookAvailable,
		event:      event,
	}
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
	
	err = controller.event.Execute(book)
	if err != nil {
		// Puedes registrar este error en los logs si el evento no se envió correctamente
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending event to RabbitMQ"})
	}
	
	



}