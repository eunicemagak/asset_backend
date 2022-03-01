package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.ci.emalify.com/roamtech/asset_be/app/http/controllers"
	// "github.com/dgrijalva/jwt-go/v4"
)

func RegisterRoutes(api fiber.Router) {

	//Auth
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)
	api.Post("/logout", controllers.Logout)

	//Users
	userController := controllers.UserController{}
	users := api.Group("/users")
	users.Get("/", userController.Index)
	users.Post("/", userController.CreateUser)
	users.Patch("/:id", userController.UpdateUser)
	users.Get("/:id", userController.GetUser)
	users.Delete("/:id", userController.DeleteUser)

	//Assets
	assetController := controllers.AssetController{}
	assets := api.Group("/assets")
	assets.Get("/", assetController.Index)
	assets.Post("/", assetController.CreateAsset)
	assets.Patch("/:id", assetController.UpdateAsset)
	assets.Get("/:id", assetController.GetAsset)
	assets.Delete("/:id", assetController.DeleteAsset)

	//Test
	// testController := controllers.TestController{}
	// tests := api.Group("/tests")
	// tests.Get("/", testController.Index)
	// tests.Post("/", testController.CreateTest)
	// tests.Patch("/:id", testController.UpdateTest)
	// tests.Get("/:id", testController.GetTest)
	// tests.Delete("/:id", testController.DeleteTest)

	//Department
	departmentController := controllers.DepartmentController{}
	departments := api.Group("/department")
	departments.Get("/", departmentController.Index)
	departments.Post("/", departmentController.CreateDepartment)
	departments.Patch("/:id", departmentController.UpdateDepartment)
	departments.Get("/:id", departmentController.GetDepartment)
	departments.Delete("/:id", departmentController.DeleteDepartment)
	//Accesories
	acccesorieController := controllers.AccesorieController{}
	acccesories := api.Group("/accesories")
	acccesories.Get("/", acccesorieController.Index)
	acccesories.Post("/", acccesorieController.CreateAccesorie)
	acccesories.Patch("/:id", acccesorieController.UpdateAccesorie)
	acccesories.Get("/:id", acccesorieController.GetAccesorie)
	acccesories.Delete("/:id", acccesorieController.DeleteAccesorie)

	//Assign Asset
	assignController := controllers.AssignController{}
	assign := api.Group("/assets")
	assign.Get("/", assignController.Index)
	assign.Post("/", assignController.AssignAsset)
	assign.Patch("/:id", assignController.UpdateAssign)
	assign.Get("/:id", assignController.GetAssign)
	assign.Delete("/:id", assetController.DeleteAsset)

}
