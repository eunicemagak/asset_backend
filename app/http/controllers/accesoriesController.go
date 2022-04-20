package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type AccesorieCreateReq struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	SerialNumber string `json:"serialnumber"`
	Description  string `json:"description"`
	PurchaseDate string `json:"purchase_date"`
	ImageType    string `json:"image_type"`
	Price        string `json:"price"`
	CategorieID  uint   `json:"categorie_id"`
}

type AccesorieController struct {
	DB *sql.DB
}

// 0722157344

func (c *AccesorieController) Index(ctx *fiber.Ctx) error {

	var acccesories []models.Accesorie
	database.DB.Order("id asc").Preload("Categories").Find(&acccesories)

	return ctx.JSON(acccesories)
}

func (c *AccesorieController) CreateAccesorie(ctx *fiber.Ctx) error {
	var accesorieReq AccesorieCreateReq

	if err := ctx.BodyParser(&accesorieReq); err != nil {
		return err
	}

	//categorie
	categorie := models.Categorie{
		ID: accesorieReq.CategorieID,
	}
	database.DB.Find(&categorie)

	//select from images where name=%accessorie.ImageType%
	var image models.Image
	// fmt.Printf(" accessorie-image type %v", accessorie.ImageType)
	database.DB.Where("image_type = ?", accesorieReq.ImageType).First(&image)
	// fmt.Printf(" image %v", image)
	accesorieReq.ImageType = image.ImageType
	// acccesorie.ImageID = image.ID

	acccesorie := models.Accesorie{
		// ID:           accesorieReq.ID,
		Title:        accesorieReq.Title,
		Description:  accesorieReq.Description,
		SerialNumber: accesorieReq.SerialNumber,
		PurchaseDate: accesorieReq.PurchaseDate,
		Price:        accesorieReq.Price,
		ImageType:    accesorieReq.ImageType,
	}

	acccesorie.ImageID = image.ID
	database.DB.Create(&acccesorie)

	database.DB.Model(&acccesorie).Association("Categories").Append(&categorie)

	return ctx.JSON(acccesorie)
}

func (c *AccesorieController) GetAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accesorie{
		ID: uint(id),
	}

	database.DB.Preload("Categories").Find(&acccesorie)

	return ctx.JSON(acccesorie)
}

func (c *AccesorieController) UpdateAccesorie(ctx *fiber.Ctx) error {
	type Update struct {
		Title        string `json:"title"`
		SerialNumber string `json:"serialnumber"`
		Description  string `json:"description"`
		IsAssigned   bool   `json:"is_assigned" gorm:"default:false"`
		IsClearedOf  bool   `json:"is_cleared_of" gorm:"default:false"`
		IsDamaged    bool   `json:"is_damaged" gorm:"default:false"`
		PurchaseDate string `json:"purchase_date"`
		AssignedTo   string `json:"assigned_to"`

		Price string `json:"price"`
	}
	db := database.DB

	var accesorie models.Accesorie
	// Read the param accessorieId
	id := ctx.Params("accessorieId")

	// Find the note with the given Id
	db.Find(&accesorie, "id = ?", id)

	// Store the body containing the updated data and return error if encountered
	var updateAccesorie Update
	err := ctx.BodyParser(&updateAccesorie)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Edit the accesorie

	accesorie.Title = updateAccesorie.Title
	accesorie.SerialNumber = updateAccesorie.SerialNumber
	accesorie.Description = updateAccesorie.Description
	accesorie.IsAssigned = updateAccesorie.IsAssigned
	accesorie.IsClearedOf = updateAccesorie.IsClearedOf
	accesorie.IsDamaged = updateAccesorie.IsDamaged
	accesorie.Price = updateAccesorie.Price
	accesorie.PurchaseDate = updateAccesorie.PurchaseDate
	if accesorie.IsClearedOf {
		accesorie.IsAssigned = false

	}

	// Save the Changes
	db.Save(&accesorie)

	// Return the updated asset
	return ctx.JSON(fiber.Map{"status": "success", "message": "accessorie Found", "data": accesorie})

}

func (c *AccesorieController) DeleteAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accesorie{
		ID: uint(id),
	}

	database.DB.Delete(&acccesorie)

	return nil
}
