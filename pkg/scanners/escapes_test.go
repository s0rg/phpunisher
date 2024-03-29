package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
)

func TestEscapes(t *testing.T) {
	t.Parallel()

	builder := func() Scanner {
		return NewEscapes(1.0)
	}

	if builder().Name() != badstrName {
		t.Error("invalid name")
	}

	testCases := []testCase{
		{Nodes: []node.Node{scalar.NewString("hello")}},
		{Nodes: []node.Node{scalar.NewString("hello\\")}},
		{Nodes: []node.Node{scalar.NewString("hello\\\\")}},
		{Nodes: []node.Node{scalar.NewString("hello\r\n\\\\")}},
		{Nodes: []node.Node{scalar.NewString("hello\\\\foo\\")}, Want: 1.0},
	}

	runCases(t, builder, testCases)
}

func TestEscapesScore(t *testing.T) {
	t.Parallel()

	type scoreCase struct {
		Bads int
		Want float64
	}

	testCases := []scoreCase{
		{Bads: 5, Want: 1.0},
		{Bads: 9, Want: 1.0},
		{Bads: 14, Want: 2.0},
		{Bads: 23, Want: 2.0},
		{Bads: 26, Want: 3.0},
		{Bads: 45, Want: 3.0},
		{Bads: 51, Want: 4.0},
		{Bads: 99, Want: 4.0},
		{Bads: 100, Want: 6.0},
		{Bads: 666, Want: 10.0},
	}

	for _, tc := range testCases {
		s := NewEscapes(1.0)
		s.scoreUp(tc.Bads)

		if s.Score() < tc.Want {
			t.Errorf("case failed: %d want: %.2f got: %.2f", tc.Bads, tc.Want, s.Score())
		}
	}
}

func TestEscapesBadValue(t *testing.T) {
	t.Parallel()

	s := NewEscapes(1.0)

	if s.EnterNode(&nonNode{}) {
		t.Error("enters bad node")
	}
}
