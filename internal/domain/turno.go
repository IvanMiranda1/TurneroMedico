package domain

import "time"

type Turno struct {
	ID         int
	MedicoID   int
	PacienteID int
	Fechayhora time.Time
	Motivo     string
}

func NewTurno(id, medicoID, pacienteID int, fechayhora time.Time, motivo string) *Turno {
	return &Turno{
		ID:         id,
		MedicoID:   medicoID,
		PacienteID: pacienteID,
		Fechayhora: fechayhora,
		Motivo:     motivo,
	}

}
