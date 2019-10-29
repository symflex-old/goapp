package routes

import "github.com/go-chi/chi"

const (
	indexRoute string = "/"
	loginRoute string = "/login"
	logoutRoute string = "/logout"


	securityRoute string = "/sec"
	securityLoginRoute string = "/login"

	//secRoute map[string]string = sli
)

type (
	ExportMethods interface {
		Export(router *chi.Mux) *chi.Mux
	}

	ExportRoutes struct {
		ExportMethods
	}
)


func (routes *ExportRoutes) exports(router *chi.Mux) *chi.Mux {

	router.

	return router
}