package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type HorarioMedicoRepository interface {
	Save(horarioMedico *domain.HorarioMedico) error
	Modificar(horarioMedicoModificado *domain.HorarioMedico) error
	FindByID(id string) (*domain.HorarioMedico, error)
	Delete(id string) error
}
