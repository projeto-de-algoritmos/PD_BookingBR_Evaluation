package transport

import "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*PD_BookingBR_Evaluation.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []PD_BookingBR_Evaluation.User `json:"users"`
		Page  int                            `json:"page"`
	}
}
