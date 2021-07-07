package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
)

func TestSingleInclude(t *testing.T) {
	t.Parallel()

	builder := func() Scanner {
		return NewSingleInclude(1.0)
	}

	if builder().Name() != sincName {
		t.Fatal("invalid name")
	}

	singleInclude := []node.Node{&expr.Include{}}
	badCase := []node.Node{
		&node.Root{Stmts: singleInclude},
	}

	badCase = append(badCase, singleInclude...)

	notSingle := []node.Node{
		&expr.Include{},
		&expr.Empty{},
	}

	goodCase := []node.Node{
		&node.Root{Stmts: notSingle},
	}

	goodCase = append(goodCase, notSingle...)

	testCases := []testCase{
		{Nodes: []node.Node{&node.Root{}}},
		{Nodes: singleInclude},
		{Nodes: goodCase},
		{Nodes: badCase, Want: 1.0},
	}

	runCases(t, builder, testCases)
}
