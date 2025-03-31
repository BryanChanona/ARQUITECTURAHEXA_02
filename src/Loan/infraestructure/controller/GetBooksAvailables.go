package controller

import (
	"ArquitecturaHexagonal_02/src/Loan/application/UseCase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBooksAvailables struct {
	useCase *UseCase.GetBooksAvailables
}
func NewGetBooksAvailablesController(useCase *UseCase.GetBooksAvailables) *GetBooksAvailables {
	return &GetBooksAvailables{useCase: useCase}
}
func (controller *GetBooksAvailables) Execute(ctx *gin.Context) {
	books, err := controller.useCase.Execute()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"books": books})

}


