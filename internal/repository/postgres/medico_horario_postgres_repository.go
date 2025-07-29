package postgres

import (
	"database/sql"
	"fmt"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
	"github.com/gofrs/uuid"
)

type HorarioMedicoPostgresRepository struct {
	db *sql.DB
}

func NewHorarioMedicoPostgresRepository(db *sql.DB) *HorarioMedicoPostgresRepository {
	return &HorarioMedicoPostgresRepository{db: db}
}

func (r *HorarioMedicoPostgresRepository) Save(mh *domain.MedicoHorario) error {
	/* Si se inserta un medico_horario para el lunes en donde le medico trabaja
	por ejemplo de 8 a 12 y por la tarde de 14 a 18 el horario []string tendra dos valores
	y por eso se usa el for, porque como en la bd no se puede
	guardar arreglos deberia hacer dos insert... */
	for _, horarioID := range mh.HorariosID {
		id := uuid.Must(uuid.NewV4()).String()
		query := `INSERT INTO medico_horario (id, medico_id, horario_id, dia_semana) VALUES ($1, $2, $3, $4)`
		_, err := r.db.Exec(query, id, mh.MedicoID, horarioID, mh.DiaSemana)
		if err != nil {
			return fmt.Errorf("error al guardar horario del medico: %w", err)
		}
	}
	return nil
}

// no se hace un upsert porque al trabajar con muchas tablas no se puede
func (r *HorarioMedicoPostgresRepository) Modificar(mh *domain.MedicoHorario) error {
	deleteQuery := "DELETE FROM medico_horario WHERE medico_id = $1 AND dia_semana = $2"
	_, err := r.db.Exec(deleteQuery, mh.MedicoID, mh.DiaSemana)

	if err != nil {
		return fmt.Errorf("error al eliminar horarios anteriores del medico: %w", err)
	}

	for _, horarioID := range mh.HorariosID {
		id := uuid.Must(uuid.NewV4()).String()
		insertQuery := "INSERT INTO medico_horario (medico_horario_id, medico_id, horario_id, dia_semana) VALUES ($1,$2,$3,$4)"
		_, err := r.db.Exec(insertQuery, id, mh.MedicoID, horarioID, mh.DiaSemana)
		if err != nil {
			return fmt.Errorf("error al insertar nuevo horario del medico :%w", err)
		}
	}
	return nil
}

func (r *HorarioMedicoPostgresRepository) FindByMedicoYDia(medicoID string, dia string) (*domain.MedicoHorario, error) {
	query := "SELECT horario_id FROM medico_horario WHERE medico_id = $1 AND dia_semana = $2"

	//Ejecutamos la query, y nos devuelve un *rows que permite iterar sobre los resultados
	rows, err := r.db.Query(query, medicoID, dia)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar consulta: %w", err)
	}
	defer rows.Close()

	var horarios []string
	//itera sobre cada fila
	for rows.Next() {
		var horarioID string
		//setea el valor de la consulta, la columna horario_id se asigna a la primera variable
		//se escanea en el orden de la consulta, osea si es SELECT horarioid, medicoid, diasemana
		//dentro del rows.scan deberia estar en el mismo orden o se asignaria mal el valor
		if err := rows.Scan(&horarioID); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}
		horarios = append(horarios, horarioID)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar resultados: %w", err)
	}

	if len(horarios) == 0 {
		return nil, domain.ErrNotFound
	}
	return &domain.MedicoHorario{
		MedicoID:   medicoID,
		HorariosID: horarios,
		DiaSemana:  dia,
	}, nil
}

/*
¿Qué es rows?
rows es una especie de cursor que te deja:
    Recorrer resultados de una query (rows.Next())
    Leer cada fila (rows.Scan(...))
    Verificar errores que ocurrieron durante el recorrido (rows.Err())
*/
