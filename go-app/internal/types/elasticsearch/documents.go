package elasticsearch

type CreateDocumentResponse struct {
	IndexID    string `json:"_index,omitempty"`
	DocumentID string `json:"_id,omitempty"`
}

type PatchDocumentResponse struct {
	IndexID    string `json:"_index,omitempty"`
	DocumentID string `json:"_id,omitempty"`
}

type Document struct {
	ID          string                 `json:"_id,omitempty"`
	IndexID     string                 `json:"_index,omitempty"`
	Type        string                 `json:"_type,omitempty"`
	Version     int                    `json:"_version,omitempty"`
	SequenceNum int                    `json:"_seq_no,omitempty"`
	PrimaryTerm int                    `json:"_primary_term,omitempty"`
	Found       bool                   `json:"found,omitempty"`
	Source      map[string]interface{} `json:"_source,omitempty"`
}

type SearchDocumentResponseMetadata struct {
	Total int `json:"value,omitempty"`
}

type SearchDocumentResponseCore struct {
	Metadata SearchDocumentResponseMetadata `json:"total,omitempty"`
	Results  []Document                     `json:"hits,omitempty"`
}

type SearchDocumentResponse struct {
	Content SearchDocumentResponseCore `json:"hits,omitempty"`
}
