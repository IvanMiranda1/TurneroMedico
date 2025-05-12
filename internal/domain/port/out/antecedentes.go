package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type AntecedentesRepository interface {
	Save(antecedentes *domain.Antecentes) error
	FindByID(id string) (*domain.Antecentes, error)
	Modificar(antecedenteModificado *domain.Antecentes) error
	Delete(id string) error
}
