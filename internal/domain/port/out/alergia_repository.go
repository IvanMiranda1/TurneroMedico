package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type AlergiaRepository interface {
	Save(alergia *domain.Alergia) error
	AlergiaFindByID(id string) (*domain.Alergia, error)
	DeleteAlergia(id string) error
	FindAlergiasByHistoriaID(id string) ([]*domain.Alergia, error)
}
