package pgsql

import (
	"github.com/go-pg/pg/v9/orm"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
	user := PD_BookingBR_Evaluation.User{Base: PD_BookingBR_Evaluation.Base{ID: id}}
	err := db.Select(&user)
	return user, err
}

// Update updates user's info
func (u User) Update(db orm.DB, user PD_BookingBR_Evaluation.User) error {
	return db.Update(&user)
}
