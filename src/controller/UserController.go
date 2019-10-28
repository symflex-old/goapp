package controller

import (
	"app/src/entity"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type UserControllerInterface interface {
	GetUser(uuid uuid.UUID) *entity.User
}

type UserController struct {
	UserControllerInterface
}
/*
func (controller UserController) GetUser(uuid uuid.UUID) *entity.User {

}*/

func UserHandler(w http.ResponseWriter, r *http.Request)  {

	fmt.Println(r.Context())
	/*w.WriteHeader(http.StatusOK)

	currentTime := time.Now()

	fmt.Fprint(w, "datetime: " + currentTime.Format("01-02-2006 15:04:05"))*/

	//return
}
