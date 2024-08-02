package routes

import (
	"fast_food_auth/internals/server/handlers"
	"net/http"
)


func CreateRoutes() *http.ServeMux {

	serverMux := http.NewServeMux();

	userHandler := handlers.NewUserHandler();

	serverMux.HandleFunc("POST /users", userHandler.CreateUser)
	serverMux.HandleFunc("POST /login", userHandler.Login)
	serverMux.HandleFunc("POST /refresh-token", userHandler.RefreshToken)

	return serverMux;
}
