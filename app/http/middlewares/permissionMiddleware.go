package middlewares

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/models"
	"gitlab.ci.emalify.com/roamtech/asset_be/database"
	"gitlab.ci.emalify.com/roamtech/asset_be/util"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Get("authorization")
	token := strings.Replace(cookie, "Bearer ", "", 1)
	Id, err := util.ParseJwt(token)

	fmt.Println(Id)

	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(Id)

	user := models.User{
		ID: uint(userId),
	}

	database.DB.Preload("Role").Find(&user)

	role := models.Role{
		ID: user.RoleId,
	}

	database.DB.Preload("Permissions").Find(&role)

	fmt.Println(role.Permissions)
	if c.Method() == "GET" {
		for _, permission := range role.Permissions {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permissions {
			fmt.Println(permission.Name)
			if permission.Name == "edit_"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
