package elasticsearchclient

import (
	"fmt"
)

type ElasticSearchClient struct {
	hostWithPort string
}

func New(protocol, host, port string) ElasticSearchClient {
	hostWithPort := fmt.Sprintf(fullHostTemplate, protocol, host, port)

	return ElasticSearchClient{
		hostWithPort: hostWithPort,
	}
}
