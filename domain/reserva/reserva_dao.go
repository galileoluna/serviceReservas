package reserva

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/galileoluna/serviceReservas/datasources"
)

const (
	queryInsertReserva   = "INSERT INTO Reserva (ID_Evento,ID_Cliente, Nombre_Cliente,Monto) VALUES (@ID_Evento,@ID_Cliente,@Nombre_Cliente,@Monto);"
	queryGetReservas     = "SELECT  *  FROM Reserva "
	queryGetReservaByID  = "SELECT  *  FROM Reserva WHERE ID=@ID;"
	queryUdpdateReservas = "UPDATE Reserva SET  Nombre_Cliente = @Nombre_Cliente , Monto = @Monto  WHERE ID = @ID;"
)

func (nReserva *Reserva) Insert() error {
	ctx := context.Background()
	var err error

	if datasources.Db == nil {
		err = errors.New("Reservas: db is null")
		return err
	}

	err = datasources.Db.PingContext(ctx)
	if err != nil {
		return err
	}

	tsql := fmt.Sprintf(queryInsertReserva)

	result, err := datasources.Db.ExecContext(
		ctx,
		tsql,
		sql.Named("ID_Evento", nReserva.ID_Evento),
		sql.Named("ID_Cliente", nReserva.ID_Cliente),
		sql.Named("Nombre_Cliente", nReserva.Nombre_Cliente),
		sql.Named("Monto", nReserva.Monto))

	fmt.Println(result)

	return nil
}

func (reser *Reserva) GetReserva(id_encuesta int64) (Reserva, error) {
	ctx := context.Background()
	var err error

	var reservaBuscada Reserva

	if datasources.Db == nil {
		err = errors.New("problema de conexion")
		return reservaBuscada, err
	}

	err = datasources.Db.PingContext(ctx)
	if err != nil {
		err = errors.New("Error interno")
		return reservaBuscada, err
	}
	tsql := fmt.Sprintf(queryGetReservaByID)

	// Execute query
	rows, err := datasources.Db.QueryContext(ctx, tsql, sql.Named("ID", id_encuesta))
	if err != nil {
		err = errors.New("Ingrese los parametros adecuados")
		return reservaBuscada, err
	}

	defer rows.Close()

	for rows.Next() {
		var Nombre_Cliente string
		var ID, ID_Evento, ID_Cliente int64
		var reservaBuscada1 Reserva
		var monto float64

		// Get values from row.
		err := rows.Scan(&ID, &ID_Evento, &ID_Cliente, &Nombre_Cliente, &monto)
		if err != nil {
			err = errors.New("Problema de lectura")
			return reservaBuscada, err
		}
		reservaBuscada1 = NewReserva(ID, ID_Evento, ID_Cliente, Nombre_Cliente, monto)
		reservaBuscada = reservaBuscada1

	}

	return reservaBuscada, nil

}

func (encuesta *Reserva) GetReservas(id_comercio int64) ([]Reserva, error) {
	ctx := context.Background()
	var reservas []Reserva
	var err error

	if datasources.Db == nil {
		err = errors.New("problema de conexion")
		return nil, err
	}

	err = datasources.Db.PingContext(ctx)
	if err != nil {
		err = errors.New("Error interno")
		return nil, err
	}

	tsql := fmt.Sprintf(queryGetReservas)

	// Execute query
	rows, err := datasources.Db.QueryContext(ctx, tsql, sql.Named("ID_Comercio", id_comercio))
	if err != nil {
		err = errors.New("Ingrese los parametros adecuados")
		return nil, err
	}

	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var Nombre_Cliente string
		var ID, ID_Evento, ID_Cliente int64
		var reservaBuscada1 Reserva
		var monto float64

		// Get values from row.
		err := rows.Scan(&ID, &ID_Evento, &ID_Cliente, &Nombre_Cliente, &monto)
		if err != nil {
			err = errors.New("Ingrese los parametros adecuados")
			return nil, err
		}
		reservaBuscada1 = NewReserva(ID, ID_Evento, ID_Cliente, Nombre_Cliente, monto)
		reservas = append(reservas, reservaBuscada1)
		fmt.Println("Entra")
	}

	return reservas, nil
}

func (reserva *Reserva) UpdateReserva(id_reserva int64) error {
	ctx := context.Background()

	var err error

	if datasources.Db == nil {
		err = errors.New("Base de datos Nula")
		return err
	}

	err = datasources.Db.PingContext(ctx)
	if err != nil {
		err = errors.New("Error interno")
		return err
	}

	tsql := fmt.Sprintf(queryUdpdateReservas)

	// Execute non-query with named parameters
	result, err := datasources.Db.ExecContext(
		ctx,
		tsql,
		sql.Named("ID", id_reserva),
		sql.Named("Nombre_Cliente", reserva.Nombre_Cliente),
		sql.Named("Monto", reserva.Monto),
	)
	if err != nil {
		err = errors.New("Problema con la query")
		return err
	}
	fmt.Println(result.RowsAffected())
	return nil
}
