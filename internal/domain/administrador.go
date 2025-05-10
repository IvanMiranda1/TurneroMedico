package domain

import "time"

type Administrador struct {
	ID       string
	Nomyape  string
	DNI      string
	Fechanac time.Time
}

func NewAdministrador(id, nomyape, dni string, fechanac time.Time) *Administrador {
	return &Administrador{
		ID:       id,
		Nomyape:  nomyape,
		DNI:      dni,
		Fechanac: fechanac,
	}
}
