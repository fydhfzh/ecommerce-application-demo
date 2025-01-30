package app

import (
	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			e-commerce application
// @version		1.0
// @description	This is ecommerce application example
//
// @BasePath		/api/v1
func NewRouter() *echo.Echo {
	// instantiate handler
	brokerHandler := handler.NewBrokerHandler()
	productHandler := handler.NewProductHandler()
	authHandler := handler.NewAuthHandler()

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/swagger/*", echoSwagger.WrapHandler)
	// apiGroup.Use(middleware.LogRequest)
	apiGroup.POST("/broker", brokerHandler.Broker)

	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/signin", authHandler.Authenticate)

	productGroup := apiGroup.Group("/products")
	productGroup.POST("", productHandler.CreateProduct)
	productGroup.PUT("/:id", productHandler.UpdateProduct)
	productGroup.GET("", productHandler.GetAllProduct)
	productGroup.GET("/:id", productHandler.GetProductById)

	return e
}
