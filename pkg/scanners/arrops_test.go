package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

func TestArrayOperations(t *testing.T) {
	builder := func() Scanner {
		return NewArrayOperations(1.0)
	}

	if builder().Name() != arropsName {
		t.Fatal("invalid name")
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
