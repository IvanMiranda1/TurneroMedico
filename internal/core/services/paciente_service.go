package services

import (
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
)

type PacienteService struct {
	repo port.PacienteRepository
}

func NewPacienteService(r port.PacienteRepository) *PacienteService {
	return &PacienteService{repo: r}
}

func (s *PacienteService) CreatePaciente(req dtos.PacienteRequest) (*dtos.PacienteResponse, error) {
	// Validación manual de campos obligatorios (opcional si ya usás validator)
	if req.Nomyape == "" || req.DNI == "" || req.Fechanac == "" || req.Email == "" || req.Direccion == "" || req.Telefono == "" || req.Sexo == "" {
		return nil, errors.New("todos los campos obligatorios deben estar completos")
	}

	// Verificar si ya existe un paciente con ese DNI
	existe, err := s.ValidateDNI(req.DNI)
	if err != nil {
		return nil, fmt.Errorf("error al validar existencia por DNI: %w", err)
	}
	if existe {
		return nil, fmt.Errorf("ya existe un paciente con el DNI %s", req.DNI)
	}
	// Verificar existencia por Email
	existeEmail, err := s.ValidateEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("error al validar existencia por email: %w", err)
	}
	if existeEmail {
		return nil, fmt.Errorf("ya existe un paciente con el email %s", req.Email)
	}

	// Convertimos el request a entidad de dominio
	paciente, err := req.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("error al convertir el request a entidad de dominio: %w", err)
	}

	// Generamos ID UUID
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error generando ID para el paciente: %w", err)
	}
	paciente.ID = id.String()

	// Guardamos
	err = s.repo.Save(paciente)
	if err != nil {
		return nil, fmt.Errorf("error al guardar el paciente: %w", err)
	}

	return dtos.PacienteFromDomain(paciente), nil
}

func (s *PacienteService) ValidateDNI(dni string) (bool, error) {
	validate, err := s.repo.ExistByDNI(dni)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de un dni igual en la bd: %w", err)
	}
	if validate {
		return false, domain.ErrAlreadyExists
	}
	if !RegexDNI.MatchString(dni) {
		return false, domain.ErrRegexMatch
	}
	return true, nil
}

func (s *PacienteService) ValidateEmail(email string) (bool, error) {
	validate, err := s.repo.ExistByEmail(email)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de un email igual en la bd: %w", err)
	}
	if validate {
		return false, domain.ErrAlreadyExists
	}
	if !RegexEmail.MatchString(email) {
		return false, domain.ErrRegexMatch
	}
	return true, nil
}
