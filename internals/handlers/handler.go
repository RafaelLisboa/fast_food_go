package handlers

import (
	"context"
	"encoding/json"
	"fast_food_auth/internals/exceptions"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/services"
	"fmt"
	"net/http"
)

type UserHandler struct {
	userService  services.UserService
	loginService services.LoginService
}

func NewUserHandler() *UserHandler {
	userService := services.NewUserService()
	loginService := services.NewLoginService()

	return &UserHandler{
		loginService: loginService,
		userService:  userService,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var userRequestBody models.User

	err := json.NewDecoder(r.Body).Decode(&userRequestBody)

	if err != nil {
		http.Error(w, "error trying decode the request body", http.StatusBadRequest)
	}

	ctx := context.WithValue(context.Background(), 0, w)

	err = uh.userService.CreateUser(ctx, userRequestBody)

	if err != nil {
		return
	}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(context.Background(), 0, w)

	var loginRequest models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		exceptions.NewError(ctx, exceptions.INTERNAL_ERROR)
	}

	token, err := uh.loginService.Login(ctx, loginRequest)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "could not login you", http.StatusInternalServerError)
	}
	tokenBytes, _ := json.Marshal(token)

	w.Write(tokenBytes)

}
