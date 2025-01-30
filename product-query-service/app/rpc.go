package app

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/product-events-service/repository"
)

type RPCServer struct {
	productRepository repository.ProductRepository
}

func NewRPCServer(productRepository repository.ProductRepository) *RPCServer {
	return &RPCServer{
		productRepository: productRepository,
	}
}

func (r *RPCServer) GetProductById(id string, resp *dto.ProductResponse) error {
	product, err := r.productRepository.GetProductById(id)
	if err != nil {
		log.Println("error getting product data")
		return err
	}

	productResponse := dto.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.String(),
		UpdatedAt:   product.UpdatedAt.String(),
	}

	*resp = productResponse

	return nil
}

func (r *RPCServer) GetAllProducts(_ int, resp *[]dto.ProductResponse) error {
	products, err := r.productRepository.GetAllProducts()
	if err != nil {
		log.Println("error getting products data")
		return err
	}

	var productsResponse []dto.ProductResponse
	for _, product := range products {
		productResponse := dto.ProductResponse{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		}

		productsResponse = append(productsResponse, productResponse)
	}

	*resp = productsResponse

	return nil
}

func rpcListen() error {
	log.Printf("Starting RPC server on port: %d", common.Config.AppConfig.RPCPort)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", common.Config.AppConfig.RPCPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}
