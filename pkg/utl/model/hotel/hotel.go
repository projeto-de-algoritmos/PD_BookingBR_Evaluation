package hotel

import (
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// Hotel model
type Hotel struct {
	PD_BookingBR_Evaluation.Base
	Name  string  `json:"name"`
	Value float64 `json:"value"`

	CompanyID int `json:"company_id"`
}
