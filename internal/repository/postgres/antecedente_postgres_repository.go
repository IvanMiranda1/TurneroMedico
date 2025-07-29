package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type AntecedentePostgresRepository struct {
	db *sql.DB
}

// Para crear instancia del repo en otro archivo
func NewAntecedentePostgresRepository(db *sql.DB) *AntecedentePostgresRepository {
	return &AntecedentePostgresRepository{db: db}
}

func (r *AntecedentePostgresRepository) Save(antecedente *domain.Antecedente) error {
	query := `
		INSERT INTO antecedente (id, historia_id, tipo, descripcion)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET
			historia_id = EXCLUDED.historia_id,
			tipo = EXCLUDED.tipo,
			descripcion = EXCLUDED.descripcion;
	`

	_, err := r.db.Exec(query,
		antecedente.ID,
		antecedente.HistoriaID,
		antecedente.Tipo,
		antecedente.Descripcion,
	)
	if err != nil {
		return fmt.Errorf("error al guardar antecedente con ID %s: %w", antecedente.ID, err)
	}

	return nil
}

func (r *AntecedentePostgresRepository) FindByID(id string) (*domain.Antecedente, error) {
	query := `
		SELECT id, historia_id, tipo, descripcion
		FROM antecedente
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)

	var a domain.Antecedente
	if err := row.Scan(&a.ID, &a.HistoriaID, &a.Tipo, &a.Descripcion); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no se encontró antecedente con ID %s", id)
		}
		return nil, fmt.Errorf("error al buscar antecedente con ID %s: %w", id, err)
	}

	return &a, nil
}

func (r *AntecedentePostgresRepository) Modificar(antecedenteModificado *domain.Antecedente) error {
	query := `
		UPDATE antecedente
		SET historia_id = $1,
			tipo = $2,
			descripcion = $3
		WHERE id = $4;
	`

	result, err := r.db.Exec(query,
		antecedenteModificado.HistoriaID,
		antecedenteModificado.Tipo,
		antecedenteModificado.Descripcion,
		antecedenteModificado.ID,
	)
	if err != nil {
		return fmt.Errorf("error al modificar antecedente con ID %s: %w", antecedenteModificado.ID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar modificación de antecedente con ID %s: %w", antecedenteModificado.ID, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró antecedente con ID %s para modificar", antecedenteModificado.ID)
	}

	return nil
}

func (r *AntecedentePostgresRepository) Delete(id string) error {
	query := `DELETE FROM antecedente WHERE id = $1;`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar antecedente con ID %s: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas al eliminar antecedente con ID %s: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró antecedente con ID %s para eliminar", id)
	}

	return nil
}

func (r *AntecedentePostgresRepository) FindAntecedentesByHistoriaID(historiaID string) ([]*domain.Antecedente, error) {
	query := `
		SELECT id, historia_id, tipo, descripcion
		FROM antecedente
		WHERE historia_id = $1;
	`

	rows, err := r.db.Query(query, historiaID)
	if err != nil {
		return nil, fmt.Errorf("error al consultar antecedentes por historia_id %s: %w", historiaID, err)
	}
	defer rows.Close()

	var antecedentes []*domain.Antecedente
	for rows.Next() {
		var a domain.Antecedente
		if err := rows.Scan(&a.ID, &a.HistoriaID, &a.Tipo, &a.Descripcion); err != nil {
			return nil, fmt.Errorf("error al escanear antecedente: %w", err)
		}
		antecedentes = append(antecedentes, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de resultados de antecedentes: %w", err)
	}

	return antecedentes, nil
}
