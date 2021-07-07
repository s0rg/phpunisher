package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

const (
	arropsName            = "array-operations"
	maxArrOpsRate float64 = 0.2
)

type ArrayOperations struct {
	stub
	step   float64
	arrOps int
	ops    int
}

func NewArrayOperations(score float64) *ArrayOperations {
	return &ArrayOperations{
		step: score,
	}
}

// EnterNode is invoked at every node in hierarchy.
func (a *ArrayOperations) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)

	if !ok {
		return false
	}

	switch n.(type) {
	case *expr.ArrayDimFetch:
		a.arrOps++
	case *scalar.Lnumber, *expr.Variable, *node.Identifier:
		// skip declarations
	default:
		a.ops++
	}

	return true
}

func (a *ArrayOperations) Score() float64 {
	rate := float64(a.arrOps) / float64(a.ops)
	if rate > maxArrOpsRate {
		return a.step * ((rate - maxArrOpsRate) * 10.0)
	}

	return 0
}

func (a *ArrayOperations) Name() string {
	return arropsName
}
