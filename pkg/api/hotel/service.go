package hotel

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/hotel/platform/pgsql"
)

// Service represents hotel application interface
type Service interface {
	Create(echo.Context, PD_BookingBR_Evaluation.Hotel) (PD_BookingBR_Evaluation.Hotel, error)
	List(echo.Context, PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.Hotel, error)
	View(echo.Context, int) (PD_BookingBR_Evaluation.Hotel, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (PD_BookingBR_Evaluation.Hotel, error)
}

// New creates new hotel application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Hotel {
	return &Hotel{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes Hotel application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Hotel {
	return New(db, pgsql.Hotel{}, rbac, sec)
}

// Hotel represents hotel application service
type Hotel struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents hotel repository interface
type UDB interface {
	Create(orm.DB, PD_BookingBR_Evaluation.Hotel) (PD_BookingBR_Evaluation.Hotel, error)
	View(orm.DB, int) (PD_BookingBR_Evaluation.Hotel, error)
	List(orm.DB, *PD_BookingBR_Evaluation.ListQuery, PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.Hotel, error)
	Update(orm.DB, PD_BookingBR_Evaluation.Hotel) error
	Delete(orm.DB, PD_BookingBR_Evaluation.Hotel) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) PD_BookingBR_Evaluation.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, PD_BookingBR_Evaluation.AccessRole, int, int) error
	IsLowerRole(echo.Context, PD_BookingBR_Evaluation.AccessRole) error
}
