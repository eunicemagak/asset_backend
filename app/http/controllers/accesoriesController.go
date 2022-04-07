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
	CategorieID  uint   `json:"categorieid"`
	StatusID     uint   `json:"status"`
}

type AccesorieController struct {
	DB *sql.DB
}

func (c *AccesorieController) Index(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB, &models.Accesorie{}, page))
}

func (c *AccesorieController) CreateAccesorie(ctx *fiber.Ctx) error {
	var accesorieReq AccesorieCreateReq
	status := models.Status{
		ID: accesorieReq.StatusID,
	}
	database.DB.Find(&status)
	categorie := models.Categorie{
		ID: accesorieReq.CategorieID,
	}
	database.DB.Find(&categorie)

	if err := ctx.BodyParser(&accesorieReq); err != nil {
		return err
	}

	//select from images where name=%asset.ImageType%
	var image models.Image
	// fmt.Printf(" asset-image type %v", asset.ImageType)
	database.DB.Where("image_type = ?", accesorieReq.ImageType).First(&image)
	// fmt.Printf(" image %v", image)
	accesorieReq.ImageType = image.ImageType
	// acccesorie.ImageID = image.ID

	acccesorie := models.Accesorie{
		ID:           accesorieReq.ID,
		Title:        accesorieReq.Title,
		Description:  accesorieReq.Description,
		SerialNumber: accesorieReq.SerialNumber,
		PurchaseDate: accesorieReq.PurchaseDate,
		Price:        accesorieReq.Price,
		ImageType:    accesorieReq.ImageType,
	}

	database.DB.Create(&acccesorie)
	database.DB.Model(&acccesorie).Association("Statuses").Append(&status)
	database.DB.Model(&acccesorie).Association("Categories").Append(&categorie)

	return ctx.JSON(&acccesorie)
}

func (c *AccesorieController) GetAccesorie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	acccesorie := models.Accesorie{
		ID: uint(id),
	}

	database.DB.Preload("Categories").Preload("Statuses").Find(&acccesorie)

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
