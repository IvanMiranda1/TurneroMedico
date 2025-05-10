package domain

type HorarioMedico struct {
	ID        string
	HorarioID string
	MedicoID  string
	DiaSemana string
}

func newHorarioMedico(id, horarioID, medicoID, diasemana string) *HorarioMedico {
	return &HorarioMedico{
		ID:        id,
		HorarioID: horarioID,
		MedicoID:  medicoID,
		DiaSemana: diasemana,
	}
}
