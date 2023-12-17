package handlers

import (
	"github.com/WildEgor/fiber-gql/internal/config"
	domains "github.com/WildEgor/fiber-gql/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	appConfig *config.AppConfig
}

func NewHealthCheckHandler(
	appConfig *config.AppConfig,
) *HealthCheckHandler {
	return &HealthCheckHandler{
		appConfig,
	}
}

func (hch *HealthCheckHandler) Handle(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"isOk": true,
		"data": &domains.StatusDomain{
			Status:      "ok",
			Version:     hch.appConfig.Version,
			Environment: hch.appConfig.GoEnv,
		},
	})
	return nil
}
