package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a NewUserRequest struct
	var newUser NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Username: %s\nPassword: %s", newUser.Username, newUser.Password)
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