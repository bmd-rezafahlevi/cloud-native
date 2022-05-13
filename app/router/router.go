package router

import (
	"github.com/bmd-rezafahlevi/cloud-native/app/app"
	"github.com/bmd-rezafahlevi/cloud-native/app/requestlog"
	"github.com/bmd-rezafahlevi/cloud-native/app/router/middleware"
	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))
	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		// Routes for books
		r.Method("GET", "/books", requestlog.NewHandler(a.HandleListBook, l))
		r.Method("POST", "/books", requestlog.NewHandler(a.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", requestlog.NewHandler(a.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", requestlog.NewHandler(a.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", requestlog.NewHandler(a.HandleDeleteBook, l))
	})

	return r
}
