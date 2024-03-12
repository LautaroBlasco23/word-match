package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ServeRoutes() {
	router := mux.NewRouter()

	r.HandleFunc("/", controllers.GetWords).Methods(http.Get)
}
