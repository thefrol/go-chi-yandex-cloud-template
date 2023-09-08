package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

var Router = chi.NewRouter()

func init() {
	Router.Get("/", rootHandler)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Добро пожаловать в яндекс облако!"))
}
