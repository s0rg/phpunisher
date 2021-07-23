package pipe

import (
	"testing"
	"testing/fstest"
)

func TestPipeRun(t *testing.T) {
	t.Parallel()

	fs := fstest.MapFS{
		"a.txt": {
			Data: []byte("a"),
		},
		"b.txt": {
			Data: []byte("b"),
		},
		"c.tzt": {
			Data: []byte("c"),
		},
	}

	var found int

	h := func(_ *File) {
		found++
	}

	p := New(1, []string{"*.txt"}, h)

	if err := p.Walk(".", fs); err != nil {
		t.Fatal("unexpected error:", err)
	}

	if found != 2 {
		t.Fatal("unexpected found value:", found)
	}
}

func TestPipeEmpty(t *testing.T) {
	t.Parallel()

	fs := fstest.MapFS{}

	var found int

	h := func(_ *File) {
		found++
	}

	p := New(1, []string{"*.*"}, h)

	if err := p.Walk(".", fs); err != nil {
		t.Fatal("unexpected error:", err)
	}

	if found != 0 {
		t.Fatal("unexpected found value:", found)
	}
}
