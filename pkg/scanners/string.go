package scanners

import (
	"strings"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

const (
	badstrName = "bad-string"
	maxStrEsc  = 2
)

type BadString struct {
	stub
	scorer
}

func NewBadString(score float64) *BadString {
	return &BadString{
		stub:   stub{badstrName},
		scorer: scorer{Step: score},
	}
}

func isBadString(s string) bool {
	return strings.Count(s, "\\") > maxStrEsc
}

// EnterNode is invoked at every node in hierarchy
func (bs *BadString) EnterNode(w walker.Walkable) bool {
	switch w.(node.Node).(type) {
	case *scalar.String:
		s := w.(*scalar.String)
		if isBadString(s.Value) {
			bs.scorer.Up()
		}
	}
	return true
}
