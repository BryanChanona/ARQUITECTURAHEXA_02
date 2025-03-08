package dependencies

import (
	"ArquitecturaHexagonal_02/src/Loan/application/UseCase"
	"ArquitecturaHexagonal_02/src/Loan/infraestructure"
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/controller"
	"ArquitecturaHexagonal_02/src/helpers"
	"log"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := helpers.ConnectDB()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL =*infraestructure.NewMySQL(db)


}

func GetBooksAvailable() *controller.SaveBookAvailableController {
	caseSaveBookAvailable := UseCase.NewSaveBookAvailableUseCase(&mySQL)
	return controller.NewSaveBookAvailableController(caseSaveBookAvailable)
}