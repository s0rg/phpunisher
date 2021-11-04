package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

func TestArrayOperations(t *testing.T) {
	t.Parallel()

	builder := func() Scanner {
		return NewArrayOperations(1.0, 0.2)
	}

	if builder().Name() != arropsName {
		t.Error("invalid name")
	}

	testCases := []testCase{
		{Nodes: []node.Node{&expr.Empty{}}},
		{
			Nodes: []node.Node{
				&expr.ArrayDimFetch{},
				&expr.ArrayDimFetch{},
				&expr.FunctionCall{},
			},
			Want: 1.0,
		},
	}

	runCases(t, builder, testCases)
}

func TestArrayOpsBadValue(t *testing.T) {
	t.Parallel()

	s := NewArrayOperations(1.0, 0.2)

	if s.EnterNode(&nonNode{}) {
		t.Error("enters bad node")
	}
}
