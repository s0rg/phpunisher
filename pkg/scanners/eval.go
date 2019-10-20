package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/walker"
)

const evalName = "eval-expr"

type EvalExpr struct {
	stub
	scorer
}

func NewEvalExpr(score float64) *EvalExpr {
	return &EvalExpr{
		scorer: scorer{Step: score, name: evalName},
	}
}

// EnterNode is invoked at every node in hierarchy
func (e *EvalExpr) EnterNode(w walker.Walkable) bool {
	switch w.(node.Node).(type) {
	case *expr.Eval:
		e.scorer.Up()
	}

	return true
}
