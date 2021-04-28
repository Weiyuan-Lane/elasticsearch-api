package elasticsearchclient

import (
	"fmt"
	"regexp"

	elasticsearchtypes "github.com/weiyuan-lane/elasticsearch-api/go-app/internal/types/elasticsearch"

	"github.com/weiyuan-lane/elasticsearch-api/go-app/internal/utils/httprequest"
)

var (
	luceneEscapeCharRegex = regexp.MustCompile(`(\&|\||\!|\(|\)|\{|\}|\[|\]|\^|\"|\~|\*|\?|\:)`)
)

func (e ElasticSearchClient) ListDocuments(
	indexID string,
	page int,
	perPage int,
	matchMap map[string]string,
	searchPropList []string,
	searchVal string,
	inputSortList [][2]string,
) (elasticsearchtypes.SearchDocumentResponse, error) {
	url := fmt.Sprintf(documentSearchPathTemplate, e.hostWithPort, indexID)
	response := elasticsearchtypes.SearchDocumentResponse{}

	offset := (page - 1) * perPage
	limit := perPage

	body := map[string]interface{}{
		"from": offset,
		"size": limit,
	}

	queryMap := e.makeQueryMap(matchMap, searchPropList, searchVal)
	if queryMap != nil {
		body["query"] = queryMap
	}

	sortConditions := e.makeSortConditions(inputSortList)
	if len(sortConditions) > 0 {
		body["sort"] = sortConditions
	}

	_, err := httprequest.GetJSON(
		url,
		map[string]string{},
		map[string]string{},
		body,
		&response,
		nil,
	)

	if err != nil {
		return elasticsearchtypes.SearchDocumentResponse{}, err
	}

	return response, nil
}

func (e ElasticSearchClient) CreateDocument(indexID string, documentID string, document map[string]interface{}) (elasticsearchtypes.CreateDocumentResponse, error) {
	url := fmt.Sprintf(createDocumentPathTemplate, e.hostWithPort, indexID, documentID)
	response := elasticsearchtypes.CreateDocumentResponse{}
	errResponse := elasticsearchtypes.ErrorResponse{}

	_, err := httprequest.PostJSON(
		url,
		map[string]string{},
		map[string]string{},
		document,
		&response,
		&errResponse,
	)

	if err != nil {
		return elasticsearchtypes.CreateDocumentResponse{}, err
	}

	return response, nil
}

func (e ElasticSearchClient) RetrieveDocument(indexID string, documentID string) (elasticsearchtypes.Document, error) {
	url := fmt.Sprintf(documentSingularPathTemplate, e.hostWithPort, indexID, documentID)
	response := elasticsearchtypes.Document{}

	_, err := httprequest.GetJSON(
		url,
		map[string]string{},
		map[string]string{},
		nil,
		&response,
		nil,
	)

	if err != nil {
		return elasticsearchtypes.Document{}, err
	}

	return response, nil
}

func (e ElasticSearchClient) escapeQueryString(query string) string {
	return luceneEscapeCharRegex.ReplaceAllString(query, `\$1`)
}

func (e ElasticSearchClient) makeQueryMap(
	matchMap map[string]string, // Property must match
	searchPropList []string, // Property can match part of in full (more score if match)
	searchVal string,

) map[string]map[string]interface{} {

	finalBoolConditions := map[string]interface{}{}
	matchConditions := []map[string]map[string]string{}
	searchConditions := []map[string]map[string]string{}

	// For getting match conditions
	if len(matchMap) != 0 {
		for key, matchVal := range matchMap {
			matchKey := fmt.Sprintf("%s.keyword", key)
			currCondition := map[string]map[string]string{
				"term": {
					matchKey: matchVal,
				},
			}

			matchConditions = append(searchConditions, currCondition)
		}
	}

	// For getting search conditions
	if len(searchPropList) != 0 && searchVal != "" {
		escapedSearchVal := fmt.Sprintf("*%s*", e.escapeQueryString(searchVal))

		for _, prop := range searchPropList {
			searchKey := fmt.Sprintf("%s.keyword", prop)
			currCondition := map[string]map[string]string{
				"wildcard": {
					searchKey: escapedSearchVal,
				},
			}

			searchConditions = append(searchConditions, currCondition)
		}
	}

	if len(searchConditions) == 0 && len(matchConditions) == 0 {
		return nil
	}

	if len(matchConditions) != 0 {
		finalBoolConditions["filter"] = matchConditions
	}

	if len(searchConditions) != 0 {
		finalBoolConditions["must"] = []map[string]map[string]interface{}{
			{
				"bool": {
					"should":               searchConditions,
					"minimum_should_match": 1,
				},
			},
		}
	}

	return map[string]map[string]interface{}{
		"bool": finalBoolConditions,
	}
}

func (e ElasticSearchClient) makeSortConditions(
	sortList [][2]string,
) []map[string]map[string]string {

	if len(sortList) == 0 {
		return nil
	}

	sortConditions := make([]map[string]map[string]string, len(sortList))
	for i, pair := range sortList {
		key, sortOrder := pair[0], pair[1]
		sortKey := fmt.Sprintf("%s.keyword", key)

		sortConditions[i] = map[string]map[string]string{
			sortKey: map[string]string{
				"order": sortOrder,
			},
		}
	}

	return sortConditions
}
