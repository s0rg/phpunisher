package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

func TestEval(t *testing.T) {
	t.Parallel()

	builder := func() Scanner {
		return NewEvals(1.0)
	}

	if builder().Name() != evalName {
		t.Error("invalid name")
	}

	testCases := []testCase{
		{Nodes: []node.Node{&expr.Empty{}}},
		{
			Nodes: []node.Node{
				&expr.ArrayDimFetch{},
				&expr.Eval{},
				&expr.FunctionCall{},
			},
			Want: 1.0,
		},
	}

	runCases(t, builder, testCases)
}

func TestEvalsBadValue(t *testing.T) {
	t.Parallel()

	s := NewEvals(1.0)

	if s.EnterNode(&nonNode{}) {
		t.Error("enters bad node")
	}
}
