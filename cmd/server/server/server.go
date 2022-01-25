package server

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type configer interface {
	Host() string
	Port() string
}

type fileManager interface {
	Create(path string) error
	Open(path string) error
}

type Server struct {
	strgr  fileManager
	client *http.Client
}

func New(cfger configer, strger fileManager) *Server {
	s := new(Server)

	srv := new(http.Server)

	srv.Addr = net.JoinHostPort(cfger.Host(), cfger.Port())

	srv.Handler = makeHandler(s)

	return s
}

func makeHandler(s *Server) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("", s.user)

	return r
}
