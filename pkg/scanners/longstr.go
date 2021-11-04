package scanners

import (
	"strings"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

const (
	longstrName = "long-str"
)

type LongStrings struct {
	visitor
	scorer
	minLen int
}

func NewLongStrings(score float64, minLen int) *LongStrings {
	return &LongStrings{
		scorer: scorer{Step: score, name: longstrName},
	}
}

// EnterNode is invoked at every node in hierarchy.
func (ls *LongStrings) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		return false
	}

	switch n.(type) {
	case *scalar.String:
		s := w.(*scalar.String).Value

		if len(s) > ls.minLen && strings.Count(s, " ") == 0 {
			ls.scorer.Up()
		}
	}

	return true
}
