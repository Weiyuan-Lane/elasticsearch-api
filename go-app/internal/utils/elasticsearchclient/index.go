package elasticsearchclient

import (
	"fmt"

	elasticsearchtypes "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/elasticsearch"
	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/httprequest"
)

func (e ElasticSearchClient) GetIndex(id string) (elasticsearchtypes.IndexMap, error) {
	url := fmt.Sprintf(indexSingularPathTemplate, e.hostWithPort, id)
	response := elasticsearchtypes.IndexMap{}

	_, err := httprequest.GetJSON(
		url,
		map[string]string{},
		map[string]string{},
		&response,
	)

	return response, err
}

func (e ElasticSearchClient) CreateIndex(id string) (elasticsearchtypes.CreatedIndexResponse, error) {
	url := fmt.Sprintf(indexSingularPathTemplate, e.hostWithPort, id)
	response := elasticsearchtypes.CreatedIndexResponse{}

	_, err := httprequest.PutJSON(
		url,
		map[string]string{},
		map[string]string{},
		&response,
	)

	return response, err
}

func (e ElasticSearchClient) IndexExists(id string) (bool, error) {
	url := fmt.Sprintf(indexSingularPathTemplate, e.hostWithPort, id)

	statusCode, err := httprequest.HeadJSON(
		url,
		map[string]string{},
		map[string]string{},
		nil,
	)
	if err != nil {
		return false, err
	}

	if statusCode == 200 {
		return true, nil
	}
	if statusCode == 404 {
		return false, nil
	}

	return false, apperrors.ErrIndexExistsInvalidStatus
}
