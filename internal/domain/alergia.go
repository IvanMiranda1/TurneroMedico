package domain

type Alergia struct {
	ID          string
	HistoriaID  string
	Descripcion string
}

func NewAlergia(id, historiaID, descripcion string) *Alergia {
	return &Alergia{
		ID:          id,
		HistoriaID:  historiaID,
		Descripcion: descripcion,
	}
}
