package domain

type Antecentes struct {
	ID          string
	HistoriaID  string
	Tipo        TipoAntecendente
	Descripcion string
}

func NewAntecedentes(id, historiaID string, tipo TipoAntecendente, descripcion string) *Antecentes {
	return &Antecentes{
		ID:          id,
		HistoriaID:  historiaID,
		Tipo:        tipo,
		Descripcion: descripcion,
	}
}
