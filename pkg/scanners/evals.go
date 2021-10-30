package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

const evalName = "evals"

type Evals struct {
	visitor
	scorer
}

func NewEvals(score float64) *Evals {
	return &Evals{
		scorer: scorer{Step: score, name: evalName},
	}
}

// EnterNode is invoked at every node in hierarchy.
func (e *Evals) EnterNode(w walker.Walkable) bool {
	switch w.(node.Node).(type) {
	case *expr.Eval:
		e.scorer.Up()
	}

	return true
}
