package httptransport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NYTimes/gziphandler"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/transports/httptransport/responseheaders"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
)

func (h HttpServer) registerRoutes(
	rtr *mux.Router,
) {

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
