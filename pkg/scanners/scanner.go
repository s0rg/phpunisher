package scanners

import (
	"github.com/z7zmey/php-parser/walker"
)

type Scanner interface {
	walker.Visitor
	Score() float64
	Name() string
}

type stub struct {
	name string
}

func (st *stub) Name() string                                 { return st.name }
func (st *stub) LeaveNode(n walker.Walkable)                  {}
func (st *stub) EnterChildNode(key string, w walker.Walkable) {}
func (st *stub) LeaveChildNode(key string, w walker.Walkable) {}
func (st *stub) EnterChildList(key string, w walker.Walkable) {}
func (st *stub) LeaveChildList(key string, w walker.Walkable) {}
