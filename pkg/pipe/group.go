package pipe

import "sync"

type group struct {
	Workers int
	Action  func()
	closer  func()
	wg      sync.WaitGroup
}

func (g *group) actor() {
	g.Action()
	g.wg.Done()
}

func (g *group) Start(closer func()) {
	g.closer = closer
	g.wg.Add(g.Workers)
	for i := 0; i < g.Workers; i++ {
		go g.actor()
	}
}

func (g *group) Wait() {
	g.closer()
	g.wg.Wait()
}
