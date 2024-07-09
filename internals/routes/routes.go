package routes

import (
	"fast_food_auth/internals/handlers"
	"net/http"
)


func CreateRoutes() *http.ServeMux {

	serverMux := http.NewServeMux();

	userHandler := handlers.NewUserHandler();

	serverMux.HandleFunc("POST /users", userHandler.CreateUser)
	serverMux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
	serverMux.HandleFunc("POST /login", userHandler.Login)

	return serverMux;
}
