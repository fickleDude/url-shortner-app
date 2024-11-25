package router

import (
	"net/http"

	controller "example.com/url-shortner/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		// AllowCredentials: true,
		// MaxAge:           300,
	}))

	router.Get("/airports", controller.GetAllAirports)
	router.Get("/airports/airport/{airport_code}", controller.GetAirportByAirportCode)
	router.Post("/airports/airport", controller.CreateAirports)
	router.Put("/airports/airport/{airport_code}", controller.UpdateAirports)
	router.Delete("/airports/departure/{airport_code}", controller.DeleteAirports)

	router.Get("/departure/{airport_code}", controller.GetDepartureFlights)
	router.Get("/arrive/{airport_code}", controller.GetArrivalFlights)

	router.Get("/aircrafts", controller.GetAllAircrafts)

	return router
}
