package router

import (
	"github.com/WildEgor/fibergo-microservice-boilerplate/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	hc *handlers.HealthCheckHandler
}

func NewRouter(hc *handlers.HealthCheckHandler) *Router {
	return &Router{
		hc: hc,
	}
}

func (r *Router) Setup(app *fiber.App) error {
	v1 := app.Group("/api/v1")

	// Server endpoint - sanity check that the server is running
	hcController := v1.Group("/health")
	hcController.Get("/check", r.hc.Handle)

	return nil
}
