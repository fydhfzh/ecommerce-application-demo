package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/constant"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/model"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/products_proto"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/repository"
	"github.com/fydhfzh/ecommerce-go-application/src/product-command-service/service"
	"google.golang.org/grpc"
)

type ProductServer struct {
	products_proto.UnimplementedProductServiceServer
	productRepository    repository.ProductRepository
	kafkaProducerService *service.KafkaProducerService
}

func (p *ProductServer) Save(ctx context.Context, request *products_proto.ProductSaveRequest) (*products_proto.ProductResponse, error) {
	// save product
	product := model.Product{
		Name:        request.GetName(),
		Description: request.GetDescription(),
		Stock:       uint(request.GetStock()),
	}

	newProduct, err := p.productRepository.Save(product)
	if err != nil {
		return nil, err
	}

	event := dto.ProductEvent{
		EventType: constant.SaveProduct,
		ProductId: newProduct.ID.String(),
		Timestamp: time.Now().Local(),
		Payload:   *newProduct,
	}

	err = p.kafkaProducerService.SendProductEvent(event)
	if err != nil {
		return nil, err
	}

	productResponse := products_proto.ProductResponse{
		Id:          newProduct.ID.String(),
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Stock:       uint32(newProduct.Stock),
		CreatedAt:   newProduct.CreatedAt.Local().String(),
		UpdatedAt:   newProduct.UpdatedAt.Local().String(),
	}

	return &productResponse, nil
}

func (p *ProductServer) Update(ctx context.Context, request *products_proto.ProductUpdateRequest) (*products_proto.ProductResponse, error) {
	// update product
	product, err := p.productRepository.GetOne(request.Id)
	if err != nil {
		return nil, err
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Stock = uint(request.Stock)

	updateProduct, err := p.productRepository.Save(*product)
	if err != nil {
		return nil, err
	}

	event := dto.ProductEvent{
		EventType: constant.UpdateProduct,
		ProductId: updateProduct.ID.String(),
		Timestamp: time.Now().Local(),
		Payload:   *updateProduct,
	}

	err = p.kafkaProducerService.SendProductEvent(event)
	if err != nil {
		return nil, err
	}

	productResponse := products_proto.ProductResponse{
		Id:          updateProduct.ID.String(),
		Name:        updateProduct.Name,
		Description: updateProduct.Description,
		Stock:       uint32(updateProduct.Stock),
		CreatedAt:   updateProduct.CreatedAt.String(),
		UpdatedAt:   updateProduct.UpdatedAt.String(),
	}

	return &productResponse, nil
}

func gRPCListen(productRepository repository.ProductRepository, kafkaProducerService *service.KafkaProducerService) {
	appConfig := common.Config.AppConfig
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	productServer := ProductServer{
		productRepository:    productRepository,
		kafkaProducerService: kafkaProducerService,
	}

	products_proto.RegisterProductServiceServer(s, &productServer)

	log.Printf("gRPC Server started on port %d", appConfig.GRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
