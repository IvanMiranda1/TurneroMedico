package services

import (
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
)

type HistoriaClinicaService struct {
	repo port.HistoriaClinicaRepository
}

func NewHistoriaClinicaService(r port.HistoriaClinicaRepository) *HistoriaClinicaService {
	return &HistoriaClinicaService{repo: r}
}

// Crear nueva historia clínica
func (s *HistoriaClinicaService) CreateHistoriaClinica(req dtos.HistoriaClinicaRequest) (*dtos.HistoriaClinicaResponse, error) {
	if req.PacienteID == "" || req.TurnoID == "" || req.Diagnostico == "" || req.Tratamiento == "" {
		return nil, errors.New("todos los campos obligatorios deben estar completos")
	}
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error generando ID para la historia clínica: %w", err)
	}

	alergias, err := s.repo.GetAlergias(req.Alergias)
	if err != nil {
		return nil, fmt.Errorf("error al obtener alergias: %w", err)
	}

	antecedentes, err := s.repo.GetAntecedentes(req.Antecedentes)
	if err != nil {
		return nil, fmt.Errorf("error al obtener antecedentes: %w", err)
	}

	newHistoriaClinica := &domain.HistoriaClinica{
		ID:           id.String(),
		PacienteID:   req.PacienteID,
		TurnoID:      req.TurnoID,
		Diagnostico:  req.Diagnostico,
		Tratamiento:  req.Tratamiento,
		Alergias:     alergias,
		Antecedentes: antecedentes,
	}
	newHistoria := dtos.HistoriaClinicaToDomain(&req, id.String())
	err = s.repo.Save(newHistoria)
	if err != nil {
		return nil, fmt.Errorf("error al guardar historia clínica: %w", err)
	}
	return dtos.HistoriaClinicaFromDomain(newHistoriaClinica), nil
}

func (s *HistoriaClinicaService) FindByPacienteID(id string) (*dtos.HistoriaClinicaResponse, error) {
	result, err := s.repo.FindByIDPaciente(id)
	if err != nil {
		return nil, fmt.Errorf("no se encontró historia clínica del paciente: %w", err)
	}
	return dtos.HistoriaClinicaFromDomain(result), nil
}

func (s *HistoriaClinicaService) ModificarHistoriaClinica(req dtos.HistoriaClinicaRequest, id string) (*dtos.HistoriaClinicaResponse, error) {
	existente, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("historia clínica no encontrada para modificar: %w", err)
	}
	modificada := dtos.HistoriaClinicaToDomain(&req, existente.ID)
	err = s.repo.Save(modificada)
	if err != nil {
		return nil, fmt.Errorf("error al modificar historia clínica: %w", err)
	}
	return dtos.HistoriaClinicaFromDomain(modificada), nil
}

func (s *HistoriaClinicaService) DeleteHistoriaClinica(id string) error {
	existente, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("la historia clínica no existe: %w", err)
	}
	err = s.repo.Delete(existente.ID)
	if err != nil {
		return fmt.Errorf("no fue posible eliminar la historia clínica: %w", err)
	}
	return nil
}
