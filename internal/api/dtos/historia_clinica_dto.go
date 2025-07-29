package dtos

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type HistoriaClinicaRequest struct {
	PacienteID   string   `json:"pacienteID" validate:"required"`
	TurnoID      string   `json:"turnoID" validate:"required"`
	Diagnostico  string   `json:"diagnostico" validate:"required"`
	Tratamiento  string   `json:"tratamiento" validate:"required"`
	Alergias     []string `json:"alergias" validate:"required"`     //[]Alergia
	Antecedentes []string `json:"antecedentes" validate:"required"` //[]Antecedente
}

func NewCrearHistoriaCinicaRequest(pacienteID, turnoID, diagnostico, tratamiento string, alergias, antecedentes []string) *HistoriaClinicaRequest {
	return &HistoriaClinicaRequest{
		PacienteID:   pacienteID,
		TurnoID:      turnoID,
		Diagnostico:  diagnostico,
		Tratamiento:  tratamiento,
		Alergias:     alergias,
		Antecedentes: antecedentes,
	}
}

type HistoriaClinicaResponse struct {
	ID           string   `json:"id"`
	PacienteID   string   `json:"pacienteID"`
	TurnoID      string   `json:"turnoID"`
	Diagnostico  string   `json:"diagnostico"`
	Tratamiento  string   `json:"tratamiento"`
	Alergias     []string `json:"alergias"`     //[]Alergia
	Antecedentes []string `json:"antecedentes"` //[]Antecedente
}


func HistoriaClinicaFromDomain(h *domain.HistoriaClinica) *HistoriaClinicaResponse {
	alergias := make([]string, 0, len(h.Alergias))
	for _, a := range h.Alergias {
		alergias = append(alergias, a.ID)
	}

	antecedentes := make([]string, 0, len(h.Antecedentes))
	for _, ant := range h.Antecedentes {
		antecedentes = append(antecedentes, ant.ID)
	}

	return &HistoriaClinicaResponse{
		ID:           h.ID,
		PacienteID:   h.PacienteID,
		TurnoID:      h.TurnoID,
		Diagnostico:  h.Diagnostico,
		Tratamiento:  h.Tratamiento,
		Alergias:     alergias,
		Antecedentes: antecedentes,
	}
}

func HistoriaClinicaToDomain(req *HistoriaClinicaRequest, id string) *domain.HistoriaClinica {
	alergias := make([]domain.Alergia, 0, len(req.Alergias))
	for _, aID := range req.Alergias {
		alergias = append(alergias, domain.Alergia{ID: aID})
	}

	antecedentes := make([]domain.Antecedente, 0, len(req.Antecedentes))
	for _, antID := range req.Antecedentes {
		antecedentes = append(antecedentes, domain.Antecedente{ID: antID})
	}

	return &domain.HistoriaClinica{
		ID:           id,
		PacienteID:   req.PacienteID,
		TurnoID:      req.TurnoID,
		Diagnostico:  req.Diagnostico,
		Tratamiento:  req.Tratamiento,
		Alergias:     alergias,
		Antecedentes: antecedentes,
	}