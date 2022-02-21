package controllers

import (
	"strconv"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/middlewares"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type UserController struct {
	DB *sql.DB
}

func (c *UserController) Index(ctx *fiber.Ctx) error {
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }

	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	// if err := middlewares.IsAuthorized(ctx, "users"); err != nil {
	// 	return err
	// }
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }

	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return ctx.JSON(user)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	if err := middlewares.IsAuthenticated(ctx); err != nil {
		return err
	}
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
