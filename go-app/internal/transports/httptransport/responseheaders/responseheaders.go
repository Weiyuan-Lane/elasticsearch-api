package responseheaders

import "net/http"

func SetJSONResponseType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
