package controllers

// import (
// 	"database/sql"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
// 	"gitlab.ci.emalify.com/roamtech/asset_be/database"
// )

// type AccessoryController struct {
// 	DB *sql.DB
// }

// func (c *AccessoryController) Index(ctx *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(ctx.Query("page", "1"))

// 	return ctx.JSON(models.Paginate(database.DB, &models.Accessory{}, page))
// }

// func (c *AccessoryController) CreateAccesorie(ctx *fiber.Ctx) error {
// 	var acccesories models.Accessory

// 	if err := ctx.BodyParser(&acccesories); err != nil {
// 		return err
// 	}

// 	database.DB.Create(&acccesories)

// 	return ctx.JSON(acccesories)
// }

// func (c *AccessoryController) GetAccesorie(ctx *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(ctx.Params("id"))

// 	acccesorie := models.Accessory{
// 		ID: uint(id),
// 	}

// 	database.DB.Find(&acccesorie)

// 	return ctx.JSON(acccesorie)
// }

// func (c *AccessoryController) UpdateAccesorie(ctx *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(ctx.Params("id"))

// 	acccesorie := models.Accessory{
// 		ID: uint(id),
// 	}

// 	if err := ctx.BodyParser(&acccesorie); err != nil {
// 		return err
// 	}

// 	database.DB.Model(&acccesorie).Updates(acccesorie)

// 	return ctx.JSON(acccesorie)
// }

// func (c *AccessoryController) DeleteAccesorie(ctx *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(ctx.Params("id"))

// 	acccesorie := models.Accessory{
// 		ID: uint(id),
// 	}

// 	database.DB.Delete(&acccesorie)

// 	return nil
// }
