package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type TurnoRepository interface {
	Save(turno *domain.Turno) error
	Modificar(turnoModificado *domain.Turno) error
	FindByID(id string) (*domain.Turno, error)
	Delete(id string) error
}
