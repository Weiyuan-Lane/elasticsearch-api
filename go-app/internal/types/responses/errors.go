package responses

type ErrorCode struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	ErrorCode ErrorCode `json:"error"`
}
