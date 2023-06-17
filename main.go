package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/identity", identityHandler)
	r.HandleFunc("/album", albumHandler)
	r.HandleFunc("/post", postHandler)

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":8080", handler)
}

func identityHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("identityHandler")
}

func albumHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("albumHandler")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postHandler")
}