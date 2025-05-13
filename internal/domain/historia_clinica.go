package domain

type HistoriaClinica struct {
	ID           string
	PacienteID   string
	TurnoID      string
	Diagnostico  string
	Tratamiento  string
	Alergias     []Alergia
	Antecedentes []Antecedente
}

func NewHistoriaClinica(id, pacienteID, turnoID, diagnostico, tratamiento string, alergias []Alergia, antecedentes []Antecedente) *HistoriaClinica {
	return &HistoriaClinica{
		ID:           id,
		PacienteID:   pacienteID,
		TurnoID:      turnoID,
		Diagnostico:  diagnostico,
		Tratamiento:  tratamiento,
		Alergias:     alergias,
		Antecedentes: antecedentes,
	}
}
