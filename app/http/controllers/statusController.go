package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type StatusController struct {
	DB *sql.DB
}

func (c *StatusController) Index(ctx *fiber.Ctx) error {
	var status []models.Status
	database.DB.Find(&status)

	return ctx.JSON(status)
	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Status{}, page))
}

func (c *StatusController) CreateStatus(ctx *fiber.Ctx) error {
	var status models.Status

	if err := ctx.BodyParser(&status); err != nil {
		return err
	}

	database.DB.Create(&status)

	return ctx.JSON(status)
}

func (c *StatusController) GetStatus(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	status := models.Tag{
		ID: uint(id),
	}

	database.DB.Find(&status)

	return ctx.JSON(status)
}

func (c *StatusController) UpdateStatus(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	status := models.Status{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&status); err != nil {
		return err
	}

	database.DB.Model(&status).Updates(status)

	return ctx.JSON(status)
}

func (c *StatusController) DeleteStatus(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	status := models.Status{
		ID: uint(id),
	}

	database.DB.Delete(&status)

	return nil
}
