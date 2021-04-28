package httptransport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NYTimes/gziphandler"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	documentsvc "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/services/documents"
	indexsvc "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/services/indices"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/transports/httptransport/responseheaders"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
)

func (h HttpServer) registerRoutes(
	rtr *mux.Router,
	indexService indexsvc.Service,
	documentService documentsvc.Service,
) {

	// Indices Routes
	rtr.Methods("POST").
		Path("/indices").
		Handler(indexsvc.CreateIndexHTTPHandler(indexService))
	rtr.Methods("GET").
		Path("/indices/{id}").
		Handler(indexsvc.ShowIndexHTTPHandler(indexService))

	// Documents Routes
	rtr.Methods("POST").
		Path("/indices/{indexID}/documents").
		Handler(documentsvc.CreateDocumentHTTPHandler(documentService))
	rtr.Methods("GET").
		Path("/indices/{indexID}/documents/{documentID}").
		Handler(documentsvc.ShowDocumentHTTPHandler(documentService))

	registerMiddlewares(rtr)
	registerFallbackRoute(rtr)
}

func registerFallbackRoute(rtr *mux.Router) {
	nullEndpoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		return nil, nil
	}
	decoder := func(_ context.Context, r *http.Request) (interface{}, error) {
		return nil, nil
	}
	encoder := func(_ context.Context, w http.ResponseWriter, _ interface{}) error {
		responseheaders.SetJSONResponseType(w)
		w.WriteHeader(http.StatusNotFound)

		return json.NewEncoder(w).Encode(responses.ErrorResponse{
			ErrorCode: responses.ErrorCode{
				ID: apperrors.ErrPathNotFound.Error(),
			},
		})
	}

	rtr.NotFoundHandler = kithttp.NewServer(
		nullEndpoint,
		decoder,
		encoder,
	)
}

func registerMiddlewares(rtr *mux.Router) {
	rtr.Use(
		gziphandler.GzipHandler,
	)
}
