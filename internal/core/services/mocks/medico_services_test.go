package service_test

import (
	"errors"
	"testing"

	"github.com/IvanMiranda1/TurneroMedico/internal/api/dtos"
	"github.com/IvanMiranda1/TurneroMedico/internal/core/services"
	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

// MockMedicoRepository es una implementación simulada de ports.MedicoRepository para tests.
type MockMedicoRepository struct {
	SaveFunc           func(medico *domain.Medico) error
	FindByIDFunc       func(id string) (*domain.Medico, error)
	FindByNameFunc     func(name string) (*domain.Medico, error)
	DeleteFunc         func(id string) error
	FindByLegajoIDFunc func(legajo string) (domain.Medico, error)
	//Exists
	ExisteLegajoIDFunc func(legajo string) (bool, error)
	ExisteByEmailFunc  func(email string) (bool, error)
	//Getters
	GetEspecialidadesFunc func(ds []string) ([]domain.Especialidad, error)
	//Validate
	ValidateTelefonoFunc func(telefono string) bool
	ValidateEmailFunc    func(email string) bool
	ValidateDNIFunc      func(i string) bool
}

// Save es el metodo mock para la interfaz save
func (m *MockMedicoRepository) Save(medico *domain.Medico) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(medico)
	}
	return nil
}

// ---demas funciones ---
func (m *MockMedicoRepository) FindByID(id string) (*domain.Medico, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	// Comportamiento por defecto: no encontrado
	return nil, errors.New("médico no encontrado por ID (comportamiento mock por defecto)")
}

// FindByName simula la búsqueda de un médico por nombre.
// Por defecto, retorna nil y un error.
func (m *MockMedicoRepository) FindByName(name string) (*domain.Medico, error) {
	if m.FindByNameFunc != nil {
		return m.FindByNameFunc(name)
	}
	// Comportamiento por defecto: no encontrado
	return nil, errors.New("médico no encontrado por nombre (comportamiento mock por defecto)")
}

func (m *MockMedicoRepository) FindByLegajoID(legajo string) (domain.Medico, error) {
	if m.FindByLegajoIDFunc != nil {
		return m.FindByLegajoIDFunc(legajo)
	}
	// Comportamiento por defecto del mock si FindByLegajoIDFunc no está asignado:
	// Retornar el valor cero de domain.Medico y un error indicando que no se encontró.
	// Asegúrate de que domain.ErrNotFound esté definido (var ErrNotFound = errors.New("recurso no encontrado"))
	return domain.Medico{}, domain.ErrNotFound
}

// Delete simula la eliminación de un médico.
// Por defecto, retorna nil (éxito).
func (m *MockMedicoRepository) Delete(id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil // Comportamiento por defecto: eliminación exitosa
}

// ExistsByEmail simula la verificación de existencia de un email.
// Por defecto, retorna 'false' (el email no existe) y nil error.
func (m *MockMedicoRepository) ExisteByEmail(email string) (bool, error) {
	if m.ExisteByEmailFunc != nil {
		return m.ExisteByEmailFunc(email)
	}
	return false, nil // Comportamiento por defecto: el email no existe
}

func (m *MockMedicoRepository) ExisteLegajoID(legajo string) (bool, error) {
	if m.ExisteLegajoIDFunc != nil {
		return m.ExisteLegajoIDFunc(legajo)
	}
	// Comportamiento por defecto: el legajo no existe, y sin error.
	return false, nil
}

// GetEspecialidades simula la obtención de especialidades.
// Por defecto, retorna un slice vacío de especialidades y nil error.
func (m *MockMedicoRepository) GetEspecialidades(ids []string) ([]domain.Especialidad, error) {
	if m.GetEspecialidadesFunc != nil {
		return m.GetEspecialidadesFunc(ids)
	}
	return []domain.Especialidad{}, nil // Comportamiento por defecto: ninguna especialidad encontrada
}

// ValidateTelefono simula la validación de un número de teléfono.
// Por defecto, retorna 'true' (teléfono válido).
func (m *MockMedicoRepository) ValidateTelefono(telefono string) bool {
	if m.ValidateTelefonoFunc != nil {
		return m.ValidateTelefonoFunc(telefono)
	}
	return true // Comportamiento por defecto: válido
}

// ValidateEmail simula la validación de un email.
// Por defecto, retorna 'true' (email válido).
func (m *MockMedicoRepository) ValidateEmail(email string) bool {
	if m.ValidateEmailFunc != nil {
		return m.ValidateEmailFunc(email)
	}
	return true // Comportamiento por defecto: válido
}

// ValidateDNI simula la validación de un DNI.
// Por defecto, retorna 'true' (DNI válido).
func (m *MockMedicoRepository) ValidateDNI(dni string) bool {
	if m.ValidateDNIFunc != nil {
		return m.ValidateDNIFunc(dni)
	}
	return true // Comportamiento por defecto: válido
}

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

// --- Test para MedicoService ---
func TestCreateMedico(t *testing.T) {
	t.Log("Test iniciado.6")
	//Escenario 1: Creacion exitosa.
	mockRepo := &MockMedicoRepository{
		SaveFunc: func(medico *domain.Medico) error {
			if medico.ID == "" {
				t.Error("ID del medico no deberia estar vacio al guardar")
			}
			return nil // guardar fue exitoso
		},
	}
	//MedicoService toma un ports.MedicoRepository en su constructor

	svc := services.NewMedicoService(mockRepo) //en vez de pasarle el repository se pasa el mock

	req := &dtos.MedicoRequest{
		Nomyape:           "Juan Perez",
		DNI:               "12345678",
		Fechanac:          "1990-01-15",
		Email:             "juan@example.com",
		Direccion:         "Luisa Rosso 972",
		Telefono:          "1122334455",
		Sexo:              "Masculino",
		EspecialidadesIDs: []string{}, // evita ambigüedades, con un [], Go lo toma como error
	}

	// pasarle un error a %v llama implicitamente al Error() de la interfaz error
	// es decir, devuelve la representacion de cadena del error
	/*
			originalErr := errors.New("algo salió mal")
		    fmt.Printf("El error es: %v\n", originalErr)
		    // Salida: El error es: algo salió mal
	*/
	resp, err := svc.CreateMedico(*req)
	if err != nil {
		t.Fatalf("error inesperado al crear medico: %v", err)
	}
	if resp == nil {
		t.Fatalf("respuesta de medico es nula: %v", domain.ErrResponseNil)
	}
	if resp.ID == "" {
		t.Errorf("id del medico no fue generado: %v", domain.ErrFailedToCreate)
	}
	// ... más aserciones sobre 'resp'

	//Escenario 2: Error al guardar  en el repositorio
	mockRepoWithError := &MockMedicoRepository{
		SaveFunc: func(medico *domain.Medico) error {
			return errors.New("error simulado de base de datos")
		},
	}
	svcWithError := services.NewMedicoService(mockRepoWithError)

	_, err = svcWithError.CreateMedico(*req)
	if err == nil {
		t.Fatal("Se esperaba un error al guardar pero no se recibio ninguno")
	}
	if err.Error() != "error al guardar el medico: error simulado de base de datos" {
		t.Errorf("Mensaje de error inesperado: %v", domain.ErrFailedToCreate)
	}
}
