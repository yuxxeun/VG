package handler

import (
	"net/http"

	"github.com/yuxxeun/VG/internal/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router.Route(w, r)
}
