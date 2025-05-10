package domain

import "fmt"

type Sexo int

const (
	Masculino = iota
	Femenino
	Otro
)

func (s Sexo) String() string {
	return [...]string{"Masculino", "Femenino", "Otro"}[s]
}

type TipoAntecendente int

const (
	Familiar = iota
	Patologico
	NoPatologico
	GinecoObstetrico
)

func (a TipoAntecendente) String() string {
	return [...]string{"Familiar", "Patologico", "NoPatologico", "GinecoObstetrico"}[a]
}

type TimeOfDay struct {
	Hour   int
	Minute int
	Second int
}

func NewTimeOfDay(hour, minute, second int) (TimeOfDay, error) {
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 || second < 0 || second > 59 {
		return TimeOfDay{}, fmt.Errorf("hora inválida: %02d:%02d:%02d", hour, minute, second)
	}
	return TimeOfDay{Hour: hour, Minute: minute, Second: second}, nil
}

func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

// comparar si una hora es posterior a otra
func (t TimeOfDay) IsAfter(other TimeOfDay) bool {
	if t.Hour != other.Hour {
		return t.Hour > other.Hour
	}
	if t.Minute != other.Minute {
		return t.Minute > other.Hour
	}
	return t.Second > other.Second

}
