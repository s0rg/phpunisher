package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

const evalName = "evals"

// Evals finds eval expression in code.
type Evals struct {
	visitor
	scorer
}

// NewEvals creates new Evals scanner.
func NewEvals(score float64) *Evals {
	return &Evals{
		scorer: scorer{Step: score, name: evalName},
	}
}

// EnterNode is invoked at every node in hierarchy.
func (e *Evals) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		return false
	}

	switch n.(type) {
	case *expr.Eval:
		e.scorer.Up()
	}

	return true
}
