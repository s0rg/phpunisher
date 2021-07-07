package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

const arrcallName = "array-call"

type ArrayCall struct {
	prev node.Node
	stub
	scorer
}

func NewArrayCall(score float64) *ArrayCall {
	return &ArrayCall{
		scorer: scorer{Step: score, name: arrcallName},
	}
}

// EnterNode is invoked at every node in hierarchy.
func (a *ArrayCall) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)

	if !ok {
		return false
	}

	switch n.(type) {
	case *expr.ArrayDimFetch:
		if _, ok := a.prev.(*expr.FunctionCall); ok {
			a.scorer.Up()
		}
	default:
		a.prev = n
	}

	return true
}
