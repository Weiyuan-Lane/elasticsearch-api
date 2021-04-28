package documents

import (
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/elasticsearch"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/responses"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/elasticsearchclient"
	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/logger"
)

type Service struct {
	Logger logger.Logger

	ElasticsearchClient elasticsearchclient.ElasticSearchClient
}

func (s Service) RetrieveDocument(indexID, documentID string) (responses.Document, error) {
	document, err := s.ElasticsearchClient.RetrieveDocument(indexID, documentID)

	if err != nil {
		return responses.Document{}, err
	}

	resDocument := s.hydrateDocument(document)

	return resDocument, nil
}

func (s Service) CreateDocument(indexID, documentID string, document map[string]interface{}) (responses.CreatedID, error) {
	documentWithID := document
	documentWithID["id"] = documentID
	createdDocRes, err := s.ElasticsearchClient.CreateDocument(indexID, documentID, document)

	if err != nil {
		return responses.CreatedID{}, err
	}

	return responses.CreatedID{
		ID: createdDocRes.DocumentID,
	}, nil
}

func (s Service) hydrateDocument(document elasticsearch.Document) responses.Document {
	return responses.Document(document.Source)
}
