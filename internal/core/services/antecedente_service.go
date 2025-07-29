package services

import (
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
)

type AntecedenteService struct {
	repo port.AntecedentesRepository
}

func NewAntecedenteService(r port.AntecedentesRepository) *AntecedenteService {
	return &AntecedenteService{repo: r}
}

func (s *AntecedenteService) CreateAntecedente(req dtos.AntecedentesRequest) (*dtos.AntecedentesResponse, error) {
	if req.Descripcion == "" || req.Tipo == "" || req.HistoriaID == "" {
		return nil, errors.New("all required fields for Antecedente are mandatory")
	}
	tipoant, err := domain.ParseTipoAntecedente(req.Tipo)
	if err != nil {
		return nil, fmt.Errorf("no pudo convertirse el tipo de antecedente: %w", err)
	}
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error al generar ID random: %w", err)
	}

	newAntecedente := &domain.Antecedente{
		ID:          id.String(),
		HistoriaID:  req.HistoriaID,
		Tipo:        tipoant,
		Descripcion: req.Descripcion,
	}
	s.repo.Save(newAntecedente)
	return dtos.AntecedenteFromDomain(newAntecedente), nil
}

func (s *AntecedenteService) DeleteAntecedente(id string) error {
	antecedente, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("el antecedente no existe: %w", err)
	}
	err = s.repo.Delete(antecedente.ID)
	if err != nil {
		return fmt.Errorf("no fue posible eliminar el antecedente: %w", err)
	}
	return nil
}

func (s *AntecedenteService) AntecedentesFindByHistoria(id string) ([]*dtos.AntecedentesResponse, error) {
	antecedentes, err := s.repo.FindAntecedentesByHistoriaID(id)
	if err != nil {
		return nil, fmt.Errorf("no se encontraron antecedentes vinculados al historial ID: %s: %w", id, err)
	}
	antecedentesDtos := make([]*dtos.AntecedentesResponse, 0, len(antecedentes))
	for _, antecedente := range antecedentes {
		antecedentesDtos = append(antecedentesDtos, dtos.AntecedenteFromDomain(&antecedente))
	}
	return antecedentesDtos, nil
}

/*
Save(antecedentes *domain.Antecedente) error
	FindByID(id string) (*domain.Antecedente, error)
	Modificar(antecedenteModificado *domain.Antecedente) error
	Delete(id string) error
*/
