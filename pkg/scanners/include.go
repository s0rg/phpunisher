package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

const sincName = "include"

// SingleInclude finds files with single operation - include.
type SingleInclude struct {
	root *node.Root
	visitor
	scorer
}

// NewSingleInclude creates new SingleInclude scanner.
func NewSingleInclude(score float64) *SingleInclude {
	return &SingleInclude{
		scorer: scorer{Step: score, name: sincName},
	}
}

// EnterNode is invoked at every node in hierarchy.
func (s *SingleInclude) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		return false
	}

	switch n.(type) {
	case *node.Root:
		s.root = w.(*node.Root)
	case *expr.Include:
		if s.root != nil && len(s.root.Stmts) == 1 {
			s.scorer.Up()
		}
	}

	return true
}
