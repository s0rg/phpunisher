package pipe

import (
	"bytes"
	"io/ioutil"
)

type File struct {
	Path string
	Body bytes.Buffer
}

func (f *File) ReadBody() error {
	bts, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	_, err = f.Body.Write(bts)
	return err
}
