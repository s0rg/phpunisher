package scanners

import (
	"testing"
)

func TestScorer(t *testing.T) {
	s := scorer{Step: 1.0}

	for i := 1; i < 100; i++ {
		s.Up()
		if s.Score() < float64(i) {
			t.Fatal("scorer - not up")
		}
	}
}
