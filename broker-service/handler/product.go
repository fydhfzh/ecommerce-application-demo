package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/broker-service/services_proto/products_proto"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type productHandler struct {
}

type ProductHandler interface {
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	GetAllProduct(c echo.Context) error
	GetProductById(c echo.Context) error
}

func NewProductHandler() ProductHandler {
	return &productHandler{}
}

func (p *productHandler) CreateProduct(c echo.Context) error {
	productServiceConfig := common.Config.ServicesConfig.ProductCommandServiceConfig

	url := fmt.Sprintf("http://%s:%d", productServiceConfig.Host, productServiceConfig.GRPCPort)

	var payload dto.ProductSaveRequest

	if err := c.Bind(&payload); err != nil {
		errorResponse := dto.NewBadRequestError("error binding request body")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	productResponse, err := p.createProductViaGRPC(url, payload)
	if err != nil {
		errorResponse := dto.NewInternalServerError("internal server error")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "product created successfully",
		Data:       productResponse,
	}

	return c.JSON(http.StatusCreated, response)
}

func (p *productHandler) UpdateProduct(c echo.Context) error {
	productCommandServiceConfig := common.Config.ServicesConfig.ProductCommandServiceConfig

	id := c.Param("id")

	url := fmt.Sprintf("%s:%d", productCommandServiceConfig.Host, productCommandServiceConfig.GRPCPort)

	var payload dto.ProductUpdateRequest

	if err := c.Bind(&payload); err != nil {
		errorResponse := dto.NewBadRequestError("error binding request body")
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	payload.Id = id

	productResponse, err := p.updateProductViaGRPC(url, payload)
	if err != nil {
		errorResponse := dto.NewInternalServerError("internal server error")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "product updated successfully",
		Data:       productResponse,
	}

	return c.JSON(http.StatusOK, response)
}

func (p *productHandler) GetAllProduct(c echo.Context) error {
	productQueryServiceConfig := common.Config.ServicesConfig.ProductQueryServiceConfig

	url := fmt.Sprintf("%s:%d", productQueryServiceConfig.Host, productQueryServiceConfig.RPCPort)
	log.Println(url)

	productsResponse, err := p.getAllProductViaRPC(url)
	if err != nil {
		log.Println(err)
		errorResponse := dto.NewInternalServerError("internal server error")

		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "products retrieved successfully",
		Data:       productsResponse,
	}

	return c.JSON(http.StatusOK, response)
}

func (p *productHandler) GetProductById(c echo.Context) error {
	productQueryServiceConfig := common.Config.ServicesConfig.ProductQueryServiceConfig

	id := c.Param("id")
	url := fmt.Sprintf("%s:%d", productQueryServiceConfig.Host, productQueryServiceConfig.RPCPort)

	productResponse, err := p.getProductByIdViaRPC(url, id)
	if err != nil {
		log.Println(err)
		errorResponse := dto.NewInternalServerError("internal server error")

		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	response := dto.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "product retrieved successfully",
		Data:       productResponse,
	}

	return c.JSON(http.StatusOK, response)
}

func (p *productHandler) createProductViaGRPC(url string, payload dto.ProductSaveRequest) (*dto.ProductResponse, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := products_proto.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	productSaveRequest := products_proto.ProductSaveRequest{
		Name:        payload.Name,
		Description: payload.Description,
		Stock:       uint32(payload.Stock),
	}

	productResponse, err := client.Save(ctx, &productSaveRequest)
	if err != nil {
		return nil, err
	}

	response := dto.ProductResponse{
		Id:          productResponse.Id,
		Name:        productResponse.Name,
		Description: productResponse.Description,
		Stock:       uint(productResponse.Stock),
		CreatedAt:   productResponse.CreatedAt,
		UpdatedAt:   productResponse.UpdatedAt,
	}

	return &response, nil
}

func (p *productHandler) updateProductViaGRPC(url string, payload dto.ProductUpdateRequest) (*dto.ProductResponse, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := products_proto.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	productUpdateRequest := products_proto.ProductUpdateRequest{
		Id:          payload.Id,
		Name:        payload.Name,
		Description: payload.Description,
		Stock:       uint32(payload.Stock),
	}

	productResponse, err := client.Update(ctx, &productUpdateRequest)
	if err != nil {
		return nil, err
	}

	response := dto.ProductResponse{
		Id:          productResponse.Id,
		Name:        productResponse.Name,
		Description: productResponse.Description,
		Stock:       uint(productResponse.Stock),
		CreatedAt:   productResponse.CreatedAt,
		UpdatedAt:   productResponse.UpdatedAt,
	}

	return &response, err
}

func (p *productHandler) getAllProductViaRPC(url string) ([]dto.ProductResponse, error) {
	conn, err := rpc.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var productsResponse []dto.ProductResponse

	err = conn.Call("RPCServer.GetAllProducts", 0, &productsResponse)
	if err != nil {
		return nil, err
	}

	return productsResponse, nil
}

func (p *productHandler) getProductByIdViaRPC(url string, id string) (*dto.ProductResponse, error) {
	conn, err := rpc.Dial("tcp", url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var productResponse dto.ProductResponse

	err = conn.Call("RPCServer.GetProductById", id, &productResponse)
	if err != nil {
		return nil, err
	}

	return &productResponse, nil
}
