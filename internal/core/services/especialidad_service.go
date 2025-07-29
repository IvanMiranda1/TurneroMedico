package services

import (
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
)

type EspecialidadService struct {
	repo port.EspecialidadRepository
}

func NewEspecialidadService(r port.EspecialidadRepository) *EspecialidadService {
	return &EspecialidadService{repo: r}
}

func (s *EspecialidadService) CreateEspecialidad(req dtos.EspecialidadRequest) (*dtos.EspecialidadResponse, error) {
	if req.Nombre == "" {
		return nil, errors.New("all required fields for Especialidad are mandatory")
	}

	_, err := s.repo.FindByName(req.Nombre)
	if err == nil {
		return nil, fmt.Errorf("el nombre de especialidad ya está registrado")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error al generar el Id random: %w", err)
	}
	newEspecialidad := &domain.Especialidad{
		ID:     id.String(),
		Nombre: req.Nombre,
	}
	s.repo.Save(newEspecialidad)
	return dtos.EspecialidadFromDomain(newEspecialidad), nil
}

func (s *EspecialidadService) FindByName(name string) (*dtos.EspecialidadResponse, error) {
	especialidad, err := s.repo.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("no se encontro especialidad por el nombre: %w", err)
	}
	especialidadDto := dtos.EspecialidadFromDomain(especialidad)
	return especialidadDto, nil
}

func (s *EspecialidadService) DeleteEspecialidad(id string) error {
	result, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("el especialidad no existe: %w", err)
	}
	err = s.repo.Delete(result.ID)
	if err != nil {
		return fmt.Errorf("no fue posible eliminar el especialidad: %w", err)
	}
	return nil
}
