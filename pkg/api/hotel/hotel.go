package hotel

import (
	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/query"
)

// Create creates a new hotel account
func (u Hotel) Create(c echo.Context, req PD_BookingBR_Evaluation.Hotel) (PD_BookingBR_Evaluation.Hotel, error) {
	if err := u.rbac.AccountCreate(c, 100, req.CompanyID, 1); err != nil {
		return PD_BookingBR_Evaluation.Hotel{}, err
	}
	return u.udb.Create(u.db, req)
}

// List returns list of hotels
func (u Hotel) List(c echo.Context, p PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.Hotel, error) {
	au := u.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return u.udb.List(u.db, q, p)
}

// View returns single hotel
func (u Hotel) View(c echo.Context, id int) (PD_BookingBR_Evaluation.Hotel, error) {
	if err := u.rbac.EnforceUser(c, id); err != nil {
		return PD_BookingBR_Evaluation.Hotel{}, err
	}
	return u.udb.View(u.db, id)
}

// Delete deletes a hotel
func (u Hotel) Delete(c echo.Context, id int) error {
	hotel, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	if err := u.rbac.IsLowerRole(c, 200); err != nil {
		return err
	}
	return u.udb.Delete(u.db, hotel)
}

// Update contains hotel's information used for updating
type Update struct {
	ID    int
	Name  string
	Value float64
}

// Update updates hotel's contact information
func (u Hotel) Update(c echo.Context, r Update) (PD_BookingBR_Evaluation.Hotel, error) {
	if err := u.rbac.EnforceUser(c, r.ID); err != nil {
		return PD_BookingBR_Evaluation.Hotel{}, err
	}

	if err := u.udb.Update(u.db, PD_BookingBR_Evaluation.Hotel{
		Base:  PD_BookingBR_Evaluation.Base{ID: r.ID},
		Name:  r.Name,
		Value: r.Value,
	}); err != nil {
		return PD_BookingBR_Evaluation.Hotel{}, err
	}

	return u.udb.View(u.db, r.ID)
}
