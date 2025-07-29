package dtos

import (
	"fmt"
	"time"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type CrearAdminRequest struct {
	Nomyape  string `json:"nomyape" validate:"required"`
	DNI      string `json:"dni" validate:"required"`
	Fechanac string `json:"fechanac" validate:"required"`
}
type AdminResponse struct {
	ID       string `json:"id"`
	Nomyape  string `json:"nomyape"`
	DNI      string `json:"dni"`
	Fechanac string `json:"fechanac"`
}

func (r *CrearAdminRequest) AdminToDomain() (*domain.Administrador, error) {
	fecha, err := time.Parse("2006/01/02", r.Fechanac)
	if err != nil {
		return nil, fmt.Errorf("formato de nacimiento invalido (esperado YYYY/MM/DD): %w", err)
	}
	return domain.NewAdministrador(
		"0",
		r.Nomyape,
		r.DNI,
		fecha,
	), nil
}

func AdminFromDomain(a *domain.Administrador) *AdminResponse {
	return &AdminResponse{
		ID:       a.DNI,
		Nomyape:  a.Nomyape,
		DNI:      a.DNI,
		Fechanac: a.Fechanac.String(),
	}
}
