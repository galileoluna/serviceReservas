package services

import "github.com/galileoluna/serviceReservas/domain/reserva"

var (
	ReservaService reservaServiceInterface = &reservaService{}
)

type reservaServiceInterface interface {
	InsertReserva(reserva reserva.Reserva) (*reserva.Reserva, error)
	GetReserva(reservaID int64) (*reserva.Reserva, error)
	UpdateReserva(reservaActualizada reserva.Reserva, reservaID int64) (*reserva.Reserva, error)
	GetReservas(id_evento int64) ([]reserva.Reserva, error)
}

type reservaService struct {
}

func (s *reservaService) InsertReserva(nuevaReserva reserva.Reserva) (*reserva.Reserva, error) {
	if err := nuevaReserva.Insert(); err != nil {
		return nil, err
	}
	return &nuevaReserva, nil
}
func (s *reservaService) GetReserva(reservaID int64) (*reserva.Reserva, error) {
	var reservaRequerida reserva.Reserva
	if _, err := reservaRequerida.GetReserva(reservaID); err != nil {
		return nil, err
	}
	reservaRequerida, _ = reservaRequerida.GetReserva(reservaID)
	return &reservaRequerida, nil
}
func (s *reservaService) UpdateReserva(reservaActualizada reserva.Reserva, reservaID int64) (*reserva.Reserva, error) {
	if err := reservaActualizada.UpdateReserva(reservaID); err != nil {
		return nil, err
	}
	return &reservaActualizada, nil
}
func (s *reservaService) GetReservas(id_evento int64) ([]reserva.Reserva, error) {
	var reservaRequerida reserva.Reserva

	reservasRequerida, _ := reservaRequerida.GetReservas(id_evento)

	return reservasRequerida, nil
}
