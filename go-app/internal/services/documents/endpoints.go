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

type patchDocumentRequest struct {
	IndexID    string
	DocumentID string
	Document   map[string]interface{}
}

type patchDocumentResponse struct{}

func makePatchDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(patchDocumentRequest)
		err = s.PatchDocument(
			assertedRequest.IndexID,
			assertedRequest.DocumentID,
			assertedRequest.Document,
		)
		return patchDocumentResponse{}, err
	}
}

type deleteDocumentRequest struct {
	IndexID    string
	DocumentID string
}

type deleteDocumentResponse struct{}

func makeDeleteDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(deleteDocumentRequest)
		err = s.DeleteDocument(
			assertedRequest.IndexID,
			assertedRequest.DocumentID,
		)
		return deleteDocumentResponse{}, err
	}
}

type listDocumentRequest struct {
	IndexID        string
	Page           int
	PerPage        int
	MatchMap       map[string]string
	SearchPropList []string
	SearchVal      string
	InputSortList  [][2]string
}

type listDocumentResponse struct {
	result responses.DocumentPage
}

func makeListDocumentEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		assertedRequest := request.(listDocumentRequest)
		res, err := s.ListDocument(
			assertedRequest.IndexID,
			assertedRequest.Page,
			assertedRequest.PerPage,
			assertedRequest.MatchMap,
			assertedRequest.SearchPropList,
			assertedRequest.SearchVal,
			assertedRequest.InputSortList,
		)
		return listDocumentResponse{res}, err
	}
}
