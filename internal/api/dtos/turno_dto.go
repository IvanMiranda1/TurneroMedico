package dtos

import (
	"fmt"
	"time"
)

type TurnoRequest struct {
	MedicoID   string `json:"medicoid" validate:"required"`
	PacienteID string `json:"pacienteid" validate:"required"`
	FechayHora string `json:"fechayhora" validate:"required"`
	Motivo     string `json:"motivo" validate:"required"`
}

type TurnoDTO struct {
	ID         string
	MedicoID   string
	PacienteID string
	FechayHora time.Time
	Motivo     string
}

func NewTurnoRequest(id, medicoid, pacienteid string, fechayhora time.Time, motivo string) *TurnoDTO {
	return &TurnoDTO{
		MedicoID:   medicoid,
		PacienteID: pacienteid,
		FechayHora: fechayhora,
		Motivo:     motivo,
	}
}

func (r *TurnoRequest) ToDomain() (*TurnoDTO, error) {
	FechayHora, err := time.Parse("2006/01/02 15:04", r.FechayHora)
	if err != nil {
		return nil, fmt.Errorf("formato de hora invalido (esperado YYYY/MM/DD HH:MM): %w", err)
	}

	return NewTurnoRequest(
		"0",
		r.MedicoID,
		r.PacienteID,
		FechayHora,
		r.Motivo,
	), nil
}
