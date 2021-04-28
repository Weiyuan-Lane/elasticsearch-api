package documents

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
)

type showDocumentRequest struct {
	IndexID    string
	DocumentID string
}

type showDocumentResponse struct {
	result responses.Document
}

func makeShowDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(showDocumentRequest)
		res, err := s.RetrieveDocument(
			assertedRequest.IndexID,
			assertedRequest.DocumentID,
		)
		return showDocumentResponse{res}, err
	}
}

type createDocumentRequest struct {
	IndexID    string
	DocumentID string
	Document   map[string]interface{}
}

type createDocumentResponse struct {
	result responses.CreatedID
}

func makeCreateDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(createDocumentRequest)
		res, err := s.CreateDocument(
			assertedRequest.IndexID,
			assertedRequest.DocumentID,
			assertedRequest.Document,
		)
		return createDocumentResponse{res}, err
	}
}
