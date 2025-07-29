package dtos

import (
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type AntecedentesRequest struct {
	HistoriaID  string `json:"historiaID" validate:"required"`
	Tipo        string `json:"tipo" validate:"required"` //TipoAntecedente
	Descripcion string `json:"descripcion" validate:"required"`
}

type AntecedentesResponse struct {
	HistoriaID  string `json:"historiaID"`
	Tipo        string `json:"tipo"` //TipoAntecedente
	Descripcion string `json:"descripcion"`
}

func (r *AntecedentesRequest) ToDomain() (*domain.Antecedente, error) {
	TipoAntecedenteDomain, err := domain.ParseTipoAntecedente(r.Tipo)
	if err != nil {
		return nil, fmt.Errorf("Valor de tipoAntecedente invalido: %w", err)
	}

	return domain.NewAntecedentes(
		"0",
		r.HistoriaID,
		TipoAntecedenteDomain,
		r.Descripcion,
	), nil
}

func AntecedenteFromDomain(a *domain.Antecedente) *AntecedentesResponse {
	return &AntecedentesResponse{
		HistoriaID:  a.HistoriaID,
		Tipo:        string(a.Tipo),
		Descripcion: a.Descripcion,
	}
}
