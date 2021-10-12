package user

import (
	"github.com/labstack/echo"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/query"
)

func (u User) Create(c echo.Context, req PD_BookingBR_Evaluation.User) (PD_BookingBR_Evaluation.User, error) {
	if err := u.rbac.AccountCreate(c, req.RoleID); err != nil {
		return PD_BookingBR_Evaluation.User{}, err
	}
	req.Password = u.sec.Hash(req.Password)
	return u.udb.Create(u.db, req)
}

// List returns list of users
func (u User) List(c echo.Context, p PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.User, error) {
	au := u.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return u.udb.List(u.db, q, p)
}

// View returns single user
func (u User) View(c echo.Context, id int) (PD_BookingBR_Evaluation.User, error) {
	if err := u.rbac.EnforceUser(c, id); err != nil {
		return PD_BookingBR_Evaluation.User{}, err
	}
	return u.udb.View(u.db, id)
}

// Delete deletes a user
func (u User) Delete(c echo.Context, id int) error {
	user, err := u.udb.View(u.db, id)
	if err != nil {
		return err
	}
	if err := u.rbac.IsLowerRole(c, user.Role.AccessLevel); err != nil {
		return err
	}
	return u.udb.Delete(u.db, user)
}

// Update contains user's information used for updating
type Update struct {
	ID        int
	FirstName string
	LastName  string
	Mobile    string
	Phone     string
	Address   string
}

// Update updates user's contact information
func (u User) Update(c echo.Context, r Update) (PD_BookingBR_Evaluation.User, error) {
	if err := u.rbac.EnforceUser(c, r.ID); err != nil {
		return PD_BookingBR_Evaluation.User{}, err
	}

	if err := u.udb.Update(u.db, PD_BookingBR_Evaluation.User{
		Base:      PD_BookingBR_Evaluation.Base{ID: r.ID},
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Mobile:    r.Mobile,
		Address:   r.Address,
	}); err != nil {
		return PD_BookingBR_Evaluation.User{}, err
	}

	return u.udb.View(u.db, r.ID)
}