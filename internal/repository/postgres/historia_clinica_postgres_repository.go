package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type HistoriaClinicaPostgresRepository struct {
	db *sql.DB
}

// Para crear instancia del repo en otro archivo
func NewHistoriaClinicaPostgresRepository(db *sql.DB) *HistoriaClinicaPostgresRepository {
	return &HistoriaClinicaPostgresRepository{db: db}
}

func (r *HistoriaClinicaPostgresRepository) Save(historia *domain.HistoriaClinica) error {
	query := `
		INSERT INTO historia_clinica (id, pacienteid, turnoid, diagnostico, tratamiento)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE SET
		pacienteid = EXCLUDE.pacienteid,
		turnoid = EXCLUDE.turnoid,
		diagnostico = EXCLUDE.diagnostico,
		tratamiento = EXCLUDE.tratamiento;
	`
	_, err := r.db.Exec(query,
		historia.ID,
		historia.PacienteID,
		historia.TurnoID,
		historia.Diagnostico,
		historia.Tratamiento,
	)
	if err != nil {
		return fmt.Errorf("error al guardar la historia clinica: %w", err)
	}

	dbTransaction, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar la transaccion para alergias: %w", err)
	}

	defer func() {
		if rv := recover(); rv != nil {
			dbTransaction.Rollback()
			panic(rv)
		} else if err != nil {
			dbTransaction.Rollback()
		}
	}()
	//elimina alergias antiguas
	deleteQuery := `DELETE FROM historia_clinica_alergia WHERE historia_clinica_id = $1;`
	_, err = dbTransaction.Exec(deleteQuery, historia.ID)
	if err != nil {
		return fmt.Errorf("error al eliminar alergias antiguos de la historia clinica %s: %w", historia.ID, err)
	}
	// buildBatchInsert crea la query de insert a la tabla de asociacion
	if len(historia.Alergias) > 0 {
		query, args := buildBatchInsert[domain.Alergia](
			historia.ID,
			historia.Alergias,
			"historia_clinica_alergia",
			"alergia_id",
			func(a domain.Alergia) string { return a.ID },
		)
		_, err = dbTransaction.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("error al insertar alergias en la tabla de historia: %w", err)
		}

	}

	//eliminamos antecedentes antiguos
	deleteQuery = `DELETE FROM historia_clinica_antecedentes WHERE historia_clinica_id = $1;`
	_, err = dbTransaction.Exec(deleteQuery, historia.ID)
	if err != nil {
		return fmt.Errorf("error al eliminar antecedentes antiguos de la historia clinica %s: %w", historia.ID, err)
	}
	if len(historia.Antecedentes) > 0 {
		query, args := buildBatchInsert[domain.Antecedente](
			historia.ID,
			historia.Antecedentes,
			"historia_clinica_antecedentes",
			"antecedente_id",
			func(a domain.Antecedente) string { return a.ID },
		)
		_, err = dbTransaction.Exec(query, args...)
		if err != nil {
			return fmt.Errorf("error al insertar antecedentes en la tabla de historia: %w", err)
		}
	}

	err = dbTransaction.Commit()
	if err != nil {
		return fmt.Errorf("error al confirmar la transaccion de antecedentes: %w", err)
	}
	return nil
}

// funcion para crear consulta de asociacion
func buildBatchInsert[T any](padreID string, hijos []T, tablename, columnName string, getID func(T) string) (string, []any) {
	valueString := make([]string, 0, len(hijos))
	valueArgs := make([]any, 0, len(hijos)*2)

	for i, h := range hijos {
		valueString = append(valueString, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, padreID, getID(h))
	}
	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (historia_clinica_id, %s)
		VALUES %s;
	`, tablename, columnName, strings.Join(valueString, ","))
	return insertQuery, valueArgs
}

func (r *HistoriaClinicaPostgresRepository) GetAntecedentes(ids []string) ([]domain.Antecedente, error) {
	if len(ids) == 0 {
		return []domain.Antecedente{}, nil
	}

	valueString := make([]string, 0, len(ids))
	valueArgs := make([]any, 0, len(ids))

	for i, id := range ids {
		valueString = append(valueString, fmt.Sprintf("$%d", i+1))
		valueArgs = append(valueArgs, id)
	}
	inClause := strings.Join(valueString, ",")

	query := fmt.Sprintf("SELECT antecedente_id, historia_id, tipo, descripcion FROM antecedente WHERE antecedente_id IN (%s)", inClause)

	rows, err := r.db.Query(query, valueArgs...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta de antecedente: %w", err)
	}
	defer rows.Close()

	var antecedentes []domain.Antecedente

	//se convierten los datos crudos del sql a las entidades para su manejo
	for rows.Next() {
		var ant domain.Antecedente
		if err := rows.Scan(&ant.ID, &ant.HistoriaID, &ant.Tipo, &ant.Descripcion); err != nil {
			return nil, fmt.Errorf("error al escanear fila de antecedente: %w", err)
		} // escanea las columnas y setea a la entidad
		antecedentes = append(antecedentes, ant)
	}

	//verifica si hubo error durante el for
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteracion de filas de antecedente: %w", err)
	}

	return antecedentes, nil
}

func (r *HistoriaClinicaPostgresRepository) GetAlergias(ids []string) ([]domain.Alergia, error) {
	if len(ids) == 0 {
		return []domain.Alergia{}, nil
	}

	valueString := make([]string, 0, len(ids))
	valueArgs := make([]any, 0, len(ids))

	for i, id := range ids {
		valueString = append(valueString, fmt.Sprintf("$%d", i+1))
		valueArgs = append(valueArgs, id)
	}
	inClause := strings.Join(valueString, ",")

	query := fmt.Sprintf("SELECT alergia_id, historia_id, descripcion FROM alergia WHERE alergia_id IN (%s)", inClause)

	rows, err := r.db.Query(query, valueArgs...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta de alergia: %w", err)
	}
	defer rows.Close()

	var alergias []domain.Alergia

	//se convierten los datos crudos del sql a las entidades para su manejo
	for rows.Next() {
		var alerg domain.Alergia
		if err := rows.Scan(&alerg.ID, &alerg.HistoriaID, &alerg.Descripcion); err != nil {
			return nil, fmt.Errorf("error al escanear fila de Alergia: %w", err)
		} // escanea las columnas y setea a la entidad
		alergias = append(alergias, alerg)
	}

	//verifica si hubo error durante el for
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteracion de filas de alergia: %w", err)
	}

	return alergias, nil
}
