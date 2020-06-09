package grpc

import (
	"context"
	"github.com/ciazhar/golang-cqrs/api-grpc/api"
	"google.golang.org/grpc"
)

type repository struct {
}

func (r repository) GetByID(ctx context.Context, request *api.SocialRequest) (*api.SocialResponse, error) {
	id := request.Id
	response := &api.SocialResponse{
		Id:     id,
		Name:   "Dummy",
		Detail: "Dummy Detail",
	}
	return response, nil
}

func NewSocialGRPCRepository(server *grpc.Server) {
	api.RegisterSocialServiceServer(server, &repository{})
}
