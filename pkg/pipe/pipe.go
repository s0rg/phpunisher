package pipe

import (
	"log"
	"os"
	"path/filepath"
)

const (
	bufSize int = 256
)

type Pipe struct {
	masks     []string
	readQ     chan *File
	readGroup group
	workQ     chan *File
	workGroup group
	handler   func(f *File)
}

func New(workers int, masks []string, handler func(f *File)) *Pipe {
	p := &Pipe{
		masks:   masks,
		handler: handler,
	}

	p.readGroup.Workers = 1
	p.readGroup.Action = p.reader

	p.workGroup.Workers = workers
	p.workGroup.Action = p.worker

	return p
}

func (p *Pipe) reader() {
	for f := range p.readQ {
		if err := f.ReadBody(); err != nil {
			log.Printf("reader: %s error: %v", f.Path, err)
			continue
		}
		p.workQ <- f
	}
}

func (p *Pipe) worker() {
	for f := range p.workQ {
		p.handler(f)
	}
}

func (p *Pipe) match(path string) {
	name := filepath.Base(path)
	for i := 0; i < len(p.masks); i++ {
		if ok, err := filepath.Match(p.masks[i], name); err == nil && ok {
			p.readQ <- &File{Path: path}
			break
		}
	}
}

func (p *Pipe) Walk(root string) (err error) {

	p.readQ = make(chan *File, bufSize)
	p.workQ = make(chan *File, p.workGroup.Workers+1)

	p.readGroup.Start(func() { close(p.readQ) })
	p.workGroup.Start(func() { close(p.workQ) })

	err = filepath.Walk(
		root,
		func(path string, f os.FileInfo, err error) error {
			switch {
			case err != nil:
			case f.IsDir():
			case !f.Mode().IsRegular():
			default:
				p.match(path)
			}
			return nil
		},
	)

	p.readGroup.Wait()
	p.workGroup.Wait()

	return
}
