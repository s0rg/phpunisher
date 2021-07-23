package pipe

import (
	"bytes"
	"io/fs"
)

type File struct {
	fsys fs.FS
	Path string
	Body bytes.Buffer
}

func FileReader(path string, fsys fs.FS) *File {
	return &File{
		fsys: fsys,
		Path: path,
	}
}

func (f *File) ReadFull() (err error) {
	var b []byte

	if b, err = fs.ReadFile(f.fsys, f.Path); err == nil {
		_, err = f.Body.Write(b)
	}

	return
}
