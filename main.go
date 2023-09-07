package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

func init() {
	router.Get("/", rootHandler)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Добро пожаловать в яндекс облако!"))
}
