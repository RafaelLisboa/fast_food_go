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
	loginService services.LoginService
}

func NewUserHandler() *UserHandler{
	userService := services.NewUserService();
	loginService := services.NewLoginService();

	return &UserHandler{
		loginService: loginService,
		userService: userService,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

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


func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	
	var userRequestBody models.User

	err := json.NewDecoder(r.Body).Decode(&userRequestBody)

	if err != nil {
		http.Error(w, "error trying decode the request body", http.StatusBadRequest);	
	}

	if userRequestBody.Password == "" || userRequestBody.Email == "" {
		http.Error(w, "user doesn't have all required arguments email | password", http.StatusBadRequest)
	}

	token, err := uh.loginService.Login(userRequestBody)

	if err != nil {
		http.Error(w, "could not login you", http.StatusInternalServerError);
	}

	w.Write([]byte(token));
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Implement Logout")
}
