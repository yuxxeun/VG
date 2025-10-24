package handler

import (
	"net/http"

	"handler/pkg/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router.Route(w, r)
}
