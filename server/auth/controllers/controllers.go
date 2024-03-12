package controllers

import "net/http"

func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func Register(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola mundo"))
}

func VerifyToken(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hola mundo"))
}
