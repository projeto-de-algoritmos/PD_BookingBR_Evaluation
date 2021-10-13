package mock

import (
	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(PD_BookingBR_Evaluation.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u PD_BookingBR_Evaluation.User) (string, error) {
	return j.GenerateTokenFn(u)
}
