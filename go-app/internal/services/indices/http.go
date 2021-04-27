package indices

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/transports/httptransport/requestbodies"
	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
)

func ShowIndexHTTPHandler(s Service) http.Handler {
	decoder := func(_ context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		id := vars["id"]

		return showIndexRequest{
			ID: id,
		}, nil
	}

	encoder := func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		assertedResponse := response.(showIndexResponse)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(assertedResponse.result)
	}

	return kithttp.NewServer(
		makeShowIndexEndpoint(s),
		decoder,
		encoder,
		apperrors.MakeGokitHTTPErrorEncoder(),
	)
}

func CreateIndexHTTPHandler(s Service) http.Handler {
	decoder := func(_ context.Context, r *http.Request) (interface{}, error) {
		jsonBodyDecoder := json.NewDecoder(r.Body)
		requestBody := requestbodies.CreateIndexBody{}
		err := jsonBodyDecoder.Decode(&requestBody)
		if err != nil {
			return nil, err
		}

		return createIndexRequest{
			ID: requestBody.ID,
		}, nil
	}

	encoder := func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		assertedResponse := response.(createIndexResponse)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		return json.NewEncoder(w).Encode(assertedResponse.result)
	}

	return kithttp.NewServer(
		makeCreateIndexEndpoint(s),
		decoder,
		encoder,
		apperrors.MakeGokitHTTPErrorEncoder(),
	)
}
