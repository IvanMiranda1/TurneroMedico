package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
	"github.com/gohugoio/hugo/config/services"
)

//importar dtos, entidades de dominio, interfaces de repositorio

// entidades de dominio
type AdminService struct {
	repo port.AdministradorRepository
}

// Inicializar AdminServices que recibe una implementacion de la interfaz AdministradorRepository
// Ejemplo de Inversion de Control (Depedency injection)
func NewAdminServices(r port.AdministradorRepository) *AdminService {
	return &AdminService{repo: r}
}

// Logica de negocio para crear una nueva entidad
// Recibe DTO de solicitud de la entidad y retorna DTO de respuesta de "entidad creada" o error
func (s *AdminService) CreateAdmin(req dtos.CrearAdminRequest) (*dtos.AdminResponse, error) {
	if req.Nomyape == "" || req.DNI == "" || req.Fechanac == "" {
		return nil, errors.New("todos los campos del administrador son obligatorios")
	}
	//Parse de la entidad DTO.fechanac a time.Time de la entidad dominio
	parsedFechanac, err := time.Parse("2006/01/02", req.Fechanac)
	if err != nil {
		return nil, fmt.Errorf("formato de fecha y hora inválido: %w", err)
	}
	if !services.MayorEdad(parsedFechanac) {
		return nil, fmt.Errorf("la edad ingresada es menor de 18 años")
	}

	//no puede repetirse el DNI, debe ser mayor de 18

	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error al generar el id")
	}
	nuevoAdmin := &domain.Administrador{
		ID:       id.String(),
		Nomyape:  req.Nomyape,
		DNI:      req.DNI,
		Fechanac: parsedFechanac,
	}
	err = s.repo.Save(nuevoAdmin)
	if err != nil {
		return nil, fmt.Errorf("no se pudo guardar el administrador: %w", err)
	}

	return dtos.AdminFromDomain(nuevoAdmin), nil
}

func (s *AdminService) FindByID(req string) (*dtos.AdminResponse, error) {
	admin, err := s.repo.FindByID(req)
	if err != nil {
		return nil, fmt.Errorf("no se encontró Administrador por ID: %w", err)
	}
	adminDto := dtos.AdminFromDomain(admin)
	return adminDto, nil
}

func (s *MedicoService) ValidateDNIAdmin(dni string) (bool, error) {
	validate, err := s.repo.ExisteDNI(dni)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de un telefono igual en la bd: %w", err)
	}
	if validate {
		return false, domain.ErrAlreadyExists
	}
	if !services.RegexDNI.MatchString(dni) {
		return false, domain.ErrRegexMatch
	}
	return true, nil
}

func (s *MedicoService) Delete(id string) error {
	result, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("admin no encontrado: %w", err)
	}
	err = s.repo.Delete(result.ID)
	if err != nil {
		return fmt.Errorf("no fue posible eliminar el admin: %w", err)
	}
	return nil
}
