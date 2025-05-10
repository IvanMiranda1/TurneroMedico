package domain

type Historia_clinica struct {
	ID          string
	PacienteID  string
	TurnoID     string
	Diagnostico string
	Tratamiento string
}

func NewHistoria_clinica(id, pacienteID, turnoID, diagnostico, tratamiento string) *Historia_clinica {
	return &Historia_clinica{
		ID:          id,
		PacienteID:  pacienteID,
		TurnoID:     turnoID,
		Diagnostico: diagnostico,
		Tratamiento: tratamiento,
	}
}
