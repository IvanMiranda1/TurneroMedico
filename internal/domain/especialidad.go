package domain

type Especialidad struct {
	ID     string
	Nombre string
}

func NewEspecialidad(id, nombre string) *Especialidad {
	return &Especialidad{
		ID:     id,
		Nombre: nombre,
	}
}
