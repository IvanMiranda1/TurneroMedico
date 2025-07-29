package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type AntecedentesRepository interface {
	Save(antecedentes *domain.Antecedente) error
	FindByID(id string) (*domain.Antecedente, error)
	Modificar(antecedenteModificado *domain.Antecedente) error
	Delete(id string) error
	FindAntecedentesByHistoriaID(id string) ([]domain.Antecedente, error)
}
