package main

import (
	"math"
	"sort"
	"testing"
)

func TestScore(t *testing.T) {
	t.Parallel()

	sc := scores{
		{Scanner: "3", Score: 3.0},
		{Scanner: "1", Score: 1.0},
		{Scanner: "2", Score: 2.0},
	}

	sort.Sort(sc)

	if sc.Len() != 3 {
		t.Fatal("unexpected len:", sc.Len())
	}

	if sc[0].Scanner != "1" {
		t.Fatal("unexpected 0 item:", sc[0])
	}

	if sc[2].Scanner != "3" {
		t.Fatal("unexpected 2 item:", sc[2])
	}

	const (
		wantSum = 6.0
		epsilon = 0.00001
	)

	if sum := sc.Sum(); math.Abs(wantSum-sum) > epsilon {
		t.Fatal("unexpected sum:", sum)
	}
}
