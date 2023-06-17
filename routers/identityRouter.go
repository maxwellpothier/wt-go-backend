package routers

import (
	"github.com/gorilla/mux"

	"github.com/maxwellpothier/wt-go-backend/handlers"
)

func SetIdentityRoutes(router *mux.Router) {
	router.HandleFunc("/identity/create", handlers.CreateNewUserHandler).Methods("POST")
	router.HandleFunc("/identity/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/identity", handlers.GetCurrentUserInfoHandler).Methods("GET")
	router.HandleFunc("/identity/forgot-password", handlers.SendForgotPasswordEmailHandler).Methods("POST")
	router.HandleFunc("/identity/reset-password", handlers.ChangePasswordHandler).Methods("POST")
	router.HandleFunc("/identity/edit", handlers.EditCurrentUserHandler).Methods("POST")
}