package query

import (
	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// List prepares data for list queries
func List(u PD_BookingBR_Evaluation.AuthUser) (*PD_BookingBR_Evaluation.ListQuery, error) {
	switch true {
	case u.Role <= PD_BookingBR_Evaluation.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == PD_BookingBR_Evaluation.CompanyAdminRole:
		return &PD_BookingBR_Evaluation.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == PD_BookingBR_Evaluation.LocationAdminRole:
		return &PD_BookingBR_Evaluation.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
