package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type TagController struct {
	DB *sql.DB
}

func (c *TagController) Index(ctx *fiber.Ctx) error {

	db := database.DB
	var tags []models.Tag

	// find all notes in the database
	db.Find(&tags)

	return ctx.JSON(&tags)
}

func (c *TagController) CreateTag(ctx *fiber.Ctx) error {
	var tag models.Tag

	if err := ctx.BodyParser(&tag); err != nil {
		return err
	}

	database.DB.Create(&tag)

	return ctx.JSON(tag)
}

func (c *TagController) GetTag(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	tag := models.Tag{
		ID: uint(id),
	}

	database.DB.Find(&tag)

	return ctx.JSON(tag)
}

func (c *TagController) UpdateTag(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	tag := models.Tag{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&tag); err != nil {
		return err
	}

	database.DB.Model(&tag).Updates(tag)

	return ctx.JSON(tag)
}

func (c *TagController) DeleteTagt(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	tag := models.Tag{
		ID: uint(id),
	}

	database.DB.Delete(&tag)

	return nil
}
