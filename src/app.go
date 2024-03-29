package main
/*
import (
	"app/src/controller"
	"app/src/middleware"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type (
	App interface {
		Bootstrap(config *ConfigYaml) *AppService
		CreateRoutes() *AppService
		CreateHttpServer() *AppService
		Process() error
	}

	AppService struct {
		App
		Config *ConfigYaml
		Router *chi.Mux
		Server *http.Server
	}
)

func (app *AppService) Bootstrap (config *ConfigYaml) *AppService {
	app.Config = config
	app.Router = chi.NewRouter()
	app.CreateHttpServer()
	app.CreateRoutes()
	return app
}

func (app *AppService) Process() error {
	return app.Server.ListenAndServe()
}

func (app *AppService) CreateHttpServer() *AppService {

	addr := fmt.Sprintf("%s:%d", app.Config.GetHost(), app.Config.GetPort())

	app.Server = &http.Server{
		Addr:    addr,
		Handler: app.Router,
	}

	return app
}

func (app *AppService) CreateRoutes() *AppService {

	app.Router.Use(middleware.JwtToken)

	app.Router.Get(indexRoute, controller.IndexController)

	app.Router.Route(securityRoute, func(r chi.Router) {
		r.HandleFunc(securityLoginRoute, controller.UserHandler)
	})

	return app
}*/