package auth

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/auth/platform/pgsql"
)

// New creates new iam service
func New(db *pg.DB, udb UserDB, j TokenGenerator, sec Securer, rbac RBAC) Auth {
	return Auth{
		db:   db,
		udb:  udb,
		tg:   j,
		sec:  sec,
		rbac: rbac,
	}
}

// Initialize initializes auth application service
func Initialize(db *pg.DB, j TokenGenerator, sec Securer, rbac RBAC) Auth {
	return New(db, pgsql.User{}, j, sec, rbac)
}

// Service represents auth service interface
type Service interface {
	Authenticate(echo.Context, string, string) (PD_BookingBR_Evaluation.AuthToken, error)
	Refresh(echo.Context, string) (string, error)
	Me(echo.Context) (PD_BookingBR_Evaluation.User, error)
}

// Auth represents auth application service
type Auth struct {
	db   *pg.DB
	udb  UserDB
	tg   TokenGenerator
	sec  Securer
	rbac RBAC
}

// UserDB represents user repository interface
type UserDB interface {
	View(orm.DB, int) (PD_BookingBR_Evaluation.User, error)
	FindByUsername(orm.DB, string) (PD_BookingBR_Evaluation.User, error)
	FindByToken(orm.DB, string) (PD_BookingBR_Evaluation.User, error)
	Update(orm.DB, PD_BookingBR_Evaluation.User) error
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(PD_BookingBR_Evaluation.User) (string, error)
}

// Securer represents security interface
type Securer interface {
	HashMatchesPassword(string, string) bool
	Token(string) string
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) PD_BookingBR_Evaluation.AuthUser
}