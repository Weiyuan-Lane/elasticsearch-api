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

func (s Service) ListDocument(
	indexID string,
	page int,
	perPage int,
	matchMap map[string]string,
	searchPropList []string,
	searchVal string,
	inputSortList [][2]string,
) (responses.DocumentPage, error) {

	parsedPage := page
	if parsedPage <= 0 {
		parsedPage = 1
	}

	parsedPerPage := perPage
	if parsedPerPage <= 0 || parsedPerPage > 50 {
		parsedPerPage = 10
	}

	searchResponse, err := s.ElasticsearchClient.ListDocuments(
		indexID,
		parsedPage,
		parsedPerPage,
		matchMap,
		searchPropList,
		searchVal,
		inputSortList,
	)

	if err != nil {
		return responses.DocumentPage{}, err
	}

	resPage := s.hydrateDocumentPage(searchResponse, page, perPage)

	return resPage, nil
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
	createdDocRes, err := s.ElasticsearchClient.CreateDocument(indexID, documentID, document)

	if err != nil {
		return responses.CreatedID{}, err
	}

	return responses.CreatedID{
		ID: createdDocRes.DocumentID,
	}, nil
}

func (s Service) PatchDocument(indexID, documentID string, document map[string]interface{}) error {
	_, err := s.ElasticsearchClient.PatchDocument(indexID, documentID, document)

	if err != nil {
		return err
	}

	return nil
}

func (s Service) hydrateDocument(document elasticsearch.Document) responses.Document {
	return responses.Document(document.Source)
}

func (s Service) hydrateDocuments(documents []elasticsearch.Document) []responses.Document {
	resDocuments := make([]responses.Document, len(documents))
	for i, document := range documents {
		resDocuments[i] = s.hydrateDocument(document)
	}

	return resDocuments
}

func (s Service) hydrateDocumentPage(searchResponse elasticsearch.SearchDocumentResponse, page, perPage int) responses.DocumentPage {
	total := searchResponse.Content.Metadata.Total

	return responses.DocumentPage{
		PageStats: responses.PageStats{
			Page:    page,
			PerPage: perPage,
			Total:   total,
		},
		Documents: s.hydrateDocuments(searchResponse.Content.Results),
	}

}
