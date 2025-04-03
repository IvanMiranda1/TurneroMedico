package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Sexo int

const (
	Masculino = iota
	Femenino
	Otro
)

// Mapea enum a string
func (s Sexo) String() string {
	return [...]string{"Masculino", "Femenino", "Otro"}[s]
}

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
	ID        int       `json:"id"`
	Nomyape   string    `json:"nomyape" validate:"required,max=255"`
	DNI       string    `json:"dni" validate:"required,len=8, regexp=^[0-9]+$"`
	Fechanac  time.Time `json:"fechanac" validate:"required,birthday"`
	Sexo      Sexo      `json:"sexo"`
	Email     string    `json:"email" validate:"email"`
	Direccion string    `json:"direccion" validate:"required,max=255"`
	Telefono  string    `json:"telefono" validate:"required,max=20"`
}

func BirthdayValidator(fl validator.FieldLevel) bool {
	fecha := fl.Field().Interface().(time.Time)
	ahora := time.Now()

	edad := ahora.Year() - fecha.Year()
	if ahora.YearDay() < fecha.YearDay() {
		edad-- // Ajuste si no ha cumplido años aún en el año actual
	}

	return edad > 0 && edad <= 150
}

func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterValidation("birthday", BirthdayValidator)
}

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
