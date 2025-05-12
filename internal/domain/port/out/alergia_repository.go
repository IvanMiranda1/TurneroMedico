package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type AlergiaRepository interface {
	Save(alergia *domain.Alergia) error
	FindByName(name string) (*domain.Alergia, error)
	Delete(id string) error
}
