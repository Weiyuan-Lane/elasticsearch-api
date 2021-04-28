package documents

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/transports/httptransport/requestbodies"
	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
)

func ShowDocumentHTTPHandler(s Service) http.Handler {
	decoder := func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		indexID := vars["indexID"]
		documentID := vars["documentID"]

		return showDocumentRequest{
			IndexID:    indexID,
			DocumentID: documentID,
		}, nil
	}

	encoder := func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		assertedResponse := response.(showDocumentResponse)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(assertedResponse.result)
	}

	return kithttp.NewServer(
		makeShowDocumentEndpoint(s),
		decoder,
		encoder,
		apperrors.MakeGokitHTTPErrorEncoder(),
	)
}

func CreateDocumentHTTPHandler(s Service) http.Handler {
	decoder := func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		indexID := vars["indexID"]

		jsonBodyDecoder := json.NewDecoder(r.Body)
		requestBody := requestbodies.CreateDocumentBody{}
		err := jsonBodyDecoder.Decode(&requestBody)
		if err != nil {
			return nil, err
		}
		if id, ok := requestBody["id"]; !ok || id == "" {
			return nil, apperrors.ErrCreateDocumentRequestBodyNoID
		}

		unparsedDocumentID := requestBody["id"]
		parsedDocumentID, ok := unparsedDocumentID.(string)
		if !ok {
			return nil, apperrors.ErrCreateDocumentRequestBodyIDIsNotString
		}

		return createDocumentRequest{
			IndexID:    indexID,
			DocumentID: parsedDocumentID,
			Document:   requestBody,
		}, nil
	}

	encoder := func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		assertedResponse := response.(createDocumentResponse)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(assertedResponse.result)
	}

	return kithttp.NewServer(
		makeCreateDocumentEndpoint(s),
		decoder,
		encoder,
		apperrors.MakeGokitHTTPErrorEncoder(),
	)
}

func ListDocumentHTTPHandler(s Service) http.Handler {
	decoder := func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		indexID := vars["indexID"]

		queryValues := r.URL.Query()

		pageStr := queryValues.Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		perPageStr := queryValues.Get("per_page")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			perPage = 10
		}

		jsonBodyDecoder := json.NewDecoder(r.Body)
		requestBody := requestbodies.SearchDocumentBody{}
		err = jsonBodyDecoder.Decode(&requestBody)
		if err != nil {
			return nil, err
		}

		return listDocumentRequest{
			IndexID:        indexID,
			Page:           page,
			PerPage:        perPage,
			MatchMap:       requestBody.MatchMap,
			SearchPropList: requestBody.SearchPropList,
			SearchVal:      requestBody.SearchVal,
			InputSortList:  requestBody.InputSortList,
		}, nil
	}

	encoder := func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		assertedResponse := response.(listDocumentResponse)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(assertedResponse.result)
	}

	return kithttp.NewServer(
		makeListDocumentEndpoint(s),
		decoder,
		encoder,
		apperrors.MakeGokitHTTPErrorEncoder(),
	)
}
