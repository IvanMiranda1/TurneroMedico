package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type MedicoRepository interface {
	Save(medico *domain.Medico) error           //save(parametro) error // como no retorna nada queda ahi
	FindByID(id string) (*domain.Medico, error) //en cambio findbyid(parametro) (return medico o un error)
	FindByName(name string) (*domain.Medico, error)
	Delete(id string) error
	ExistsByEmail(email string) (bool, error)
}
