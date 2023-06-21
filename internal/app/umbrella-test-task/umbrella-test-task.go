package app

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/spatecon/echo-middleware-assessment/internal/pkg/endpoint"
	"github.com/spatecon/echo-middleware-assessment/internal/pkg/mw"
	"github.com/spatecon/echo-middleware-assessment/internal/pkg/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	log.Info().Msg("server running")

	err := a.echo.Start(":8080")
	if err != nil {
		return err
	}

	return nil
}
