package pgsql

import (
	"github.com/go-pg/pg/v9/orm"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db orm.DB, id int) (PD_BookingBR_Evaluation.User, error) {
	var user PD_BookingBR_Evaluation.User
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(&user, sql, id)
	return user, err
}

// FindByUsername queries for single user by username
func (u User) FindByUsername(db orm.DB, uname string) (PD_BookingBR_Evaluation.User, error) {
	var user PD_BookingBR_Evaluation.User
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."username" = ? and deleted_at is null)`
	_, err := db.QueryOne(&user, sql, uname)
	return user, err
}

// FindByToken queries for single user by token
func (u User) FindByToken(db orm.DB, token string) (PD_BookingBR_Evaluation.User, error) {
	var user PD_BookingBR_Evaluation.User
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."token" = ? and deleted_at is null)`
	_, err := db.QueryOne(&user, sql, token)
	return user, err
}

// Update updates user's info
func (u User) Update(db orm.DB, user PD_BookingBR_Evaluation.User) error {
	return db.Update(&user)
}
