package main

import (
	"TimeMachineApi/service"
	json2 "encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func handleRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Hello! Root Path serves no content"))
		if err != nil {
			return
		}
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("API is up and ready"))
		if err != nil {
			return
		}
	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {

		products := service.GetProducts()
		json, _ := json2.Marshal(products)

		_, err := w.Write(json)
		if err != nil {
			return
		}
	})
}

func serveHttp(router *chi.Mux) {

	err := http.ListenAndServe(":3100", router)
	if err != nil {
		return
	}
}

func main() {

	fmt.Println("BEGIN TimeMachineApi")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	handleRoutes(r)
	serveHttp(r)

	fmt.Println("END TimeMachineApi")
}
