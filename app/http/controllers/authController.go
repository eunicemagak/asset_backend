package controllers

import (
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"

	"strconv"
	"time"

	"gitlab.ci.emalify.com/roamtech/asset_be/util"

	"github.com/gofiber/fiber/v2"
)

////REGISTER
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//registration validations
	if data["first_name"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Firstname is required!",
		})
	}
	if data["last_name"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Lastname is required!",
		})
	}
	if data["email"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Email is required!",
		})
	}
	if data["password"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password is required!",
		})
	}
	if len(data["password"]) <= 6 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password should be more then 6 chars",
		})
	}
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(data["password"])

	database.DB.Create(&user)

	return c.JSON(user)
}

////LOGIN
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//Login validations
	if data["email"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Email is required!",
		})
	}
	if data["password"] == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password is required!",
		})
	}
	if len(data["password"]) <= 6 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password should be more then 6 chars",
		})
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Invalid Email or Password",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Get("authorization")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Get("authorization")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		ID:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match",
		})
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		ID: uint(userId),
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}
