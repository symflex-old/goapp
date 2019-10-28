package controller

import (
	"fmt"
	"net/http"
	"time"
)

func IndexController(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)

	currentTime := time.Now()

	fmt.Fprint(w, "datetime: " + currentTime.Format("01-02-2006 15:04:05"))
}
