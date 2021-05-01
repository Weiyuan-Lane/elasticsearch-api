package responses

type Index struct {
	ID       string      `json:"id"`
	Aliases  interface{} `json:"aliases"`
	Mappings interface{} `json:"mappings"`
	Settings interface{} `json:"settings"`
}
