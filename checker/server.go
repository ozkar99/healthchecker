package checker

import (
	"fmt"
	"strings"
)

type Server struct {
	Name, Domain string
	Error error
}

func (s *Server) String() string {
	return fmt.Sprintf("Name: %s, Domain: %s, Error: %s", s.Name, s.Domain, s.Error)
}

func (s *Server) SchemaDomain() string {
	prefix := "http://"
	
	if strings.HasPrefix(s.Domain, prefix) {
		return s.Domain
	}
	 
	return prefix + s.Domain
}