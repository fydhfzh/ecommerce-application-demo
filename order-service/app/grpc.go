package app

import (
	"context"
	"fmt"
	"net"

	"github.com/fydhfzh/ecommerce-go-application/src/order-service/entity"
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/proto_files/order_proto"
	"github.com/fydhfzh/ecommerce-go-application/src/order-service/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type OrderServer struct {
	orderRepository repository.OrderRepository
	order_proto.UnimplementedOrderServiceServer
}

func (o *OrderServer) CreateOrder(ctx context.Context, req *order_proto.OrderRequest) (*order_proto.OrderResponse, error) {
	productId := uuid.MustParse(req.ProductId)
	itemPriceExample := 1000 // TODO

	order := entity.Order{
		ProductId:  productId,
		Quantity:   uint(req.Quantity),
		TotalPrice: uint(req.Quantity) * uint(itemPriceExample),
	}

	newOrder, err := o.orderRepository.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	response := order_proto.OrderResponse{
		ProductId:  newOrder.ID.String(),
		Quantity:   uint32(newOrder.Quantity),
		TotalPrice: uint32(newOrder.TotalPrice),
		Status:     newOrder.Status,
	}

	return &response, nil
}

// func (o *OrderServer) getProduct(id string)
func grpcListen(gRpcPort uint, orderServer OrderServer) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gRpcPort))
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	order_proto.RegisterOrderServiceServer(s, &orderServer)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
