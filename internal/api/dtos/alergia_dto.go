package dtos

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type AlergiaRequest struct {
	HistoriaID  string `json:"historiaID" validate:"required"`
	Descripcion string `json:"descripcion" validate:"required"`
}

type AlergiaResponse struct {
	ID          string `json:"id"`
	HistoriaID  string `json:"historiaID"`
	Descripcion string `json:"descripcion"`
}

func (r *AlergiaRequest) ToDomain() (*domain.Alergia, error) {
	return domain.NewAlergia(
		"0",
		r.HistoriaID,
		r.Descripcion,
	), nil
}

func AlergiaFromDomain(a *domain.Alergia) *AlergiaResponse {
	return &AlergiaResponse{
		ID:          a.ID,
		HistoriaID:  a.HistoriaID,
		Descripcion: a.Descripcion,
	}
}
