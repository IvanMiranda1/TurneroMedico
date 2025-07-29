package port

import "github.com/IvanMiranda1/TurneroMedico/internal/domain"

type HistoriaClinicaRepository interface {
	Save(historiaClinica *domain.HistoriaClinica) error
	Modificar(historiaClinica *domain.HistoriaClinica) error
	Delete(id string) error
	FindByID(id string) (*domain.HistoriaClinica, error)
	FindByIDPaciente(id string) (*domain.HistoriaClinica, error)
	GetAntecedentes(ids []string) ([]domain.Antecedente, error)
	GetAlergias(ids []string) ([]domain.Alergia, error)
}
