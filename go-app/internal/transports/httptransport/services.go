package httptransport

import (
	"github.com/gorilla/mux"
	documentsvc "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/services/documents"
	indexsvc "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/services/indices"
)

func (h HttpServer) registerServices(router *mux.Router) {
	indexService := indexsvc.Service{
		Logger:              h.Logger,
		ElasticsearchClient: h.ElasticsearchClient,
	}

	documentService := documentsvc.Service{
		Logger:              h.Logger,
		ElasticsearchClient: h.ElasticsearchClient,
	}

	h.registerRoutes(
		router,
		indexService,
		documentService,
	)
}
