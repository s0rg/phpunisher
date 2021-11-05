package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/walker"
)

const (
	badstrName = "escapes"
)

// Escapes finds too many escapes in string literals.
type Escapes struct {
	visitor
	scorer
}

// NewEscapes creates new Escapes scanner.
func NewEscapes(score float64) *Escapes {
	return &Escapes{
		scorer: scorer{Step: score, name: badstrName},
	}
}

// countBadEscapes finds number or escaped symbols in string, that are not in (\n, \r, \t).
func countBadEscapes(s string) (result int) {
	var afterSlash bool

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '\\':
			afterSlash = true
		case 'n', 'r', 't':
			// skip lf, cr, tab
		default:
			if afterSlash {
				afterSlash = false

				result++
			}
		}
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (bs *Escapes) scoreUp(count int) {
	var ups int

	switch {
	case count < 10:
		ups = 1
	case count < 25:
		ups = 2
	case count < 50:
		ups = 3
	case count < 100:
		ups = 4
	default:
		ups = 5 + min(count/100, 5)
	}

	for i := 0; i < ups; i++ {
		bs.scorer.Up()
	}
}

// EnterNode is invoked at every node in hierarchy.
func (bs *Escapes) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		return false
	}

	switch n.(type) {
	case *scalar.String:
		if bad := countBadEscapes(w.(*scalar.String).Value); bad > 0 {
			bs.scoreUp(bad)
		}
	}

	return true
}
