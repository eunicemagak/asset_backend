package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/controllers"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/middlewares"
	// "github.com/dgrijalva/jwt-go/v4"
)

/*
func RegisterAssetsRoutes(api fiber.Router) {
	assetController := controllers.AssetController{}
	api.Get("/assets/index", assetController.Index)
}
*/

func RegisterRoutes(api fiber.Router) {

	//Auth
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)

	//Users
	userController := controllers.UserController{}
	users := api.Group("/users", middlewares.IsAuthenticated)
	users.Get("/", userController.Index)
	users.Post("/", userController.CreateUser)
	users.Patch("/:id", userController.UpdateUser)
	users.Get("/:id", userController.GetUser)
	users.Delete("/:id", userController.DeleteUser)

	//Assets
	assetController := controllers.AssetController{}
	assets := api.Group("/assets", middlewares.IsAuthenticated)
	assets.Get("/", assetController.Index)
	assets.Post("/", assetController.CreateAsset)
	assets.Patch("/:id", assetController.UpdateAsset)
	assets.Get("/:id", assetController.GetAsset)
	assets.Delete("/:id", assetController.DeleteAsset)

	//Complaints
	//complaintController := controllers.ComplaintController{}
	// api.Get("/complaints", complaintController.Index)
	// api.Post("/complaints", complaintController.CreateComplaint)
	// api.Patch("/complaints/:id", complaintController.UpdateComplaint)
	// api.Get("/complaints/:id", complaintController.GetComplaint)
	// api.Delete("/complaints/:id", complaintController.DeleteComplaint)
}
