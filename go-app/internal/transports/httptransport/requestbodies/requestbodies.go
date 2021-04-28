package requestbodies

type CreateIndexBody struct {
	ID string `json:"id"`
}

type CreateDocumentBody map[string]interface{}
