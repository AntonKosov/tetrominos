package ticker

import (
	"time"
)

type ID int

type Group struct {
	idGen   ID
	tickers map[ID]*ticker
}

func NewGroup() *Group {
	g := Group{tickers: map[ID]*ticker{}}
	return &g
}

func (g *Group) NewTicker(d time.Duration) (ID, <-chan struct{}) {
	id := g.idGen
	g.idGen++
	t := newTicker(d)
	g.tickers[id] = t
	return id, t.Tick
}

func (g *Group) Reset(id ID, d time.Duration) {
	t := g.tickers[id]
	t.Reset(d)
}

func (g *Group) DeleteTicker(id ID) {
	t := g.tickers[id]
	t.Done()
	delete(g.tickers, id)
}

func (g *Group) Pause() {
	for _, t := range g.tickers {
		t.Pause()
	}
}

func (g *Group) Resume() {
	for _, t := range g.tickers {
		t.Resume()
	}
}

func (g *Group) Done() {
	ids := make([]ID, 0, len(g.tickers))
	for id := range g.tickers {
		ids = append(ids, id)
	}
	for _, id := range ids {
		g.DeleteTicker(id)
	}
}
