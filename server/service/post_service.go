package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/golanguzb70/grpc-load-balancer-example/server/genproto/post_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PostService ...
type PostService struct {
	Posts map[string]*pb.Post
	pb.UnimplementedPostServiceServer
}

// NewPostService ...
func NewPostService() *PostService {
	return &PostService{
		Posts: map[string]*pb.Post{},
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.Post, error) {
	s.Posts[req.Id] = &pb.Post{
		Id:        req.Id,
		Slug:      req.Slug,
		Title:     req.Title,
		Body:      req.Body,
		OwnerId:   req.OwnerId,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	return s.Posts[req.Id], nil
}

func (s *PostService) GetByKey(ctx context.Context, req *pb.KeyRequest) (*pb.Post, error) {
	v, ok := s.Posts[req.GetId()]
	if !ok {
		return &pb.Post{}, status.Error(codes.NotFound, fmt.Sprintf("item with id [%s] not found", req.Id))
	}

	return v, nil
}

func (s *PostService) Find(ctx context.Context, req *pb.Filter) (*pb.Posts, error) {
	fmt.Println("Find request received")

	response := &pb.Posts{}

	for _, v := range s.Posts {
		if strings.Contains(v.Title, req.Search) || req.Search == "" {
			response.Count++
			response.Items = append(response.Items, v)
		}
	}

	if (req.Page-1)*req.Limit+req.Limit <= response.Count {
		response.Items = response.Items[(req.Page-1)*req.Limit:]
	} else {
		response.Items = response.Items[(req.Page-1)*req.Limit:]
	}

	if len(response.Items) > int(req.Limit) {
		response.Items = response.Items[:req.Limit]
	}

	return response, nil
}

func (s *PostService) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Post, error) {
	v, ok := s.Posts[req.GetId()]
	if !ok {
		return &pb.Post{}, status.Error(codes.NotFound, fmt.Sprintf("item with id [%s] not found", req.Id))
	}

	v.Title = req.Title
	v.Body = req.Body
	v.UpdatedAt = time.Now().Format(time.RFC3339)

	s.Posts[req.Id] = v

	return s.Posts[req.Id], nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.KeyRequest) (*empty.Empty, error) {
	_, ok := s.Posts[req.GetId()]
	if !ok {
		return &emptypb.Empty{}, status.Error(codes.NotFound, fmt.Sprintf("item with id [%s] not found", req.Id))
	}

	delete(s.Posts, req.GetId())

	return &emptypb.Empty{}, nil
}
