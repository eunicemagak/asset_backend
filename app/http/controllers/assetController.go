package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type AssetController struct {
	DB *sql.DB
}

func (c *AssetController) Index(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"message": "Hello world",
	})
}
