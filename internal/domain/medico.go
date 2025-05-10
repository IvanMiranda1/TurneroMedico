package domain

import (
	"time"
)

type Medico struct {
	ID        string
	Nomyape   string
	DNI       string
	Fechanac  time.Time
	Email     string
	Direccion string
	Telefono  string
	Legajo    int
	Sexo      Sexo
	CreatedAt time.Time
	UpdateAt  time.Time
}

// medico debe ser mayor de 23 años para ejercer
func (m Medico) EsMayorDeEdad() bool {
	ageThreshold := time.Now().AddDate(-23, 0, 0)
	return !m.Fechanac.After(ageThreshold)
}

func NewMedico(id, nomyape, dni string, fechanac time.Time, email, direccion, telefono string, legajo int, sexo Sexo) *Medico {
	return &Medico{
		ID:        id,
		Nomyape:   nomyape,
		DNI:       dni,
		Fechanac:  fechanac,
		Email:     email,
		Direccion: direccion,
		Telefono:  telefono,
		Legajo:    legajo,
		Sexo:      sexo,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}

/*
type Medico struct {
	ID        int       `json:"id"`
	Nomyape   string    `json:"nomyape" validate:"required,max=255"`
	DNI       string    `json:"dni" validate:"required,len=8, regexp=^[0-9]+$"`
	Fechanac  time.Time `json:"fechanac" validate:"required, birthday"`
	Email     string    `json:"email" validate:"email"`
	Direccion string    `json:"direccion" validate:"required,max=255"`
	Telefono  string    `json:"telefono" validate:"required,max=20"`
	Legajo    int       `json:"legajo" validate:"required,len=4"`
} */
