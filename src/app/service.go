package app

import (
	. "app/src/config"
	"app/src/routes"
	"app/src/controller"
	"app/src/middleware"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"net/http"
)

type (
	Service interface {
		Bootstrap(config *ConfigYaml) *AppService
		CreateDbConnection() *AppService
		CreateRoutes() *AppService
		CreateHttpServer() *AppService
		Process() error
	}

	GetConfigSection interface {
		GetDbSection() *DbSection
		GetHttpSection() *HttpSection
	}

	AppService struct {
		Service
		Config *ConfigYaml
		Router *chi.Mux
		Db *pg.DB
		Server *http.Server
	}
)

func (app *AppService) Bootstrap (config *ConfigYaml) *AppService {
	app.Config = config
	app.Router = chi.NewRouter()
	app.CreateDbConnection()
	app.CreateHttpServer()
	app.CreateRoutes()
	return app
}

func (app *AppService) CreateDbConnection() *AppService {
	var config = *app.Config.GetDb()
	db := pg.Connect(&pg.Options{
		Addr: config.GetHost(),
		User: config.GetUser(),
		Password: config.GetPassword(),
		Database: config.GetBase(),
	})
	defer db.Close()
	app.Db = db
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

	routes.e

	/*
	app.Router.Use(middleware.JwtToken)

	app.Router.Get(routes., controller.IndexController)

	app.Router.Route(securityRoute, func(r chi.Router) {
		r.HandleFunc(securityLoginRoute, controller.UserHandler)
	})*/

	return app
}
