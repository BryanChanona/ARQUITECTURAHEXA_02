package dependencies

import (
	service "ArquitecturaHexagonal_02/src/Loan/application/Service"
	"ArquitecturaHexagonal_02/src/Loan/application/UseCase"
	"ArquitecturaHexagonal_02/src/Loan/infraestructure"
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/adapters"
	"ArquitecturaHexagonal_02/src/Loan/infraestructure/controller"
	"ArquitecturaHexagonal_02/src/helpers"
	"log"
)

var (
	mySQL infraestructure.MySQL
	eventService *service.Event 
)

func Init() {
	db, err := helpers.ConnectDB()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL =*infraestructure.NewMySQL(db)
	rabbit := adapters.NewRabbit() // Crear la conexión Rabbit
	eventService = service.NewEvent(rabbit) // Inicializar el servicio Event con Rabbit


}

func SaveBooksAvailable() *controller.SaveBookAvailableController {
	caseSaveBookAvailable := UseCase.NewSaveBookAvailableUseCase(&mySQL)
	return controller.NewSaveBookAvailableController(caseSaveBookAvailable,eventService)
}
func GetBooksAvailables() *controller.GetBooksAvailables {
	caseGetBooksAvailables := UseCase.NewGetBooksAvailableUseCase(&mySQL)
	return controller.NewGetBooksAvailablesController(caseGetBooksAvailables)
}
func GetNewBookIdAddedController() *controller.BookIsAddedController{
	useCase := UseCase.NewNewBookIsAddedUc(&mySQL)
	return controller.NewBookIsAddedController(useCase)
}