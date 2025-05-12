package domain

type HistoriaClinica struct {
	ID          string
	PacienteID  string
	TurnoID     string
	Diagnostico string
	Tratamiento string
}

func NewHistoriaClinica(id, pacienteID, turnoID, diagnostico, tratamiento string) *HistoriaClinica {
	return &HistoriaClinica{
		ID:          id,
		PacienteID:  pacienteID,
		TurnoID:     turnoID,
		Diagnostico: diagnostico,
		Tratamiento: tratamiento,
	}
}
