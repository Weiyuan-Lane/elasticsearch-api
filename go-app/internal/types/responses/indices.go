package responses

type Index struct {
	Aliases  interface{} `json:"aliases"`
	Mappings interface{} `json:"mappings"`
	Settings interface{} `json:"settings"`
}
