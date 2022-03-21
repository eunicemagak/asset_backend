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
	admins := api.Group("/admin")
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
	// assets.Post("/", assetController.UploadImage)
	assets.Post("/", assetController.CreateAsset)

	//Images
	imageController := controllers.ImageController{}
	images := api.Group("/images", middlewares.IsAuthenticated)
	images.Get("/", imageController.Index)
	images.Post("/", imageController.Upload)

	//Complaints
	//complaintController := controllers.ComplaintController{}
	// api.Get("/complaints", complaintController.Index)
	// api.Post("/complaints", complaintController.CreateComplaint)
	// api.Patch("/complaints/:id", complaintController.UpdateComplaint)
	// api.Get("/complaints/:id", complaintController.GetComplaint)
	// api.Delete("/complaints/:id", complaintController.DeleteComplaint)

	//Department
	departmentController := controllers.DepartmentController{}
	departments := api.Group("/department")
	departments.Get("/", departmentController.Index)
	departments.Post("/", departmentController.CreateDepartment)
	departments.Patch("/:id", departmentController.UpdateDepartment)
	departments.Get("/:id", departmentController.GetDepartment)
	departments.Delete("/:id", departmentController.DeleteDepartment)

	//Accesories
	// acccesorieController := controllers.AccessoryController{}
	// acccesories := api.Group("/accessories")
	// acccesories.Get("/", acccesorieController.Index)
	// acccesories.Post("/", acccesorieController.CreateAccesorie)
	// acccesories.Patch("/:id", acccesorieController.UpdateAccesorie)
	// acccesories.Get("/:id", acccesorieController.GetAccesorie)
	// acccesories.Delete("/:id", acccesorieController.DeleteAccesorie)

	// //Gmail
	// complainController := controllers.getClient{}
	// complains := api.Group("/complains")

}
