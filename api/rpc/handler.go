package rpc

import (
	"context"
	"ir.safari.shortlink/api/gen"
	"ir.safari.shortlink/model"
	"ir.safari.shortlink/repository"
)

type ServiceRpc interface {
	CreateShortLink(context.Context, *gen.CreateShortLinkRequest) (*gen.CreateShortLinkResponse, error)
	Ping(context.Context, *gen.PingRequest) (*gen.PingResponse, error)
}

type ServiceRpcImpl struct {
	originalUrlRepository repository.OriginalUrlRepository
}

func NewServiceRpcImpl(originalUrlRepository repository.OriginalUrlRepository) *ServiceRpcImpl {
	return &ServiceRpcImpl{
		originalUrlRepository: originalUrlRepository,
	}
}

func (s *ServiceRpcImpl) CreateShortLink(_ context.Context, request *gen.CreateShortLinkRequest) (*gen.CreateShortLinkResponse, error) {
	orb := model.NewUrlBuilder()
	orb.SetOriginalUrl(request.OriginalUrl)

	return nil, nil
}

func (s *ServiceRpcImpl) Ping(_ context.Context, request *gen.PingRequest) (*gen.PingResponse, error) {
	return &gen.PingResponse{
		Outcome: request.GetIncome() * 2,
	}, nil
}
