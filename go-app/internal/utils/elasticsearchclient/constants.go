package elasticsearchclient

const (
	fullHostTemplate = "%s://%s:%s"

	indexSingularPathTemplate = "%s/%s"

	createDocumentPathTemplate   = "%s/%s/_create/%s"
	documentSingularPathTemplate = "%s/%s/_doc/%s"
	documentSearchPathTemplate   = "%s/%s/_search"
)
