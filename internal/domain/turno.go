package domain

import (
	"time"
)

type Turno struct {
	ID         string
	MedicoID   *Medico
	PacienteID *Paciente
	Fechayhora time.Time
	Motivo     string
}

func NewTurno(id string, medicoID *Medico, pacienteID *Paciente, fechayhora time.Time, motivo string) *Turno {
	return &Turno{
		ID:         id,
		MedicoID:   medicoID,
		PacienteID: pacienteID,
		Fechayhora: fechayhora,
		Motivo:     motivo,
	}

}
