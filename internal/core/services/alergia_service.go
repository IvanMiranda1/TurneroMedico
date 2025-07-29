package services

import (
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
)

type AlergiaService struct {
	repo port.AlergiaRepository
}

func NewAlergiaService(r port.AlergiaRepository) *AlergiaService {
	return &AlergiaService{repo: r}
}

// --- --- Logica de Negocio --- ---
func (s *AlergiaService) CreateAlergia(req dtos.AlergiaRequest) (*dtos.AlergiaResponse, error) {
	if req.HistoriaID == "" || req.Descripcion == "" {
		return nil, errors.New("all required fields for alergia are mandatory")
	}
	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error al generar ID random: %w", err)
	}
	newAlergia := &domain.Alergia{
		ID:          id.String(),
		HistoriaID:  req.HistoriaID,
		Descripcion: req.Descripcion,
	}
	s.repo.Save(newAlergia)
	return dtos.AlergiaFromDomain(newAlergia), nil
}

func (s *AlergiaService) FindAlergiasByHistoriaID(id string) ([]*dtos.AlergiaResponse, error) {
	alergias, err := s.repo.FindAlergiasByHistoriaID(id)
	if err != nil {
		return nil, fmt.Errorf("no se encontraron alergias vinculadas al historial ID: %s: %w", id, err)
	}
	alergiasDtos := make([]*dtos.AlergiaResponse, 0, len(alergias))
	for _, alergia := range alergias {
		alergiasDtos = append(alergiasDtos, dtos.AlergiaFromDomain(alergia))
	}
	return alergiasDtos, nil
}

func (s *AlergiaService) Delete(id string) error {
	result, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("alergia no encontrada: %w", err)
	}
	err = s.repo.Delete(result.ID)
	if err != nil {
		return fmt.Errorf("no fue posible eliminar la alergia: %w", err)
	}
}
