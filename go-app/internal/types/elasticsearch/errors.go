package elasticsearch

type ErrorCorePayload struct {
	RootCause interface{} `json:"root_cause,omitempty"`
	Type      string      `json:"type,omitempty"`
	Reason    string      `json:"reason,omitempty"`
	IndexID   string      `json:"index,omitempty"`
}

type ErrorResponse struct {
	ErrorCore ErrorCorePayload `json:"error,omitempty"`
	Status    int              `json:"status,omitempty"`
}
