package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gustasousaghs/carros-pcd-service/controllers"
)

func InitRouter() *chi.Mux {
	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Post("/admin/api/cars/create-car", controllers.CreateCarController)
	app.Put("/admin/api/cars/{id}", controllers.UpdateCarController)
	app.Delete("/admin/api/cars/{id}", controllers.DeleteCarController)
	app.Get("/api/cars", controllers.GetAllCars)
	app.Get("/api/cars/{id}", controllers.GetCarById)
	return app
}
