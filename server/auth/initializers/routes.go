package initializers

import (
	"match-words/controllers"

	"github.com/gorilla/mux"
)

func ServeRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/checkToken", controllers.VerifyToken).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	return router
}
