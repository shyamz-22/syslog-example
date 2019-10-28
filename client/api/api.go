package api

import (
	"fmt"
	"github.com/shyamz-22/syslog-client/logger"
	"net/http"
	"net/http/httputil"
)

type Server struct {
	log *logger.Logger
}

func NewServer(l *logger.Logger) *Server {
	return &Server{log: l}
}

func (s *Server) Home(rw http.ResponseWriter, req *http.Request) {

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		s.log.Error(err)
	}
	s.log.Info(string(requestDump))

	_, _ = fmt.Fprint(rw, "UP!")
}

func (s *Server) Health(rw http.ResponseWriter, req *http.Request) {

	s.log.Debug("this is an Debug message")
	s.log.Info("this is an Info message")
	s.log.Warn("this is an Warn message")
	s.log.Error("this is an Error message")

	_, _ = fmt.Fprint(rw, "UP!")
}

func (s *Server) Fatal(rw http.ResponseWriter, req *http.Request) {

	s.log.Fatal("this is an Fatal message")

	_, _ = fmt.Fprint(rw, "UP!")
}
