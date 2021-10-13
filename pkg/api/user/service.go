package user

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	PD_BookingBR_Evaluation "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/user/platform/pgsql"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, PD_BookingBR_Evaluation.User) (PD_BookingBR_Evaluation.User, error)
	List(echo.Context, PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.User, error)
	View(echo.Context, int) (PD_BookingBR_Evaluation.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (PD_BookingBR_Evaluation.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.User{}, rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, PD_BookingBR_Evaluation.User) (PD_BookingBR_Evaluation.User, error)
	View(orm.DB, int) (PD_BookingBR_Evaluation.User, error)
	List(orm.DB, *PD_BookingBR_Evaluation.ListQuery, PD_BookingBR_Evaluation.Pagination) ([]PD_BookingBR_Evaluation.User, error)
	Update(orm.DB, PD_BookingBR_Evaluation.User) error
	Delete(orm.DB, PD_BookingBR_Evaluation.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) PD_BookingBR_Evaluation.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, PD_BookingBR_Evaluation.AccessRole, int, int) error
	IsLowerRole(echo.Context, PD_BookingBR_Evaluation.AccessRole) error
}
