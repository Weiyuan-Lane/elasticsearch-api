package indices

import (
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/elasticsearch"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/elasticsearchclient"
	apperrors "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/errors"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/logger"
)

type Service struct {
	Logger logger.Logger

	ElasticsearchClient elasticsearchclient.ElasticSearchClient
}

func (s Service) RetrieveIndex(id string) (responses.Index, error) {
	indexExists, err := s.ElasticsearchClient.IndexExists(id)
	if err != nil {
		return responses.Index{}, err
	}
	if !indexExists {
		return responses.Index{}, apperrors.ErrIndexNotFound
	}

	indexMap, err := s.ElasticsearchClient.GetIndex(id)

	if err != nil {
		return responses.Index{}, err
	}

	resIndex := s.hydrateIndex(id, indexMap)
	return resIndex, nil
}

func (s Service) CreateIndex(id string) (responses.CreatedID, error) {
	indexExists, err := s.ElasticsearchClient.IndexExists(id)
	if err != nil {
		return responses.CreatedID{}, err
	}
	if indexExists {
		return responses.CreatedID{}, apperrors.ErrIndexAlreadyCreated
	}

	createdIndexRes, err := s.ElasticsearchClient.CreateIndex(id)

	if err != nil {
		return responses.CreatedID{}, err
	}

	if !createdIndexRes.Acknowledged {
		return responses.CreatedID{}, apperrors.ErrCreateIndexNotAcknowledged
	} else if createdIndexRes.IndexID != id {
		return responses.CreatedID{}, apperrors.ErrCreateIndexInvalidID
	}

	return responses.CreatedID{
		ID: createdIndexRes.IndexID,
	}, nil
}

func (s Service) hydrateIndex(id string, indexMap elasticsearch.IndexMap) responses.Index {
	if val, ok := indexMap[id]; ok {
		return responses.Index{
			Aliases:  val.Aliases,
			Mappings: val.Mappings,
			Settings: val.Settings,
		}
	}

	return responses.Index{}
}
