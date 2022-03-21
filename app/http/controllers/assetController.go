package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AssetController struct {
	DB *sql.DB
}

func (c *AssetController) Index(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB, &models.Asset{}, page))
}

func (c *AssetController) CreateAsset(ctx *fiber.Ctx) error {
	var asset models.Asset

	if err := ctx.BodyParser(&asset); err != nil {
		return err
	}

	//select from images where name=%asset.ImageType%
	var image models.Image
	// fmt.Printf(" asset-image type %v", asset.ImageType)
	database.DB.Where("image_type = ?", asset.ImageType).First(&image)
	// fmt.Printf(" image %v", image)
	asset.ImageType = image.ImageType
	asset.ImageID = image.ID
	database.DB.Create(&asset)

	return ctx.JSON(asset)
}

func (c *AssetController) GetAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Find(&asset)

	return ctx.JSON(asset)
}

func (c *AssetController) UpdateAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&asset); err != nil {
		return err
	}

	database.DB.Model(&asset).Updates(asset)

	return ctx.JSON(asset)
}

func (c *AssetController) DeleteAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Delete(&asset)

	return nil
}
