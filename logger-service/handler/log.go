package handler

import (
	"net/http"

	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/logger-service/service"
	"github.com/labstack/echo/v4"
)

type logHandler struct {
	logService service.LogService
}

type LogHandler interface {
	Save(c echo.Context) error
	GetAll(c echo.Context) error
}

func NewLogHandler(logService service.LogService) LogHandler {
	return &logHandler{
		logService: logService,
	}
}

func (l *logHandler) Save(c echo.Context) error {
	var logRequest dto.LogRequest

	if err := c.Bind(&logRequest); err != nil {
		return err
	}

	logResponse, err := l.logService.Save(logRequest)
	if err != nil {
		return err
	}

	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "log created successfully",
		Data:       logResponse,
	}

	return c.JSON(http.StatusCreated, response)
}

func (l *logHandler) GetAll(c echo.Context) error {
	logsResponse, err := l.logService.GetAll()
	if err != nil {
		return err
	}

	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "logs retrieved successfully",
		Data:       logsResponse,
	}

	return c.JSON(http.StatusOK, response)
}
