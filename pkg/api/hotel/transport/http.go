package transport

import (
	"net/http"
	"strconv"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/hotel"

	"github.com/labstack/echo"
)

// HTTP represents hotel http service
type HTTP struct {
	svc hotel.Service
}

// NewHTTP creates new hotel http service
func NewHTTP(svc hotel.Service, r *echo.Group) {
	h := HTTP{svc}
	ur := r.Group("/hotels")
	// swagger:route POST /v1/hotel hotel hotelCreate
	// Creates new hotel.
	// responses:
	//  200: hotelResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/hotel hotel listhotel
	// ---
	// summary: Returns list of hotel.
	// description: Returns list of hotel. Depending on the user role requesting it, it may return all hotel for SuperAdmin/Admin hotel, all company/location hotel for Company/Location admins, and an error for non-admin hotel.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/userListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/hotel/{id} hotel getUser
	// ---
	// summary: Returns a single user.
	// description: Returns a single user by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/hotelResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/:id", h.view)

	// swagger:operation PATCH /v1/hotel/{id} hotel hotelUpdate
	// ---
	// summary: Updates hotel's contact information
	// description: Updates hotel's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/userUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/hotel/{id} hotel userDelete
	// ---
	// summary: Deletes a user
	// description: Deletes a user with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of hotel
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.DELETE("/:id", h.delete)
}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

// User create request
// swagger:model userCreate
type createReq struct {
	Name  string  `json:"name" validate:"required"`
	Value float64 `json:"value" validate:"required"`

	CompanyID int                                `json:"company_id" validate:"required"`
	RoleID    PD_BookingBR_Evaluation.AccessRole `json:"role_id" validate:"required"`
}

func (h HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}

	if r.RoleID < PD_BookingBR_Evaluation.SuperAdminRole || r.RoleID > PD_BookingBR_Evaluation.UserRole {
		return PD_BookingBR_Evaluation.ErrBadRequest
	}

	htl, err := h.svc.Create(c, PD_BookingBR_Evaluation.Hotel{
		Name:      r.Name,
		Value:     r.Value,
		CompanyID: r.CompanyID,
		RoleID:    r.RoleID,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, htl)
}

type listResponse struct {
	Hotels []PD_BookingBR_Evaluation.Hotel `json:"hotels"`
	Page   int                             `json:"page"`
}

func (h HTTP) list(c echo.Context) error {
	var req PD_BookingBR_Evaluation.PaginationReq
	if err := c.Bind(&req); err != nil {
		return err
	}

	result, err := h.svc.List(c, req.Transform())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result, req.Page})
}

func (h HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return PD_BookingBR_Evaluation.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// User update request
// swagger:model userUpdate
type updateReq struct {
	ID    int     `json:"-"`
	Name  string  `json:"name" validate:"omitempty,min=4"`
	Value float64 `json:"value" validate:"omitempty,min=2"`
}

func (h HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return PD_BookingBR_Evaluation.ErrBadRequest
	}

	req := new(updateReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	htl, err := h.svc.Update(c, hotel.Update{
		ID:    id,
		Name:  req.Name,
		Value: req.Value,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, htl)
}

func (h HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return PD_BookingBR_Evaluation.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
