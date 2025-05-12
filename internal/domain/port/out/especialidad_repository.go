package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type EspecialidadRepository interface {
	Save(especialidad *domain.Especialidad) error
	Modificar(especialidadModificada *domain.Especialidad) error
	Delete(id string) error
	FindByID(id string) (*domain.Especialidad, error)
	FindByName(name string) (*domain.Especialidad, error)
}
