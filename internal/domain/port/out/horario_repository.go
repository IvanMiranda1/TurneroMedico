package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type HorarioRepository interface {
	Save(horario *domain.Horario) error
	Modificar(horarioModificado *domain.Horario) error
	FindByID(id string) (*domain.Horario, error)
	Delete(id string) error
}
