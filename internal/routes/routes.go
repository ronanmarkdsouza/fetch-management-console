package routes

import (
	"fetch/management-console/internal/controllers"
	"fetch/management-console/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRoute(port string) {
	router := gin.Default()

	//Test endpoints
	router.GET("/test", controllers.Test)

	router.Use(middleware.AuthenticateMiddleware())

	//Hosts endpoints
	router.GET("/:apikey/hosts", controllers.ViewHosts)
	router.GET("/:apikey/hosts/:id", controllers.ViewHost)
	router.POST("/:apikey/hosts/create", controllers.CreateHost)
	router.PUT("/:apikey/hosts/edit/:id", controllers.EditHost)
	router.DELETE("/:apikey/hosts/delete/:id", controllers.DeleteHost)

	//Machine Type endpoints
	router.GET("/:apikey/machinetypes", controllers.ViewMachineTypes)
	router.GET("/:apikey/machinetypes/:id", controllers.ViewMachineType)
	router.POST("/:apikey/machinetypes/create", controllers.CreateMachineType)
	router.PUT("/:apikey/machinetypes/edit/:id", controllers.EditMachineType)
	router.DELETE("/:apikey/machinetypes/delete/:id", controllers.DeleteMachineType)

	//Products endpoints
	router.GET("/:apikey/products", controllers.ViewProducts)
	router.GET("/:apikey/products/:id", controllers.ViewProduct)
	router.POST("/:apikey/products/create", controllers.CreateProduct)
	router.PUT("/:apikey/products/edit/:id", controllers.EditProduct)
	router.DELETE("/:apikey/products/delete/:id", controllers.DeleteProduct)

	//Vending Machine endpoints
	router.GET("/:apikey/vendingmachines", controllers.ViewVendingMachines)
	router.GET("/:apikey/vendingmachines/:id", controllers.ViewVendingMachine)
	router.POST("/:apikey/vendingmachines/create", controllers.CreateVendingMachine)
	router.PUT("/:apikey/vendingmachines/edit/:id", controllers.EditVendingMachine)
	router.DELETE("/:apikey/vendingmachines/delete/:id", controllers.DeleteVendingMachine)

	router.Run(":" + port)
}
