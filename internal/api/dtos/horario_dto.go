package dtos

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IvanMiranda1/TurneroMedico/internal/domain"
)

type HorarioRequest struct {
	HoraInicio string `json:"horainicio" validate:"required"`
	HoraFin    string `json:"horafin" validate:"required"`
}

func (r *HorarioRequest) ToDomain() (*domain.Horario, error) {
	partsInicio := strings.Split(r.HoraInicio, ":")
	partsFin := strings.Split(r.HoraFin, ":")
	if len(partsInicio) != 2 || len(partsFin) != 2 {
		return nil, fmt.Errorf("formato de hora invalido (esperado HHs:MM)")
	}

	hourInicio, err := strconv.Atoi(partsInicio[0])
	if err != nil {
		return nil, fmt.Errorf("Error al formatear horario")
	}

	minuteInicio, err := strconv.Atoi(partsInicio[1])
	if err != nil {
		return nil, fmt.Errorf("Error al formatear horario")
	}
	HoraInicio, err := domain.NewTimeOfDay(hourInicio, minuteInicio)
	if err != nil {
		return nil, fmt.Errorf("Error al formatear horario")
	}

	hourFin, err := strconv.Atoi(partsFin[0])
	if err != nil {
		return nil, fmt.Errorf("Error al formatear horario")
	}

	minuteFin, err := strconv.Atoi(partsFin[1])
	if err != nil {
		return nil, fmt.Errorf("Error al formatear horario")
	}
	HoraFin, err := domain.NewTimeOfDay(hourFin, minuteFin)
	if err != nil {
		return nil, fmt.Errorf("Error al formatear horario")
	}

	return domain.NewHorario(
		"0",
		HoraInicio,
		HoraFin,
	), nil
}

type HorarioResponse struct {
	ID         string `json:"id"`
	HoraInicio string `json:"horainicio"`
	HoraFin    string `json:"horafin"`
}

func HorarioFromDomain(h *domain.Horario) *HorarioResponse {
	return &HorarioResponse{
		ID:         h.ID,
		HoraInicio: h.HoraInicio.String(),
		HoraFin:    h.HoraFin.String(),
	}
}
