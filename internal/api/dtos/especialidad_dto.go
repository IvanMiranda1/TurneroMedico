package dtos

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type EspecialidadRequest struct {
	Nombre string `json:"nombre" validate:"required"`
}

func (r *EspecialidadRequest) EspecialidadToDomain() (*domain.Especialidad, error) {
	return domain.NewEspecialidad(
		"0",
		r.Nombre,
	), nil
}

type EspecialidadResponse struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

func EspecialidadFromDomain(e *domain.Especialidad) *EspecialidadResponse {
	if e == nil {
		return nil
	}
	return &EspecialidadResponse{
		ID:     e.ID,
		Nombre: e.Nombre,
	}
}
