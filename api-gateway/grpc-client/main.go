package grpcclient

import (
	"fmt"

	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway/config"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway/genproto/post_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PostService() post_service.PostServiceClient
}

type client struct {
	postService post_service.PostServiceClient
}

func New(cfg *config.Config) (ServiceManager, error) {
	postServiceConn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &client{
		postService: post_service.NewPostServiceClient(postServiceConn),
	}, nil
}

func (c *client) PostService() post_service.PostServiceClient {
	return c.postService
}
