package scanners

import (
	"github.com/z7zmey/php-parser/walker"
)

type Scanner interface {
	walker.Visitor
	Score() float64
	Name() string
}

type stub struct{}

func (s *stub) LeaveNode(n walker.Walkable)                  {}
func (s *stub) EnterChildNode(key string, w walker.Walkable) {}
func (s *stub) LeaveChildNode(key string, w walker.Walkable) {}
func (s *stub) EnterChildList(key string, w walker.Walkable) {}
func (s *stub) LeaveChildList(key string, w walker.Walkable) {}
