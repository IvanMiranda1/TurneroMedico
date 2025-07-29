package services

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	port "github.com/IvanMiranda1/TurneroMedico/internal/domain/port/out"
	"github.com/gofrs/uuid"
)

type MedicoService struct {
	repo port.MedicoRepository
}

func NewMedicoService(r port.MedicoRepository) *MedicoService {
	return &MedicoService{repo: r}
}

// regex = expresion regular
var RegexTel = regexp.MustCompile(`^\d{10}$`)
var RegexDNI = regexp.MustCompile(`^\d{8}$`)
var RegexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// --- --- Logica de negocio --- ---
func (s *MedicoService) CreateMedico(req dtos.MedicoRequest) (*dtos.MedicoResponse, error) {
	if err := validarCamposObligatorios(req); err != nil {
		return nil, err
	}

	if ok, err := s.validarCampoUnico(req.DNI, RegexDNI, s.repo.ExisteDNI, "DNI"); !ok {
		return nil, err
	}
	if ok, err := s.validarCampoUnico(req.Email, RegexEmail, s.repo.ExisteEmail, "email"); !ok {
		return nil, err
	}
	if ok, err := s.validarCampoUnico(req.Telefono, RegexTel, s.repo.ExisteTelefono, "teléfono"); !ok {
		return nil, err
	}

	fechaNacimiento, err := time.Parse("2006/01/02", req.Fechanac)
	if err != nil {
		return nil, fmt.Errorf("fecha inválida (formato esperado YYYY/MM/DD): %w", err)
	}
	if !esMayorDeEdad(fechaNacimiento) {
		return nil, errors.New("el médico debe ser mayor de edad")
	}

	sexo, err := domain.ParseSexo(req.Sexo)
	if err != nil {
		return nil, fmt.Errorf("sexo inválido: %w", err)
	}

	especialidades, err := s.repo.GetEspecialidades(req.EspecialidadesIDs)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo especialidades: %w", err)
	}

	legajo, err := generarLegajoUnico(s.repo)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("error generando UUID: %w", err)
	}

	medico := &domain.Medico{
		ID:           id.String(),
		Nomyape:      req.Nomyape,
		DNI:          req.DNI,
		Fechanac:     fechaNacimiento,
		Email:        req.Email,
		Telefono:     req.Telefono,
		Legajo:       legajo,
		Sexo:         sexo,
		Especialidad: especialidades,
	}

	if err := s.repo.Save(medico); err != nil {
		return nil, fmt.Errorf("error al guardar el médico: %w", err)
	}

	return dtos.MedicoFromDomain(medico), nil
}

func (s *MedicoService) FindByID(req string) (*dtos.MedicoResponse, error) {
	medico, err := s.repo.FindByID(req)
	if err != nil {
		return nil, fmt.Errorf("no se encontro Medico por ID: %w", err)
	}
	medicoDto := dtos.MedicoFromDomain(medico)
	return medicoDto, nil
}

func (s *MedicoService) FindByName(name string) (*dtos.MedicoResponse, error) {
	medico, err := s.repo.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("no se encontro Medico por el nombre: %w", err)
	}
	medicoDto := dtos.MedicoFromDomain(medico)
	return medicoDto, nil
}

func (s *MedicoService) DeleteMedico(id string) error {
	result, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("el medico no existe: %w", err)
	}
	err = s.repo.Delete(result.ID)
	if err != nil {
		return fmt.Errorf("no fue posible eliminar el medico: %w", err)
	}
	return nil
}

func (s *MedicoService) FindByLegajoID(legajo string) (*dtos.MedicoResponse, error) {
	medico, err := s.repo.FindByLegajoID(legajo)
	if err != nil {
		return nil, fmt.Errorf("no se encontron medico: %w", err)
	}
	medicoDTO := dtos.MedicoFromDomain(&medico)
	return medicoDTO, nil
}

// Existe
func (s *MedicoService) ExisteLegajoID(legajo string) (bool, error) {
	exist, err := s.repo.ExisteLegajoID(legajo)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de legajo: %w", err)
	}
	if exist {
		return true, domain.ErrAlreadyExists
	}
	return false, nil
}

// Getters
// para crear el medico solo se pasan los ids, por eso necesito los ids, y retornar domain.especialidad

func (s *MedicoService) GetEspecialidades(ds []string) ([]domain.Especialidad, error) {
	// return []especialidad
	especialidades, err := s.repo.GetEspecialidades(ds) //esto debe retornar un array de domain
	if err != nil {
		return nil, fmt.Errorf("error obteniendo especialidades: %w", err)
	}
	//transformar las especialidades a dtoResponse
	return especialidades, nil
}

// Validates
func (s *MedicoService) ValidateTelefono(telefono string) (bool, error) {
	validate, err := s.repo.ExisteTelefono(telefono)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de un telefono igual en la bd: %w", err)
	}
	if validate { // ya existe un telefono asi
		return false, domain.ErrAlreadyExists
	}
	if !RegexTel.MatchString(telefono) {
		return false, domain.ErrRegexMatch
	}
	return true, nil
}

func (s *MedicoService) ValidateEmail(email string) (bool, error) {
	validate, err := s.repo.ExisteEmail(email)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de un telefono igual en la bd: %w", err)
	}
	if validate {
		return false, domain.ErrAlreadyExists
	}
	if !RegexEmail.MatchString(email) {
		return false, domain.ErrRegexMatch
	}
	return true, nil
}

func (s *MedicoService) ValidateDNI(dni string) (bool, error) {
	validate, err := s.repo.ExisteDNI(dni)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia de un telefono igual en la bd: %w", err)
	}
	if validate {
		return false, domain.ErrAlreadyExists
	}
	if !RegexDNI.MatchString(dni) {
		return false, domain.ErrRegexMatch
	}
	return true, nil
}

// --- --- FIN Logica de negocio --- ---

// Funcs
func MayorEdad(fechanac time.Time) bool {
	ahora := time.Now()
	mayoriaEdad := ahora.AddDate(-18, 0, 0)
	return !fechanac.After(mayoriaEdad)
}

func generarLegajoUnico(repo port.MedicoRepository) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const intentosMaximos = 100

	for i := 0; i < intentosMaximos; i++ {
		legajo := fmt.Sprintf("%04d", r.Intn(9000)+1000)

		_, err := repo.ExisteLegajoID(legajo)
		if errors.Is(err, domain.ErrNotFound) {
			return legajo, nil
		}
		if err != nil {
			return "", fmt.Errorf("error verificando legajo único: %w", err)
		}
	}

	return "", errors.New("no se pudo generar un legajo único luego de varios intentos")
}

func validarCamposObligatorios(req dtos.MedicoRequest) error {
	if req.Nomyape == "" || req.DNI == "" || req.Fechanac == "" || req.Email == "" || req.Telefono == "" || len(req.EspecialidadesIDs) == 0 {
		return errors.New("todos los campos obligatorios deben estar completos")
	}
	return nil
}

func esMayorDeEdad(fecha time.Time) bool {
	return fecha.Before(time.Now().AddDate(-18, 0, 0))
}

func (s *MedicoService) validarCampoUnico(
	valor string,
	regex *regexp.Regexp,
	checkExist func(string) (bool, error),
	campo string,
) (bool, error) {
	if !regex.MatchString(valor) {
		return false, fmt.Errorf("formato inválido de %s", campo)
	}
	existe, err := checkExist(valor)
	if err != nil {
		return false, fmt.Errorf("error verificando %s en la base de datos: %w", campo, err)
	}
	if existe {
		return false, fmt.Errorf("ya existe un médico con ese %s", campo)
	}
	return true, nil
}
