package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/walker"
)

const bfName = "bad-func"

var (
	badFuncs = []string{
		"str_rot13",
		"base64_decode",
	}
)

type BadFunc struct {
	stub
	scorer
}

func NewBadFunc(score float64) *BadFunc {
	return &BadFunc{
		stub:   stub{bfName},
		scorer: scorer{Step: score},
	}
}

func isBadFunc(name string) bool {
	for i := 0; i < len(badFuncs); i++ {
		if name == badFuncs[i] {
			return true
		}
	}
	return false
}

// EnterNode is invoked at every node in hierarchy
func (bf *BadFunc) EnterNode(w walker.Walkable) bool {

	switch w.(node.Node).(type) {
	case *expr.FunctionCall:
		fc := w.(*expr.FunctionCall)
		if n, ok := fc.Function.(*name.Name); ok {
			for _, p := range n.Parts {
				if np, ok := p.(*name.NamePart); ok && isBadFunc(np.Value) {
					bf.scorer.Up()
				}
			}
		}
	}
	return true
}
