package services

import (
	"fmt"
	"time"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

// dto para la creacion del paciente, no tiene ID
type CrearPacienteRequest struct {
	Nomyape   string `json:"nomyape" validate:"required, min=3, max=255"`
	DNI       string `json:"dni" validate:"required,len=8"`
	Fechanac  string `json:"fechanac" validate:"required,datetime=2006/01/02"`
	Email     string `json:"email" validate:"required,email"`
	Direccion string `json:"direccion" validate:"required"`
	Telefono  string `json:"telefono" validate:"required"`
	Sexo      string `json:"sexo" validate:"required,oneof=Masculino Femenino Otro"`
}

// toDomain parse - para que siga el struct del domain
func (r *CrearPacienteRequest) ToDomain() (*domain.Paciente, error) {
	fechanaci, err := time.Parse("2004/02/01", r.Fechanac)
	if err != nil {
		return nil, fmt.Errorf("formato de nacimiento invalido (esperado YYYY/MM/DD): %w", err)
	}
	sexoDomain, err := domain.ParseSexo(r.Sexo)
	if err != nil {
		return nil, fmt.Errorf("valor de sexo invalido: %w", err)
	}

	return domain.NewPaciente(
		"0",
		r.Nomyape,
		r.DNI,
		fechanaci,
		r.Email,
		r.Direccion,
		r.Telefono,
		sexoDomain,
	), nil
}

// PacienteResponse es un DTO de salida para el caso de uso(logica de negocio) (ej. devolver el paciente creado).
type PacienteResponse struct {
	ID        string `json:"id"`
	Nomyape   string `json:"nomyape"`
	DNI       string `json:"dni"`
	Fechanac  string `json:"fechanac"` // Puede ser un string para la salida también
	Email     string `json:"email"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Sexo      string `json:"sexo"`
}

// FromDomain convierte una entidad de dominio a un DTO de Response.
func FromDomain(p *domain.Paciente) *PacienteResponse {
	return &PacienteResponse{
		ID:        p.ID,
		Nomyape:   p.Nomyape,
		DNI:       p.DNI,
		Fechanac:  p.Fechanac.Format("2006/01/02"), // Formatear a string para la salida
		Email:     p.Email,
		Direccion: p.Direccion,
		Telefono:  p.Telefono,
		Sexo:      p.Sexo.String(),
	}
}

// Alergia
type AlergiaDTO struct {
	Descripcion string `json:"descripcion"`
}

type AntecedenteDTO struct {
	Tipo        string `json:"tipo"`
	Descripcion string `json:"descripcion"`
}

type HistoriaClinicaRepsonse struct {
	ID           string           `json:"id"`
	Diagnostico  string           `json:"diagnostico"`
	Tratamiento  string           `json:"tratamiento"`
	Alergias     []AlergiaDTO     `json:"alergias"`
	Antecedentes []AntecedenteDTO `json:"antecedentes"`
}
