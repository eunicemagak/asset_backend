package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/controllers"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/middlewares"
	// "github.com/dgrijalva/jwt-go/v4"
)

func RegisterRoutes(api fiber.Router) {

	// Auth
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)
	api.Post("/logout", controllers.Logout)

	//Admin
	adminController := controllers.AdminController{}
	admins := api.Group("/admin", middlewares.IsAuthenticated)
	admins.Get("/", adminController.Index)
	admins.Post("/", adminController.CreateAdmin)
	admins.Patch("/:id", adminController.UpdateAdmin)
	admins.Get("/:id", adminController.GetAdmin)
	admins.Delete("/:id", adminController.DeleteAdmin)
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

	//UnAssignedAsset
	unassignedAssetController := controllers.UnAssignedAssetController{}
	unassignedAsset := api.Group("/unassignedAsset", middlewares.IsAuthenticated)
	unassignedAsset.Get("/", unassignedAssetController.Index)

	//Department
	departmentController := controllers.DepartmentController{}
	departments := api.Group("/department", middlewares.IsAuthenticated)
	departments.Get("/", departmentController.Index)
	departments.Post("/", departmentController.CreateDepartment)
	departments.Patch("/:id", departmentController.UpdateDepartment)
	departments.Get("/:id", departmentController.GetDepartment)
	departments.Delete("/:id", departmentController.DeleteDepartment)
	//Accesories
	acccesorieController := controllers.AccesorieController{}
	acccesories := api.Group("/accessories", middlewares.IsAuthenticated)
	acccesories.Get("/", acccesorieController.Index)
	acccesories.Post("/", acccesorieController.CreateAccesorie)
	acccesories.Patch("/:id", acccesorieController.UpdateAccesorie)
	acccesories.Get("/:id", acccesorieController.GetAccesorie)
	acccesories.Delete("/:id", acccesorieController.DeleteAccesorie)
	//Categories
	categoriesController := controllers.CategorieController{}
	categories := api.Group("/categories", middlewares.IsAuthenticated)
	categories.Get("/", categoriesController.Index)
	categories.Post("/", categoriesController.CreateCategorie)
	categories.Patch("/:id", categoriesController.UpdateCategories)
	categories.Get("/:id", categoriesController.GetCategorie)
	categories.Delete("/:id", categoriesController.DeleteCategorie)

	//UnAssignedAccesorie
	unassignedAccesorieController := controllers.UnAssignedAccesorieController{}
	unassignedAccesorie := api.Group("/unassignedAccesorie", middlewares.IsAuthenticated)
	unassignedAccesorie.Get("/", unassignedAccesorieController.Index)
	// un.Post("/", tagController.CreateTag)
	// tags.Patch("/:id", tagController.UpdateTag)
	// tags.Get("/:id", tagController.GetTag)
	// tags.Delete("/:id", tagController.DeleteTagt)

	//Status
	statusController := controllers.StatusController{}
	status := api.Group("/status", middlewares.IsAuthenticated)
	status.Get("/", statusController.Index)
	status.Post("/", statusController.CreateStatus)
	status.Patch("/:id", statusController.UpdateStatus)
	status.Get("/:id", statusController.GetStatus)
	status.Delete("/:id", statusController.DeleteStatus)

	//Image
	imageController := controllers.ImageController{}
	images := api.Group("/images", middlewares.IsAuthenticated)
	images.Get("/", imageController.Index)
	images.Post("/", imageController.Upload)
	images.Delete("/:id", imageController.Delete)

}
