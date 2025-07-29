package dtos

type MedicoHorarioRequest struct {
	HorariosID []string `json:"horariosID" validate:"required"` //[]*Horario
	MedicoID   string   `json:"medicoID" validate:"required"`
	DiaSemana  string   `json:"diasemana" validate:"required"`
}

func NewMedicoHorarioRequest(horariosID []string, medicoID, diasemana string) *MedicoHorarioRequest {
	return &MedicoHorarioRequest{
		HorariosID: horariosID,
		MedicoID:   medicoID,
		DiaSemana:  diasemana,
	}
}

type MedicoHorarioResponse struct {
	ID         string   `json:"id"`
	HorariosID []string `json:"horariosID"`
	MedicoID   string   `json:"medicoID"`
	DiaSemana  string   `json:"diasemana"`
}

func NewMedicoHorarioResponse(id string, horariosID []string, medicoID, diasemana string) *MedicoHorarioResponse {
	return &MedicoHorarioResponse{
		ID:         id,
		HorariosID: horariosID,
		MedicoID:   medicoID,
		DiaSemana:  diasemana,
	}
}
