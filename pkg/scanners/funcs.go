package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/walker"

	"github.com/s0rg/phpunisher/pkg/set"
)

const bfName = "funcs"

// FuncsBlacklist finds blacklisted functions calls.
type FuncsBlacklist struct {
	visitor
	scorer
	list set.Strings
}

// NewFuncsBlacklist creates new FuncsBlacklist scanner.
func NewFuncsBlacklist(score float64, list []string) *FuncsBlacklist {
	bf := &FuncsBlacklist{
		scorer: scorer{Step: score, name: bfName},
		list:   make(set.Strings),
	}

	bf.list.FromList(list)

	return bf
}

// EnterNode is invoked at every node in hierarchy.
func (bf *FuncsBlacklist) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		return false
	}

	switch n.(type) {
	case *expr.FunctionCall:
		if nm, ok := w.(*expr.FunctionCall).Function.(*name.Name); ok {
			for _, p := range nm.Parts {
				if np, ok := p.(*name.NamePart); ok {
					if bf.list.Has(np.Value) {
						bf.scorer.Up()
					}
				}
			}
		}
	}

	return true
}
