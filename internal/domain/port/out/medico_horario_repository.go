package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type MedicoHorarioRepository interface {
	Save(horarioMedico *domain.MedicoHorario) error
	Modificar(horarioMedicoModificado *domain.MedicoHorario) error
	FindByID(id string) (*domain.MedicoHorario, error)
	Delete(id string) error
	FindByMedicoYDia(medicoID string, dia string) ([]domain.MedicoHorario, error)
}
