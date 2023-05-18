package injectors

import (
	"AutomobilesAPI/configs"
	"AutomobilesAPI/controllers"
	"AutomobilesAPI/repository"
	"AutomobilesAPI/services"
	"database/sql"
	"log"
)

var (

	//repos
	autosRepository repository.IAutosRepository

	// services
	autoServices services.IAutoServices

	//controllers
	autosController *controllers.AutosController
)

func loadInjectors() {

	connection, errs := configs.NewConnection()
	if errs != nil {
		log.Fatalf("The connection are si broken!! %s", errs)
	}
	initRepository(connection)
	initServices()
	initControls()
}

func initRepository(connection *sql.DB) {

	autosRepository = repository.NewAutosRepository(connection)
}

func initServices() {
	autoServices = services.NewAutoServices(autosRepository)
}

func initControls() {
	autosController = controllers.NewAutosController(autoServices)
}

// It loads all the injectors in the `injectors` directory and then calls the `Init` function on each
// of them

func InitApp() {
	loadInjectors()
}
