package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type PacienteRepository interface {
	Save(paciente *domain.Paciente) error
	FindByID(id string) (*domain.Paciente, error)
	FindByDNI(dni string) (*domain.Paciente, error)
	Modificar(pacienteModificado *domain.Paciente) error
	Delete(id string) error
	ExistByEmail(email string) (bool, error)
	ExistByDNI(dni string) (bool, error)
}
