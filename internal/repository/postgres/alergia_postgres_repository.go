package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type AlergiaPostgresRepository struct {
	db *sql.DB
}

func NewAlergiaPostgresRepository(db *sql.DB) *AlergiaPostgresRepository {
	return &AlergiaPostgresRepository{db: db}
}

func (r *AlergiaPostgresRepository) Save(alergia *domain.Alergia) error {
	query := `
		INSERT INTO alergia (id, historia_id, descripcion)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET
			historia_id = EXCLUDED.historia_id,
			descripcion = EXCLUDED.descripcion;
	`

	_, err := r.db.Exec(query,
		alergia.ID,
		alergia.HistoriaID,
		alergia.Descripcion,
	)
	if err != nil {
		return fmt.Errorf("error al guardar la alergia con ID %s: %w", alergia.ID, err)
	}

	return nil
}

func (r *AlergiaPostgresRepository) FindByID(id string) (*domain.Alergia, error) {
	query := `
		SELECT id, historia_id, descripcion
		FROM alergia
		WHERE id = $1;
	`

	row := r.db.QueryRow(query, id)

	var a domain.Alergia
	if err := row.Scan(&a.ID, &a.HistoriaID, &a.Descripcion); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no se encontró alergia con ID %s", id)
		}
		return nil, fmt.Errorf("error al buscar alergia con ID %s: %w", id, err)
	}

	return &a, nil
}

func (r *AlergiaPostgresRepository) DeleteAlergia(id string) error {
	query := `DELETE FROM alergia WHERE id = $1;`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar alergia con ID %s: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener filas afectadas al eliminar alergia con ID %s: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró alergia con ID %s para eliminar", id)
	}

	return nil
}

func (r *AlergiaPostgresRepository) FindAlergiasByHistoriaID(historiaID string) ([]*domain.Alergia, error) {
	query := `
		SELECT id, historia_id, descripcion
		FROM alergia
		WHERE historia_id = $1;
	`
	//consulta sql

	rows, err := r.db.Query(query, historiaID) // el id se inserta al placeholder $1 de la consulta
	if err != nil {
		return nil, fmt.Errorf("error al consultar alergias por historia_id %s: %w", historiaID, err)
	}
	//defer pospone la funcion y la ejecuta cuando termine el bloque de codigo que la contiene
	//en este caso posponemos Close
	defer rows.Close()
	// se cierra la consulta, porque Rows es un puntero activo,
	// Go no carga todos los resultados de una vez, en su caso :
	// * Crea un puntero activo en la base de datos
	// * el puntero permanece abierto mientras se recorren los resultados con rows.Next()
	// * si no se cierra rows.close(), esta conexion sigue ocupada aunque termine la funcion y puede llevar a
	// 1) fuga de recursos
	// 2) llegar al limite de conexiones del pool
	// 3) Comportamiento impredecible si se reusa la conexión sin haber cerrado el cursor

	var alergias []*domain.Alergia
	for rows.Next() {
		var a domain.Alergia
		if err := rows.Scan(&a.ID, &a.HistoriaID, &a.Descripcion); err != nil {
			return nil, fmt.Errorf("error al escanear alergia: %w", err)
		}
		alergias = append(alergias, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en la iteración de resultados de alergias: %w", err)
	}

	return alergias, nil
}
