package requestbodies

type CreateIndexBody struct {
	ID string `json:"id"`
}

type CreateDocumentBody map[string]interface{}

type PatchDocumentBody map[string]interface{}

type SearchDocumentBody struct {
	MatchMap       map[string]string `json:"match,omitempty"`
	SearchPropList []string          `json:"search_fields,omitempty"`
	SearchVal      string            `json:"search_value,omitempty"`
	InputSortList  []InputSortField  `json:"sort,omitempty"`
}

type InputSortField struct {
	Property string `json:"property,omitempty"`
	Order    string `json:"order,omitempty"`
}
