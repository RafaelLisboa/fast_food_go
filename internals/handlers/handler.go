package handlers

import (
	"context"
	"encoding/json"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/services"
	"fmt"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler() *UserHandler{
	service := services.NewUserService();
	
	return &UserHandler{
		userService: service,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.GetBody()

	var userRequestBody models.User

	err := json.NewDecoder(r.Body).Decode(&userRequestBody)

	if err != nil {
		http.Error(w, "error trying decode the request body", http.StatusBadRequest);	
	}

	if userRequestBody.Password == "" || userRequestBody.Username == "" {
		http.Error(w, "User doesn't have all required arguments username | password", http.StatusBadRequest)
	}

	uh.userService.CreateUser(context.Background() ,userRequestBody);
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PathValue("id"))
	fmt.Println("Implement delete user")
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implement Login")
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implement Logout")
}
