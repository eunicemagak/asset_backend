package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
	"gorm.io/gorm"
)

type AssignController struct {
	DB gorm.DB
}

func (c *AssignController) Index(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB, &models.Assign{}, page))
}

func (c AssignController) AssignAsset(ctx *fiber.Ctx) error {

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	var asset models.Asset

	if err := ctx.BodyParser(&asset); err != nil {
		return err
	}

	var assign = models.Assign{
		Asset: asset,
		User:  user,
	}
	if err := ctx.BodyParser(&assign); err != nil {
		return err
	}

	database.DB.Create(&assign)

	return ctx.JSON(assign)
}

func (c *AssignController) GetAssign(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var asset models.Asset
	var user models.User

	assign := models.Assign{
		ID:    uint(id),
		Asset: asset,
		User:  user,
	}

	database.DB.Find(&assign)

	return ctx.JSON(assign)
}
func (c *AssignController) UpdateAssign(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	assign := models.Assign{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&assign); err != nil {
		return err
	}

	database.DB.Model(&assign).Updates(assign)

	return ctx.JSON(assign)
}

func (c *AssignController) DeleteAssign(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	assign := models.Assign{
		ID: uint(id),
	}

	database.DB.Delete(&assign)

	return nil
}
