package server

import (
	"log"
	"net/http"

	"github.com/DanielDanteDosSantosViana/darth_vader/routes"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen(port string) {
	s.init()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error init Server : ", err)
	}
}

func (s *Server) init() {
	routes.StartRoutes()
}
