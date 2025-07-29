package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type AdministradorRepository interface {
	Save(administrador *domain.Administrador) error
	FindByID(id string) (*domain.Administrador, error)
	Delete(id string) error
	ExisteDNI(dni string) (bool, error)
}
