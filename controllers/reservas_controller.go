package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/galileoluna/serviceReservas/domain/reserva"
	"github.com/galileoluna/serviceReservas/gateways"
	"github.com/galileoluna/serviceReservas/services"
	"github.com/gorilla/mux"
)

var (
	ReservaController reservaControllerInterface = &reservaController{}
)

type reservaControllerInterface interface {
	InsertReserva(w http.ResponseWriter, r *http.Request)
	GetReserva(w http.ResponseWriter, r *http.Request)
	UpdateReserva(w http.ResponseWriter, r *http.Request)
	GetReservas(w http.ResponseWriter, r *http.Request)
	GetEventosDisponibles(w http.ResponseWriter, r *http.Request)
}

type reservaController struct {
}

func (c *reservaController) GetEventosDisponibles(w http.ResponseWriter, r *http.Request) {
	var eventos gateways.Evento
	fmt.Println("eventos")
	if err := json.NewDecoder(r.Body).Decode(&eventos); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	gateways.DecodeEventos()
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&eventos); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (c *reservaController) InsertReserva(w http.ResponseWriter, r *http.Request) {
	var reservaNueva reserva.Reserva
	err := json.NewDecoder(r.Body).Decode(&reservaNueva)
	fmt.Println("insertReserva")
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	services.ReservaService.InsertReserva(reservaNueva)
	w.WriteHeader(http.StatusCreated)
}
func (c *reservaController) GetReserva(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID_Reserva, _ := strconv.ParseInt(id, 10, 64)
	fmt.Println("getReserva")
	reservaRequerida, _ := services.ReservaService.GetReserva(ID_Reserva)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reservaRequerida); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (c *reservaController) UpdateReserva(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID64, _ := strconv.ParseInt(id, 10, 64)
	var reservaActualizada reserva.Reserva
	fmt.Println("actualizar")
	if ID64 < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&reservaActualizada); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	services.ReservaService.UpdateReserva(reservaActualizada, ID64)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&reservaActualizada); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
func (c *reservaController) GetReservas(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	ID64, _ := strconv.ParseInt(id, 10, 64)
	var reservaActualizada reserva.Reserva
	fmt.Println("reservas")
	if ID64 < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&reservaActualizada); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	services.ReservaService.GetReservas(ID64)
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&reservaActualizada); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
