package controller

import (
	"fmt"
	"log"
	"net/http"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"
	_ "github.com/volatiletech/authboss/v3/auth"

	// "github.com/volatiletech/authboss/v3/lock"
	_ "github.com/volatiletech/authboss/v3/logout"

	// _ "github.com/volatiletech/authboss/v3/recover"
	_ "github.com/volatiletech/authboss/v3/register"

	"github.com/MatthewBehnke/exampleGoApi/internal/delivery/middleware"
	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"
	"github.com/MatthewBehnke/exampleGoApi/pkg/httpserver"
)

func New(i *do.Injector) (*Controller, error) {
	c := &Controller{
		config:          do.MustInvoke[*entity.Config](i),
		log:             do.MustInvoke[*usecase.LoggerUseCase](i).WithSubsystem("controller"),
		authBossUseCase: do.MustInvoke[*usecase.AuthBossUseCase](i),
		httpV1:          do.MustInvoke[*HttpV1](i),
	}
	return c, nil
}

type Controller struct {
	config          *entity.Config
	log             usecase.Logger
	authBossUseCase usecase.AuthBoss
	httpV1          *HttpV1
	httpServer      *httpserver.Server
}

func (c *Controller) Shutdown() error {
	err := c.httpServer.Shutdown()
	if err != nil {
		return fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
	log.Print("controller service shutdown")
	return nil
}

func (c *Controller) Run() {
	// HTTP Server
	mux := chi.NewRouter()
	mux.Use(chimiddleware.RequestID)
	mux.Use(chimiddleware.RealIP)
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middleware.NewStructuredLogger(c.log))

	// loading usecase's onto context
	mux.Use(middleware.LoggerCtx(c.log))

	ab := newAuthentication(c.config, c.authBossUseCase, c.log)

	mux.Use(ab.LoadClientStateMiddleware)
	// mux.Use(remember.Middleware(ab))

	mux.Get("/metrics", promhttp.Handler().(http.HandlerFunc))

	// Mount the router to a path (this should be the same as the Mount path above)
	// mux in this example is a chi router, but it could be anything that can route to
	// the Core.Router.
	mux.Group(func(mux chi.Router) {
		mux.Use(authboss.ModuleListMiddleware(ab))
		mux.Mount("/auth", http.StripPrefix("/auth", ab.Config.Core.Router))
	})

	// Protected Routes
	mux.Group(func(r chi.Router) {
		r.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized))
		// r.Use(lock.Middleware(ab))

		// Openapi generated interfaces
		HandlerFromMuxWithBaseURL(c.httpV1, r, "/v1")
	})

	c.httpServer = httpserver.New(mux, httpserver.Port(c.config.HTTP.Port))
}
