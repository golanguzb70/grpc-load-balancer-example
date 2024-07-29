package grpcclient

import (
	"fmt"

	grpclb "github.com/golanguzb70/grpc-lb"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway-client-lb/config"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway-client-lb/genproto/post_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PostService() post_service.PostServiceClient
}

type client struct {
	conn *grpc.ClientConn
	lb   grpclb.LB
}

func New(cfg *config.Config) (ServiceManager, error) {
	postServiceConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	factory := func() (*grpc.ClientConn, error) {
		return grpc.NewClient(
			fmt.Sprintf("%s:%s", cfg.PostServiceHost, cfg.PostServicePort),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	}

	logger := func(msg string) {
		fmt.Println(msg)
	}

	lb, err := grpclb.New(6, 5, factory, logger)
	if err != nil {
		return nil, err
	}

	return &client{
		conn: postServiceConn,
		lb:   lb,
	}, nil
}

func (c *client) PostService() post_service.PostServiceClient {
	return post_service.NewPostServiceClient(c.lb.Get())
}
