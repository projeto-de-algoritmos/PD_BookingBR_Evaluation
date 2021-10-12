package query

import (
	"github.com/labstack/echo"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

func List(u PD_BookingBR_Evaluation.AuthUser) (*PD_BookingBR_Evaluation.ListQuery, error) {
	switch true {
	case u.Role <= PD_BookingBR_Evaluation.AdminRole:
		return nil, nil
	case u.Role <= PD_BookingBR_Evaluation.CompanyAdminRole:
		return &PD_BookingBR_Evaluation.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil

	default:
		return nil, echo.ErrForbidden
	}
}
