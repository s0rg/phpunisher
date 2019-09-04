package scanners

type scorer struct {
	Step  float64
	score float64
	name  string
}

func (sc *scorer) Up()            { sc.score += sc.Step }
func (sc *scorer) Score() float64 { return sc.score }
func (sc *scorer) Name() string   { return sc.name }
