package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type HorarioPostgresRepository struct {
	db *sql.DB
}

// Para crear instancia del repo en otro archivo
func NewHorarioPostgresRepository(db *sql.DB) *HorarioPostgresRepository {
	return &HorarioPostgresRepository{db: db}
}

func (r *HorarioPostgresRepository) Save(horario *domain.Horario) error {
	query := `
		INSERT INTO horario (id, horainicio, horafin)
		VALUES($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET
		horainicio = EXCLUDE.horainicio,
		horafin = EXCLUDE.horafin;
	`
	_, err := r.db.Exec(query,
		horario.ID,
		horario.HoraInicio,
		horario.HoraFin,
	)
	if err != nil {
		return fmt.Errorf("error al guardar horario: %w", err)
	}
	return nil
}

func (r *HorarioPostgresRepository) FindByID(id string) (*domain.Horario, error) {
	query := `
		SELECT id, horainicio, horafin
		FROM horario
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)

	var a domain.Horario
	if err := row.Scan(&a.ID, &a.HoraInicio, &a.HoraFin); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no se encontró horario con ID %s", id)
		}
		return nil, fmt.Errorf("error al buscar horario con ID %s: %w", id, err)
	}
	return &a, nil
}

func (r *HorarioPostgresRepository) Delete(id string) error {
	query := `DELETE FROM horario WHERE id = $1;`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar Horario con ID %s: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas al eliminar Horario con ID %s: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró Horario con ID %s para eliminar", id)
	}

	return nil
}
