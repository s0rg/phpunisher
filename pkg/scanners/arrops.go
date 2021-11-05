package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

const (
	arropsName = "array-ops"
)

// ArrayOperations finds too many array operations in file.
type ArrayOperations struct {
	visitor
	step    float64
	maxRate float64
	arrOps  int
	ops     int
}

// NewArrayOperations creates new ArrayOperations scanner.
func NewArrayOperations(score, rate float64) *ArrayOperations {
	return &ArrayOperations{
		step:    score,
		maxRate: rate,
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
	if rate > a.maxRate {
		return a.step * ((rate - a.maxRate) * 10.0)
	}

	return 0
}

func (a *ArrayOperations) Name() string {
	return arropsName
}
