package checker

import "fmt"

type Server struct {
	Name, Domain string
}

func (s *Server) String() string {
	return fmt.Sprintf("Name: %s, Domain: %s", s.Name, s.Domain)
}

func (s *Server) SchemaDomain() string {
	return "http://" + s.Domain
}