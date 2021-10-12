package api

import (
	"crypto/sha1"
	"os"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/zlog"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/auth"
	al "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/auth/logging"
	at "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/auth/transport"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/password"
	pl "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/password/logging"
	pt "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/password/transport"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/user"
	ul "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/user/logging"
	ut "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/api/user/transport"

	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/config"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/jwt"
	authMw "github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/middleware/auth"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/postgres"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/rbac"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/secure"
	"github.com/projeto-de-algoritmos/PD_BookingBR_Evaluation/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := postgres.New(os.Getenv("DATABASE_URL"), cfg.DB.Timeout, cfg.DB.LogQueries)
	if err != nil {
		return err
	}

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.Service{}
	jwt, err := jwt.New(cfg.JWT.SigningAlgorithm, os.Getenv("JWT_SECRET"), cfg.JWT.DurationMinutes, cfg.JWT.MinSecretLength)
	if err != nil {
		return err
	}

	log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	authMiddleware := authMw.Middleware(jwt)

	at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, authMiddleware)

	v1 := e.Group("/v1")
	v1.Use(authMiddleware)

	ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
