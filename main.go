package main

import (
	"TimeMachineApi/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func handleRoutes(router *chi.Mux) {

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Hello! Root Path serves no content"))
		if err != nil {
			log.Error("Error trying to query /")
			return
		}
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("API is up and ready"))
		if err != nil {
			log.Error("Error trying to query /health")
			return
		}
	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {

		products := service.GetProducts()

		render.Status(r, http.StatusOK)
		render.JSON(w, r, products)
	})
}

func serveHttp(router *chi.Mux) {

	err := http.ListenAndServe(":3100", router)
	if err != nil {
		return
	}
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.Info("BEGIN TimeMachineApi")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	handleRoutes(r)
	serveHttp(r)

	log.Info("END TimeMachineApi")
}
