package errors

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
)

func MakeGokitHTTPErrorEncoder() kithttp.ServerOption {
	errorEncoder := func(_ context.Context, err error, w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")

		if wrappedErr, ok := err.(WrappedError); ok {
			w.WriteHeader(wrappedErr.StatusCode())
			json.NewEncoder(w).Encode(responses.ErrorResponse{
				ErrorCode: responses.ErrorCode{
					ID: wrappedErr.Error(),
				},
			})

			return
		}

		// Default error return
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responses.ErrorResponse{
			ErrorCode: responses.ErrorCode{
				ID: err.Error(),
			},
		})

		return
	}

	return kithttp.ServerErrorEncoder(errorEncoder)
}
