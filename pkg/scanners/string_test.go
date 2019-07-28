package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
)

func TestBadString(t *testing.T) {

	builder := func() Scanner {
		return NewBadString(1.0)
	}

	if builder().Name() != badstrName {
		t.Fatal("invalid name")
	}

	testCases := []testCase{
		{Nodes: []node.Node{scalar.NewString("hello")}},
		{Nodes: []node.Node{scalar.NewString("hello\\")}},
		{Nodes: []node.Node{scalar.NewString("hello\\\\")}},
		{Nodes: []node.Node{scalar.NewString("hello\\\\\\")}, Want: 1.0},
	}

	runCases(t, builder, testCases)
}
