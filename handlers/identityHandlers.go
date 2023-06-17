package handlers

import (
	"fmt"
	"net/http"
)

func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	message := "Create User: /identity/create"
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(message))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login: /identity/login")
}

func GetCurrentUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Current User Info: /identity")
}

func SendForgotPasswordEmailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Send Forgot Password Email: /identity/forgot-password")
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Change Password: /identity/reset-password")
}

func EditCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Edit Current User: /identity/edit")
}