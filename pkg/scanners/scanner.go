package scanners

import (
	"github.com/z7zmey/php-parser/walker"
)

// Scanner is what every scanner must do.
type Scanner interface {
	walker.Visitor
	Score() float64
	Name() string
}

type visitor struct{}

func (v *visitor) LeaveNode(walker.Walkable)              {}
func (v *visitor) EnterChildNode(string, walker.Walkable) {}
func (v *visitor) LeaveChildNode(string, walker.Walkable) {}
func (v *visitor) EnterChildList(string, walker.Walkable) {}
func (v *visitor) LeaveChildList(string, walker.Walkable) {}
