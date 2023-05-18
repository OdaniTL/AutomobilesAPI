package injectors

import "AutomobilesAPI/controllers"

func GetAutosController() *controllers.AutosController {
	return autosController
}
