package backend

import (
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Run(port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return server.httpServer.ListenAndServe()
}
