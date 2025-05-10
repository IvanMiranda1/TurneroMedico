package domain

type Horario struct {
	ID         string
	HoraInicio TimeOfDay
	HoraFin    TimeOfDay
}

func NewHorario(id string, horaInicio, horaFin TimeOfDay) *Horario {
	return &Horario{
		ID:         id,
		HoraInicio: horaInicio,
		HoraFin:    horaFin,
	}
}
