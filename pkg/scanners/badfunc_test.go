package scanners

import (
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
)

func buildFuncNodes(fn string) []node.Node {
	return []node.Node{
		&expr.FunctionCall{
			Function: &name.Name{
				Parts: []node.Node{
					name.NewNamePart(fn),
				},
			},
		},
	}
}

func TestBadFunc(t *testing.T) {
	builder := func() Scanner {
		return NewBadFunc(1.0)
	}

	if builder().Name() != bfName {
		t.Fatal("invalid name")
	}

	testCases := []testCase{
		{Nodes: []node.Node{&node.Root{}}},
		{Nodes: buildFuncNodes("foo")},
		{Nodes: buildFuncNodes("php_info")},
		{Nodes: buildFuncNodes("str_rot13"), Want: 1.0},
		{Nodes: buildFuncNodes("base64_decode"), Want: 1.0},
	}

	runCases(t, builder, testCases)
}
