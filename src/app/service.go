package app

import (
	. "app/src/config"
	"app/src/controller"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"net/http"
)

type (
	GetDbMethod interface {
		GetDb() *pg.DB
	}

	Service interface {
		Bootstrap(config *ConfigYaml) *AppService
		CreateDbConnection() *AppService
		CreateRoutes() *AppService
		CreateHttpServer() *AppService
		Process() error
	}

	AppService struct {
		Service
		GetDbMethod
		Config *ConfigYaml
		Router *chi.Mux
		Db *pg.DB
		Server *http.Server
	}
)

func (app *AppService) GetDb() *pg.DB {
	return app.Db
}

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
		User: config.GetUser(),
		Password: config.GetPassword(),
		Database: config.GetBase(),
	})
	app.Db = db
	return app
}

func (app *AppService) App() *AppService {
	return app
}

func (app *AppService) Process() error {
	return app.Server.ListenAndServe()
}

func (app *AppService) CreateHttpServer() *AppService {

	addr := fmt.Sprintf("%s:%d", app.Config.GetHttp().GetHost(), app.Config.GetHttp().GetPort())

	app.Server = &http.Server{
		Addr:    addr,
		Handler: app.Router,
	}

	return app
}

func (app *AppService) CreateRoutes() *AppService {

	//routes.Export(app.Router, app.Db)

	app.Router.Get("/test", controller.IndexControllerFactory(app.Db))

	/*
	app.Router.Use(middleware.JwtToken)

	app.Router.Get(routes., controller.IndexController)

	app.Router.Route(securityRoute, func(r chi.Router) {
		r.HandleFunc(securityLoginRoute, controller.UserHandler)
	})*/

	return app
}
