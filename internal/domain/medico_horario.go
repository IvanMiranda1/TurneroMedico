package domain

type MedicoHorario struct {
	ID         string
	HorariosID []string
	MedicoID   string
	DiaSemana  string
}

func NewMedicoHorario(id string, horariosID []string, medicoID, diaSemana string) *MedicoHorario {
	return &MedicoHorario{
		ID:         id,
		HorariosID: horariosID,
		MedicoID:   medicoID,
		DiaSemana:  diaSemana,
	}
}
