package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type MedicoPostgresRepository struct {
	db *sql.DB
}

// Para crear instancia del repo en otro archivo
func NewMedicoPostgresRepository(db *sql.DB) *MedicoPostgresRepository {
	return &MedicoPostgresRepository{db: db}
}

func (r *MedicoPostgresRepository) Save(medico *domain.Medico) error {
	//LOGICA PARA INSERT Y UPDATE
	query := `INSERT INTO medico (id, nomyape, dni, fechanac, email, telefono, legajo, sexo)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			  ON CONFLICT (id) DO UPDATE SET
			  nomyape = EXCLUDE.nomyape,
			  dni = EXCLUDE.dni,
			  fechanac = EXCLUDE.fechanac,
			  email = EXCLUDE.email,
			  telefono = EXCLUDE.telefono,
			  legajo = EXCLUDE.legajo,
			  sexo = EXCLUDE.sexo;
			  `
	_, err := r.db.Exec(query,
		medico.ID,
		medico.Nomyape,
		medico.DNI,
		medico.Fechanac,
		medico.Email,
		medico.Telefono,
		medico.Legajo,
		medico.Sexo,
	)
	if err != nil {
		return fmt.Errorf("error al guardar el medico: %w", err)
	}

	dbTransaction, err := r.db.Begin() // 'dbTransaction' es tu 'tx' original
	if err != nil {
		return fmt.Errorf("error al iniciar la transacción para especialidades: %w", err)
	}
	//Asegurar que se revierta la transaccion si algo sale mal
	defer func() { // funcion diferida se ejecuta antes de que termine la funcion (sea por return o error (panic o error capturado))
		if rv := recover(); r != nil { //panic
			dbTransaction.Rollback() //Deshace cambios
			panic(rv)
		} else if err != nil { // error capturado
			dbTransaction.Rollback()
		}

	}()

	//Elimina especialidades antiguas
	deleteQuery := `DELETE FROM medico_especialidad WHERE medico_id = $1;`
	_, err = dbTransaction.Exec(deleteQuery, medico.ID)
	if err != nil {
		return fmt.Errorf("error al eliminar especialidades antiguas del medico %s: %w", medico.ID, err)
	}

	//Crear nuevas asociaciones
	if len(medico.Especialidad) > 0 {
		//                  //array   //cantidad inicial // espacio en memoria guardado (la cantidad de especialidad)
		//Crea un slice vaciov pero aunque aun no tenga nada, el espacio en memoria es suficiente para len(medico.especialidad)
		//evita re-asignaciones y copias innecesarias de memoria cuando se añaden dentro de un for
		valueString := make([]string, 0, len(medico.Especialidad))
		valueArgs := make([]any, 0, len(medico.Especialidad)*2)

		//valueString = ["($1, $2)", "($3, $4)", "($5, $6)"]
		//

		for i, especialidad := range medico.Especialidad {
			valueString = append(valueString, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
			//$ indica parametro de postgres y con %d insertamos un el valor de i
			//logrando crear ($1,$2) ($3,$4) para la consulta sql
			//postgres permite hacer muchas inserciones
			valueArgs = append(valueArgs, medico.ID, especialidad.ID)
		}
		/*
			tx.Exec(insertQuery, valueArgs...). El ... (operador de puntos suspensivos) en valueArgs... "desempaca" el slice valueArgs y pasa cada uno de sus elementos como un argumento individual a la función tx.Exec.
			entonces se insertan los valores a los placeholder en el orden que estan $1,$2,$3,$4,$5,$6
		*/

		insertQuery := `
			INSERT INTO medico_especialidad (medico_id, especialidad_id)
			VALUES %s;
			`
		insertQuery = fmt.Sprintf(insertQuery, strings.Join(valueString, ","))

		_, err = dbTransaction.Exec(insertQuery, valueArgs...)
		if err != nil {
			return fmt.Errorf("error al insertar nuevas especialidades para el medico %s: %w", medico.ID, err)
		}

	}
	err = dbTransaction.Commit()
	if err != nil {
		return fmt.Errorf("error al confirmar la transaccion  de especialidades: %w", err)
	}
	return nil //exito
}

// func (sera extension de) nombreDeFunc(name typevar) (todo lo que retorna...) -> en este caso (bool, error)
func (r *MedicoPostgresRepository) ExisteLegajoID(legajo string) (bool, error) {
	query := `SELECT COUNT(*) FROM medico WHERE legajo = $1;`

	var count int
	err := r.db.QueryRow(query, legajo).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia del legajo en la base de datos: %w", err)
	}

	return count > 0, nil
}

func (r *MedicoPostgresRepository) GetEspecialidades(ids []string) ([]domain.Especialidad, error) {
	if len(ids) == 0 {
		return []domain.Especialidad{}, nil
	}

	//ejemplo "SELECT id, nombre FROM especialidades WHERE id IN (?, ?, ?)"
	placeholders := make([]string, 0, len(ids))
	args := make([]interface{}, 0, len(ids)) //Para almacenar los ids como interface
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}
	inClause := strings.Join(placeholders, ",")

	query := fmt.Sprintf("SELECT id, nombre FROM especialidades WHERE id IN (%s)", inClause)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta de especialidades: %w", err)
	}
	defer rows.Close()

	var especialidades []domain.Especialidad

	//se convierten los datos crudos del sql a las entidades para su manejo
	for rows.Next() {
		var esp domain.Especialidad
		err := rows.Scan(&esp.ID, &esp.Nombre) // escanea las columnas y setea a la entidad
		if err != nil {
			return nil, fmt.Errorf("error al escanear fila de especialidad: %w", err)
		}
		especialidades = append(especialidades, esp)
	}

	//verifica si hubo error durante el for
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteracion de filas de especialidad: %w", err)
	}

	return especialidades, nil
}

func (r *MedicoPostgresRepository) ExisteTelefono(telefono string) (bool, error) {
	query := `SELECT COUNT(*) FROM medico WHERE telefono = $1;`

	var count int
	err := r.db.QueryRow(query, telefono).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia del telefono en la base de datos: %w", err)
	}

	return count > 0, nil
}

func (r *MedicoPostgresRepository) ExisteEmail(email string) (bool, error) {
	query := `SELECT COUNT(*) FROM medico WHERE email = $1;`

	var count int
	err := r.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar la existencia del email en la base de datos: %w", err)
	}

	return count > 0, nil
}
