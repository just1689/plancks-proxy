package v2

import (
	"github.com/just1689/plancks-proxy/config"
	"sync"
)

var server = NewServer()

func NewServer() *Server {
	result := &Server{}
	return result
}

type Server struct {
	sync.Mutex
	Config config.Config
}

func (s *Server) UpdateConfig(c config.Config) {
	s.Lock()
	defer s.Unlock()
	s.Config = c
	s.NotifyChange()
}

func (s *Server) NotifyChange() {
	//TODO: implement
}
