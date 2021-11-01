package server

import (
	"net"
	"net/ftp"
	"net/http"
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
	// cfger configer
	srv    *ftp.Server
	strgr  fileManager
	client *http.Client
}

func New(cfger configer, strger fileManager) *Server {
	srv := new(http.Server)

	srv.Addr = net.JoinHostPort(cfger.Host(), cfger.Port())

	return &Server{
		srv: &http.Server{},
	}
}
