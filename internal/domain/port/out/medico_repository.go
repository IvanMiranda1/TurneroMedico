package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type MedicoRepository interface {
	Save(medico *domain.Medico) error           //save(parametro) error // como no retorna nada queda ahi
	FindByID(id string) (*domain.Medico, error) //en cambio findbyid(parametro) (return medico o un error)
	FindByName(name string) (*domain.Medico, error)
	Delete(id string) error
	FindByLegajoID(legajo string) (domain.Medico, error)
	//existe
	ExisteLegajoID(legajo string) (bool, error)
	ExisteTelefono(telefono string) (bool, error)
	ExisteEmail(email string) (bool, error)
	ExisteDNI(i string) (bool, error)
	//Getters
	GetEspecialidades(ds []string) ([]domain.Especialidad, error) // return []especialidad

}
