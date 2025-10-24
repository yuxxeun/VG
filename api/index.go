package handler

import (
	"net/http"

	"github.com/yuxxeun/VG/pkg/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router.Route(w, r)
}
