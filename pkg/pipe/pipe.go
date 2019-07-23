package pipe

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	bufSize int = 256
)

type File struct {
	Path string
	Body bytes.Buffer
}

type Pipe struct {
	masks    []string
	read_q   chan *File
	read_grp group
	work_q   chan *File
	work_grp group
	handler  func(f *File)
}

func New(workers int, masks []string, handler func(f *File)) *Pipe {
	p := &Pipe{
		masks:   masks,
		handler: handler,
	}

	p.read_grp.Workers = 1
	p.read_grp.Action = p.reader

	p.work_grp.Workers = workers
	p.work_grp.Action = p.worker

	return p
}

func (p *Pipe) reader() {
	for f := range p.read_q {
		bts, err := ioutil.ReadFile(f.Path)
		if err != nil {
			log.Printf("reader: %s error: %v", f.Path, err)
			continue
		}
		f.Body.Write(bts)
		p.work_q <- f
	}
}

func (p *Pipe) worker() {
	for f := range p.work_q {
		p.handler(f)
	}
}

func (p *Pipe) match(path string) {
	name := filepath.Base(path)
	for i := 0; i < len(p.masks); i++ {
		if ok, err := filepath.Match(p.masks[i], name); err == nil && ok {
			p.read_q <- &File{Path: path}
			break
		}
	}
}

func (p *Pipe) walker(path string, f os.FileInfo, err error) error {
	switch {
	case err != nil:
	case f.IsDir():
	case !f.Mode().IsRegular():
	default:
		p.match(path)
	}
	return nil
}

func (p *Pipe) Walk(root string) (err error) {

	p.read_q = make(chan *File, bufSize)
	p.work_q = make(chan *File, p.work_grp.Workers+1)

	p.read_grp.Start(func() { close(p.read_q) })
	p.work_grp.Start(func() { close(p.work_q) })

	err = filepath.Walk(root, p.walker)

	p.read_grp.Drain()
	p.work_grp.Drain()
	return
}
