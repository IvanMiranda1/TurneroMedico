package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type EspecialidadPostgresRepository struct {
	db *sql.DB
}

// Para crear instancia del repo en otro archivo
func NewEspecialidadPostgresRepository(db *sql.DB) *EspecialidadPostgresRepository {
	return &EspecialidadPostgresRepository{db: db}
}

func (r *EspecialidadPostgresRepository) Save(especialidad *domain.Especialidad) error {
	query := `
		INSERT INTO especialidad (id, nombre)
		VALUES ($1, $2)
		ON CONFLICT (id) DO UPDATE SET
			nombre = EXCLUDED.nombre;
	`
	_, err := r.db.Exec(query,
		especialidad.ID,
		especialidad.Nombre)
	if err != nil {
		return fmt.Errorf("error al guardar la especialidad con ID %s: %w", especialidad.ID, err)
	}
	return nil
}

func (r *EspecialidadPostgresRepository) Delete(id string) error {
	query := `DELETE FROM especialidad WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar la especialidad con el ID %s: %w", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas al eliminar especialidad con ID %s: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontro especialidad con ID %s para eliminar", id)
	}
	return nil
}

func (r *EspecialidadPostgresRepository) FindByID(id string) (*domain.Especialidad, error) {
	query := `
		SELECT id,nombre
		FROM especialidad
		WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var e domain.Especialidad
	if err := row.Scan(&e.ID, &e.Nombre); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no se encontro especialidad con ID %s", id)
		}
		return nil, fmt.Errorf("error al buscar especialidad con ID %s: %w", id, err)
	}
	return &e, nil
}

func (r *EspecialidadPostgresRepository) FindByName(nombre string) (*domain.Especialidad, error) {
	query := `
		SELECT id,nombre
		FROM especialidad
		WHERE nombre = $1`
	row := r.db.QueryRow(query, nombre)

	var e domain.Especialidad
	if err := row.Scan(&e.ID, &e.Nombre); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no se encontro especialidad con Nombre %s", nombre)
		}
		return nil, fmt.Errorf("error al buscar especialidad con Nombre %s: %w", nombre, err)
	}
	return &e, nil
}

/*
type EspecialidadRepository interface {
	Save(especialidad *domain.Especialidad) error
	Modificar(especialidadModificada *domain.Especialidad) error
	Delete(id string) error
	FindByID(id string) (*domain.Especialidad, error)
	FindByName(name string) (*domain.Especialidad, error)
}

*/
