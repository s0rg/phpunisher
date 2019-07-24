package pipe

import (
	"testing"
)

func runGroupN(n int) (count int, closed bool) {
	ch := make(chan struct{}, n)
	g := group{
		Workers: n,
		Action:  func() { ch <- struct{}{} },
	}
	g.Start(func() { closed = true })
	g.Wait()

	for _ = range ch {
		if count++; count == n {
			break
		}
	}
	close(ch)

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
