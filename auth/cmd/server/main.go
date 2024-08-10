package main

import (
	"fast_food_auth/internals/server/routes"
	"log"
	"net/http"
)

func main() {
	startServer();
}


func startServer() {
	routes := routes.CreateRoutes();
	log.Println("Service Running on port 8080")
	http.ListenAndServe(":8080", routes);

}
