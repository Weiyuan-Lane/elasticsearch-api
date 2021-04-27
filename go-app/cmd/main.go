package main

import (
	"strconv"

	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/transports/httptransport"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/config"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/elasticsearchclient"
)

func main() {
	appConfig := config.New()

	esClient := elasticsearchclient.New(
		appConfig.ElasticsearchProtocol,
		appConfig.ElasticsearchHost,
		strconv.Itoa(appConfig.ElasticsearchPort),
	)

	httptransport.Init(
		appConfig,
		esClient,
	)
}
