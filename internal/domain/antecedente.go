package domain

type Antecedente struct {
	ID          string
	HistoriaID  string
	Tipo        TipoAntecedente
	Descripcion string
}

func NewAntecedentes(id, historiaID string, tipo TipoAntecedente, descripcion string) *Antecedente {
	return &Antecedente{
		ID:          id,
		HistoriaID:  historiaID,
		Tipo:        tipo,
		Descripcion: descripcion,
	}
}
