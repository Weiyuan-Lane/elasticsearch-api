package httptransport

import "github.com/gorilla/mux"

func (h HttpServer) registerServices(router *mux.Router) {
	h.registerRoutes(
		router,
	)
}
