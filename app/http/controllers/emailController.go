package controllers

// import (
// 	"database/sql"
// 	"strconv"

// 	"github.com/gofiber/fiber"
// 	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
// 	"gitlab.ci.emalify.com/roamtech/asset_be/database"
// )

// type EmailController struct {
// 	DB *sql.DB
// }

// func (c *EmailController) Index(ctx *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(ctx.Query("page 1"))
// 	return ctx.JSON(models.Paginate(database.DB, &models.Email{}, page))

// }
// func (c *EmailController) CreateEmail(ctx *fiber.Ctx) error {
// 	var emails models.Email

// 	if err := ctx.BodyParser(&emails); err != nil {
// 		return err
// 	}

// 	database.DB.Create(&emails)

// 	return ctx.JSON(emails)
// }
// func (c *EmailController) GetEmails(ctx *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(ctx.Params("id"))

// 	email := models.Email{
// 		ID: uint(id),
// 	}

// 	database.DB.Find(&email)

// 	return ctx.JSON(email)
// }

// func (c *EmailController) UpdateEmail(ctx *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(ctx.Params("id"))

// 	gmail := models.Email{
// 		ID: uint(id),
// 	}

// 	if err := ctx.BodyParser(&gmail); err != nil {
// 		return err
// 	}

// 	database.DB.Model(&gmail).Updates(gmail)

// 	return ctx.JSON(gmail)
// }

// func (c *EmailController) DeleteEmail(ctx *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(ctx.Params("id"))

// 	email := models.Ema{
// 		ID: uint(id),
// 	}

// 	database.DB.Delete(&email)

// 	return nil
// }
