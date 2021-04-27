package main

import (
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/transports/httptransport"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/config"
)

func main() {
	appConfig := config.New()
	httptransport.Init(appConfig)
}
