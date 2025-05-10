package domain

import (
	"fmt"
	"time"
)

// String a enum
func ParseSexo(s string) (Sexo, error) {
	switch s {
	case "Masculino":
		return Masculino, nil
	case "Femenino":
		return Femenino, nil
	case "Otro":
		return Otro, nil
	default:
		return -1, fmt.Errorf("Sexo no valido: %s", s)
	}
}

type Paciente struct {
	ID        int
	Nomyape   string
	DNI       string
	Fechanac  time.Time
	Sexo      Sexo
	Email     string
	Direccion string
	Telefono  string
}

/*
type Paciente struct {
	ID        int       `json:"id"`
	Nomyape   string    `json:"nomyape" validate:"required,max=255"`
	DNI       string    `json:"dni" validate:"required,len=8, regexp=^[0-9]+$"`
	Fechanac  time.Time `json:"fechanac" validate:"required"`
	Sexo      Sexo      `json:"sexo"`
	Email     string    `json:"email" validate:"email"`
	Direccion string    `json:"direccion" validate:"required,max=255"`
	Telefono  string    `json:"telefono" validate:"required,max=20"`
}
*/

/*
type Turno struct {
    ID        int       `json:"id"`
    PacienteID int       `json:"paciente_id"`
    Fecha     time.Time `json:"fecha"`
}
Usa PacienteID int en la DB y convierte a Paciente solo en la capa de servicio.

¿Cómo hacer una Primary Key compuesta con 3 IDs?
En Go, no hay una forma nativa de definir Primary Keys compuestas en los structs, pero puedes hacer una estructura que actúe como clave:

type MedicoEspecialidad struct {
    MedicoID       int `json:"medico_id"`
    EspecialidadID int `json:"especialidad_id"`
    ClinicaID      int `json:"clinica_id"`
}

En la base de datos, esos tres campos juntos serán la clave primaria. En Go, simplemente defines una estructura y la manejas con gorm o el ORM que uses.


*/
