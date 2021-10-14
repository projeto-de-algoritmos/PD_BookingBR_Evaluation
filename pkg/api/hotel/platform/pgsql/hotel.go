package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg/v9"

	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
)

// Hotel represents the client for hotel table
type Hotel struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Hotel already exists.")
)

// Create creates a new hotel on database
func (u Hotel) Create(db orm.DB, usr PD_BookingBR_Evaluation.Hotel) (PD_BookingBR_Evaluation.Hotel, error) {
	var hotel = new(PD_BookingBR_Evaluation.Hotel)
	err := db.Model(hotel).Where("lower(name) = ? or value = ? and deleted_at is null",
		strings.ToLower(hotel.Name), hotel.Value).Select()
	if err == nil || err != pg.ErrNoRows {
		return PD_BookingBR_Evaluation.Hotel{}, ErrAlreadyExists
	}

	err = db.Insert(&usr)
	return *hotel, err
}

// View returns single hotel by ID
func (u Hotel) View(db orm.DB, id int) (PD_BookingBR_Evaluation.Hotel, error) {
	var hotel PD_BookingBR_Evaluation.Hotel
	sql := `SELECT "hotel".*, "name"."value"."CompanyID"
	FROM "hotels" AS "hotel"
	WHERE ("hotel"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(&hotel, sql, id)
	return hotel, err
}

// Update updates hotel's contact info
func (u Hotel) Update(db orm.DB, hotel PD_BookingBR_Evaluation.Hotel) error {
	_, err := db.Model(&hotel).WherePK().UpdateNotZero()
	return err
}

// List returns list of all hotels retrievable for the current user, depending on role
func (u Hotel) List(db orm.DB, qp *PD_BookingBR_Evaluation.ListQuery, p PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.Hotel, error) {
	var hotels []PD_BookingBR_Evaluation.Hotel
	q := db.Model(&hotels).Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("hotel.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	err := q.Select()
	return hotels, err
}

// Delete sets deleted_at for a hotel
func (u Hotel) Delete(db orm.DB, hotel PD_BookingBR_Evaluation.Hotel) error {
	return db.Delete(&hotel)
}
