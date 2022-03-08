package controllers

import (
	"strconv"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/middlewares"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type createUserReq struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	DepartmentID uint   `json:"department_id"`
	Name         string `json:"name"`
	AssetID      uint   `json:"assetid"`
	AccesorieID  uint   `json:"acccesorie_id"`
}

type UserController struct {
	DB *sql.DB
}

func (c *UserController) Index(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	return ctx.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {

	var userReq createUserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}
	asset := models.Asset{
		ID: userReq.AssetID,
	}
	database.DB.Find(&asset)

	user := models.User{
		ID:           userReq.ID,
		Name:         userReq.Name,
		Email:        userReq.Email,
		DepartmentID: userReq.DepartmentID,
	}

	// user.SetPassword("1234")
	database.DB.Create(&user)

	database.DB.Model(&user).Association("Assets").Append(&asset)

	return ctx.JSON(&user)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Find(&user)

	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Preload("Department").Preload("Asset").Updates(user)

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

	database.DB.Preload("Department").Preload("Asset").Delete(&user)

	return nil
}
