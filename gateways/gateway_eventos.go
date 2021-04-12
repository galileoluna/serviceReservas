package gateways

import (
	"encoding/json"
	"net/http"
)

const (
	apiEventos = "http://localhost:8081/evento/all"
)

type Evento struct {
	ID_Evento          int64  `json:"id_evento"`
	Nombre             string `json:"evento"`
	Descripcion        string `json:"descripcion"`
	HoraDeInicio       string `json:"hora_inicio"`
	HoraDeFinalizacion string `json:"hora_finalizacion"`
}

func DecodeEventos() (*[]Evento, error) {
	r, err := http.Get(apiEventos)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body := r.Body
	dec := json.NewDecoder(body)
	eventos := []Evento{}
	final := dec.Decode(&eventos)
	return &eventos, final
}
