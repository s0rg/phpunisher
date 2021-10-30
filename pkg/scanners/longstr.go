package scanners

import (
	"strings"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

const (
	longstrName = "long-str"
	minCheckLen = 64
)

type LongStrings struct {
	visitor
	scorer
}

func NewLongStrings(score float64) *LongStrings {
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
		s, ok := w.(*scalar.String)

		if !ok {
			return false
		}

		if len(s.Value) > minCheckLen && strings.Count(s.Value, " ") == 0 {
			ls.scorer.Up()
		}
	}

	return true
}
