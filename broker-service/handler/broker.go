package handler

import (
	"net/http"

	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/dto"
	"github.com/labstack/echo/v4"
)

type brokerHandler struct{}

type BrokerHandler interface {
	Broker(c echo.Context) error
}

func NewBrokerHandler() BrokerHandler {
	return &brokerHandler{}
}

// Broker godoc
//
//	@Summary		Test endpoint
//	@Description	This endpoint can be hit as a test
//	@Tags			broker
//	@Accept			json
//	@Product		json
//
//	@Param			action	body		dto.BrokerRequest	true "Broker Action"
//
//	@Success		200		{object}	dto.APIResponse
//	@Failure		400		{object}	dto.APIResponse
//	@Router			/broker [post]
func (b *brokerHandler) Broker(c echo.Context) error {
	var brokerRequest dto.BrokerRequest

	if err := c.Bind(&brokerRequest); err != nil {
		brokerResponse := dto.NewBadRequestError("error binding request body")

		return c.JSON(http.StatusBadRequest, &brokerResponse)
	}

	brokerResponse := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "hit the broker!",
	}

	return c.JSON(http.StatusOK, &brokerResponse)
}
