package transport

import "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"

// Hotel model response
// swagger:response userResp
type swaggHotelResponse struct {
	// in:body
	Body struct {
		*PD_BookingBR_Evaluation.Hotel
	}
}

// Hotels model response
// swagger:response userListResp
type swaggHotelListResponse struct {
	// in:body
	Body struct {
		Users []PD_BookingBR_Evaluation.Hotel `json:"hotels"`
		Page  int                             `json:"page"`
	}
}
