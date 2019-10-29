package controller

import (
	"fmt"
	"github.com/go-pg/pg"
	"net/http"
	"time"
)

func IndexController(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)



	/*var users []User
	users = db.Model(&users).Select()
	fmt.Println(users)
*/

	currentTime := time.Now()

	fmt.Fprint(w, "datetime: " + currentTime.Format("01-02-2006 15:04:05"))
}

func IndexControllerFactory(db *pg.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, db.Options())
	}
}
/*
func (app *app) IndexAction(w http.ResponseWriter, r *http.Request, db *pg.DB)  {



	var users []entity.User

	db.Model(&users).Select()

	for _, value := range users {
		fmt.Println(value.Name)
	}
}
*/