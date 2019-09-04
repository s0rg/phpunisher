package scanners

import (
	"testing"
)

func TestScorerStep(t *testing.T) {
	s := scorer{Step: 1.0}

	for i := 1; i < 100; i++ {
		s.Up()
		if s.Score() < float64(i) {
			t.Fatal("scorer - not up")
		}
	}
}

func TestScorerName(t *testing.T) {
	n := "test-name"
	s := scorer{name: n}
	if s.Name() != n {
		t.Fatal("unexpected name")
	}
}
