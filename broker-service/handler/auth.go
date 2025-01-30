package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"log"

	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/dto"
	auth_proto "github.com/fydhfzh/ecommerce-go-application/src/broker-service/services_proto/auth-proto"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authHandler struct{}

type AuthHandler interface {
	Authenticate(c echo.Context) error
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (a *authHandler) Authenticate(c echo.Context) error {
	var authRequest dto.AuthenticationRequest

	log.Printf("accept incoming request")
	if err := c.Bind(&authRequest); err != nil {
		log.Printf("error binding request body: %v", err)
		errorResponse := dto.NewBadRequestError("error binding request body: " + err.Error())

		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	url := fmt.Sprintf("%s:%d", common.Config.ServicesConfig.AuthServiceConfig.Host, common.Config.ServicesConfig.AuthServiceConfig.GRPCPort)

	log.Printf("Dialing gRpc server with url: %v", url)
	authResponse, err := a.authenticateViaGRPC(authRequest, url)
	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") {
			log.Printf("invalid credentials")
			errorResponse := dto.NewBadRequestError("invalid credentials")

			return c.JSON(http.StatusBadRequest, errorResponse)
		}
		log.Printf("error dialing auth service: %v", err)
		errorResponse := dto.NewInternalServerError("internal server error")

		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	log.Printf("%v", authResponse)
	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "user authenticated successfully",
		Data:       authResponse,
	}

	return c.JSON(http.StatusOK, response)
}

func (a *authHandler) authenticateViaGRPC(authRequest dto.AuthenticationRequest, url string) (*dto.AuthenticationResponse, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := auth_proto.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	grpcAuthRequest := auth_proto.AuthRequest{
		Email:    authRequest.Email,
		Password: authRequest.Password,
	}

	grpcAuthResponse, err := client.Authenticate(ctx, &grpcAuthRequest)
	if err != nil {
		return nil, err
	}

	authResponse := dto.AuthenticationResponse{
		Jwt: grpcAuthResponse.Jwt,
	}

	return &authResponse, nil
}
