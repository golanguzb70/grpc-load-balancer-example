package api

import grpcclient "github.com/golanguzb70/grpc-load-balancer-example/api-gateway-client-lb/grpc-client"

type Handler struct {
	services grpcclient.ServiceManager
}

func NewHandler(serviceManager grpcclient.ServiceManager) *Handler {
	return &Handler{
		services: serviceManager,
	}
}
