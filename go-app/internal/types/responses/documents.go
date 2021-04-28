package responses

type Document map[string]interface{}
type DocumentPage struct {
	PageStats PageStats  `json:"page_stats"`
	Documents []Document `json:"results"`
}
