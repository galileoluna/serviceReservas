package reserva

type Reserva struct {
	ID             int64   `json:"id_reserva"`
	ID_Evento      int64   `json:"id_evento"`
	ID_Cliente     int64   `json:"id_cliente"`
	Nombre_Cliente string  `json:"nombre_cliente"`
	Monto          float64 `json:"monto_reserva"`
}

func NewReserva(id int64, id_evento int64, id_cliente int64, nombre_cliente string, monto float64) Reserva {
	return Reserva{
		ID:             id,
		ID_Evento:      id_evento,
		ID_Cliente:     id_cliente,
		Nombre_Cliente: nombre_cliente,
		Monto:          monto,
	}
}
