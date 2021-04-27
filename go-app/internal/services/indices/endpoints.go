package indices

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
)

type showIndexRequest struct {
	ID string
}

type showIndexResponse struct {
	result responses.Index
}

func makeShowIndexEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(showIndexRequest)
		res, err := s.RetrieveIndex(assertedRequest.ID)
		return showIndexResponse{res}, err
	}
}

type createIndexRequest struct {
	ID string
}

type createIndexResponse struct {
	result responses.CreatedID
}

func makeCreateIndexEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(createIndexRequest)
		res, err := s.CreateIndex(assertedRequest.ID)
		return createIndexResponse{res}, err
	}
}
