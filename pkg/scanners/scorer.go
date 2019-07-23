package scanners

type scorer struct {
	Step  float64
	score float64
}

func (sc *scorer) Up()            { sc.score += sc.Step }
func (sc *scorer) Score() float64 { return sc.score }
