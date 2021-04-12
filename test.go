package main

import (
	"fmt"

	"github.com/galileoluna/serviceReservas/datasources"
	"github.com/galileoluna/serviceReservas/domain/reserva"
	"github.com/galileoluna/serviceReservas/gateways"
	"github.com/galileoluna/serviceReservas/services"
)

func ain() {
	datasources.Init()
	reserv1 := reserva.NewReserva(1, 1, 2, "astor", 10.5)
	services.ReservaService.InsertReserva(reserv1)
	fmt.Println(gateways.DecodeEventos())
	fmt.Println(services.ReservaService.GetReserva(2))
	fmt.Println(services.ReservaService.GetReservas(1))
}
