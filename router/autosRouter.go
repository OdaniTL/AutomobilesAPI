package router

import (
	"AutomobilesAPI/injectors"
	"net/http"
)

func RouterProcut() {
	actrl := injectors.GetAutosController()
	http.HandleFunc("/api/autos/getAll", actrl.GetAll)
	http.HandleFunc("/api/autos/create", actrl.Create)
	http.HandleFunc("/api/autos/update", actrl.Update)
	http.HandleFunc("/api/autos/delete", actrl.Delete)
	http.HandleFunc("/api/autos/findById", actrl.FindById)
	http.HandleFunc("/api/autos/findByName", actrl.FindByName)
}
