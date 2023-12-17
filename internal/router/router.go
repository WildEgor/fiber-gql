package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/WildEgor/fiber-gql/internal/handlers"
	handlers_gql "github.com/WildEgor/fiber-gql/internal/handlers/gql"
	"github.com/WildEgor/fiber-gql/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
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

	pg := playground.AltairHandler("GraphQL playground", "/graphql")

	app.Get("/graphql", func(ctx *fiber.Ctx) error {

		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			pg(writer, request)
		})(ctx.Context())

		return nil
	})

	app.Post("/altair", func(ctx *fiber.Ctx) error {

		h := handler.NewDefaultServer(server.NewExecutableSchema(server.Config{Resolvers: &handlers_gql.Resolver{}}))

		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			h.ServeHTTP(writer, request)

		})(ctx.Context())

		return nil
	})

	v1 := app.Group("/api/v1")

	// Server endpoint - sanity check that the server is running
	hcController := v1.Group("/health")
	hcController.Get("/check", r.hc.Handle)

	return nil
}
