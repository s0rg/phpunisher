package scanners

import (
	"strings"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
)

func TestLongStrings(t *testing.T) {
	t.Parallel()

	const minCheckLen = 32

	builder := func() Scanner {
		return NewLongStrings(1.0, minCheckLen)
	}

	if builder().Name() != longstrName {
		t.Error("invalid name")
	}

	var ballast = strings.Repeat("A", minCheckLen-1)

	testCases := []testCase{
		{Nodes: []node.Node{scalar.NewString(ballast)}},
		{Nodes: []node.Node{scalar.NewString(ballast + "A")}},
		{Nodes: []node.Node{scalar.NewString(ballast + " " + ballast)}},
		{Nodes: []node.Node{scalar.NewString(ballast + ballast)}, Want: 1.0},
	}

	runCases(t, builder, testCases)
}

func TestLongStringsBadValue(t *testing.T) {
	t.Parallel()

	s := NewLongStrings(1.0, 1)

	if s.EnterNode(&nonNode{}) {
		t.Error("enters bad node")
	}
}
