package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type UnAssignedAccesorieController struct {
	DB *sql.DB
}

func (c *UnAssignedAccesorieController) Index(ctx *fiber.Ctx) error {

	var acccesorie []models.Accesorie
	// Get first matched record
	database.DB.Where("is_assigned = ?", false).Preload("Images").Find(&acccesorie)

	return ctx.JSON(&acccesorie)
}
