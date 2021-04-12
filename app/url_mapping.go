package app

import (
	"net/http"

	"github.com/galileoluna/serviceReservas/controllers"
)

func mapUrls() {

	resevasRouter := router.PathPrefix("/reserva").Subrouter()
	resevasRouter.Path("").Methods(http.MethodPost).HandlerFunc(controllers.ReservaController.InsertReserva)
	resevasRouter.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controllers.ReservaController.UpdateReserva)
	resevasRouter.Path("/data/{id}").Methods(http.MethodGet).HandlerFunc(controllers.ReservaController.GetReserva)
	resevasRouter.Path("/evento/{id}").Methods(http.MethodGet).HandlerFunc(controllers.ReservaController.GetReservas)
	resevasRouter.Path("/evento/disponibles").Methods(http.MethodGet).HandlerFunc(controllers.ReservaController.GetEventosDisponibles)
}
