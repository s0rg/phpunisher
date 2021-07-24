package pipe

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

const bufSize int = 256

// Pipe allows to create a 'pipeline' with several goroutines.
type Pipe struct {
	masks   []string
	rq      chan *File
	rg      group
	wq      chan *File
	wg      group
	handler func(f *File)
}

// New construct new Pipe.
func New(workers int, masks []string, handler func(f *File)) *Pipe {
	p := &Pipe{
		masks:   masks,
		handler: handler,
	}

	p.rg.Workers = 1
	p.rg.Action = p.reader

	p.wg.Workers = workers
	p.wg.Action = p.worker

	return p
}

func (p *Pipe) reader() {
	for f := range p.rq {
		if err := f.ReadFull(); err != nil {
			log.Printf("reader: %s error: %v", f.Path, err)

			continue
		}

		p.wq <- f
	}
}

func (p *Pipe) worker() {
	for f := range p.wq {
		p.handler(f)
	}
}

func (p *Pipe) match(path string, fsys fs.FS) {
	name := filepath.Base(path)

	for i := 0; i < len(p.masks); i++ {
		if ok, err := filepath.Match(p.masks[i], name); err == nil && ok {
			p.rq <- FileReader(path, fsys)

			break
		}
	}
}

func (p *Pipe) start() {
	p.rq = make(chan *File, bufSize)
	// we need +1 here, to prevent consumer starvation
	p.wq = make(chan *File, p.wg.Workers+1)

	p.rg.Start(func() { close(p.rq) })
	p.wg.Start(func() { close(p.wq) })
}

func (p *Pipe) stop() {
	p.rg.Wait()
	p.wg.Wait()
}

// Walk runs pipeline for given path.
func (p *Pipe) Walk(root string, fsys fs.FS) error {
	p.start()
	defer p.stop()

	if err := fs.WalkDir(
		fsys,
		root,
		func(path string, d fs.DirEntry, err error) error {
			switch {
			case err != nil:
				return err
			case d.Type().IsDir():
			case !d.Type().IsRegular():
			default:
				p.match(path, fsys)
			}

			return nil
		},
	); err != nil {
		return fmt.Errorf("walk error: %w", err)
	}

	return nil
}
