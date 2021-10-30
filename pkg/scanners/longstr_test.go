package scanners

import (
	"strings"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
)

func TestBadLongString(t *testing.T) {
	t.Parallel()

	builder := func() Scanner {
		return NewLongStrings(1.0)
	}

	if builder().Name() != longstrName {
		t.Fatal("invalid name")
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
