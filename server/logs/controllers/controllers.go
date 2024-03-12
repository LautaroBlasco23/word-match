package controllers

import "net/http"

func GetWords(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola mundo"))
}
