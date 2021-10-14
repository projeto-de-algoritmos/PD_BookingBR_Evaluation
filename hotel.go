package PD_BookingBR_Evaluation

// Hotel model
type Hotel struct {
	Base
	Name  string  `json:"name"`
	Value float64 `json:"value"`

	CompanyID int        `json:"company_id"`
	RoleID    AccessRole `json:"-"`
}
