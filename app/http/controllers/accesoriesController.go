package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AccesorieController struct {
	DB *sql.DB
}

func (c *AccesorieController) Index(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB.Preload("Images"), &models.Accesorie{}, page))
}

func (c *AccesorieController) CreateAccesorie(ctx *fiber.Ctx) error {
	var acccesories models.Accesorie

	if err := ctx.BodyParser(&acccesories); err != nil {
		return err
	}
	//select from images where name=%asset.ImageType%
	var image models.Image
	// fmt.Printf(" asset-image type %v", asset.ImageType)
	database.DB.Where("image_type = ?", acccesories.ImageType).First(&image)
	// fmt.Printf(" image %v", image)
	acccesories.ImageType = image.ImageType
	acccesories.ImageID = image.ID

	database.DB.Create(&acccesories)

	return ctx.JSON(acccesories)
}

func (c *AccesorieController) GetAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accesorie{
		ID: uint(id),
	}

	database.DB.Find(&acccesorie)

	return ctx.JSON(acccesorie)
}

func (c *AccesorieController) UpdateAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accesorie{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&acccesorie); err != nil {
		return err
	}

	database.DB.Model(&acccesorie).Updates(acccesorie)

	return ctx.JSON(acccesorie)
}

func (c *AccesorieController) DeleteAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accesorie{
		ID: uint(id),
	}

	database.DB.Delete(&acccesorie)

	return nil
}
