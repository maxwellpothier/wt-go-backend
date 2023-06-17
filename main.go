package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxwellpothier/wt-go-backend/routers"
)

func main() {
	r := mux.NewRouter()
	r.Use(RequestLogger)
	routers.SetIdentityRoutes(r)

	http.ListenAndServe(":8080", r)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}