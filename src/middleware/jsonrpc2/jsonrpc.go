package jsonrpc2

import (
	"encoding/json"
	"net/http"
)

const (
	ErrCodeInvalidParams = 40001
	ErrCodeNotFaund      = 40401
	ErrCodeInternal      = 50001
)

type Server struct {
	handlers map[string]Handler
}

type Handler interface {
	ParseParams(*json.RawMessage) (interface{}, error)
	Exec() interface{}
}

func (s *Server) Register(name string, handler Handler) {

}

func (s *Server) handleSingleRequest() {

}

func (s *Server) hundleBatchRequest() {

}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {

}
