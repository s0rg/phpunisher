package pipe

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
	"time"
)

func TestPipeRun(t *testing.T) {
	t.Parallel()

	var found int

	tfs := fstest.MapFS{
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

	p := New(1, []string{"*.txt"}, func(*File) {
		found++
	})

	if err := p.Walk(".", tfs); err != nil {
		t.Fatal("unexpected error:", err)
	}

	if found != 2 {
		t.Fatal("unexpected found value:", found)
	}
}

func TestPipeEmpty(t *testing.T) {
	t.Parallel()

	var found int

	p := New(1, []string{"*.*"}, func(*File) {
		found++
	})

	tfs := fstest.MapFS{}

	if err := p.Walk(".", tfs); err != nil {
		t.Fatal("unexpected error:", err)
	}

	if found != 0 {
		t.Fatal("unexpected found value:", found)
	}
}

type testFS struct {
	err  error
	impl fs.File
}

func (tfs *testFS) Open(name string) (fs.File, error) {
	return tfs.impl, tfs.err
}

type testFile struct {
	statErr error
	readErr error
}

func (tf *testFile) Stat() (fs.FileInfo, error) {
	nfo := &testFileInfo{}

	return nfo, tf.statErr
}

func (tf *testFile) Read(_ []byte) (int, error) {
	return 0, tf.readErr
}

func (tf *testFile) Close() error { return nil }

type testFileInfo struct{}

func (ti *testFileInfo) Name() string {
	return ""
}

func (ti *testFileInfo) Size() int64 {
	return 0
}

func (ti *testFileInfo) Mode() fs.FileMode {
	return fs.ModePerm
}

func (ti *testFileInfo) ModTime() time.Time {
	return time.Time{}
}

func (ti *testFileInfo) IsDir() bool {
	return false
}

func (ti *testFileInfo) Sys() interface{} {
	return nil
}

func TestPipeErrorFs(t *testing.T) {
	t.Parallel()

	var testError = errors.New("test fs error")

	tfs := &testFS{err: testError}

	p := New(1, []string{"*.*"}, func(*File) {})

	if err := p.Walk(".", tfs); err == nil {
		t.Error("no error")
	}
}

func TestPipeErrorFile(t *testing.T) {
	t.Parallel()

	var (
		testStatErr = errors.New("test stat error")
		testReadErr = errors.New("test read error")
	)

	tfile := &testFile{statErr: testStatErr}
	tfs := &testFS{impl: tfile}

	p := New(1, []string{"*.*"}, func(*File) {})

	err := p.Walk(".", tfs)
	if err == nil {
		t.Error("stat - no error")
	}

	if !errors.Is(err, testStatErr) {
		t.Error("not stat error")
	}

	tfile.statErr = nil
	tfile.readErr = testReadErr

	if err = p.Walk(".", tfs); err != nil {
		t.Error("walk - read error")
	}
}
