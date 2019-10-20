package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

type testCase struct {
	Nodes []node.Node
	Want  float64
}

func runCases(t *testing.T, builder func() Scanner, cases []testCase) {
	for _, tc := range cases {
		s := builder()

		for _, n := range tc.Nodes {
			if !s.EnterNode(n) {
				t.Fatalf("EnterNode returns false for: %v", n)
			}
		}

		if score := s.Score(); score < tc.Want {
			t.Fatalf("failed on case: %+v score: %.1f", tc, score)
		}
	}
}

func TestArrayCall(t *testing.T) {
	builder := func() Scanner {
		return NewArrayCall(1.0)
	}

	if builder().Name() != arrcallName {
		t.Fatal("invalid name")
	}

	testCases := []testCase{
		{Nodes: []node.Node{&expr.Empty{}}},
		{Nodes: []node.Node{&expr.ArrayDimFetch{}, &expr.FunctionCall{}}},
		{Nodes: []node.Node{&expr.FunctionCall{}, &expr.ArrayDimFetch{}}, Want: 1.0},
		{
			Nodes: []node.Node{
				&expr.FunctionCall{},
				&expr.ArrayDimFetch{},
				&expr.Empty{},
				&expr.FunctionCall{},
				&expr.ArrayDimFetch{},
			},
			Want: 2.0,
		},
	}

	runCases(t, builder, testCases)
}
