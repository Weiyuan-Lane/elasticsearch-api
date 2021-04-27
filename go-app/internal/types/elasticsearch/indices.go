package elasticsearch

type Index struct {
	Aliases  interface{} `json:"aliases,omitempty"`
	Mappings interface{} `json:"mappings,omitempty"`
	Settings interface{} `json:"settings,omitempty"`
}

type IndexMap map[string]Index

type CreatedIndexResponse struct {
	Acknowledged       bool   `json:"acknowledged"`
	ShardsAcknowledged bool   `json:"shards_acknowledged"`
	IndexID            string `json:"index"`
}
