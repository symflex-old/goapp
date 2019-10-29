package routes

import (
	"app/src/controller"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"net/http"
)

const (
	indexRoute string = "/"
	loginRoute string = "/login"
	logoutRoute string = "/logout"


	securityRoute string = "/sec"
	securityLoginRoute string = "/login"

	//secRoute map[string]string = sli
)

func Export(router *chi.Mux, db *pg.DB) *chi.Mux  {
	router.Get(indexRoute, func(writer http.ResponseWriter, request *http.Request) {
		controller.IndexAction(writer, request, db)
	})

	return router
}