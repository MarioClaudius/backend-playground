package rest

import (
	"backend/go/rest/handler"

	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"
)

type APIServer struct {
	listenAddr string
	handler    *handler.Handlers
}

func NewAPIServer(listenAddr string, h *handler.Handlers) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		handler:    h,
	}
}

func (s *APIServer) Run() {
	router := httprouter.New()

	router.GET("/blogs/:id", s.handler.GetBlogByID)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}
