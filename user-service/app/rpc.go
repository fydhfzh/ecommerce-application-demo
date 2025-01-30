package app

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/fydhfzh/ecommerce-go-application/src/user-service/common"
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/dto"
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/entity"
	"github.com/fydhfzh/ecommerce-go-application/src/user-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type RPCServer struct {
	userRepository repository.UserRepository
}

func NewRPCServer(userRepository repository.UserRepository) RPCServer {
	return RPCServer{
		userRepository: userRepository,
	}
}

func (r *RPCServer) Register(req dto.RegisterRequest, reply *dto.RegisterResponse) error {
	user := entity.User{
		Email:    req.Email,
		Password: req.Email,
		Fullname: req.Fullname,
		Age:      req.Age,
	}

	err := user.HashPassword()
	if err != nil {
		return err
	}

	newUser, err := r.userRepository.SaveUser(user)
	if err != nil {
		return err
	}

	userResponse := dto.RegisterResponse{
		Email:     newUser.Email,
		Fullname:  newUser.Fullname,
		Age:       newUser.Age,
		Active:    newUser.Active,
		Role:      newUser.Role,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	*reply = userResponse

	return nil
}

func (r *RPCServer) Login(userCredentials dto.LoginRequest, reply *bool) error {
	user, err := r.userRepository.GetUserByEmail(userCredentials.Email)
	if err != nil {
		log.Printf("error retrieving user: %v", err)

		return err
	}

	err = user.ComparePassword(userCredentials.Password)
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			log.Printf("invalid credentials")

			*reply = false

			return errors.New("invalid credentials")
		default:
			log.Printf("error comparing password: %v", err)

			*reply = false

			return err
		}
	}

	*reply = true

	return nil
}

func rpcListen(appConfig common.ApplicationConfig) error {
	log.Printf("Serving rpc on port %d...", appConfig.RpcPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.RpcPort))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		rpcConn, err := listener.Accept()
		fmt.Printf("Connection coming from %s", rpcConn.RemoteAddr())
		if err != nil {
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}
