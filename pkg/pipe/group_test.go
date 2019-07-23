package pipe

import (
	"testing"
)

func runGroupN(n int) (count int, closed bool) {
	ch_in := make(chan struct{}, n)
	ch_out := make(chan struct{}, n)

	action := func() {
		for _ = range ch_in {
			ch_out <- struct{}{}
		}
	}

	closer := func() {
		close(ch_in)
		closed = true
	}

	g := group{
		Workers: n,
		Action:  action,
	}
	g.Start(closer)

	for i := 0; i < n; i++ {
		ch_in <- struct{}{}
	}

	g.Wait()

	for _ = range ch_out {
		if count++; count == n {
			break
		}
	}
	close(ch_out)

	return
}

func TestGroup(t *testing.T) {

	tests := []int{1, 8, 16, 32}

	for i := 0; i < len(tests); i++ {
		n := tests[i]
		rv, ok := runGroupN(n)
		if !ok {
			t.Fatalf("test failed for %d - not ok", n)
		}
		if rv != n {
			t.Fatalf("test failed for %d - got %d as result", n, rv)
		}
	}
}
