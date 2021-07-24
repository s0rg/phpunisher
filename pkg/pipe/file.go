package pipe

import (
	"bytes"
	"io/fs"
)

// File represents a single file.
type File struct {
	fsys fs.FS
	Path string
	Body bytes.Buffer
}

// FileReader returns *File with given path in given FS.
func FileReader(path string, fsys fs.FS) *File {
	return &File{
		fsys: fsys,
		Path: path,
	}
}

// ReadFull fills Body with file content.
func (f *File) ReadFull() (err error) {
	var b []byte

	if b, err = fs.ReadFile(f.fsys, f.Path); err == nil {
		_, err = f.Body.Write(b)
	}

	return
}
