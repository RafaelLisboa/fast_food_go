package main

import (
	"fast_food_auth/internals/routes"
	"net/http"
)

func main() {
	startServer();
}


func startServer() {
	routes := routes.CreateRoutes();

	http.ListenAndServe(":8080", routes);
}
