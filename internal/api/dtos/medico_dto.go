package dtos

import (
	"fmt"
	"time"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

//medico + horario_medico + especialidad_medico

//realizar medicoDTO; medicoConHorariosDTO;

// Cree este dto y newMedicoDTO donde se pasan todos los datos formateados menos los ids especialidad, que debo acceder desde la BBDD desde la infraestructura
type MedicoDTO struct {
	ID           string
	Nomyape      string
	DNI          string
	Fechanac     time.Time
	Email        string
	Direccion    string
	Telefono     string
	Legajo       string
	Sexo         domain.Sexo
	Especialidad []string
}

func newMedicoDTO(id, nomyape, dni string, fechanac time.Time, email, direccion, telefono, legajo string, sexo domain.Sexo, especialidades []string) *MedicoDTO {
	return &MedicoDTO{
		ID:           id,
		Nomyape:      nomyape,
		DNI:          dni,
		Fechanac:     fechanac,
		Email:        email,
		Direccion:    direccion,
		Telefono:     telefono,
		Legajo:       legajo,
		Sexo:         sexo,
		Especialidad: especialidades,
	}
}

//______________________________________________________

type MedicoRequest struct {
	Nomyape           string   `json:"nomyape" validate:"required, min=3, max=255"`
	DNI               string   `json:"dni" validate:"required, len=8"`
	Fechanac          string   `json:"fechanac" validate:"required, datetime=2006/01/02"`
	Email             string   `json:"email" validate:"required,email"`
	Direccion         string   `json:"direccion" validate:"required"`
	Telefono          string   `json:"telefono" validate:"required"`
	Legajo            string   `json:"legajo" validate:"required"`
	Sexo              string   `json:"sexo" validate:"required"`
	EspecialidadesIDs []string `json:"especialidadesIds"`
}
type MedicoResponse struct {
	ID                string   `json:"id"`
	Nomyape           string   `json:"nomyape"`
	DNI               string   `json:"dni"`
	Fechanac          string   `json:"fechanac"`
	Email             string   `json:"email"`
	Direccion         string   `json:"direccion"`
	Telefono          string   `json:"telefono"`
	Legajo            string   `json:"legajo"`
	Sexo              string   `json:"sexo"`
	EspecialidadesIDs []string `json:"especialidadesIds"`
}

func (r *MedicoRequest) MedicoToDomain() (*MedicoDTO, error) {
	fechanaci, err := time.Parse("2006/01/02", r.Fechanac)
	if err != nil {
		return nil, fmt.Errorf("formato de nacimiento invalido (esperado YYYY/MM/DD): %w", err)
	}
	sexoDomain, err := domain.ParseSexo(r.Sexo)
	if err != nil {
		return nil, fmt.Errorf("valor de sexo invalido: %w", err)
	}

	return newMedicoDTO(
		"0",
		r.Nomyape,
		r.DNI,
		fechanaci,
		r.Email,
		r.Direccion,
		r.Telefono,
		r.Legajo,
		sexoDomain,
		r.EspecialidadesIDs,
	), nil
}

func MedicoFromDomain(m *domain.Medico) *MedicoResponse {
	especialidades := make([]string, 0, len(m.Especialidad))
	for _, especialidad := range m.Especialidad {
		especialidades = append(especialidades, especialidad.ID)
	}
	return &MedicoResponse{
		ID:                m.ID,
		Nomyape:           m.Nomyape,
		DNI:               m.DNI,
		Fechanac:          m.Fechanac.Format("2006/01/02"),
		Email:             m.Email,
		Direccion:         m.Direccion,
		Telefono:          m.Telefono,
		Legajo:            m.Legajo,
		Sexo:              m.Sexo.String(),
		EspecialidadesIDs: especialidades,
	}
}

// Relacion especialdad_medico
type RelEspecialidadMedico struct {
	MedicoId       string `json:"medicoid" validate:"required"`
	EspecialidadID string `json:"especialidadid" validate:"required"`
}

func newRelEspecialidadMedico(medicoid, especialidadid string) *RelEspecialidadMedico {
	return &RelEspecialidadMedico{
		MedicoId:       medicoid,
		EspecialidadID: especialidadid,
	}
}

// Horario_Medico
type CrearRelHorarioMedico struct {
	HorarioIDs []string `json:"horarioid" validate:"required"`
	MedicoID   string   `json:"medicoid" validate:"required"`
	DiaSemana  string   `json:"diasemana" validate:"required"`
}

func newHorarioMedicoDTO(horarioids []string, medicoid, diasemana string) *CrearRelHorarioMedico {
	return &CrearRelHorarioMedico{
		HorarioIDs: horarioids,
		MedicoID:   medicoid,
		DiaSemana:  diasemana,
	}
}

func (r *CrearRelHorarioMedico) ToDomain() (*CrearRelHorarioMedico, error) {
	return newHorarioMedicoDTO(
		r.HorarioIDs,
		r.MedicoID,
		r.DiaSemana,
	), nil
}

//Horarios se crearan desde el admin_dto
