package controllers

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
)

type DepartmentController struct {
	DB *sql.DB
}

func (c *DepartmentController) Index(ctx *fiber.Ctx) error {
	db := database.DB
	var departments []models.Department

	// find all notes in the database
	db.Find(&departments)

	return ctx.JSON(&departments)
}

func (c *DepartmentController) CreateDepartment(ctx *fiber.Ctx) error {
	var department models.Department

	if err := ctx.BodyParser(&department); err != nil {
		return err
	}

	database.DB.Create(&department)

	return ctx.JSON(department)
}

func (c *DepartmentController) GetDepartment(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	department := models.Department{
		ID: uint(id),
	}

	database.DB.Find(&department)

	return ctx.JSON(department)
}

func (c *DepartmentController) UpdateDepartment(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	department := models.Department{
		ID: uint(id),
	}

	if err := ctx.BodyParser(&department); err != nil {
		return err
	}

	database.DB.Model(&department).Updates(department)

	return ctx.JSON(department)
}

func (c *DepartmentController) DeleteDepartment(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	department := models.Department{
		ID: uint(id),
	}

	database.DB.Delete(&department)

	return nil
}
