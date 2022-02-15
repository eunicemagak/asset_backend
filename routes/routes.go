package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/controllers"
)

func RegisterAssetsRoutes(api fiber.Router) {
	assetController := controllers.AssetController{}
	api.Get("/assets/index", assetController.Index)
}
