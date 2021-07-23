package pipe

import (
	"testing"
	"testing/fstest"
)

func TestFileOK(t *testing.T) {
	t.Parallel()

	const body byte = 'a'

	fs := fstest.MapFS{
		"a.txt": {
			Data: []byte{body},
		},
	}

	f := FileReader("a.txt", fs)

	if err := f.ReadFull(); err != nil {
		t.Fatal("unexpected error:", err)
	}

	b := f.Body.Bytes()

	if len(b) != 1 {
		t.Fatal("unexpected len:", len(b))
	}

	if b[0] != body {
		t.Fatal("unexpected body:", b)
	}
}

func TestFileErr(t *testing.T) {
	t.Parallel()

	fs := fstest.MapFS{
		"a.txt": {
			Data: []byte("a"),
		},
	}

	f := FileReader("b.txt", fs)

	if err := f.ReadFull(); err == nil {
		t.Fatal("unexpected no-error")
	}
}
