package transport

import (
	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// Login request
// swagger:parameters login
type swaggLoginReq struct {
	// in:body
	Body credentials
}

// Login response
// swagger:response loginResp
type swaggLoginResp struct {
	// in:body
	Body struct {
		*PD_BookingBR_Evaluation.AuthToken
	}
}

// Token refresh response
// swagger:response refreshResp
type swaggRefreshResp struct {
	// in:body
	Body struct {
		*PD_BookingBR_Evaluation.RefreshToken
	}
}
