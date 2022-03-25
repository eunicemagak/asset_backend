package controllers

import (
	"fmt"
	"strconv"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type createUserReq struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`

	Name    string `json:"name"`
	AssetID uint   `json:"asset_id"`
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
	fmt.Printf(" Userreq %v userreq \n", userReq)

	asset := models.Asset{
		ID: userReq.AssetID,
	}
	database.DB.Find(&asset)
	fmt.Printf("log asset %v", asset)

	user := models.User{
		ID:    userReq.ID,
		Name:  userReq.Name,
		Email: userReq.Email,
	}

	// Update with conditions and model value

	asset.IsAssigned = true
	result := database.DB.Model(&asset).Where("id = ?", userReq.AssetID).Update("is_assigned", true)

	if result.Error != nil {
		fmt.Printf("Error in updating %v", result.Error)
		//return result.Error
	}
	fmt.Printf(" asset after update %v", asset)
	database.DB.Create(&user)

	database.DB.Model(&user).Association("Assets").Append(&asset)

	return ctx.JSON(&user)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Preload("Assets").Preload("Department").Preload(" Accesorie").Preload("Tag").Preload("").Find(&user)
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
	// if err := middlewares.IsAuthenticated(ctx); err != nil {
	// 	return err
	// }
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Preload("Department").Preload("Asset").Preload("Accesorie").Delete(&user)

	return nil
}
