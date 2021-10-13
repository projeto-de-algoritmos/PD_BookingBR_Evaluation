package mockdb

import (
	"github.com/go-pg/pg/v9/orm"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, PD_BookingBR_Evaluation.User) (PD_BookingBR_Evaluation.User, error)
	ViewFn           func(orm.DB, int) (PD_BookingBR_Evaluation.User, error)
	FindByUsernameFn func(orm.DB, string) (PD_BookingBR_Evaluation.User, error)
	FindByTokenFn    func(orm.DB, string) (PD_BookingBR_Evaluation.User, error)
	ListFn           func(orm.DB, *PD_BookingBR_Evaluation.ListQuery, PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.User, error)
	DeleteFn         func(orm.DB, PD_BookingBR_Evaluation.User) error
	UpdateFn         func(orm.DB, PD_BookingBR_Evaluation.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr PD_BookingBR_Evaluation.User) (PD_BookingBR_Evaluation.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (PD_BookingBR_Evaluation.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (PD_BookingBR_Evaluation.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *PD_BookingBR_Evaluation.ListQuery, p PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr PD_BookingBR_Evaluation.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr PD_BookingBR_Evaluation.User) error {
	return u.UpdateFn(db, usr)
}
