package routes

import (
	"app/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// user
	e.GET("/users", controllers.GetUsersControllers)
	e.POST("/users", controllers.CreateUsersControllers)
	e.GET("/users/:userid", controllers.GetUserByID)
	e.PUT("/users/:userid", controllers.UpdateUserControllers)
	e.DELETE("/users/:userid", controllers.DeleteUserControllrs)

	// categori
	e.GET("/category", controllers.GetCategoryControllers)
	e.POST("/category", controllers.CreateCategoryControllers)
	e.GET("/category/:categoryid", controllers.GetCategoryByID)
	e.PUT("/category/:categoryid", controllers.UpdateCategoryControllers)
	e.DELETE("/category/:categoryid", controllers.DeleteCategoryControllers)

	// product
	e.GET("/product", controllers.GetProductControllers)
	e.POST("/product", controllers.CreateProductControllers)
	e.GET("/product/:productid", controllers.GetProductByID)
	e.PUT("/product/:productid", controllers.UpdateProductControllers)
	e.DELETE("/product/:productid", controllers.DeleteProductControllers)

	// methods manual atau otomatis
	e.GET("/methods", controllers.GetMethodsControllers)
	e.POST("/methods", controllers.CreateMethodsControllers)
	e.GET("/methods/:methodsid", controllers.GetMethodsByID)
	e.PUT("/methods/:methodsid", controllers.UpdateMethodsControllers)
	e.DELETE("/methods/:methodsid", controllers.DeleteMethodsControllers)

	// rekening
	e.GET("/rekening", controllers.GetRekeningControllers)
	e.POST("/rekening", controllers.CreateRekeningControllers)
	e.GET("/rekening/:rekeningid", controllers.GetRekeningByID)
	e.PUT("/rekening/:rekeningid", controllers.UpdateRekeningControllers)
	e.DELETE("/rekening/:rekeningid", controllers.DeleteRekeningControllers)

	// cart

	// order

	// payment

	return e
}
