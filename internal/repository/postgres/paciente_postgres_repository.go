package postgres

import (
	"database/sql"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type PacientePostgresRepository struct {
	db *sql.DB
}

// Para crear instancia del repo en otro archivo
func NewPacientePostgresRepository(db *sql.DB) *PacientePostgresRepository {
	return &PacientePostgresRepository{db: db}
}

func (r *PacientePostgresRepository) Save(paciente *domain.Paciente) error {
	query := `INSERT INTO paciente (id, nomyape, dni, fechanac, email, direccion, telefono, sexo)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	ON CONFLICT (id) DO UPDATE SET
	nomyape = EXCLUDE.nomyape,
	dni = EXCLUDE.dni,
	fechanac = EXCLUDE.fechanac,
	email = EXCLUDE.email,
	direccion = EXCLUDE.direccion,
	telefono = EXCLUDE.telefono,
	sexo = EXCLUDE.sexo;
	`
	_, err := r.db.Exec(query,
		paciente.ID,
		paciente.Nomyape,
		paciente.DNI,
		paciente.Fechanac,
		paciente.Email,
		paciente.Direccion,
		paciente.Telefono,
		paciente.Sexo,
	)
	if err != nil {
		return fmt.Errorf("error al guardar el paciente: %w", err)
	}

}

func (r *PacientePostgresRepository) ExistsByDNI(dni string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM pacientes WHERE dni = $1)", dni).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Verifica si existe un paciente con ese Email
func (r *PacientePostgresRepository) ExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM pacientes WHERE email = $1)", email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

/*

chatgpt me dijo que asi no, porque es menos eficiente:
		Tu implementación funciona perfectamente y no está mal, especialmente si el volumen de datos no es enorme.
		Pero si estás buscando eficiencia y claridad, es mejor usar EXISTS porque:

			Es semánticamente más expresivo: "¿Existe?", no "¿Cuántos hay?"

			Es más eficiente: el motor de la base de datos se detiene cuando encuentra la primera coincidencia.

func (r *PacientePostgresRepository) ExisteEmail(email string) (bool, error) {
	query := `SELECT COUNT(*) FROM paciente WHERE email = $1;`

	var count int
	err := r.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia del email en la base de datos: %w", err)
	}

	return count > 0, nil
}

func (r *PacientePostgresRepository) ExisteDNI(dni string) (bool, error) {
	query := `SELECT COUNT(*) FROM paciente WHERE dni = $1;`

	var count int
	err := r.db.QueryRow(query, dni).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia del dni en la base de datos: %w", err)
	}

	return count > 0, nil
}
*/
