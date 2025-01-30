package app

import (
	"context"
	"fmt"
	"net"
	"net/rpc"

	auth_proto "github.com/fydhfzh/ecommerce-go-application/src/auth-service/auth_proto"
	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/logger"
	"github.com/fydhfzh/ecommerce-go-application/src/auth-service/service"
	"google.golang.org/grpc"
)

type AuthServer struct {
	jwtService        service.JwtService
	userServiceConfig common.UserServiceConfig
	auth_proto.UnimplementedAuthServiceServer
}

func (a *AuthServer) Authenticate(ctx context.Context, req *auth_proto.AuthRequest) (*auth_proto.AuthResponse, error) {
	logger.Logger.Infof("authenticate endpoint hit with data: %v", req)

	err := a.ValidateUserCredentials(req.Email, req.Password)
	if err != nil {
		logger.Logger.Errorf("error validating user credentials: %v", err)

		return nil, err
	}

	token, err := a.jwtService.GenerateToken(req.Email)
	if err != nil {
		logger.Logger.Errorf("error generating token: %v", err)
		return nil, err
	}

	authResponse := auth_proto.AuthResponse{
		Jwt: token,
	}

	return &authResponse, nil
}

func (a *AuthServer) ValidateUserCredentials(email, password string) error {
	url := fmt.Sprintf("%s:%d", a.userServiceConfig.Host, a.userServiceConfig.RpcPort)

	loginRequest := dto.LoginRequest{
		Email:    email,
		Password: password,
	}

	valid, err := validateUserViaRPC(loginRequest, url)
	if err != nil && !valid {
		logger.Logger.Errorf("error validating user: %v", err)

		return err
	} else if err == nil && !valid {
		logger.Logger.Errorf("invalid credentials")

		return err
	}

	return nil
}

func validateUserViaRPC(loginRequest dto.LoginRequest, url string) (bool, error) {
	conn, err := rpc.Dial("tcp", url)
	if err != nil {
		logger.Logger.Errorf("error dialing user service: %v", err)

		return false, err
	}
	defer conn.Close()

	var valid bool

	err = conn.Call("RPCServer.Login", loginRequest, &valid)
	if err != nil && !valid {
		return false, err
	} else if err == nil && !valid {
		return false, nil
	}

	return true, nil
}

func gRpcListen(gRpcPort uint, authServer AuthServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gRpcPort))
	if err != nil {
		logger.Logger.Fatalf("error listening to port %v", gRpcPort)
	}

	s := grpc.NewServer()

	auth_proto.RegisterAuthServiceServer(s, &authServer)

	logger.Logger.Infof("gRPC server started on port %d", gRpcPort)

	if err := s.Serve(lis); err != nil {
		logger.Logger.Fatalf("error listening for grpc")
	}
}
