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
type CreateAssetReq struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	SerialNumber string `json:"serialnumber"`
	Description  string `json:"description"`
	PurchaseDate string `json:"purchase_date"`
	ImageType    string `json:"image_type"`
	Price        string `json:"price"`
	CategorieID  uint   `json:"categorie_id"`
	// ImageID      uint   `json:"image_id"`
}

func (c *AssetController) Index(ctx *fiber.Ctx) error {

	var assets []models.Asset
	database.DB.Preload("Categories").Find(&assets)

	return ctx.JSON(assets)
	// page, _ := strconv.Atoi(ctx.Query("page", "1"))

	// return ctx.JSON(models.Paginate(database.DB, &models.Asset{}, page))

}

func (c *AssetController) CreateAsset(ctx *fiber.Ctx) error {

	var assetReq CreateAssetReq

	if err := ctx.BodyParser(&assetReq); err != nil {
		return err
	}
	categorie := models.Categorie{
		ID: assetReq.CategorieID,
	}
	database.DB.Find(&categorie)

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

	return ctx.JSON(asset)
}

func (c *AssetController) GetAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Preload("Categories").Find(&asset)

	return ctx.JSON(asset)
}

func (c *AssetController) UpdateAsset(ctx *fiber.Ctx) error {
	type Update struct {
		Title        string `json:"title"`
		SerialNumber string `json:"serialnumber"`
		Description  string `json:"description"`
		Price        string `json:"price"`
		PurchaseDate string `json:"purchase_date"`
		AssignedTo   string `json:"assigned_to"`
		IsAssigned   bool   `json:"is_assigned" gorm:"default:false"`
		IsClearedOf  bool   `json:"is_cleared_of" gorm:"default:false"`
		IsDamaged    bool   `json:"is_damaged" gorm:"default:false"`
	}
	db := database.DB

	var asset models.Asset
	// Read the param noteId
	id := ctx.Params("assetId")

	// Find the note with the given Id
	db.Find(&asset, "id = ?", id)

	// Store the body containing the updated data and return error if encountered
	var updateAsset Update
	err := ctx.BodyParser(&updateAsset)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Edit the note
	asset.Title = updateAsset.Title
	asset.SerialNumber = updateAsset.SerialNumber
	asset.Description = updateAsset.Description
	asset.IsAssigned = updateAsset.IsAssigned
	asset.IsClearedOf = updateAsset.IsClearedOf
	asset.IsDamaged = updateAsset.IsDamaged
	asset.Price = updateAsset.Price
	asset.PurchaseDate = updateAsset.PurchaseDate
	if asset.IsClearedOf {
		asset.IsAssigned = false

	}

	// Save the Changes
	db.Save(&asset)

	// Return the updated note
	return ctx.JSON(fiber.Map{"status": "success", "message": "Asset Found", "data": asset})

}

func (c *AssetController) DeleteAsset(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	asset := models.Asset{
		ID: uint(id),
	}

	database.DB.Delete(&asset)

	return nil
}
