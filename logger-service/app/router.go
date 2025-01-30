package app

import (
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/db"
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/handler"
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/repository"
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/service"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	logRepository := repository.NewLogRepository(db.LogCollection)
	logService := service.NewLogService(logRepository)
	logHandler := handler.NewLogHandler(logService)

	e := echo.New()
	apiGroup := e.Group("/api/v1")
	apiGroup.GET("", logHandler.GetAll)
	apiGroup.POST("", logHandler.Save)

	return e
}
