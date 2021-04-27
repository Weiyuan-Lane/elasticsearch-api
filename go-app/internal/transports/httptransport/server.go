package httptransport

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/config"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/elasticsearchclient"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/logger"
)

func Init(appConfig config.ApplicationConfig, esClient elasticsearchclient.ElasticSearchClient) {
	server := HttpServer{
		Logger:                  appConfig.Logger,
		Port:                    strconv.Itoa(appConfig.HTTPPort),
		GracefulShutdownSeconds: int64(appConfig.HTTPGracefulShutdownSeconds),
		ElasticsearchClient:     esClient,
	}

	server.ListenAndServe()
}

type HttpServer struct {
	Logger                  logger.Logger
	Port                    string
	GracefulShutdownSeconds int64
	ElasticsearchClient     elasticsearchclient.ElasticSearchClient
}

func (h HttpServer) ListenAndServe() {
	h.initHTTPServer()
}

func (h HttpServer) initHTTPServer() {
	router := mux.NewRouter()
	errs := make(chan error)
	address := ":" + h.Port
	h.registerServices(router)
	server := h.makeHttpServerFrom(address, router)

	go func() {
		errs <- server.ListenAndServe()
	}()

	h.Logger.Infow("Serving from port " + h.Port)
	h.Logger.Infow((<-errs).Error())

	gracefulShutdownTime := time.Duration(h.GracefulShutdownSeconds) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTime)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		h.Logger.Infow(fmt.Sprintf("Server Shutdown Failed:%+v", err))
	} else {
		h.Logger.Infow("Graceful shutdown completed")
	}
}

func (h HttpServer) makeHttpServerFrom(address string, handler http.Handler) *http.Server {
	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    address,
		Handler: h2c.NewHandler(handler, h2s),
	}

	return server
}
