package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxwellpothier/wt-go-backend/routers"
)

func main() {
	r := mux.NewRouter()
	routers.SetIdentityRoutes(r)

	http.ListenAndServe(":8080", r)
}