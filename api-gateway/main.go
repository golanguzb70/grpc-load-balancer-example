package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway/api"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway/config"
	grpcclient "github.com/golanguzb70/grpc-load-balancer-example/api-gateway/grpc-client"

	_ "github.com/golanguzb70/grpc-load-balancer-example/api-gateway/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Post CRUD API Gateway v1
// @version 1.0
// @description This is a Post CRUD API Gateway server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()

	serviceManager, err := grpcclient.New(cfg)
	if err != nil {
		panic(err)
	}

	handler := api.NewHandler(serviceManager)

	router := gin.Default()

	router.POST("/post", handler.CreatePost)
	router.PUT("/post/:key", handler.UpdatePost)
	router.GET("/post/:key", handler.GetPost)
	router.DELETE("/post/:key", handler.DeletePost)
	router.GET("/post", handler.ListPost)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":" + cfg.HTTPPort)
}
