package domain

import (
	"fmt"
	"strconv"
	"strings"
)

// Tipo de datos reutilizables, utilizados en la logica de negocio

type Sexo int

const (
	Masculino = iota
	Femenino
	Otro
)

// enum a string
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

type TipoAntecedente int

const (
	Familiar = iota
	Patologico
	NoPatologico
	GinecoObstetrico
)

// String a enum
func ParseTipoAntecedente(s string) (TipoAntecedente, error) {
	switch s {
	case "Familiar":
		return Familiar, nil
	case "Patologico":
		return Patologico, nil
	case "No Patologico":
		return NoPatologico, nil
	case "GinecoObstetrico":
		return GinecoObstetrico, nil
	default:
		return -1, fmt.Errorf("TipoAntecedente no valido: %s", s)
	}
}

func (a TipoAntecedente) String() string {
	return [...]string{"Familiar", "Patologico", "NoPatologico", "GinecoObstetrico"}[a]
}

// Me sirve para poder usar un rango de horarios para las jornadas laborales
// tambien puede ser horario para rango de horarios de los turnos es decir, horarios de 45min?
type TimeOfDay struct {
	Hour   int
	Minute int
}

func NewTimeOfDay(hour, minute int) (TimeOfDay, error) {
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return TimeOfDay{}, fmt.Errorf("hora inválida: %02d:%02d", hour, minute)
	}
	return TimeOfDay{Hour: hour, Minute: minute}, nil
}

func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
}

// comparar si una hora es posterior a otra
func (t TimeOfDay) IsAfter(other TimeOfDay) bool {
	if t.Hour != other.Hour {
		return t.Hour > other.Hour
	}
	return t.Minute > other.Hour

}

func ParseTimeOfDay(input string) (TimeOfDay, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return TimeOfDay{}, fmt.Errorf("formato inválido, se esperaba 'HH:MM'")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return TimeOfDay{}, fmt.Errorf("hora inválida: %w", err)
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return TimeOfDay{}, fmt.Errorf("minuto inválido: %w", err)
	}

	return NewTimeOfDay(hour, minute)
}
