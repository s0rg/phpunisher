package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

const sincName = "single-include"

type SingleInclude struct {
	root *node.Root
	stub
	scorer
}

func NewSingleInclude(score float64) *SingleInclude {
	return &SingleInclude{
		scorer: scorer{Step: score, name: sincName},
	}
}

// EnterNode is invoked at every node in hierarchy
func (s *SingleInclude) EnterNode(w walker.Walkable) bool {
	switch w.(node.Node).(type) {
	case *node.Root:
		s.root = w.(*node.Root)
	case *expr.Include:
		if s.root != nil && len(s.root.Stmts) == 1 {
			s.scorer.Up()
		}
	}

	return true
}
