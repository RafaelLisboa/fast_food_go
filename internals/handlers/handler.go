package handlers

import (
	"context"
	"encoding/json"
	"fast_food_auth/internals/models"
	"fast_food_auth/internals/services"
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
		return
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
		return	
	}

	token, err := uh.loginService.Login(ctx, loginRequest)

	if err != nil {
		return
	}

	tokenBytes, _ := json.Marshal(token)

	w.Write(tokenBytes)
}
