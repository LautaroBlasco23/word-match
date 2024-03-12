package handlers

import "net/http"

func getWords(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola mundo"))
}
