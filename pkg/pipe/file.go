package pipe

import (
	"bytes"
	"io/ioutil"
)

type File struct {
	Path string
	Body bytes.Buffer
}

func (f *File) ReadBody() (err error) {
	var b []byte
	if b, err = ioutil.ReadFile(f.Path); err == nil {
		_, err = f.Body.Write(b)
	}
	return
}
