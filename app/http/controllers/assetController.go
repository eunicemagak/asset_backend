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
type UpdateAssetReq struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	SerialNumber string `json:"serialnumber"`
	Description  string `json:"description"`
	PurchaseDate string `json:"purchase_date"`
	ImageType    string `json:"image_type"`
	Price        string `json:"price"`
	StatusID     uint   `json:"status_id"`
	CategorieID  uint   `json:"categorie_id"`
	// ImageID      uint   `json:"image_id"`
}

func (c *AssetController) Index(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB, &models.Asset{}, page))

}

func (c *AssetController) CreateAsset(ctx *fiber.Ctx) error {

	var assetReq UpdateAssetReq
	status := models.Status{
		ID: assetReq.StatusID,
	}
	database.DB.Find(&status)
	categorie := models.Categorie{
		ID: assetReq.CategorieID,
	}
	database.DB.Find(&categorie)

	if err := ctx.BodyParser(&assetReq); err != nil {
		return err
	}

	//select from images where name=%asset.ImageType%
	var image models.Image
	// fmt.Printf(" asset-image type %v", asset.ImageType)
	database.DB.Where("image_type = ?", assetReq.ImageType).First(&image)
	// fmt.Printf(" image %v", image)
	assetReq.ImageType = image.ImageType

	asset := models.Asset{
		ID:           assetReq.ID,
		Title:        assetReq.Title,
		Description:  assetReq.Description,
		SerialNumber: assetReq.SerialNumber,
		PurchaseDate: assetReq.PurchaseDate,
		Price:        assetReq.Price,
		ImageType:    assetReq.ImageType,
	}
	asset.ImageID = image.ID

	database.DB.Create(&asset)
	database.DB.Model(&asset).Association("Categories").Append(&categorie)
	database.DB.Model(&asset).Association("Statuses").Append(&status)

	return ctx.JSON(asset)
}

func (c *AssetController) GetAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Preload("Statuses").Preload("Categories").Find(&asset)

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
