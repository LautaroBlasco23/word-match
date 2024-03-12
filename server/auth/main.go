package main

import (
	"match-words/initializers"
	"net/http"
)

func main() {
	router := initializers.ServeRoutes()

	http.ListenAndServe(":8080", router)
}
