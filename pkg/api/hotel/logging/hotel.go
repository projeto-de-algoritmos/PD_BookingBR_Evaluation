package user

import (
	"time"

	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/hotel"
)

// New creates new user logging service
func New(svc hotel.Service, logger PD_BookingBR_Evaluation.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents user logging service
type LogService struct {
	hotel.Service
	logger PD_BookingBR_Evaluation.Logger
}

const name = "hotel"

// Create logging
func (ls *LogService) Create(c echo.Context, req PD_BookingBR_Evaluation.Hotel) (resp PD_BookingBR_Evaluation.Hotel, err error) {
	defer func(begin time.Time) {
		req.Name = "xxx-redacted-xxx"
		ls.logger.Log(
			c,
			name, "Create hotel request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// List logging
func (ls *LogService) List(c echo.Context, req PD_BookingBR_Evaluation.Pagination) (resp []PD_BookingBR_Evaluation.Hotel, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List hotel request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, req)
}

// View logging
func (ls *LogService) View(c echo.Context, req int) (resp PD_BookingBR_Evaluation.Hotel, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View hotel request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, req int) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete hotel request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req hotel.Update) (resp PD_BookingBR_Evaluation.Hotel, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update user request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}
