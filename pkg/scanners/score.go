package scanners

// Score holds scanner name and total no-zero score.
type Score struct {
	Scanner string
	Score   float64
}

// Scores holds bunch of Score.
type Scores []*Score

func (s Scores) Len() int           { return len(s) }
func (s Scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Scores) Less(i, j int) bool { return s[i].Score < s[j].Score }

// Sum returs overall scores sum.
func (s Scores) Sum() (rv float64) {
	for i := 0; i < len(s); i++ {
		rv += s[i].Score
	}

	return
}
