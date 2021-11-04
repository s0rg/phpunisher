package scanners

import (
	"testing"
)

func TestScorerStep(t *testing.T) {
	t.Parallel()

	s := scorer{Step: 1.0}

	for i := 1; i < 100; i++ {
		s.Up()

		if s.Score() < float64(i) {
			t.Error("scorer - not up")
		}
	}
}

func TestScorerName(t *testing.T) {
	t.Parallel()

	n := "test-name"

	s := scorer{name: n}
	if s.Name() != n {
		t.Error("unexpected name")
	}
}
